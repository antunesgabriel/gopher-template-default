.PHONY: new-migration

include .env

ENV_FILE := .env

# Função para ler as variáveis de ambiente do arquivo .env
define load-env
@echo "Loading .env file: $(ENV_FILE)"
env $$(cat .env | grep -v '^#' | xargs)
endef


build:
	docker compose up --build

up:
	docker compose up -d

dev:
	$(call load-env) go run ./cmd/server

# action maybe be create, up or down
migrate:
	go run ./cmd/migrate --action=$(action)

migrate-dev:
	$(call load-env) go run ./cmd/migrate --action=$(action)

new-migration:
	go run ./cmd/migrate --action=$(action) --name=$(name)

generate-wire:
	go generate ./cmd/server

run-test:
	go test ./... -cover

# $$(go list -f '{{if or .TestGoFiles .Xgit add .TestGoFiles}}{{.ImportPath}}{{end}}' ./...)