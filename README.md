# Go Gin Template


## Installation

### Install dependencies

```bash
$ go mod download
```

### Startup database in Docker

```bash
$ make docker-database-up
```

### Run project

```bash
$ make dev
```

### Install golangci-lint for local use

```bash
$ brew install golangci-lint
```

After you can run:

```bash
$ golangci-lint run ./...
```

## Other commands

```bash
$ make help
build                          Build a version
clean                          Remove temporary files
dev                            Go Run
swag                           Update swagger.json
swag-fmt                       Formatter for GoDoc (Swagger)
docker-up                      Start Docker-Compose Container with app & database
docker-down                    Down Docker-Compose Containers
docker-database-up             Start Docker-compose Container with only database service
```

## Project Structure

```text
📂cmd/
├─ 📂app
│  ├─ 📄main.go     // Main package of the application, containing minimal logic, only responsible for launching the application
📂internal/
├─ 📂app/           // Core application package. Dependencies are initialized here, main goroutines are started, and the web server is launched
├─ 📂database/      // Database configuration and migrations
├─ 📂config/        // Application configuration
├─ 📂server/        // Server configuration (router), API entry points description
├─ 📂entity/        // Entities and repositories
├─ 📂service        // Business logic layer
│  ├─ 📂 user       // Users
├─ 📂utils          // Utility functions used across all layers of the application

```

## Addresses

1. [Localhost 8080](http://localhost:8080) (for dev)
2. [Swagger API Doc](http://localhost:8080/swagger/index.html)