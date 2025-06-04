package dockercontainer

import (
	"io"

	"github.com/docker/docker/api/types/build"
)

// FromDockerfile represents the parameters needed to build an image from a Dockerfile
// rather than using a pre-built one
type FromDockerfile struct {
	// BuildArgs enable user to pass build args to docker daemon
	BuildArgs map[string]*string

	// BuildLogWriter for output of build log, defaults to io.Discard
	BuildLogWriter io.Writer

	// BuildOptionsModifier Modifier for the build options before image build. Use it for
	// advanced configurations while building the image. Please consider that the modifier
	// is called after the default build options are set.
	BuildOptionsModifier func(*build.ImageBuildOptions)

	// ContextArchive the tar archive file to send to docker that contains the build context
	ContextArchive io.ReadSeeker

	// 16-byte aligned fields (strings)
	// Context the path to the context of the docker build
	Context string

	// Dockerfile the path from the context to the Dockerfile for the image, defaults to "Dockerfile"
	Dockerfile string

	// Repo the repo label for image, defaults to UUID
	Repo string

	// Tag the tag label for image, defaults to UUID
	Tag string

	// 1-byte aligned fields
	// KeepImage describes whether DockerContainer.Terminate should not delete the
	// container image. Useful for images that are built from a Dockerfile and take a
	// long time to build. Keeping the image also Docker to reuse it.
	KeepImage bool
}
