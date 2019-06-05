package builder

import (
	"github.com/kevinrizza/offline-cataloger/pkg/appregistry"
)

func NewHandler() (Handler, error) {
	workingDirectory := "./manifests/"
	decoder, err := appregistry.NewManifestDecoder(workingDirectory)
	if err != nil {
		return nil, err
	}
	return &handler{
		downloader:       NewDownloader(),
		imageBuilder:     NewImageBuilder(workingDirectory),
		manifestDecoder:  *decoder,
		workingDirectory: workingDirectory,
	}, nil
}

type Handler interface {
	Handle(request *BuildRequest) error
}

type handler struct {
	downloader       Downloader
	imageBuilder     ImageBuilder
	manifestDecoder  appregistry.ManifestDecoder
	workingDirectory string
}

func (h *handler) Handle(request *BuildRequest) error {
	// download files
	manifests, err := h.downloader.GetManifests(request)
	if err != nil {
		return err
	}

	// decode binary and parse yaml
	_, err = h.manifestDecoder.Decode(manifests)
	if err != nil {
		return err
	}

	// parse yaml for additional images

	// create dockerfile pointing to the yaml
	// docker build
	err = h.imageBuilder.Build(request.Image)
	if err != nil {
		return err
	}

	// pull additional images to local registry

	return nil
}
