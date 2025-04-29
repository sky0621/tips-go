
.PHONY: redoc-lint
redoc-lint:
	npx @redocly/cli lint ./schema/openapi/parts/root.yaml

.PHONY: redoc
redoc: redoc-lint
	npx @redocly/cli bundle ./schema/openapi/parts/root.yaml -o ./schema/openapi/api.yaml

.PHONY: generate-api
generate-api: redoc
	go tool oapi-codegen --config=oapi-codegen-config.yaml -o internal/handler/interfaces/generated_handler.go schema/openapi/api.yaml

.PHONY: migrate-up
migration-up:
	migrate -path schema/db -database "mysql://root:yuckyjuice@tcp(127.0.0.1:22221)/architecturemodelmysql01?parseTime=true&charset=utf8mb4" up

.PHONY: migrate-down
migration-down:
	migrate -path schema/db -database "mysql://root:yuckyjuice@tcp(127.0.0.1:22221)/architecturemodelmysql01?parseTime=true&charset=utf8mb4" down

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
