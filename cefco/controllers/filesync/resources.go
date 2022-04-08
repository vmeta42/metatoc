package filesync

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	filesyncv1alpha1 "github.com/inspursoft/cefco/pkg/apis/filesync/v1alpha1"
)

// volInfo is used to create a pvc or hostpath volume
type volInfo struct {
	name    string
	path    *string
	pvc     *string
	subPath *string
}

// createVolInfoMap create a vol map
func createVolInfoMap(fs *filesyncv1alpha1.FileSync, isReader bool) (vm map[string]*volInfo) {
	var rangeKeys []string
	if isReader {
		vm = map[string]*volInfo{
			"src": {
				name: "src",
				path: &fs.Spec.Reader.SrcPath,
				pvc:  fs.Spec.Reader.SrcPathPVC,
			},
			"shadow": {
				name: "shadow",
				path: fs.Spec.Reader.ShadowPath,
				pvc:  fs.Spec.Reader.ShadowPathPVC,
			},
			"log": {
				name: "log",
				path: fs.Spec.Reader.LogPath,
				pvc:  fs.Spec.Reader.LogPathPVC,
			},
			"trash": {
				name: "trash",
				path: fs.Spec.Reader.TrashPath,
				pvc:  fs.Spec.Reader.TrashPathPVC,
			},
		}
		// 顺序同 vm 的 key
		rangeKeys = []string{"src", "shadow", "log", "trash"}
	} else {
		vm = map[string]*volInfo{
			"rcv": {
				name: "rcv",
				path: &fs.Spec.Writer.RcvPath,
				pvc:  fs.Spec.Writer.RcvPathPVC,
			},
			"log": {
				name: "log",
				path: fs.Spec.Writer.LogPath,
				pvc:  fs.Spec.Writer.LogPathPVC,
			},
			"trash": {
				name: "trash",
				path: fs.Spec.Writer.TrashPath,
				pvc:  fs.Spec.Writer.TrashPathPVC,
			},
		}
		// 顺序同 vm 的 key
		rangeKeys = []string{"rcv", "log", "trash"}
	}

	// if there is duplicate PVC, add subpath to it
	pvcCount := map[string]*volInfo{}
	// map 遍历顺序是随机的, 所以通过遍历索引 slice 来唯一确定结果
	for _, k := range rangeKeys {
		v := vm[k]
		if v.pvc != nil {
			if vi, exist := pvcCount[*v.pvc]; exist {
				vi.subPath = &vi.name
				p := k
				v.subPath = &p
				v.name = vi.name
			} else {
				pvcCount[*v.pvc] = v
			}
		}
	}
	return
}

