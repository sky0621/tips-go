# architecture_model

## tool

```
go get -tool github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen
```

```
go get -tool github.com/golang-migrate/migrate/v4/cmd/migrate
```

```
go get -tool github.com/sqlc-dev/sqlc/cmd/sqlc
```

```
go get -u tool
```

## exec

### golang-migrate

```
go tool migrate create -ext sql -dir schema/db -seq create_contents_table
```

```
go tool migrate create -ext sql -dir schema/db -seq create_courses_table
```

```
go tool migrate create -ext sql -dir schema/db -seq create_programs_table
```
