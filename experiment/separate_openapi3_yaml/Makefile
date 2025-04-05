.PHONY: gen
gen:
	go tool oapi-codegen --config=config.yml ./schema/openapi.yml

.PHONY: build
build:
	go build -o ./bin/openapi3 ./cmd/main.go

.PHONY: run
run: build
	./bin/openapi3
