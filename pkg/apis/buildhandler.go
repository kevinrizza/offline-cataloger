package apis

import (
	"io/ioutil"
	"os"

	"github.com/kevinrizza/offline-cataloger/pkg/appregistry"
	"github.com/kevinrizza/offline-cataloger/pkg/builder"
	"github.com/kevinrizza/offline-cataloger/pkg/downloader"
)

// BuildRequest is a struct to describe the API used by
// the command line package to make requests to the builder
// handler.
type BuildRequest struct {
	AuthorizationToken string
	Endpoint           string
	Namespace          string
	Image              string
	ImageBuildArgs     string
}

// NewHandler is a constructor for the Handler interface
func NewBuildHandler() (BuildHandler, error) {
	decoder, err := appregistry.NewManifestDecoder()
	if err != nil {
		return nil, err
	}
	return &buildhandler{
		downloader:      downloader.NewDownloader(),
		imageBuilder:    builder.NewImageBuilder(),
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
type BuildHandler interface {
	Handle(request *BuildRequest) error
}

type buildhandler struct {
	downloader      downloader.Downloader
	imageBuilder    builder.ImageBuilder
	manifestDecoder appregistry.ManifestDecoder
}

func (h *buildhandler) Handle(request *BuildRequest) error {
	// Create temporary working directory for manifests
	workingDirectory, err := ioutil.TempDir(".", "manifests-")
	if err != nil {
		return err
	}
	defer os.RemoveAll(workingDirectory)

	// Download files from App Registry
	manifests, err := h.downloader.GetManifests(request.AuthorizationToken, request.Endpoint, request.Namespace)
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
