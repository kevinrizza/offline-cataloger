.PHONY: build

REPO = github.com/kevinrizza/offline-cataloger
BUILD_PATH = $(REPO)/cmd/offline-cataloger
PKG_PATH = $(REPO)/pkg

# Mocks
MOCKS_PATH = ./pkg/mocks
BUILDER_MOCKS_PKG = builder_mocks
MANIFESTCLIENT_MOCKS_PKG = manifestclient_mocks

ifeq ($(GOBIN),)
mockgen = $(GOPATH)/bin/mockgen
else
mockgen = $(GOBIN)/mockgen
endif

all: build

build:
	# build binary
	./build/build.sh

install:
	go install $(BUILD_PATH)

unit: generate-mocks unit-test

unit-test:
	go test -v $(PKG_PATH)/...

generate-mocks:
	# Build mockgen from the same version used by gomock. This ensures that 
	# gomock and mockgen are never out of sync.
	go install ./vendor/github.com/golang/mock/mockgen
	
	@echo making sure directory for mocks exists
	mkdir -p $(MOCKS_PATH)

	# $(mockgen) -destination=<Path/file where the mock is generated> -package=<The package that the generated mock files will belong to> -mock_names=<Original Interface name>=<Name of Generated mocked Interface> <Go package path of the original interface> <comma seperated list of the interface you want to mock>

	# builder package
	$(mockgen) -destination=$(MOCKS_PATH)/$(BUILDER_MOCKS_PKG)/mock_manifestclient.go -package=$(BUILDER_MOCKS_PKG) -mock_names=Downloader=Downloader $(PKG_PATH)/manifestclient Downloader
	$(mockgen) -destination=$(MOCKS_PATH)/$(BUILDER_MOCKS_PKG)/mock_imagebuilder.go -package=$(BUILDER_MOCKS_PKG) -mock_names=ImageBuilder=ImageBuilder $(PKG_PATH)/builder ImageBuilder
	$(mockgen) -destination=$(MOCKS_PATH)/$(BUILDER_MOCKS_PKG)/mock_manifestdecoder.go -package=$(BUILDER_MOCKS_PKG) -mock_names=ManifestDecoder=ManifestDecoder $(PKG_PATH)/appregistry ManifestDecoder

	# manifestclient package
	$(mockgen) -destination=$(MOCKS_PATH)/$(MANIFESTCLIENT_MOCKS_PKG)/mock_appregistry.go -package=$(MANIFESTCLIENT_MOCKS_PKG) -mock_names=ClientFactory=AppRegistryClientFactory,Client=AppRegistryClient $(PKG_PATH)/apprclient ClientFactory,Client

clean-mocks:
	@echo cleaning mock folder
	rm -rf $(MOCKS_PATH)

clean: clean-mocks
