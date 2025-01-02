# test

## シンプル

```
go test
```

### 詳細ログ

```
go test -v
```

## スキップコードを入れた分をスキップ

```
go test -v -short
```

## カバレッジ

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
## 自動でランダム値をインプットしてひたすらテスト

### バグを見つけてクラッシュするまでエンドレスで実行し続ける

```
go test -v -fuzz .
```

### テストケース指定

```
go test -v -fuzz FuzzDoSomething
```

### 実行時間を指定

```
go test -v -fuzz FuzzDoSomething -fuzztime 10s
```

## ベンチマーク

```
go test -v -bench .
```

### テストケース指定

```
go test -v -bench BenchmarkDoSomething
```

### メモリアロケーション量

```
go test -v -benchmem -bench BenchmarkDoSomething
```

### 回数を指定

```
go test -v -count 5 -bench BenchmarkDoSomething
```
