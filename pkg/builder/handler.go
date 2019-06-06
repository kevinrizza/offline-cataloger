package builder

import (
	"io/ioutil"
	"os"

	"github.com/kevinrizza/offline-cataloger/pkg/appregistry"
	"github.com/kevinrizza/offline-cataloger/pkg/manifestclient"
)

func NewHandler() (Handler, error) {
	decoder, err := appregistry.NewManifestDecoder()
	if err != nil {
		return nil, err
	}
	return &handler{
		downloader:      manifestclient.NewDownloader(),
		imageBuilder:    NewImageBuilder(),
		manifestDecoder: decoder,
	}, nil
}

type Handler interface {
	Handle(request *BuildRequest) error
}

type handler struct {
	downloader      manifestclient.Downloader
	imageBuilder    ImageBuilder
	manifestDecoder appregistry.ManifestDecoder
}

func (h *handler) Handle(request *BuildRequest) error {
	// create temp directory for manifests
	workingDirectory, err := ioutil.TempDir(".", "manifests-")
	if err != nil {
		return err
	}
	defer os.RemoveAll(workingDirectory)

	// download files
	manifests, err := h.downloader.GetManifests(request.Endpoint, request.Namespace)
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
