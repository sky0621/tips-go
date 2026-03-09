# sqlparser

`sqlc` のクエリファイルを対象に、`SELECT` / `UPDATE` / `DELETE` が `schools.id` 条件を持つかを検査するサンプルです。

## 使っているもの

- `sqlc` で SQL ファイルを管理
- `github.com/pingcap/parser` で SQL を AST 解析

`github.com/pingcap/parser` の GitHub リポジトリは `pingcap/tidb` 側へ移動済みですが、このサンプルではユーザー指定の import path を使っています。

## 構成

- `sqlc.yaml`: `sqlc` の設定
- `queries/*.sql`: `sqlc` 管理の SQL
- `validator/validator.go`: SQL 検査ロジック
- `main.go`: CLI エントリポイント

## 実行

```bash
go run .
```

失敗が 1 件でもあると終了コード 1 で終了します。

特定の SQL 名を検査対象外にする場合は、`sqlc` の `-- name:` 名を `--exclude` で指定します。

```bash
go run . --exclude=UnsafeListSchools
```

複数指定する場合はカンマ区切りです。

## 判定ルール

- 対象は `SELECT`, `UPDATE`, `DELETE`
- 対象 SQL に `schools` テーブルが含まれる場合、その SQL の `WHERE` 句に `schools.id` を絞る条件が必要
- `schools` の別名指定も許可
- `UPDATE schools ... WHERE id = ?` のような、対象テーブルが 1 つだけで列名が `id` の場合も許可
