# pman

> A lightweight, declarative container environment manager built on the Podman REST API.

`pman` lets you describe an application environment in a `pman.yaml` file, then build and run it through Podman with a simple, consistent CLI.

> ⚠️ **Work in progress:** pman is currently in early development. The v0.1 milestone focuses on proving the core build and run workflow for a single service.

## Table of Contents

- [What is pman?](#what-is-pman)
- [Project Status](#project-status)
  - [Current focus: v0.1](#current-focus-v01)

- [End Goal](#end-goal)
- [Requirements](#requirements)
- [Configuration](#configuration)
  - [Minimal example](#minimal-example)
  - [Configuration fields](#configuration-fields)

- [Commands](#commands)
  - [Build](#build)
  - [Up](#up)

- [Roadmap](#roadmap)
  - [Phase 1: Configuration](#phase-1-configuration)
  - [Phase 2: Podman API Client](#phase-2-podman-api-client)
  - [Phase 3: Build Workflow](#phase-3-build-workflow)
  - [Phase 4: Up Workflow](#phase-4-up-workflow)
  - [Phase 5: Environment Management](#phase-5-environment-management)
  - [Phase 6: Multi-Service Support](#phase-6-multi-service-support)
  - [Future Ideas](#future-ideas)

- [Non-Goals for v0.1](#non-goals-for-v01)
- [Contributing](#contributing)
- [License](#license)

---

## What is pman?

pman aims to make local container environments easier to define and manage without requiring a large orchestration layer.

Instead of manually running Podman commands, you define your application once:

```yaml
version: "1"

name: go-app

services:
  app:
    image: golang:1.24
    build:
      context: .
```

Then use pman to build and start it:

```bash
pman build -f pman.yaml
pman up -f pman.yaml
```

The intended workflow is:

```text
pman.yaml
    ↓
pman build
    ↓
Podman builds an image
    ↓
pman up
    ↓
Podman creates and starts a container
```

At its core, pman translates declarative YAML configuration into Podman API calls:

```text
YAML → pman → Podman REST API → Image → Container
```

---

## Project Status

pman is an experimental project and is not ready for production workloads.

### Current focus: v0.1

The first release is intentionally small. Its purpose is to validate the basic developer experience:

- Define one service in `pman.yaml`
- Specify a base image and build context
- Build that service through Podman
- Create a deterministic image name
- Start a container from that image

The initial commands are:

```bash
pman build -f pman.yaml
pman up -f pman.yaml
```

---

## End Goal

The long-term goal is for pman to provide a simple, Podman-native way to manage local multi-container application environments.

A completed pman workflow should make it easy to:

- Define services, builds, ports, volumes, networks, and environment variables in one file
- Build and run an entire local application stack with predictable names
- Start, stop, inspect, and troubleshoot environments through a small CLI
- Use Podman directly rather than introducing another container runtime
- Keep configuration readable, portable, and easy to version with an application repository

Example of the eventual direction:

```yaml
version: "1"

name: example-app

services:
  api:
    image: golang:1.24
    build:
      context: ./api
    ports:
      - "8080:8080"
    environment:
      DATABASE_URL: postgres://db:5432/app
    depends_on:
      - db

  db:
    image: postgres:latest
    ports:
      - "5432:5432"
    volumes:
      - postgres-data:/var/lib/postgresql/data
```

---

## Requirements

pman currently requires:

- [Go](https://go.dev/)
- [Podman](https://podman.io/)
- Podman's REST API exposed over TCP

> v0.1 communicates with Podman over TCP. Unix socket support is planned for a future release.

---

## Configuration

pman uses a `pman.yaml` file to define an application environment.

### Minimal example

```yaml
version: "1"

name: go-app

services:
  app:
    image: golang:1.24
    build:
      context: .
```

### Configuration fields

| Field                           | Description                                      |
| ------------------------------- | ------------------------------------------------ |
| `version`                       | Version of the pman configuration format         |
| `name`                          | Project name used when generating resource names |
| `services`                      | Application services managed by pman             |
| `services.<name>.image`         | Base image used by the service                   |
| `services.<name>.build.context` | Directory used as the container build context    |

The configuration file is the source of truth for the environment. pman translates it into Podman images and containers.

The `image` field describes the image used as the base for the build. pman generates the final image name itself.

For example:

```yaml
services:
  app:
    image: golang:1.24
    build:
      context: .
```

can produce:

```text
pman-go-app-app:latest
```

The v0.1 naming format is:

```text
pman-<project>-<service>:latest
```

---

## Commands

### Build

Build all services that define a `build` configuration:

```bash
pman build -f pman.yaml
```

For the example configuration above, pman generates a deterministic image name:

```text
pman-go-app-app:latest
```

The `image` field describes the base image:

```text
golang:1.24
```

The generated pman image is:

```text
pman-go-app-app:latest
```

The v0.1 naming format is:

```text
pman-<project>-<service>:latest
```

### Up

Create and start containers for the configured services:

```bash
pman up -f pman.yaml
```

The expected flow is:

```text
pman build
    ↓
pman-go-app-app:latest
    ↓
pman up
    ↓
pman-go-app-app container
```

---

## Roadmap

### Phase 1: Configuration

Build the `pman.yaml` configuration layer.

- [ ] Define the `pman.yaml` schema
- [ ] Parse YAML into Go structs
- [ ] Validate required configuration fields
- [ ] Parse service definitions
- [ ] Parse image configuration
- [ ] Parse build configuration
- [ ] Generate deterministic project and service names

### Phase 2: Podman API Client

Connect pman directly to Podman.

- [ ] Connect to Podman over TCP
- [ ] Create an HTTP API client
- [ ] Verify Podman API connectivity
- [ ] Implement image API requests
- [ ] Implement container API requests
- [ ] Improve API error messages

### Phase 3: Build Workflow

Implement:

```bash
pman build -f pman.yaml
```

- [ ] Read and validate configuration
- [ ] Find services with a `build` section
- [ ] Read the service base image
- [ ] Generate deterministic image names
- [ ] Send build requests to Podman
- [ ] Tag generated images as `:latest`
- [ ] Verify that images were created successfully

### Phase 4: Up Workflow

Implement:

```bash
pman up -f pman.yaml
```

- [ ] Read and validate configuration
- [ ] Find the corresponding pman image
- [ ] Create containers from configured services
- [ ] Apply basic container configuration
- [ ] Start created containers
- [ ] Report container status and failures clearly

### Phase 5: Environment Management

Add basic commands for inspecting and removing managed resources.

- [ ] `pman ps`
- [ ] `pman images`
- [ ] `pman down`
- [ ] Show project-scoped containers
- [ ] Remove project containers safely

### Phase 6: Multi-Service Support

Support complete local development environments.

- [ ] Multiple services
- [ ] Port mappings
- [ ] Environment variables
- [ ] Volumes
- [ ] Networking
- [ ] Service dependencies
- [ ] Health checks
- [ ] Startup ordering

### Future Ideas

- [ ] `pman logs`
- [ ] `pman exec`
- [ ] `.env` file support
- [ ] Custom image tags
- [ ] `pman publish`
- [ ] Container registry support
- [ ] Unix socket support
- [ ] Kubernetes manifest export
- [ ] Compose file migration/import support
- [ ] Development mode with file watching

---

## Non-Goals for v0.1

To keep the first version focused, v0.1 does not aim to support:

- Full Docker Compose compatibility
- Multi-service networking
- Volume management
- Container orchestration for production
- Kubernetes deployment management
- Registry publishing
- Unix socket connections
- Custom image tags

The goal of v0.1 is:

```text
pman.yaml
    ↓
pman build
    ↓
Image
    ↓
pman up
    ↓
Container
```

---

## Contributing

pman is early in development, so contributions, ideas, and feedback are welcome.

Useful areas to contribute include:

- YAML schema design
- Podman REST API integration
- CLI ergonomics
- Error handling and diagnostics
- Documentation and examples
- Tests for configuration parsing and naming behavior

Before opening a larger change, consider opening an issue to discuss the design and scope first.

---

## License

License information has not been added yet.
