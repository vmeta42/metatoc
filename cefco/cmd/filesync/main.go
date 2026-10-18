// Copyright 2024 Neolink.LLC
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the
// Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

package main

import (
	"flag"
	"fmt"
	"time"

	kubeinformers "k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"

	// Uncomment the following line to load the gcp plugin (only required to authenticate against GKE clusters).
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"

	fsCtl "github.com/inspursoft/cefco/controllers/filesync"
	clientset "github.com/inspursoft/cefco/pkg/generated/clientset/versioned"
	informers "github.com/inspursoft/cefco/pkg/generated/informers/externalversions"
	"github.com/inspursoft/cefco/pkg/signals"
)

var (
	Name      = "FileSync-controller"
	BuildTime = "build-time-not-set"
	GitInfo   = "gitinfo-not-set"
	Version   = "version-not-set"
)

var (
	masterURL  string
	kubeconfig string
)

func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	flag.StringVar(&masterURL, "master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
}

func main() {
	var (
		err                 error
		cfg                 *rest.Config
		kubeClient          *kubernetes.Clientset
		fsClient            *clientset.Clientset
		kubeInformerFactory kubeinformers.SharedInformerFactory
		fsInformerFactory   informers.SharedInformerFactory
		controller          *fsCtl.Controller
	)

	klog.InitFlags(nil)
	flag.Parse()

	klog.Info("Starting ", Name)
	klog.Info(fmt.Sprintf("        Version:    %s", Version))
	klog.Info(fmt.Sprintf("        BuildTime:  %s", BuildTime))
	klog.Info(fmt.Sprintf("        GitInfo:    %s", GitInfo))

	// set up signals so we handle the first shutdown signal gracefully
	stopCh := signals.SetupSignalHandler()
	if kubeconfig == "" {
		cfg, err = rest.InClusterConfig()
		if err != nil {
			klog.Fatalf("Can not read cluster config from the ServiceAccount: %s", err.Error())
		}
	} else {
		cfg, err = clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
		if err != nil {
			klog.Fatalf("Error building kubeconfig: %s", err.Error())
		}
	}

	kubeClient, err = kubernetes.NewForConfig(cfg)
	if err != nil {
		klog.Fatalf("Error building kubernetes clientset: %s", err.Error())
	}

	fsClient, err = clientset.NewForConfig(cfg)
	if err != nil {
		klog.Fatalf("Error building filesync clientset: %s", err.Error())
	}

	kubeInformerFactory = kubeinformers.NewSharedInformerFactory(kubeClient, time.Second*30)
	fsInformerFactory = informers.NewSharedInformerFactory(fsClient, time.Second*30)

	controller = fsCtl.NewController(kubeClient, fsClient,
		kubeInformerFactory.Apps().V1().Deployments(),
		kubeInformerFactory.Core().V1().Secrets(),
		kubeInformerFactory.Core().V1().ConfigMaps(),
		fsInformerFactory.Idx().V1alpha1().FileSyncs())

	// notice that there is no need to run Start methods in a separate goroutine. (i.e. go kubeInformerFactory.Start(stopCh)
	// Start method is non-blocking and runs all registered informers in a dedicated goroutine.
	kubeInformerFactory.Start(stopCh)
	fsInformerFactory.Start(stopCh)

	if err = controller.Run(2, stopCh); err != nil {
		klog.Fatalf("Error running controller: %s", err.Error())
	}
}
