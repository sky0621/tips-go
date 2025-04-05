.PHONY: gen
gen:
	go tool sqlc generate

.PHONY: migrate-up
migration-up:
	migrate -path migrations -database "mysql://root:yuckyjuice@tcp(127.0.0.1:18351)/sqlcmigratemysql01?parseTime=true&charset=utf8mb4" up

.PHONY: migrate-down
migration-down:
	migrate -path migrations -database "mysql://root:yuckyjuice@tcp(127.0.0.1:18351)/sqlcmigratemysql01?parseTime=true&charset=utf8mb4" down

.PHONY: build
build:
	go build -o ./bin/sqlc ./main.go

.PHONY: run
run: build
	./bin/sqlc
