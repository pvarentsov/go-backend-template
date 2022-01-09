# Constants

DB_URL = 'postgres://go-backend-template:go-backend-template@localhost:5454/go-backend-template?sslmode=disable'

# Help

.SILENT: help
help:
	@echo "Usage: make [command]"
	@echo
	@echo "Commands:"
	@echo "   mcreate name={migration_name}    Create migration"
	@echo "   mup                              Up migrations"
	@echo "   mdown                            Down last migration"
	@echo ""
	@echo "   fmt                              Format source code"

# Create migration

.SILENT: "mcreate"
mcreate:
	@migrate create -ext sql -dir ./migrations -seq $(name)

# Up migration

.SILENT: "mup"
mup:
	@migrate -database $(DB_URL) -path ./migrations up

# Down migration

.SILENT: "mdown"
mdown:
	@migrate -database $(DB_URL) -path ./migrations down 1

# Format

.SILENT: fmt
fmt:
	@go fmt ./...

# Default

.DEFAULT_GOAL := help
