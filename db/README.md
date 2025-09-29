# MySQL with Docker Compose

This setup spins up a MySQL container for local development using the sample schemas under `dbmysql01`–`dbmysql03`.

## Getting started

1. Move to this directory.
2. Adjust `.env` if you want different credentials, port, or sample data set.
3. Start the container with:

   ```bash
   docker compose up -d
   ```

Once the container transitions to `healthy`, MySQL is reachable on `localhost` (default port `3306`) with the credentials from `.env`.

## Switching the bootstrap data set

The `MYSQL_DATASET` entry in `.env` controls which folder is bind-mounted to `/docker-entrypoint-initdb.d` during the first start. The default is `dbmysql01`. Available choices today:

- `dbmysql01`
- `dbmysql02`
- `dbmysql03`

If you change the data set, remove the named volume so that the initialization scripts run again:

```bash
docker compose down -v
docker compose up -d
```

## Extras

- Custom MySQL settings live in `mysql/conf.d/charset.cnf` and enforce `utf8mb4`. Add more `.cnf` files in the same folder if you need further tweaks.
- The container defines a health check using `mysqladmin ping`, which helps orchestrators wait until MySQL is accepting connections.
