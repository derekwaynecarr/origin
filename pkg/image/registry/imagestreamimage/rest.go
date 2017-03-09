package imagestreamimage

import (
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	apirequest "k8s.io/apiserver/pkg/endpoints/request"

	"github.com/openshift/origin/pkg/image/api"
	"github.com/openshift/origin/pkg/image/registry/image"
	"github.com/openshift/origin/pkg/image/registry/imagestream"
)

// REST implements the RESTStorage interface in terms of an image registry and
// image stream registry. It only supports the Get method and is used
// to retrieve an image by id, scoped to an ImageStream. REST ensures
// that the requested image belongs to the specified ImageStream.
type REST struct {
	imageRegistry       image.Registry
	imageStreamRegistry imagestream.Registry
}

// NewREST returns a new REST.
func NewREST(imageRegistry image.Registry, imageStreamRegistry imagestream.Registry) *REST {
	return &REST{imageRegistry, imageStreamRegistry}
}

// New is only implemented to make REST implement RESTStorage
func (r *REST) New() runtime.Object {
	return &api.ImageStreamImage{}
}

// parseNameAndID splits a string into its name component and ID component, and returns an error
// if the string is not in the right form.
func parseNameAndID(input string) (name string, id string, err error) {
	name, id, err = api.ParseImageStreamImageName(input)
	if err != nil {
		err = errors.NewBadRequest("ImageStreamImages must be retrieved with <name>@<id>")
	}
	return
}

// Get retrieves an image by ID that has previously been tagged into an image stream.
// `id` is of the form <repo name>@<image id>.
func (r *REST) Get(ctx apirequest.Context, id string, options *metav1.GetOptions) (runtime.Object, error) {
	name, imageID, err := parseNameAndID(id)
	if err != nil {
		return nil, err
	}

	// TODO(rebase): are the GetOptions related?
	repo, err := r.imageStreamRegistry.GetImageStream(ctx, name, &metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	if repo.Status.Tags == nil {
		return nil, errors.NewNotFound(api.Resource("imagestreamimage"), id)
	}

	event, err := api.ResolveImageID(repo, imageID)
	if err != nil {
		return nil, err
	}

	imageName := event.Image
	// TODO(rebase): are the GetOptions related?
	image, err := r.imageRegistry.GetImage(ctx, imageName, &metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	if err := api.ImageWithMetadata(image); err != nil {
		return nil, err
	}
	image.DockerImageManifest = ""
	image.DockerImageConfig = ""

	isi := api.ImageStreamImage{
		ObjectMeta: metav1.ObjectMeta{
			Namespace:         apirequest.NamespaceValue(ctx),
			Name:              api.MakeImageStreamImageName(name, imageID),
			CreationTimestamp: image.ObjectMeta.CreationTimestamp,
			Annotations:       repo.Annotations,
		},
		Image: *image,
	}

	return &isi, nil
}
