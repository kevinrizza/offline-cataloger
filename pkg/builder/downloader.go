package builder

import (
	"fmt"

	"github.com/kevinrizza/offline-cataloger/pkg/apprclient"
)

func NewDownloader() Downloader {
	return &downloader{
		registryClientFactory: apprclient.NewClientFactory(),
	}
}

type Downloader interface {
	GetManifests(request *BuildRequest) error
}

type downloader struct {
	registryClientFactory apprclient.ClientFactory
}

func (d *downloader) GetManifests(request *BuildRequest) error {
	options := apprclient.Options{
		Source: request.Endpoint,
	}

	client, err := d.registryClientFactory.New(options)
	if err != nil {
		return err
	}

	manifests, err := client.RetrieveAll(request.Namespace)
	if err != nil {
		return err
	}

	for _, manifest := range manifests {
		fmt.Println(manifest.RegistryMetadata.Name)
		//converted := string(manifest.RawYAML)
		//fmt.Println(converted)
	}

	return nil
}
