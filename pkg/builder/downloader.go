package builder

import (
	"github.com/kevinrizza/offline-cataloger/pkg/apprclient"
)

func NewDownloader() Downloader {
	return &downloader{
		registryClientFactory: apprclient.NewClientFactory(),
	}
}

type Downloader interface {
	GetManifests(request *BuildRequest) ([]*apprclient.OperatorMetadata, error)
}

type downloader struct {
	registryClientFactory apprclient.ClientFactory
}

func (d *downloader) GetManifests(request *BuildRequest) ([]*apprclient.OperatorMetadata, error) {
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
