# ddd_aggregate

Sample for persisting a DDD aggregate (Order) in an RDB and managing it in a transaction.

## Setup

```bash
cd experiment/ddd_aggregate

# Start MySQL
docker compose up -d

# Apply schema (sqldef)
go run ./cmd/migrate

# Generate sqlc code
sqlc generate
# or: go tool sqlc generate

# Run
go run .
```

## Environment variables

- Docker
  - `MYSQL_PORT` (default: 34507)
  - `MYSQL_DATABASE` (default: ddd_aggregate)
  - `MYSQL_USER` (default: app)
  - `MYSQL_PASSWORD` (default: app)
- Application
  - `DDD_AGG_DB_HOST` (default: 127.0.0.1)
  - `DDD_AGG_DB_PORT` (default: 34507)
  - `DDD_AGG_DB_USER` (default: app)
  - `DDD_AGG_DB_PASSWORD` (default: app)
  - `DDD_AGG_DB_NAME` (default: ddd_aggregate)
- Migration (sqldef)
  - `SQLDEF_MYSQL_HOST` (default: 127.0.0.1)
  - `SQLDEF_MYSQL_PORT` (default: 34507)
  - `SQLDEF_MYSQL_USER` (default: app)
  - `SQLDEF_MYSQL_PASSWORD` (default: app)
  - `SQLDEF_MYSQL_DATABASE` (default: ddd_aggregate)
