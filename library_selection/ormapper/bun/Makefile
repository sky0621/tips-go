.PHONY: build-migration
build-migration:
	go build -o ./bin/migration ./migration/cmd/main.go

.PHONY: run-migration
run-migration: build-migration
	./bin/migration -path=./migrations up

.PHONY: build
build:
	go build -o ./bin/bun ./main.go

.PHONY: exec-select
exec-select: build
	./bin/bun

.PHONY: exec-insert
exec-insert: build
	./bin/bun C

.PHONY: exec-update
exec-update: build
	./bin/bun U

.PHONY: exec-delete
exec-delete: build
	./bin/bun D

.PHONY: exec-various
exec-various: build
	./bin/bun V
