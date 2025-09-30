# sqldef

`docker compose up -d` で MySQL 8.0 を立ち上げ、`sample` ディレクトリから `mysqldef` を呼び出してスキーマ同期を試す最小構成です。

## セットアップ

```bash
cd library_selection/db_migration/sqldef

# MySQL を起動
docker compose up -d

# サンプルを実行
cd sample
go run .
```

必要に応じて `SQLDEF_MYSQL_*` 系の環境変数を上書きすれば、別ホストやユーザにも接続できます（既定値は docker-compose に合わせて `host=127.0.0.1`, `port=3306`, `user=app`, `password=app`, `database=sample`）。
