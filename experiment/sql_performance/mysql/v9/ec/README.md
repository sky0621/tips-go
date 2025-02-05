# experiment/sql_performance/mysql/v9/ec

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
migrate create -ext sql -dir migrations -seq create_all_tables
```
