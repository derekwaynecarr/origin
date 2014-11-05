package api

import (
	"github.com/GoogleCloudPlatform/kubernetes/pkg/api"
	"github.com/openshift/origin/pkg/api2"
)

func init() {
	api.Scheme.AddKnownTypes("",
		&Project{},
		&ProjectList{},
	)
	api2.Scheme.AddKnownTypes("",
		&Project{},
		&ProjectList{},
	)
}

func (*Project) IsAnAPIObject()     {}
func (*ProjectList) IsAnAPIObject() {}
