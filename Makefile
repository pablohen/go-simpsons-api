# Binary output directory (gitignored via /bin/ in .gitignore)
BIN_DIR := bin
BINARY := $(BIN_DIR)/go-simpsons-api

SWAG := go run github.com/swaggo/swag/cmd/swag@latest
AIR := $(shell command -v air 2>/dev/null)
AIR_RUN := $(if $(AIR),air,go run github.com/air-verse/air@latest)

.DEFAULT_GOAL := help

.PHONY: help
help: ## Show this help
	@echo "Targets:"
	@grep -hE '^[a-zA-Z0-9_.-]+:.*?## ' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "  %-14s %s\n", $$1, $$2}'

.PHONY: run
run: ## Run the API (go run .)
	go run .

.PHONY: build
build: ## Build binary to bin/go-simpsons-api
	@mkdir -p $(BIN_DIR)
	go build -o $(BINARY) .

.PHONY: test
test: ## Run tests
	go test ./...

.PHONY: vet
vet: ## Run go vet
	go vet ./...

.PHONY: tidy
tidy: ## go mod tidy
	go mod tidy

.PHONY: swagger
swagger: ## Regenerate OpenAPI / docs package (swag)
	$(SWAG) init --parseDependency -g main.go -o docs

.PHONY: dev
dev: ## Run with live reload (Air: install with go install github.com/air-verse/air@latest, or uses go run)
	$(AIR_RUN)

.PHONY: clean
clean: ## Remove build artifacts (bin/, tmp/, air log)
	rm -rf $(BIN_DIR) tmp build-errors.log
