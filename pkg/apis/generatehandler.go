package apis

import (
	"io/ioutil"

	"github.com/kevinrizza/offline-cataloger/pkg/appregistry"
	"github.com/kevinrizza/offline-cataloger/pkg/downloader"
)

type GenerateManifestsRequest struct {
	AuthorizationToken string
	Endpoint           string
	Namespace          string
}

func NewGenerateHandler() (GenerateHandler, error) {
	decoder, err := appregistry.NewManifestDecoder()
	if err != nil {
		return nil, err
	}
	return &generatehandler{
		downloader:      downloader.NewDownloader(),
		manifestDecoder: decoder,
	}, nil
}

// GenerateHandler is an interface that is implemented by structs
// that implement the Handle method. A GenerateHandler takes GenerateManifests requests
// as input and downloads operator manifests from an appregistry namespace.
type GenerateHandler interface {
	Handle(request *GenerateManifestsRequest) error
}

type generatehandler struct {
	downloader      downloader.Downloader
	manifestDecoder appregistry.ManifestDecoder
}

func (h *generatehandler) Handle(request *GenerateManifestsRequest) error {
	// Create manifest directory
	workingDirectory, err := ioutil.TempDir(".", "manifests-")
	if err != nil {
		return err
	}

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

	return nil
}
