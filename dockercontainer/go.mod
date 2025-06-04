module github.com/docker/go-sdk/dockercontainer

go 1.23.6

replace github.com/docker/go-sdk/dockerconfig => ../dockerconfig

require (
	github.com/containerd/errdefs v1.0.0
	github.com/docker/docker v28.2.2+incompatible
	github.com/docker/go-connections v0.5.0
	github.com/docker/go-sdk/dockerconfig v0.1.0
	github.com/stretchr/testify v1.10.0
	golang.org/x/sys v0.32.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/docker/go-units v0.5.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/google/go-cmp v0.7.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/moby/docker-image-spec v1.3.1 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/objx v0.5.2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gotest.tools/v3 v3.5.2 // indirect
)
