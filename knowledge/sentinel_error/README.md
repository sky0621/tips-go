# sentinel_error

## ref

https://100go.co/#comparing-an-error-value-inaccurately-51

- センチネルエラーはグローバル変数として定義されたエラーのこと
- 慣例として Err で始め、その後にエラー型を続けます

ガイドライン

- 予期されるエラーはエラー値（センチネルエラー）として設計する必要があります
```
var ErrFoo =errors.New("foo")
```
- 予期しないエラーはエラー型として設計する必要があります
```
BarError は error インタフェースを実装した上で
type BarError struct { ... }
```

## sample

https://cs.opensource.google/go/go/+/refs/tags/go1.22.0:src/archive/zip/reader.go;l=27
