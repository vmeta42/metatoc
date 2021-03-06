/*
Copyright 2021 The idx Authors.

http://git.inspur.com/middleware/idx-component

### It needs to be supplemented ###
### It needs to be supplemented ###
### It needs to be supplemented ###
### It needs to be supplemented ###
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/inspursoft/cefco/pkg/generated/clientset/versioned/typed/filesync/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeIdxV1alpha1 struct {
	*testing.Fake
}

func (c *FakeIdxV1alpha1) FileSyncs(namespace string) v1alpha1.FileSyncInterface {
	return &FakeFileSyncs{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeIdxV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
