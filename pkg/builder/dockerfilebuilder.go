package builder

func NewDockerfileBuilder(workingDirectory string) DockerfileBuilder {
	return &dockerfilebuilder{
		workingDirectory: workingDirectory,
	}
}

type DockerfileBuilder interface {
	BuildDockerfile() string
}

type dockerfilebuilder struct {
	workingDirectory string
}

func (d *dockerfilebuilder) BuildDockerfile() string {
	return registryDockerfileTemplate
}

const registryDockerfileTemplate = `
FROM python:3 as manifests

RUN pip3 install operator-courier==2.1.0
WORKDIR /usr/src
COPY manifests /usr/src/upstream-community-operators
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
