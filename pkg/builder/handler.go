package builder

import (
	"io/ioutil"
	"os"

	"github.com/kevinrizza/offline-cataloger/pkg/apis"
	"github.com/kevinrizza/offline-cataloger/pkg/appregistry"
)

// NewHandler is a constructor for the Handler interface
func NewHandler() (Handler, error) {
	decoder, err := appregistry.NewManifestDecoder()
	if err != nil {
		return nil, err
	}
	return &handler{
		downloader:      NewDownloader(),
		imageBuilder:    NewImageBuilder(),
		manifestDecoder: decoder,
	}, nil
}

// Handler is an interface that is implemented by structs
// that implement the Handle method. A Handler takes BuildRequests
// as input and builds an operator-registry image from that input.
//
// It downloads operator manifests from a specified app registry,
// decodes them into files and then calls docker build to generate
// the operator-registry image.
type Handler interface {
	Handle(request *apis.BuildRequest) error
}

type handler struct {
	downloader      Downloader
	imageBuilder    ImageBuilder
	manifestDecoder appregistry.ManifestDecoder
}

func (h *handler) Handle(request *apis.BuildRequest) error {
	// create temp directory for manifests
	workingDirectory, err := ioutil.TempDir(".", "manifests-")
	if err != nil {
		return err
	}
	defer os.RemoveAll(workingDirectory)

	// download files
	manifests, err := h.downloader.GetManifests(request)
	if err != nil {
		return err
	}

	// decode binary and parse yaml
	_, err = h.manifestDecoder.Decode(manifests, workingDirectory)
	if err != nil {
		return err
	}

	// parse yaml for additional images

	// create dockerfile pointing to the yaml
	// docker build
	err = h.imageBuilder.Build(request.Image, workingDirectory)
	if err != nil {
		return err
	}

	// pull additional images to local registry

	return nil
}
