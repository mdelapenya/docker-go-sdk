package dockercontainer

import "context"

type LifecycleHooks struct {
	PreBuilds      []DefinitionHook
	PostBuilds     []DefinitionHook
	PreCreates     []DefinitionHook
	PostCreates    []ContainerHook
	PreStarts      []ContainerHook
	PostStarts     []ContainerHook
	PostReadies    []ContainerHook
	PreStops       []ContainerHook
	PostStops      []ContainerHook
	PreTerminates  []ContainerHook
	PostTerminates []ContainerHook
}

// DefinitionHook is a hook that will be called before a container is started.
// It can be used to modify the container definition on container creation,
// using the different lifecycle hooks that are available:
// - Building
// - Creating
// For that, it will receive a Definition, modify it and return an error if needed.
type DefinitionHook func(ctx context.Context, req Definition) error

// ContainerHook is a hook that is called after a container is created
// It can be used to modify the state of the container after it is created,
// using the different lifecycle hooks that are available:
// - Created
// - Starting
// - Started
// - Readied
// - Stopping
// - Stopped
// - Terminating
// - Terminated
// It receives a [Container], modify it and return an error if needed.
type ContainerHook func(ctx context.Context, ctr Container) error
