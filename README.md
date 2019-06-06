# Offline Cataloger
`offline-cataloger` is a command line tool that provides a toolkit for generating [`CatalogSources`](https://github.com/operator-framework/operator-lifecycle-manager#discovery-catalogs-and-automated-upgrades) by building [`operator-registry`](https://github.com/operator-framework/operator-registry) container images and then making them available to a local container registry. At that point, a Kubernetes cluster that does not have access to the outside internet but does have access to that local registry is able to access these catalogs.

This project is currently in a pre-alpha state.

## Installation
Currently, you can install the offline-cataloger by building and installing from source:
`make install`

## Usage
To build an operator-registry image:
`offline-cataloger build-image quay.io/$NAMESPACE/example-registry-image:latest -n "$NAMESPACE"`

For help message:
`offline-cataloger -h`
