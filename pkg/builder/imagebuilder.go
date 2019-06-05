package builder

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func NewImageBuilder() ImageBuilder {
	return &imageBuilder{
		dockerfilebuilder: NewDockerfileBuilder(),
	}
}

type ImageBuilder interface {
	Build(image, workingDirectory string) error
}

type imageBuilder struct {
	dockerfilebuilder DockerfileBuilder
}

func (i *imageBuilder) Build(image, workingDirectory string) error {
	fmt.Println(fmt.Sprintf("Building the image %s", image))

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
