.PHONY: build

REPO = github.com/kevinrizza/offline-cataloger
BUILD_PATH = $(REPO)/cmd/offline-cataloger
PKG_PATH = $(REPO)/pkg

all: build

build:
	# build binary
	./build/build.sh

install:
	go install $(BUILD_PATH)

unit:
	go test $(PKG_PATH)/...
