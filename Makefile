.PHONY: build
build: ## Build a version
	make swag
	go build -v ./cmd/app

.PHONY: clean
clean: ## Remove temporary files
	go clean

.PHONY: dev
dev: ## Go Run
	env $(shell cat .env) go run cmd/app/main.go

.PHONY: test
test: ## Go Tests
	go test ./... -v

.PHONY: lint
lint: ## Go Lint
	golangci-lint run ./...



.PHONY: swag
swag: ## Update swagger.json
	swag init -g ./cmd/app/main.go

.PHONY: swag-fmt
swag-fmt: ## Formatter for GoDoc (Swagger)
	swag fmt -g ./cmd/app/main.go

.PHONY: docker-up
docker-up: ## Start Docker-Compose Container with app & database
	docker-compose -f build/docker-compose.yml up -d --build

.PHONY: docker-down
docker-down: ## Down Docker-Compose Containers
	docker-compose -f build/docker-compose.yml down

.PHONY: docker-database-up
docker-database-up: ## Start Docker-compose Container with only database service
	docker-compose -f build/docker-compose.yml up database -d --build

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := build