package strategy

import (
	"github.com/GoogleCloudPlatform/kubernetes/pkg/api"
)

// setupDockerSocket configures the pod to support the host's Docker socket
func setupDockerSocket(podSpec *api.Pod) {
	dockerSocketVolume := api.Volume{
		Name: "docker-socket",
		Source: &api.VolumeSource{
			HostDirectory: &api.HostDirectory{
				Path: "/var/run/docker.sock",
			},
		},
	}

	dockerSocketVolumeMount := api.VolumeMount{
		Name:      "docker-socket",
		MountPath: "/var/run/docker.sock",
	}

	podSpec.DesiredState.Manifest.Volumes = append(podSpec.DesiredState.Manifest.Volumes,
		dockerSocketVolume)
	podSpec.DesiredState.Manifest.Containers[0].VolumeMounts =
		append(podSpec.DesiredState.Manifest.Containers[0].VolumeMounts,
			dockerSocketVolumeMount)
}
