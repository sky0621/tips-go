# sqlc_with_golang_migrate

## get

### sqlc
```
go get -tool github.com/sqlc-dev/sqlc/cmd/sqlc
```

### golang-migrate
```
go get -tool github.com/golang-migrate/migrate/v4/cmd/migrate
```

## update

```
go get -u tool
```

## exec

### golang-migrate
```
go tool migrate create -ext sql -dir migrations -seq create_users_table
```
```
go tool migrate create -ext sql -dir migrations -seq add_age_to_users
```

## check

```
go tool
```
