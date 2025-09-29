# mysqldef sample

`sqldef` を Go から呼び出してスキーマを同期するシンプルなサンプルです。`schema/schema.sql` に記述した DDL を `mysqldef` コマンドへ標準入力で渡し、MySQL のスキーマを最新化します。

## 前提

- `mysqldef` がインストールされていること（例: `go install github.com/k0kubun/sqldef/...@latest`）。
- 参照する MySQL/互換エンジンが起動済みで、接続情報が環境変数で指定できること。

## 環境変数

| 変数名 | 必須 | 役割 |
| --- | --- | --- |
| `SQLDEF_MYSQL_HOST` | 任意 | 接続先ホスト。未指定時は `127.0.0.1` |
| `SQLDEF_MYSQL_PORT` | 任意 | 接続先ポート。未指定時は `3306` |
| `SQLDEF_MYSQL_USER` | 必須 | 接続ユーザ |
| `SQLDEF_MYSQL_PASSWORD` | 任意 | 接続パスワード |
| `SQLDEF_MYSQL_DATABASE` | 必須 | 対象データベース名 |
| `SQLDEF_DRY_RUN` | 任意 | `true`/`1` などを指定すると `mysqldef --dry-run` で差分のみ表示 |

## 実行方法

```bash
export SQLDEF_MYSQL_HOST=127.0.0.1
export SQLDEF_MYSQL_PORT=3306
export SQLDEF_MYSQL_USER=root
export SQLDEF_MYSQL_PASSWORD=secret
export SQLDEF_MYSQL_DATABASE=app_db

# 差分を確認したい場合
export SQLDEF_DRY_RUN=true

# スキーマ適用
cd library_selection/db_migration/sqldef/sample
go run .
```

実行すると `schema/schema.sql` の内容が `mysqldef` に送られ、差分が存在すれば適用（または `SQLDEF_DRY_RUN=true` の場合は SQL のみを出力）されます。

## スキーマの更新

`schema/schema.sql` を編集した後に再度 `go run .` を実行すると、`mysqldef` が既存のデータベーススキーマとの差分を計算し、必要な DDL を実行します。テーブルを削除・変更する操作も `sqldef` が適切に差分を算出してくれるため、スキーマのドリフト防止に活用できます。
