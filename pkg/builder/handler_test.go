package builder

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kevinrizza/offline-cataloger/pkg/apis"
	"github.com/kevinrizza/offline-cataloger/pkg/apprclient"
	"github.com/kevinrizza/offline-cataloger/pkg/appregistry"
	mocks "github.com/kevinrizza/offline-cataloger/pkg/mocks/builder_mocks"
	"github.com/stretchr/testify/assert"
)

func TestHandleNormalCase(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockDownloader := mocks.NewDownloader(controller)
	mockImageBuilder := mocks.NewImageBuilder(controller)
	mockManifestDecoder := mocks.NewManifestDecoder(controller)

	handler := &handler{
		downloader:      mockDownloader,
		imageBuilder:    mockImageBuilder,
		manifestDecoder: mockManifestDecoder,
	}

	request := &apis.BuildRequest{
		Endpoint:  "fake.io/testendpoint",
		Namespace: "fakenamespace",
	}

	returnedManifests := make([]*apprclient.OperatorMetadata, 0)

	mockDownloader.EXPECT().GetManifests(request).Return(returnedManifests, nil)

	result := &appregistry.Result{}

	mockManifestDecoder.EXPECT().Decode(returnedManifests, gomock.Any()).Return(*result, nil)
	mockImageBuilder.EXPECT().Build(request.Image, gomock.Any()).Return(nil)

	err := handler.Handle(request)
	assert.Nil(t, err)
}
