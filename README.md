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

## ref

https://100go.co/
