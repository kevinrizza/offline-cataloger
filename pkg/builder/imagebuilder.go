package builder

import (
	"fmt"

	"github.com/kevinrizza/offline-cataloger/pkg/apprclient"
)

func NewImageBuilder() ImageBuilder {
	return &imageBuilder{}
}

type ImageBuilder interface {
	Build(image string) error
}

type imageBuilder struct {
	registryClientFactory apprclient.ClientFactory
}

func (i *imageBuilder) Build(image string) error {
	fmt.Println(fmt.Sprintf("Building the image %s", image))

	return nil
}