// newDeployment creates new Deployments for a FileSync resource. It also sets
// the appropriate OwnerReferences on the resources so handleObject can discover
// the FileSync resource that 'owns' them.
func newDeployment(fs *filesyncv1alpha1.FileSync, isReader bool) (*appsv1.Deployment, error) {
	createDeploy := func(name, image string, selectedLabels, fullLabels, nodeSelector map[string]string,
		envs []corev1.EnvVar, vl []corev1.Volume, vml []corev1.VolumeMount,
		tol []corev1.Toleration, ps []corev1.LocalObjectReference) *appsv1.Deployment {
		return &appsv1.Deployment{
			ObjectMeta: metav1.ObjectMeta{
				Name:      name,
				Namespace: fs.Namespace,
				OwnerReferences: []metav1.OwnerReference{
					*metav1.NewControllerRef(fs, filesyncv1alpha1.SchemeGroupVersion.WithKind("FileSync")),
				},
			},
			Spec: appsv1.DeploymentSpec{
				Selector: &metav1.LabelSelector{
					MatchLabels: selectedLabels,
				},
				Template: corev1.PodTemplateSpec{
					ObjectMeta: metav1.ObjectMeta{
						Labels: fullLabels,
					},
					Spec: corev1.PodSpec{
						NodeSelector: nodeSelector,
						Containers: []corev1.Container{{
							Image: image,
							Name:  name,
							Env:   envs,
							Ports: []corev1.ContainerPort{{
								ContainerPort: 4222,
								Name:          "nats",
								Protocol:      corev1.ProtocolTCP,
							}},
							SecurityContext: &corev1.SecurityContext{
								Capabilities: &corev1.Capabilities{
									Add: []corev1.Capability{"SYS_PTRACE"},
								},
							},
							VolumeMounts: vml,
						}},
						Volumes:          vl,
						Tolerations:      tol,
						ImagePullSecrets: ps,
					},
				},
			},
		}
	}
	createFullLabels := func(selectedLabels map[string]string) map[string]string {
		var sourceLabels map[string]string
		fullLabels := make(map[string]string)
		// deep copy from selectedLabels
		for k, v := range selectedLabels {
			fullLabels[k] = v
		}
		if isReader {
			sourceLabels = fs.Spec.Reader.Labels
		} else {
			sourceLabels = fs.Spec.Writer.Labels
		}
		if len(sourceLabels) > 0 {
			fullLabels["custom-labels"] = "true"
			for k, v := range sourceLabels {
				if _, exist := fullLabels[k]; !exist {
					fullLabels[k] = v
				}
			}
		}
		return fullLabels
	}
	selectedLabels := map[string]string{
		"app":        "filesync",
		"controller": fs.Name,
	}
	if isReader {
		if fs.Spec.Reader != nil {
			selectedLabels["type"] = "reader"
			fullLabels := createFullLabels(selectedLabels)
			name := fs.Name + "-r"
			vmap := createVolInfoMap(fs, isReader)
			vl, vml := createVolumesForDeployment(fs, vmap)
			envs := []corev1.EnvVar{
				{Name: "START_CMD", Value: "FileReader -m start -f /filesync -t inc_always -D"},
				{Name: "STOP_CMD", Value: "FileReader -m stop"},
				{Name: "PROCESS_NAME", Value: "FileReader"},
			}
			return createDeploy(name, fs.Spec.Reader.Image, selectedLabels, fullLabels, fs.Spec.Reader.NodeSelector,
				envs, vl, vml, fs.Spec.Reader.Tolerations, fs.Spec.Reader.ImagePullSecrets), nil
		} else {
			return nil, fmt.Errorf("expect to create a reader deployment, but can not get reader in filesync object: %s", fs.Name)
		}
	} else {
		if fs.Spec.Writer != nil {
			selectedLabels["type"] = "writer"
			fullLabels := createFullLabels(selectedLabels)
			name := fs.Name + "-w"
			vmap := createVolInfoMap(fs, isReader)
			vl, vml := createVolumesForDeployment(fs, vmap)
			envs := []corev1.EnvVar{
				{Name: "START_CMD", Value: "FileWriter -m start -f /filesync"},
				{Name: "STOP_CMD", Value: "FileWriter -m stop -f /filesync"},
				{Name: "PROCESS_NAME", Value: "FileWriter"},
			}
			return createDeploy(name, fs.Spec.Writer.Image, selectedLabels, fullLabels, fs.Spec.Writer.NodeSelector,
				envs, vl, vml, fs.Spec.Writer.Tolerations, fs.Spec.Writer.ImagePullSecrets), nil
		} else {
			return nil, fmt.Errorf("expect to create a writer deployment, but can not get writer in filesync object: %s", fs.Name)
		}
	}
}

type volumeType string

var (
	configmapVolume     volumeType = "configmap"
	secretVolume        volumeType = "secret"
	hostpathOrPVCVolume volumeType = "hostpathOrPVC"
)

func createVolumesForDeployment(fs *filesyncv1alpha1.FileSync, volumes map[string]*volInfo) (vl []corev1.Volume, vml []corev1.VolumeMount) {
	hostVolumeType := corev1.HostPathDirectoryOrCreate
	fsName := fs.Name
	createVS := func(vType volumeType, vol *volInfo) corev1.VolumeSource {
		var res corev1.VolumeSource
		switch vType {
		case configmapVolume:
			res = corev1.VolumeSource{
				ConfigMap: &corev1.ConfigMapVolumeSource{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: fsName,
					},
				},
			}
		case secretVolume:
			res = corev1.VolumeSource{
				Secret: &corev1.SecretVolumeSource{
					SecretName: *fs.Spec.NATS.CredsSecret,
				},
			}
		case hostpathOrPVCVolume:
			if vol.pvc == nil {
				// PVC does not exist, create a hostpath
				res = corev1.VolumeSource{
					HostPath: &corev1.HostPathVolumeSource{
						Path: *vol.path,
						Type: &hostVolumeType,
					},
				}
			} else {
				// PVC exists, create a pvc
				res = corev1.VolumeSource{
					PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
						ClaimName: *vol.pvc,
					},
				}
			}
		}
		return res
	}
	createV := func(name string, vType volumeType, vol *volInfo) corev1.Volume {
		return corev1.Volume{
			Name:         name,
			VolumeSource: createVS(vType, vol),
		}
	}
	createVM := func(name string, vol *volInfo) corev1.VolumeMount {
		vm := corev1.VolumeMount{
			Name:      name,
			MountPath: *vol.path,
		}
		if vol.subPath != nil {
			vm.SubPath = *vol.subPath
		}
		return vm
	}
	vl = append(vl, createV(fsName, configmapVolume, nil))
	cmPath := "/filesync/cfg"
	vml = append(vml, createVM(fsName, &volInfo{path: &cmPath}))
	if fs.Spec.NATS.CredsSecret != nil {
		vl = append(vl, createV(*fs.Spec.NATS.CredsSecret, secretVolume, nil))
		vml = append(vml, createVM(*fs.Spec.NATS.CredsSecret, &volInfo{path: &credsBasePath}))
	}
	for k, v := range volumes {
		if v.path != nil {
			vl = append(vl, createV(k, hostpathOrPVCVolume, v))
			vml = append(vml, createVM(v.name, v))
		}
	}
	return
}

