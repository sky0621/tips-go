# golang-migrate

## install

```
brew install golang-migrate
```

### version

```
migrate --version 
v4.18.2
```

## create

```
migrate create -ext sql -dir migrations -seq create_users_table
```

## migrate

### up

```
migrate -path migrations -database "mysql://root:yuckyjuice@tcp(127.0.0.1:18111)/mgmysql01?parseTime=true&charset=utf8mb4" up
```

### down

```
migrate -path migrations -database "mysql://root:yuckyjuice@tcp(127.0.0.1:18111)/mgmysql01?parseTime=true&charset=utf8mb4" down
```
