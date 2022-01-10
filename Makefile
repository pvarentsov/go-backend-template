# Constants

DB_URL = 'postgres://go-backend-template:go-backend-template@localhost:5454/go-backend-template?sslmode=disable'

# Help

.SILENT: help
help:
	@echo
	@echo "Usage: make [command]"
	@echo
	@echo "Commands:"
	@echo " build-http                    Build http server"
	@echo
	@echo " migration-create name={name}  Create migration"
	@echo " migration-up                  Up migrations"
	@echo " migration-down                Down last migration"
	@echo
	@echo " docker-up                     Up docker services"
	@echo " docker-down                   Down docker services"
	@echo
	@echo " fmt                           Format source code"
	@echo
	@echo "Requirements:"
	@echo " docker-compose                Docker Compose CLI: https://docs.docker.com/compose/reference"
	@echo " migrate                       Migration CLI tool: https://github.com/golang-migrate/migrate"
	@echo

# Build

.SILENT: build-http
build-http:
	@go build -o ./bin/http-server ./cmd/http/main.go
	@echo executable file \"http-server\" saved in ./bin/http-server

# Create migration

.SILENT: migration-create
migration-create:
	@migrate create -ext sql -dir ./migrations -seq $(name)

# Up migration

.SILENT: migration-up
migration-up:
	@migrate -database $(DB_URL) -path ./migrations up

# Down migration

.SILENT: "migration-down"
migration-down:
	@migrate -database $(DB_URL) -path ./migrations down 1

# Docker

.SILENT: docker-up
docker-up:
	@docker-compose up -d

.SILENT: docker-down
docker-down:
	@docker-compose down

# Format

.SILENT: fmt
fmt:
	@go fmt ./...

# Default

.DEFAULT_GOAL := help
