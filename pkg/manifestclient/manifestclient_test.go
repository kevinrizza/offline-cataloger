package manifestclient

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kevinrizza/offline-cataloger/pkg/apprclient"
	mocks "github.com/kevinrizza/offline-cataloger/pkg/mocks/manifestclient_mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetManifestsNormalCase(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	endpoint := "fake.io/endpoint"
	namespace := "fakenamespace"

	factory := mocks.NewAppRegistryClientFactory(controller)

	registryClient := mocks.NewAppRegistryClient(controller)

	optionsWant := apprclient.Options{Source: endpoint}
	factory.EXPECT().New(optionsWant).Return(registryClient, nil).Times(1)

	expectedMetaResult := make([]*apprclient.OperatorMetadata, 0)
	metaItem := apprclient.OperatorMetadata{
		RegistryMetadata: apprclient.RegistryMetadata{
			Name:      "test",
			Namespace: "test",
		},
	}
	expectedMetaResult = append(expectedMetaResult, &metaItem)

	registryClient.EXPECT().RetrieveAll(namespace).Return(expectedMetaResult, nil)

	downloader := &downloader{
		registryClientFactory: factory,
	}

	actualMetaResult, err := downloader.GetManifests(endpoint, namespace)

	assert.Nil(t, err)
	assert.Equal(t, expectedMetaResult, actualMetaResult)
}
