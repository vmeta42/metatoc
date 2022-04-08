package filesync

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/nats-io/nats.go"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	appsinformers "k8s.io/client-go/informers/apps/v1"
	coreinformers "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	typedcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	appslisters "k8s.io/client-go/listers/apps/v1"
	corelisters "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog/v2"

	filesyncv1alpha1 "github.com/inspursoft/cefco/pkg/apis/filesync/v1alpha1"
	clientset "github.com/inspursoft/cefco/pkg/generated/clientset/versioned"
	filesyncscheme "github.com/inspursoft/cefco/pkg/generated/clientset/versioned/scheme"
	informers "github.com/inspursoft/cefco/pkg/generated/informers/externalversions/filesync/v1alpha1"
	listers "github.com/inspursoft/cefco/pkg/generated/listers/filesync/v1alpha1"
)

const controllerAgentName = "filesync-controller"

const (
	// OnCreating is used as part of the Event 'reason' when a FileSync is creating its own resources
	OnCreating = "Creating"
	// OnUpdating is used as part of the Event 'reason' when a FileSync is updating its own resources
	OnUpdating = "Updating"
	// SuccessSynced is used as part of the Event 'reason' when a FileSync is synced
	SuccessSynced = "Synced"
	// FailSynced is used as part of the Event 'reason' when a FileSync synced failed
	FailSynced = "Failed"
	// ErrResourceExists is used as part of the Event 'reason' when a FileSync fails
	// to sync due to a Deployment of the same name already existing.
	ErrResourceExists = "ErrResourceExists"

	// MessageResourceExists is the message used for Events when a resource
	// fails to sync due to a Deployment already existing
	MessageResourceExists = "Resource %q already exists and is not managed by FileSync"
	// MessageResourceSynced is the message used for an Event fired when a FileSync
	// is synced successfully
	MessageResourceSynced = "FileSync synced successfully"

	MessageResourceModifyFailed = "Can not modify an exist Filesync object"
)

// Controller is the controller implementation for FileSync resources
type Controller struct {
	// kubeclientset is a standard kubernetes clientset
	kubeclientset kubernetes.Interface
	// filesyncclientset is a clientset for our own API group
	filesyncclientset clientset.Interface

	deploymentsLister appslisters.DeploymentLister
	deploymentsSynced cache.InformerSynced

	secretsLister corelisters.SecretLister

	configmapsLister corelisters.ConfigMapLister
	configmapsSynced cache.InformerSynced

	filesyncsLister listers.FileSyncLister
	filesyncsSynced cache.InformerSynced

	// workqueue is a rate limited work queue. This is used to queue work to be
	// processed instead of performing it as soon as a change happens. This
	// means we can ensure we only process a fixed amount of resources at a
	// time, and makes it easy to ensure we are never processing the same item
	// simultaneously in two different workers.
	workqueue workqueue.RateLimitingInterface
	// recorder is an event recorder for recording Event resources to the
	// Kubernetes API.
	recorder record.EventRecorder
}

