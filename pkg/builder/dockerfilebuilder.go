package builder

import (
	"bytes"
	"text/template"
)

// NewDockerfileBuilder is a constructor for the DockerfileBuilder interface
func NewDockerfileBuilder() DockerfileBuilder {
	return &dockerfilebuilder{}
}

// DockerfileTemplate is a struct that describes the fields needed to render
// the operator-registry dockerfile.
type DockerfileTemplate struct {
	WorkingDirectory string
}

// DockerfileBuilder is an interface that is implemented by structs
// that implement the Render method. The Render method takes a DockerfileTemplate
// and creates a custom Dockerfile that creates an operator-registry image
// based on the template's input.
type DockerfileBuilder interface {
	Render(dockerfileTemplate DockerfileTemplate) (string, error)
}

type dockerfilebuilder struct {
}

func (d *dockerfilebuilder) Render(dockerfileTemplate DockerfileTemplate) (string, error) {
	templ, err := template.New("dockerfile").Parse(registryDockerfile)
	if err != nil {
		return "", err
	}

	templateBuffer := &bytes.Buffer{}
	err = templ.Execute(templateBuffer, dockerfileTemplate)
	if err != nil {
		return "", err
	}

	return templateBuffer.String(), nil
}

const registryDockerfile = `
FROM python:3 as manifests

RUN pip3 install operator-courier==2.1.0
WORKDIR /usr/src
COPY {{.WorkingDirectory}} /usr/src/upstream-community-operators
RUN for file in /usr/src/upstream-community-operators/*; do operator-courier nest $file /manifests/$(basename $file); done

FROM quay.io/operator-framework/upstream-registry-builder:v1.1.0 as builder
COPY --from=manifests /manifests manifests
RUN ./bin/initializer -o ./bundles.db

FROM scratch
COPY --from=builder /build/bundles.db /bundles.db
COPY --from=builder /build/bin/registry-server /registry-server
COPY --from=builder /bin/grpc_health_probe /bin/grpc_health_probe
EXPOSE 50051
ENTRYPOINT ["/registry-server"]
CMD ["--database", "bundles.db"]
`
