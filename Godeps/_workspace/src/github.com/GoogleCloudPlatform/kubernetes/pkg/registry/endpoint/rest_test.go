/*
Copyright 2014 Google Inc. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package endpoint

import (
	"reflect"
	"testing"

	"github.com/GoogleCloudPlatform/kubernetes/pkg/api"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/api/errors"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/labels"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/registry/registrytest"
)

func TestGetEndpoints(t *testing.T) {
	registry := &registrytest.ServiceRegistry{
		Endpoints: api.Endpoints{
			JSONBase:  api.JSONBase{ID: "foo"},
			Endpoints: []string{"127.0.0.1:9000"},
		},
	}
	storage := NewREST(registry)
	obj, err := storage.Get("foo")
	if err != nil {
		t.Fatalf("unexpected error: %#v", err)
	}
	if !reflect.DeepEqual([]string{"127.0.0.1:9000"}, obj.(*api.Endpoints).Endpoints) {
		t.Errorf("unexpected endpoints: %#v", obj)
	}
}

func TestGetEndpointsMissingService(t *testing.T) {
	registry := &registrytest.ServiceRegistry{
		Err: errors.NewNotFound("service", "foo"),
	}
	storage := NewREST(registry)

	// returns service not found
	_, err := storage.Get("foo")
	if !errors.IsNotFound(err) || !reflect.DeepEqual(err, errors.NewNotFound("service", "foo")) {
		t.Errorf("expected NotFound error, got %#v", err)
	}

	// returns empty endpoints
	registry.Err = nil
	registry.Service = &api.Service{
		JSONBase: api.JSONBase{ID: "foo"},
	}
	obj, err := storage.Get("foo")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if obj.(*api.Endpoints).Endpoints != nil {
		t.Errorf("unexpected endpoints: %#v", obj)
	}
}

func TestEndpointsRegistryList(t *testing.T) {
	registry := registrytest.NewServiceRegistry()
	storage := NewREST(registry)
	registry.EndpointsList = api.EndpointsList{
		JSONBase: api.JSONBase{ResourceVersion: 1},
		Items: []api.Endpoints{
			{JSONBase: api.JSONBase{ID: "foo"}},
			{JSONBase: api.JSONBase{ID: "bar"}},
		},
	}
	s, _ := storage.List(labels.Everything(), labels.Everything())
	sl := s.(*api.EndpointsList)
	if len(sl.Items) != 2 {
		t.Fatalf("Expected 2 endpoints, but got %v", len(sl.Items))
	}
	if e, a := "foo", sl.Items[0].ID; e != a {
		t.Errorf("Expected %v, but got %v", e, a)
	}
	if e, a := "bar", sl.Items[1].ID; e != a {
		t.Errorf("Expected %v, but got %v", e, a)
	}
	if sl.ResourceVersion != 1 {
		t.Errorf("Unexpected resource version: %#v", sl)
	}
}
