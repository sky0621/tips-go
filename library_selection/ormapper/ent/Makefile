.PHONY: build-migration
build-migration:
	go build -o ./bin/migration ./migration/cmd/main.go

.PHONY: run-migration
run-migration: build-migration
	./bin/migration -path=./migrations up

.PHONY: build
build:
	go build -o ./bin/sqlc ./main.go

.PHONY: run
run: build
	./bin/sqlc