// NewController returns a new filesync controller
func NewController(
	kubeclientset kubernetes.Interface,
	filesyncclientset clientset.Interface,
	deploymentInformer appsinformers.DeploymentInformer,
	secretInformer coreinformers.SecretInformer,
	configmapInformer coreinformers.ConfigMapInformer,
	filesyncInformer informers.FileSyncInformer) *Controller {

	// Create event broadcaster
	// Add sample-controller types to the default Kubernetes Scheme so Events can be
	// logged for sample-controller types.
	utilruntime.Must(filesyncscheme.AddToScheme(scheme.Scheme))
	klog.V(4).Info("Creating event broadcaster")
	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartStructuredLogging(0)
	eventBroadcaster.StartRecordingToSink(&typedcorev1.EventSinkImpl{Interface: kubeclientset.CoreV1().Events("")})
	recorder := eventBroadcaster.NewRecorder(scheme.Scheme, corev1.EventSource{Component: controllerAgentName})

	controller := &Controller{
		kubeclientset:     kubeclientset,
		filesyncclientset: filesyncclientset,

		deploymentsLister: deploymentInformer.Lister(),
		deploymentsSynced: deploymentInformer.Informer().HasSynced,

		secretsLister: secretInformer.Lister(),

		configmapsLister: configmapInformer.Lister(),
		configmapsSynced: configmapInformer.Informer().HasSynced,

		filesyncsLister: filesyncInformer.Lister(),
		filesyncsSynced: filesyncInformer.Informer().HasSynced,

		workqueue: workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "FileSyncs"),
		recorder:  recorder,
	}

	klog.Info("Setting up event handlers")
	// Set up an event handler for when FileSync resources change
	filesyncInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: controller.enqueueFileSync,
		UpdateFunc: func(old, new interface{}) {
			controller.enqueueFileSync(new)
		},
	})

	// Set up an event handler for when Deployment resources change. This
	// handler will lookup the owner of the given Deployment, and if it is
	// owned by a FileSync resource will enqueue that FileSync resource for
	// processing. This way, we don't need to implement custom logic for
	// handling Deployment resources. More info on this pattern:
	// https://github.com/kubernetes/community/blob/8cafef897a22026d42f5e5bb3f104febe7e29830/contributors/devel/controllers.md
	deploymentInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: controller.handleObject,
		UpdateFunc: func(old, new interface{}) {
			newDepl := new.(*appsv1.Deployment)
			oldDepl := old.(*appsv1.Deployment)
			if newDepl.ResourceVersion == oldDepl.ResourceVersion {
				// Periodic resync will send update events for all known Deployments.
				// Two different versions of the same Deployment will always have different RVs.
				return
			}
			controller.handleObject(new)
		},
		DeleteFunc: controller.handleObject,
	})

	// as same as Deployment
	configmapInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: controller.handleObject,
		UpdateFunc: func(old, new interface{}) {
			newCM := new.(*corev1.ConfigMap)
			oldCM := old.(*corev1.ConfigMap)
			if newCM.ResourceVersion == oldCM.ResourceVersion {
				return
			}
			controller.handleObject(new)
		},
		DeleteFunc: controller.handleObject,
	})

	return controller
}

// Run will set up the event handlers for types we are interested in, as well
// as syncing informer caches and starting workers. It will block until stopCh
// is closed, at which point it will shutdown the workqueue and wait for
// workers to finish processing their current work items.
func (c *Controller) Run(threadiness int, stopCh <-chan struct{}) error {
	defer utilruntime.HandleCrash()
	defer c.workqueue.ShutDown()

	// Start the informer factories to begin populating the informer caches
	klog.Info("FileSync controller is running")

	// Wait for the caches to be synced before starting workers
	klog.Info("Waiting for informer caches to sync")
	if ok := cache.WaitForCacheSync(stopCh, c.deploymentsSynced, c.configmapsSynced, c.filesyncsSynced); !ok {
		return fmt.Errorf("failed to wait for caches to sync")
	}

	klog.Info("Starting workers")
	// Launch two workers to process Filesync resources
	for i := 0; i < threadiness; i++ {
		go wait.Until(c.runWorker, time.Second, stopCh)
	}

	klog.Info("Started workers")
	<-stopCh
	klog.Info("Shutting down workers")

	return nil
}

// runWorker is a long-running function that will continually call the
// processNextWorkItem function in order to read and process a message on the
// workqueue.
func (c *Controller) runWorker() {
	for c.processNextWorkItem() {
	}
}

