package dockercontainer

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/docker/docker/api/types/container"
	cexec "github.com/docker/go-sdk/dockercontainer/exec"
	"github.com/docker/go-sdk/dockercontainer/wait"
)

// Container represents a container
type Container struct {
	// ID the Container ID
	ID string

	// WaitingFor the waiting strategy to use for the container.
	WaitingFor wait.Strategy

	// Image the image to use for the container.
	Image string

	// exposedPorts a reference to the container's requested exposed ports.
	// It allows checking they are ready before any wait strategy
	exposedPorts []string

	// isRunning whether the container is running
	isRunning bool

	// imageWasBuilt whether the image was built
	imageWasBuilt bool

	// keepBuiltImage makes Terminate not remove the image if imageWasBuilt.
	keepBuiltImage bool

	// sessionID the session ID for the container
	sessionID string

	// terminationSignal the termination signal for the container
	terminationSignal chan bool

	// consumers the log consumers for the container
	consumers []LogConsumer

	// TODO: Remove locking and wait group once the deprecated StartLogProducer and
	// StopLogProducer have been removed and hence logging can only be started and
	// stopped once.

	// logProductionCancel is used to signal the log production to stop.
	logProductionCancel context.CancelCauseFunc
	logProductionCtx    context.Context

	logProductionTimeout *time.Duration
	logger               log.Logger
	lifecycleHooks       []LifecycleHooks
}

// Exec executes a command in the current container.
// It returns the exit status of the executed command, an [io.Reader] containing the combined
// stdout and stderr, and any encountered error. Note that reading directly from the [io.Reader]
// may result in unexpected bytes due to custom stream multiplexing headers.
// Use [cexec.Multiplexed] option to read the combined output without the multiplexing headers.
// Alternatively, to separate the stdout and stderr from [io.Reader] and interpret these headers properly,
// [github.com/docker/docker/pkg/stdcopy.StdCopy] from the Docker API should be used.
func (c *Container) Exec(ctx context.Context, cmd []string, options ...cexec.ProcessOption) (int, io.Reader, error) {
	return 0, nil, fmt.Errorf("not implemented")
}
