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


.PHONY: run-local-mysql
run-local-mysql:
	docker compose up mysql -d --wait

.PHONY: stop-local-mysql
stop-local-mysql:
	docker compose down mysql

.PHONY: migrate-up
migrate-up:
	go run cmd/migration/main.go --up

.PHONY: migrate-down
migrate-down:
	go run cmd/migration/main.go --down

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
