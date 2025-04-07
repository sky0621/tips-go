.PHONY: build-migration
build-migration:
	go build -o ./bin/migration ./migration/cmd/main.go

.PHONY: run-migration
run-migration: build-migration
	./bin/migration -path=./migrations up

.PHONY: build
build:
	go build -o ./bin/gorm ./main.go

.PHONY: exec-select
exec-select: build
	./bin/gorm

.PHONY: exec-insert
exec-insert: build
	./bin/gorm C

.PHONY: exec-update
exec-update: build
	./bin/gorm U

.PHONY: exec-delete
exec-delete: build
	./bin/gorm D
