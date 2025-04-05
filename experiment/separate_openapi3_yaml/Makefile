.PHONY: redoc-lint
redoc-lint:
	npx @redocly/cli lint ./schema/root.yml

.PHONY: redoc
redoc: redoc-lint
	npx @redocly/cli bundle ./schema/root.yml -o ./schema/openapi.yaml

.PHONY: gen
gen: redoc
	go tool oapi-codegen --config=config.yml ./schema/openapi.yaml

.PHONY: build
build:
	go build -o ./bin/openapi3 ./cmd/main.go

.PHONY: run
run: build
	./bin/openapi3
