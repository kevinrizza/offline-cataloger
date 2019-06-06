package manifestclient

import (
	"github.com/kevinrizza/offline-cataloger/pkg/apprclient"
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
	GetManifests(endpoint, namespace string) ([]*apprclient.OperatorMetadata, error)
}

type downloader struct {
	registryClientFactory apprclient.ClientFactory
}

func (d *downloader) GetManifests(endpoint, namespace string) ([]*apprclient.OperatorMetadata, error) {
	options := apprclient.Options{
		Source: endpoint,
	}

	client, err := d.registryClientFactory.New(options)
	if err != nil {
		return nil, err
	}

	manifests, err := client.RetrieveAll(namespace)
	if err != nil {
		return nil, err
	}

	return manifests, nil
}
