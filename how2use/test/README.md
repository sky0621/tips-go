# test

## simple

```
go test
```

### 詳細ログ

```
go test -v
```

## skip

```
go test -v -short
```

## coverage

```
go test -v -cover
```

### output

```
go test -v -cover -coverprofile=cover.out
```

#### to HTML

```
go tool cover -html=cover.out -o cover.html
```
## Fuzzing

### バグを見つけてクラッシュするまでエンドレスで実行し続ける

```
go test -v -fuzz FuzzDoSomething
```

### 実行時間を指定

```
go test -v -fuzz FuzzDoSomething -fuzztime 10s
```
