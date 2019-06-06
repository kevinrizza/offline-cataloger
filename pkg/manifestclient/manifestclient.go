package manifestclient

import (
	"github.com/kevinrizza/offline-cataloger/pkg/apprclient"
)

func NewDownloader() Downloader {
	return &downloader{
		registryClientFactory: apprclient.NewClientFactory(),
	}
}

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
