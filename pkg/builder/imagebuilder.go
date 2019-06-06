package builder

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

// NewImageBuilder is a constructor for the ImageBuilder interface
func NewImageBuilder() ImageBuilder {
	return &imageBuilder{
		dockerfilebuilder: NewDockerfileBuilder(),
	}
}

// ImageBuilder is an interface that is implemented by structs that
// implement the Build method. Build takes an image name and a path
// which contains operator manifests and builds an operator-registry
// container image using docker build.
type ImageBuilder interface {
	Build(image, workingDirectory string) error
}

type imageBuilder struct {
	dockerfilebuilder DockerfileBuilder
}

func (i *imageBuilder) Build(image, workingDirectory string) error {
	// Generate the dockerfile
	template := &DockerfileTemplate{WorkingDirectory: workingDirectory}
	dockerfileText, err := i.dockerfilebuilder.Render(*template)
	if err != nil {
		return err
	}

	dockerfile, err := ioutil.TempFile(".", "Dockerfile-")
	if err != nil {
		return err
	}
	defer os.Remove(dockerfile.Name())

	_, err = dockerfile.WriteString(dockerfileText)
	if err != nil {
		return err
	}

	// Create the docker command
	var args []string
	args = append(args, "build", "-f", dockerfile.Name(), "-t", image, ".")
	cmd := exec.Command("docker", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Exec the build
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to exec %#v: %v", cmd.Args, err)
	}

	return nil
}
