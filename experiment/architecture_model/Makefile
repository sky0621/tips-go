OPENAPI_SCHEMA_DIR := ./schema/openapi
OPENAPI_PARTS_DIR := $(OPENAPI_SCHEMA_DIR)/parts
OPENAPI_PARTS_ROOT_YAML := $(OPENAPI_PARTS_DIR)/root.yaml
OPENAPI_SCHEMA_YAML := $(OPENAPI_SCHEMA_DIR)/api.yaml

.PHONY: redoc-lint
redoc-lint:
	npx @redocly/cli lint $(OPENAPI_PARTS_ROOT_YAML)

.PHONY: redoc
redoc: redoc-lint
	npx @redocly/cli bundle $(OPENAPI_PARTS_ROOT_YAML) -o $(OPENAPI_SCHEMA_YAML)

.PHONY: generate-api
generate-api: redoc
	go tool oapi-codegen --config=oapi-codegen-config.yaml -o internal/api/generated.go $(OPENAPI_SCHEMA_YAML)

DB_USER := root
DB_PASSWORD := yuckyjuice
DB_HOST := 127.0.0.1
DB_PORT := 22221
DB_NAME := architecturemodelmysql01
DSN := "mysql://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)?parseTime=true&charset=utf8mb4"

.PHONY: migrate-up
migration-up:
	migrate -path schema/db -database $(DSN) up

.PHONY: migrate-down
migration-down:
	migrate -path schema/db -database $(DSN) down

.PHONY: generate-sqlc
generate-sqlc:
	go tool sqlc generate

.PHONY: test
test:
	go test ./internal/...

.PHONY: build
build:
	mkdir -p bin
	go build -o bin/server ./cmd

.PHONY: run
run: build
	./bin/server
