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

package client2

import (
	"github.com/openshift/origin/pkg/api2"
	"github.com/openshift/origin/pkg/labels"
	"github.com/openshift/origin/pkg/util/wait"
)

// ControllerHasDesiredReplicas returns a condition that will be true iff the desired replica count
// for a controller's ReplicaSelector equals the Replicas count.
func (c *Client) ControllerHasDesiredReplicas(controller api2.ReplicationController) wait.ConditionFunc {
	return func() (bool, error) {
		pods, err := c.Pods(controller.Namespace).List(labels.Set(controller.DesiredState.ReplicaSelector).AsSelector())
		if err != nil {
			return false, err
		}
		return len(pods.Items) == controller.DesiredState.Replicas, nil
	}
}
