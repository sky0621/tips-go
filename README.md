# tips-go

```
go mod init github.com/sky0621/tips-go/xxxx
```

## 3rd party library の実ソースを保持しておきたい場合

```
go mod vendor
```

## 100 Go Mistakes and How to Avoid Them

### interface

https://100go.co/5-interface-pollution/

When to use interfaces
- Common behavior
- Decoupling
- Restricting behavior

```
Rob Pike
The bigger the interface, the weaker the abstraction.
```

```
Rob Pike
Don’t design with interfaces, discover them.
```

### generics

https://100go.co/9-generics/

### 誤ったプロジェクト構成 (プロジェクト構造とパッケージ構成)

- プロジェクトが過度に複雑になる可能性があるため、時期尚早なパッケージ化は避けるべき
- 1 つまたは 2 つのファイルだけを含む数十のナノパッケージを作成することは避けるべき
- パッケージ名は短く、簡潔で、表現力豊かで、慣例により単一の小文字にする必要があります
- エクスポートする必要があるものをできる限り最小限に抑える必要があります

https://go.dev/doc/modules/layout

### linter

- [Go言語の標準コードアナライザー](https://golang.org/cmd/vet)
- [エラーチェッカー](https://github.com/kisielk/errcheck)
- [循環的複雑度アナライザー](https://github.com/fzipp/gocyclo)
- [複数回使用文字列アナライザー](https://github.com/jgautheron/goconst)
- [golangci-lint](https://github.com/golangci/golangci-lint)

### formatter

- [Go言語の標準コードフォーマッター](https://golang.org/cmd/gofmt)
- [Go言語の標準インポートフォーマッター](https://godoc.org/golang.org/x/tools/cmd/goimports)

### 整数オーバーフローのチェック

https://github.com/teivah/100-go-mistakes/blob/master/src/03-data-types/18-integer-overflows/main.go

### Slice

https://100go.co/20-slice/

### Receiver

https://100go.co/#not-knowing-which-type-of-receiver-to-use-42

```
値レシーバーとポインタレシーバーのどちらを使用するかは、
どの型なのか、変化させる必要があるかどうか、コピーできないフィールドが含まれているかどうか、
オブジェクトはどれくらい大きいのか、などの要素に基づいて決定する必要があります。

通常は、そうしない正当な理由がない限り、値レシーバーを使用して間違いありません。
分からない場合は、ポインタレシーバを使用してください。
```

```
ポインタレシーバーでなければならない とき

・メソッドがレシーバーを変化させる必要がある場合。
　このルールは、受信側がスライスであり、メソッドが要素を追加する必要がある場合にも有効です。
・メソッドレシーバーにコピーできないフィールドが含まれている場合。
```

```
ポインタレシーバーであるべき とき

・レシーバーが大きなオブジェクトの場合。
　ポインタを使用すると、大規模なコピーの作成が防止されるため、呼び出しがより効率的になります。
```

```
値レシーバーでなければならない とき

・レシーバーの不変性を強制する必要がある場合。
・レシーバーがマップ、関数、チャネルの場合。
　それ以外の場合はコンパイルエラーが発生します。
```

```
値レシーバーであるべき とき

・レシーバーが変化させる必要のないスライスの場合。
・レシーバーが、time.Time などの小さな配列または構造体で、可変フィールドを持たない値型である場合。
・レシーバーが int、float64、または string などの基本型の場合。
```

### Concurrency

https://100go.co/56-concurrency-faster/

## ref

https://100go.co/
