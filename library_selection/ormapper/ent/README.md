# ent

## Modelのガワ生成

```
go run -mod=mod entgo.io/ent/cmd/ent new User
go run -mod=mod entgo.io/ent/cmd/ent new Post
go run -mod=mod entgo.io/ent/cmd/ent new Comment
```

## ガワの編集後

```
go generate ./ent
```