// processNextWorkItem will read a single work item off the workqueue and
// attempt to process it, by calling the syncHandler.
func (c *Controller) processNextWorkItem() bool {
	obj, shutdown := c.workqueue.Get()

	if shutdown {
		return false
	}

	// We wrap this block in a func so we can defer c.workqueue.Done.
	err := func(obj interface{}) error {
		// We call Done here so the workqueue knows we have finished
		// processing this item. We also must remember to call Forget if we
		// do not want this work item being re-queued. For example, we do
		// not call Forget if a transient error occurs, instead the item is
		// put back on the workqueue and attempted again after a back-off
		// period.
		defer c.workqueue.Done(obj)
		var key string
		var ok bool
		// We expect strings to come off the workqueue. These are of the
		// form namespace/name. We do this as the delayed nature of the
		// workqueue means the items in the informer cache may actually be
		// more up to date that when the item was initially put onto the
		// workqueue.
		if key, ok = obj.(string); !ok {
			// As the item in the workqueue is actually invalid, we call
			// Forget here else we'd go into a loop of attempting to
			// process a work item that is invalid.
			c.workqueue.Forget(obj)
			utilruntime.HandleError(fmt.Errorf("expected string in workqueue but got %#v", obj))
			return nil
		}
		// Run the syncHandler, passing it the namespace/name string of the
		// FileSync resource to be synced.
		if err := c.syncHandler(key); err != nil {
			// Put the item back on the workqueue to handle any transient errors.
			c.workqueue.AddRateLimited(key)
			return fmt.Errorf("error syncing '%s': %s, requeuing", key, err.Error())
		}
		// Finally, if no error occurs we Forget this item so it does not
		// get queued again until another change happens.
		c.workqueue.Forget(obj)
		klog.V(4).Infof("Successfully synced '%s'", key)
		return nil
	}(obj)

	if err != nil {
		utilruntime.HandleError(err)
		return true
	}

	return true
}

