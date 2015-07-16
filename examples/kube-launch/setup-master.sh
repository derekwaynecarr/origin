#!/bin/bash

# This script initializes a vanilla OpenShift Origin to run Kubernetes Launch demo
# This script should be run as root after ssh-ing into a Vagrant box

# Create a service account to run the docker registry in the default namespace
echo '{"kind":"ServiceAccount","apiVersion":"v1","metadata":{"name":"registry"}}' | openshift cli create -n default -f -

# Recreate the privileged security context constraints to include the registry
openshift cli delete scc privileged -n default

echo 'allowHostDirVolumePlugin: true
allowPrivilegedContainer: true
apiVersion: v1
groups:
- system:cluster-admins
- system:nodes
kind: SecurityContextConstraints
metadata:
  name: privileged
runAsUser:
  type: RunAsAny
seLinuxContext:
  type: RunAsAny
users:
- system:serviceaccount:openshift-infra:build-controller
- system:serviceaccount:default:registry' | openshift cli create -n default -f -

# Create the docker registry
openshift admin registry --service-account=registry \
    --credentials=/openshift.local.config/master/openshift-registry.kubeconfig \
    --mount-host=/openshift.local.registry

# Start OpenShift Router
openshift admin router router --create \
     --credentials=/openshift.local.config/master/openshift-router.kubeconfig

# Populate Image Streams
pushd /data/src/github.com/openshift/origin/examples/image-streams
openshift cli create -f ./image-streams-centos7.json -n openshift
popd
pushd /data/src/github.com/openshift/origin/examples/kube-launch/application-templates
openshift cli create -f ./jboss-image-streams.json -n openshift
popd

# Populate Shared Templates
pushd /data/src/github.com/openshift/origin/examples/db-templates
openshift cli create -f . -n openshift
popd
pushd /data/src/github.com/openshift/origin/examples/kube-launch/application-templates
openshift cli create -f amq -f eap -f webserver -n openshift
popd