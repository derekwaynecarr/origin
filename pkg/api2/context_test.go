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

package api2_test

import (
	"testing"

	"github.com/openshift/origin/pkg/api2"
)

// TestNamespaceContext validates that a namespace can be get/set on a context object
func TestNamespaceContext(t *testing.T) {
	ctx := api2.NewDefaultContext()
	result, ok := api2.NamespaceFrom(ctx)
	if !ok {
		t.Errorf("Error getting namespace")
	}
	if api2.NamespaceDefault != result {
		t.Errorf("Expected: %v, Actual: %v", api2.NamespaceDefault, result)
	}

	ctx = api2.NewContext()
	result, ok = api2.NamespaceFrom(ctx)
	if ok {
		t.Errorf("Should not be ok because there is no namespace on the context")
	}
}

// TestValidNamespace validates that namespace rules are enforced on a resource prior to create or update
func TestValidNamespace(t *testing.T) {
	ctx := api2.NewDefaultContext()
	namespace, _ := api2.NamespaceFrom(ctx)
	resource := api2.ReplicationController{}
	if !api2.ValidNamespace(ctx, &resource.ObjectMeta) {
		t.Errorf("expected success")
	}
	if namespace != resource.Namespace {
		t.Errorf("expected resource to have the default namespace assigned during validation")
	}
	resource = api2.ReplicationController{ObjectMeta: api2.ObjectMeta{Namespace: "other"}}
	if api2.ValidNamespace(ctx, &resource.ObjectMeta) {
		t.Errorf("Expected error that resource and context errors do not match because resource has different namespace")
	}
	ctx = api2.NewContext()
	if api2.ValidNamespace(ctx, &resource.ObjectMeta) {
		t.Errorf("Expected error that resource and context errors do not match since context has no namespace")
	}

	ctx = api2.NewContext()
	ns := api2.Namespace(ctx)
	if ns != "" {
		t.Errorf("Expected the empty string")
	}
}
