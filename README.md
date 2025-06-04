<<<<<<< HEAD
# ⚠️ Disclaimer

**This repository is Work in Progress (WIP)**. We are actively developing and improving this SDK. While the current functionality is stable and continuously tested in the CI, the API may change as we continue to enhance and refine the codebase until we reach a `v1.0.0` release. We recommend:

- Using specific version tags in your dependencies.
- Reviewing the changelog before upgrading.
- Testing thoroughly in your environment.
- Reporting any issues you encounter.

# Docker SDK for Go

A lightweight, modular SDK for interacting with Docker configuration and context data in Go.

This project is designed to be:
- **Extensible**: Built with composability in mind to support additional Docker-related features.
- **Lightweight**: No unnecessary dependencies; only what's needed to manage Docker configurations.
- **Go-native**: Idiomatic Go modules for clean integration in CLI tools and backend services.

## Features

- Parse and load Docker CLI config (`~/.docker/config.json`)
- Handle credential helpers
- Read and manage Docker contexts

## Installation

```bash
go get github.com/docker/go-sdk
```

## Usage

### dockerconfig

```go
cfg, err := dockerconfig.Load()
if err != nil {
    log.Fatalf("failed to load config: %v", err)
}

auth, ok := cfg.AuthConfigs["https://index.docker.io/v1/"]
if ok {
    fmt.Println("Username:", auth.Username)
}
```

### dockercontext

```go
dockerHost, err := dockercontext.CurrentDockerHost()
if err != nil {
    log.Fatalf("failed to get current docker host: %v", err)
}
```

More usage examples are coming soon!

## Contributing

We welcome contributions! Please read the [CONTRIBUTING](./CONTRIBUTING.md) file and open issues or submit pull requests once you're ready. Make sure your changes are well-tested and documented.

## Licensing

This project is licensed under the [Apache License 2.0](./LICENSE).

It includes portions of code derived from the other open source projects which are licensed under the MIT License. Their original licenses are preserved [here](./third_party), and attribution is provided in the [NOTICE](./NOTICE) file.

Modifications have been made to this code as part of its integration into this project.
=======
# Testcontainers

[![Main pipeline](https://github.com/testcontainers/testcontainers-go/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/testcontainers/testcontainers-go/actions/workflows/ci.yml)
[![GoDoc Reference](https://pkg.go.dev/badge/github.com/testcontainers/testcontainers-go.svg)](https://pkg.go.dev/github.com/testcontainers/testcontainers-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/testcontainers/testcontainers-go)](https://goreportcard.com/report/github.com/testcontainers/testcontainers-go)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=testcontainers_testcontainers-go&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=testcontainers_testcontainers-go)
[![License](https://img.shields.io/badge/license-MIT-blue)](https://github.com/testcontainers/testcontainers-go/blob/main/LICENSE)

[![Open in GitHub Codespaces](https://github.com/codespaces/badge.svg)](https://github.com/codespaces/new?hide_repo_select=true&ref=main&repo=141451032&machine=standardLinux32gb&devcontainer_path=.devcontainer%2Fdevcontainer.json&location=EastUs)

[![Join our Slack](https://img.shields.io/badge/Slack-4A154B?logo=slack)](https://testcontainers.slack.com/)

_Testcontainers for Go_ is a Go package that makes it simple to create and clean up container-based dependencies for
automated integration/smoke tests. The clean, easy-to-use API enables developers to programmatically define containers
that should be run as part of a test and clean up those resources when the test is done.

You can find more information about _Testcontainers for Go_ at [golang.testcontainers.org](https://golang.testcontainers.org), which is rendered from the [./docs](./docs) directory.

## Using _Testcontainers for Go_

Please visit [the quickstart guide](https://golang.testcontainers.org/quickstart) to understand how to add the dependency to your Go project.
>>>>>>> tcgo/main
