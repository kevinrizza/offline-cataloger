package builder

import (
	"github.com/kevinrizza/offline-cataloger/pkg/apis"
	"github.com/kevinrizza/offline-cataloger/pkg/apprclient"

	log "github.com/sirupsen/logrus"
)

// NewDownloader is a constructor for the Downloader interface
func NewDownloader() Downloader {
	return &downloader{
		registryClientFactory: apprclient.NewClientFactory(),
	}
}

// Downloader is an interface that is implemented by structs that
// implement the GetManifests method. GetManifests takes data about where
// an appregistry namespace is located, and downloads the manifests
// at that namespace.
type Downloader interface {
	GetManifests(request *apis.BuildRequest) ([]*apprclient.OperatorMetadata, error)
}

type downloader struct {
	registryClientFactory apprclient.ClientFactory
}

func (d *downloader) GetManifests(request *apis.BuildRequest) ([]*apprclient.OperatorMetadata, error) {
	log.Debugf("Downloading manifests from %s at namespace %s", request.Endpoint, request.Namespace)

	options := apprclient.Options{
		Source: request.Endpoint,
	}

	client, err := d.registryClientFactory.New(options)
	if err != nil {
		return nil, err
	}

	manifests, err := client.RetrieveAll(request.Namespace)
	if err != nil {
		return nil, err
	}

	return manifests, nil
}
