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
	// Create temporary working directory for manifests
	workingDirectory, err := ioutil.TempDir(".", "manifests-")
	if err != nil {
		return err
	}
	defer os.RemoveAll(workingDirectory)

	// Download files from App Registry
	manifests, err := h.downloader.GetManifests(request)
	if err != nil {
		return err
	}

	// Decode binary and parse yaml
	_, err = h.manifestDecoder.Decode(manifests, workingDirectory)
	if err != nil {
		return err
	}

	// Build the operator-registry container
	err = h.imageBuilder.Build(request.Image, workingDirectory, request.ImageBuildArgs)
	if err != nil {
		return err
	}

	return nil
}
