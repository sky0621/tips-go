.PHONY: gen
gen:
	go tool sqlc generate

.PHONY: build-migration
build-migration:
	go build -o ./bin/migration ./migration/cmd/main.go

.PHONY: run-migration
run-migration: build-migration
	./bin/migration -path=./migrations up

.PHONY: build
build:
	go build -o ./bin/sqlc ./main.go

.PHONY: run-select
run-select: build
	./bin/sqlc

.PHONY: run-complex
run-complex: build
	./bin/sqlc CM

.PHONY: run-other-setup
run-other-setup: build
	./bin/sqlc OS

.PHONY: run-other
run-other: build
	./bin/sqlc O

.PHONY: run-search
run-search: build
	./bin/sqlc S

.PHONY: run-school-setup
run-school-setup: build
	./bin/sqlc SCS

.PHONY: run-school
run-school: build
	./bin/sqlc SC

.PHONY: run-school2
run-school2: build
	./bin/sqlc SC2