// syncHandler compares the actual state with the desired, and attempts to
// converge the two. It then updates the Status block of the FileSync resource
// with the current status of the resource.
func (c *Controller) syncHandler(key string) error {
	// Convert the namespace/name string into a distinct namespace and name
	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("invalid resource key: %s", key))
		return nil
	}

	// Get the FileSync resource with this namespace/name
	fs, err := c.filesyncsLister.FileSyncs(namespace).Get(name)
	if err != nil {
		// The FileSync resource may no longer exist, in which case we stop
		// processing.
		if errors.IsNotFound(err) {
			utilruntime.HandleError(fmt.Errorf("filesync '%s' in work queue no longer exists", key))
			return nil
		}

		return err
	}

	// check resource version
	if fs.Status.ObservedGeneration != 0 && fs.Generation != fs.Status.ObservedGeneration {
		c.recorder.Event(fs, corev1.EventTypeWarning, FailSynced, MessageResourceModifyFailed)
		return nil
	}

	if fs.Spec.Reader == nil && fs.Spec.Writer == nil {
		return fmt.Errorf("cannot create a filesync object without reader or writer")
	}

	natsServer := fs.Spec.NATS.IP
	jetStr := ""
	jetCon := ""

	deploys := []string{}
	if fs.Spec.Reader != nil {
		deploys = append(deploys, fs.Name+"-r")
		if !fs.Spec.Reader.DontCheckDataTopic {
			jetStr = fs.Spec.Reader.DataTopic
			if len(strings.Split(jetStr, ".")) > 1 {
				return fmt.Errorf("unable to identify Stream from topic %q", jetStr)
			}
		}
	}
	if fs.Spec.Writer != nil {
		deploys = append(deploys, fs.Name+"-w")
		if !fs.Spec.Writer.DontCheckDataTopic {
			strCon := strings.Split(fs.Spec.Writer.DataTopic, ".")
			if len(strCon) != 2 {
				return fmt.Errorf("unable to identify Stream and Consumer from topic %q", fs.Spec.Writer.DataTopic)
			}
			if jetStr != "" && strCon[0] != jetStr {
				return fmt.Errorf("the reader and the writer are in different streams: reader: %q, writer: %q", jetStr, strCon[0])
			}
			if strCon[0] != strCon[1] {
				return fmt.Errorf("stream and consumer are not the same: stream: %q, consumer: %q", strCon[0], strCon[1])
			}
			jetStr = strCon[0]
			jetCon = strCon[1]
		}
	}

	credsFile, err := c.getCredsPathFromFileSyncObj(fs)
	if err != nil {
		return err
	}

	if jetStr != "" {
		opts := make([]nats.Option, 0)
		if credsFile != "" {
			opts = append(opts, nats.UserCredentials(credsFile))
		} else if fs.Spec.NATS.User != "" || fs.Spec.NATS.Password != "" {
			opts = append(opts, nats.UserInfo(fs.Spec.NATS.User, fs.Spec.NATS.Password))
		}

		jc, err := NewJsmClient(natsServer, opts...)
		if err != nil {
			return fmt.Errorf("can not create a new jetstream client: %w", err)
		}
		subjs := []string{jetStr}
		if ok, err := jc.CreateDefaultStreamIfNotExist(jetStr, subjs...); err == nil {
			if ok {
				klog.V(4).Infof("The target stream does not exist, created a new stream: %s", jetStr)
			} else {
				klog.V(4).Infof("Using an exist stream: %s", jetStr)
			}
			if jetCon != "" {
				if ok, err = jc.CreateDefaultConsumerIfNotExist(jetStr, jetCon, jetCon); err == nil {
					if ok {
						klog.V(4).Infof("The target consumer does not exist, created a new consumer %q", jetCon)
					} else {
						klog.V(4).Infof("Using an exist consumer: %s", jetCon)
					}
				} else {
					return fmt.Errorf("can not create or use an exist consumer %q: %w", jetCon, err)
				}
			}
		} else {
			return fmt.Errorf("can not create or use an exist stream %q: %w", jetStr, err)
		}
		jc.Close()
	}

	// Get the comfigmap with the name specified in FileSync.spec
	cm, err := c.configmapsLister.ConfigMaps(fs.Namespace).Get(fs.Name)
	if errors.IsNotFound(err) {
		var newCM *corev1.ConfigMap
		newCM, err = newComfigMap(fs)
		if err != nil {
			return fmt.Errorf("create new comfigmap error: %v", err)
		}
		cm, err = c.kubeclientset.CoreV1().ConfigMaps(fs.Namespace).Create(context.TODO(), newCM, metav1.CreateOptions{})
	}
	if err != nil {
		return err
	}

	if !metav1.IsControlledBy(cm, fs) {
		msg := fmt.Sprintf(MessageResourceExists, cm.Name)
		c.recorder.Event(fs, corev1.EventTypeWarning, ErrResourceExists, msg)
		return fmt.Errorf(msg)
	}

	deployList := []*appsv1.Deployment{}
	for _, deploymentName := range deploys {
		// Get the deployment with the name specified in FileSync.spec
		deployment, err := c.deploymentsLister.Deployments(fs.Namespace).Get(deploymentName)
		// If the resource doesn't exist, we'll create it
		if errors.IsNotFound(err) {
			isReader := true
			if deploymentName == fs.Name+"-w" {
				isReader = false
			}
			var deploy *appsv1.Deployment
			deploy, err = newDeployment(fs, isReader)
			if err != nil {
				return fmt.Errorf("create new deployment %s error: %v", deploymentName, err)
			}
			deployment, err = c.kubeclientset.AppsV1().Deployments(fs.Namespace).Create(context.TODO(), deploy, metav1.CreateOptions{})
		}

		// If an error occurs during Get/Create, we'll requeue the item so we can
		// attempt processing again later. This could have been caused by a
		// temporary network failure, or any other transient reason.
		if err != nil {
			return err
		}

		// If the Deployment is not controlled by this FileSync resource, we should log
		// a warning to the event recorder and return error msg.
		if !metav1.IsControlledBy(deployment, fs) {
			msg := fmt.Sprintf(MessageResourceExists, deployment.Name)
			c.recorder.Event(fs, corev1.EventTypeWarning, ErrResourceExists, msg)
			return fmt.Errorf(msg)
		}
		deployList = append(deployList, deployment)
	}

	// Finally, we update the status block of the FileSync resource to reflect the
	// current state of the world
	err = c.updateFileSyncStatus(fs, cm, deployList)
	if err != nil {
		return err
	}
	c.recorder.Event(fs, corev1.EventTypeNormal, SuccessSynced, MessageResourceSynced)
	return nil
}

