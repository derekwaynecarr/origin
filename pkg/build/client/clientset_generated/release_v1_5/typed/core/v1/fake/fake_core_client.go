package fake

import (
	v1 "github.com/openshift/origin/pkg/build/client/clientset_generated/release_v1_5/typed/core/v1"
	restclient "k8s.io/client-go/rest"
	core "k8s.io/kubernetes/pkg/client/testing/core"
)

type FakeCoreV1 struct {
	*core.Fake
}

func (c *FakeCoreV1) Builds(namespace string) v1.BuildInterface {
	return &FakeBuilds{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeCoreV1) RESTClient() restclient.Interface {
	var ret *restclient.RESTClient
	return ret
}
