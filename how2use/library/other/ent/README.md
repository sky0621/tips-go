# ent

## ドキュメント
https://entgo.io/ja/docs/getting-started/

## スキーマのガワを自動生成

```
go run -mod=mod entgo.io/ent/cmd/ent new User
```

## 実装したスキーマに該当するソース群を自動生成 

```
go generate ./ent
```

## DB起動

```
docker compose up
```

## DBマイグレーション

```
go run main.go
```

