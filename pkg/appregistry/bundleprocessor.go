package appregistry

import (
	"archive/tar"
	"errors"
	"io"
	"os"
	"path/filepath"
)

const (
	directoryPerm = 0755
	fileFlag      = os.O_CREATE | os.O_RDWR
)

// NewBundleProcessor is a bundleProcessor constructor
func NewBundleProcessor(manifestsDirectory string) (*bundleProcessor, error) {
	if manifestsDirectory == "" {
		return nil, errors.New("folder to store downloaded operator bundle has not been specified")
	}

	return &bundleProcessor{
		manifestsDirectory:   manifestsDirectory,
	}, nil
}

type bundleProcessor struct {
	manifestsDirectory   string
}

// Process takes an item of the tar ball and writes it to the underlying file
// system.
func (w *bundleProcessor) Process(header *tar.Header, manifestName string, reader io.Reader) (done bool, err error) {

	namedManifestDirectory := filepath.Join(w.manifestsDirectory, manifestName)
	target := filepath.Join(namedManifestDirectory, header.Name)

	if header.Typeflag == tar.TypeDir {
		if _, err = os.Stat(target); err == nil {
			return
		}

		err = os.MkdirAll(target, directoryPerm)
		return
	}

	if header.Typeflag != tar.TypeReg {
		return
	}

	// It's a file.
	f, err := os.OpenFile(target, fileFlag, os.FileMode(header.Mode))
	if err != nil {
		return
	}

	defer f.Close()

	_, err = io.Copy(f, reader)
	return
}