func newComfigMap(fs *filesyncv1alpha1.FileSync) (*corev1.ConfigMap, error) {
	tmpl, err := template.New("filereader").Parse(fileReaderConfig)
	if err != nil {
		return nil, fmt.Errorf("can not read template form fileReaderConfig: %v", err)
	}
	tmpl, err = tmpl.New("filewriter").Parse(fileWriterConfig)
	if err != nil {
		return nil, fmt.Errorf("can not read template form fileWriterConfig: %v", err)
	}
	data := map[string]string{
		"reader.cfg": "",
		"writer.cfg": "",
	}
	buf := new(bytes.Buffer)
	if fs.Spec.Reader != nil {
		err = tmpl.ExecuteTemplate(buf, "filereader", fs.Spec)
		if err != nil {
			return nil, fmt.Errorf("can not prase templates form filereader.tpl: %v", err)
		}
		data["reader.cfg"] = buf.String()
		buf = new(bytes.Buffer)
	}
	if fs.Spec.Writer != nil {
		err = tmpl.ExecuteTemplate(buf, "filewriter", fs.Spec)
		if err != nil {
			return nil, fmt.Errorf("can not prase templates form filewriter.tpl: %v", err)
		}
		data["writer.cfg"] = buf.String()
	}
	return &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      fs.Name,
			Namespace: fs.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(fs, filesyncv1alpha1.SchemeGroupVersion.WithKind("FileSync")),
			},
		},
		Data: data,
	}, nil
}

// getCredsPathFromFileSyncObj will store the creds to controller and
// if the NATS SecretKey does not exist, this func will update it.
func (c *Controller) getCredsPathFromFileSyncObj(fs *filesyncv1alpha1.FileSync) (credsFile string, err error) {
	if fs.Spec.NATS.CredsSecret != nil {
		var (
			secret *corev1.Secret
			value  []byte
		)
		isFileIsExist := func(filename string) bool {
			_, e := os.Stat(filename)
			return e == nil
		}

		// by default, it was '/nats-accounts/creds/{namespace}/{name}/'
		credsPath := credsBasePath + fs.Namespace + "/" + fs.Name + "/"
		if !isFileIsExist(credsPath) {
			err = os.MkdirAll(credsPath, 0666)
			if err != nil {
				err = fmt.Errorf("cannot creat folder: %w", err)
				return
			}
		}
		secret, err = c.secretsLister.Secrets(fs.Namespace).Get(*fs.Spec.NATS.CredsSecret)
		if err != nil {
			err = fmt.Errorf("cannot get the target secret %q in namespace %q: %w", *fs.Spec.NATS.CredsSecret, fs.Namespace, err)
			return
		}
		if fs.Spec.NATS.SecretKey != "" {
			credsFile = credsPath + fs.Spec.NATS.SecretKey
			v, exist := secret.Data[fs.Spec.NATS.SecretKey]
			if !exist {
				err = fmt.Errorf("cannot found target key %q in secret %q", fs.Spec.NATS.SecretKey, *fs.Spec.NATS.CredsSecret)
				return
			}
			value = v
		} else if len(secret.Data) == 1 {
			for k, v := range secret.Data {
				credsFile = credsPath + k
				value = v
				fs.Spec.NATS.SecretKey = k
			}
		} else {
			err = fmt.Errorf("too much secrets in %q", *fs.Spec.NATS.CredsSecret)
			return
		}

		err = ioutil.WriteFile(credsFile, value, 0666)
		if err != nil {
			err = fmt.Errorf("cannot create creds for nats: %w", err)
			return
		}
	}
	return
}
