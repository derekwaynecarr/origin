package v1beta1

import (
	"github.com/GoogleCloudPlatform/kubernetes/pkg/api"
	"github.com/openshift/origin/pkg/api2"
)

func init() {
	api.Scheme.AddKnownTypes("v1beta1",
		&Project{},
		&ProjectList{},
	)
	api2.Scheme.AddKnownTypes("v1beta1",
		&Project{},
		&ProjectList{},
	)

}

func (*Project) IsAnAPIObject()     {}
func (*ProjectList) IsAnAPIObject() {}