func (c *Controller) updateFileSyncStatus(fs *filesyncv1alpha1.FileSync, cm *corev1.ConfigMap, deployList []*appsv1.Deployment) error {
	// NEVER modify objects from the store. It's a read-only, local cache.
	// You can use DeepCopy() to make a deep copy of original object and modify this copy
	// Or create a copy manually for better performance
	fsCopy := fs.DeepCopy()
	fsCopy.Status.ObservedGeneration = fs.Generation
	if cm != nil {
		fsCopy.Status.ConfigMapStatus = SuccessSynced
	}
	for _, deploy := range deployList {
		status := OnCreating
		if deploy.Status.Replicas == deploy.Status.AvailableReplicas {
			status = SuccessSynced
		}
		if deploy.Name == fs.Name+"-r" {
			fsCopy.Status.ReaderStatus = &status
		} else {
			fsCopy.Status.WriterStatus = &status
		}
	}

	// If the CustomResourceSubresources feature gate is not enabled,
	// we must use Update instead of UpdateStatus to update the Status block of the FileSync resource.
	// UpdateStatus will not allow changes to the Spec of the resource,
	// which is ideal for ensuring nothing other than resource status has been updated.
	_, err := c.filesyncclientset.IdxV1alpha1().FileSyncs(fs.Namespace).UpdateStatus(context.TODO(), fsCopy, metav1.UpdateOptions{})
	return err
}

// enqueueFileSync takes a FileSync resource and converts it into a namespace/name
// string which is then put onto the work queue. This method should *not* be
// passed resources of any type other than FileSync.
func (c *Controller) enqueueFileSync(obj interface{}) {
	var key string
	var err error
	if key, err = cache.MetaNamespaceKeyFunc(obj); err != nil {
		utilruntime.HandleError(err)
		return
	}
	c.workqueue.Add(key)
}

// handleObject will take any resource implementing metav1.Object and attempt
// to find the FileSync resource that 'owns' it. It does this by looking at the
// objects metadata.ownerReferences field for an appropriate OwnerReference.
// It then enqueues that FileSync resource to be processed. If the object does not
// have an appropriate OwnerReference, it will simply be skipped.
func (c *Controller) handleObject(obj interface{}) {
	var object metav1.Object
	var ok bool
	if object, ok = obj.(metav1.Object); !ok {
		tombstone, ok := obj.(cache.DeletedFinalStateUnknown)
		if !ok {
			utilruntime.HandleError(fmt.Errorf("error decoding object, invalid type"))
			return
		}
		object, ok = tombstone.Obj.(metav1.Object)
		if !ok {
			utilruntime.HandleError(fmt.Errorf("error decoding object tombstone, invalid type"))
			return
		}
		klog.V(4).Infof("Recovered deleted object '%s' from tombstone", object.GetName())
	}
	klog.V(4).Infof("Processing object: %s", object.GetName())
	if ownerRef := metav1.GetControllerOf(object); ownerRef != nil {
		// If this object is not owned by a FileSync, we should not do anything more
		// with it.
		if ownerRef.Kind != "FileSync" {
			return
		}

		fsl, err := c.filesyncsLister.FileSyncs(object.GetNamespace()).Get(ownerRef.Name)
		if err != nil {
			klog.V(4).Infof("ignoring orphaned object '%s' of filesync '%s'", object.GetSelfLink(), ownerRef.Name)
			return
		}

		c.enqueueFileSync(fsl)
		return
	}
}
