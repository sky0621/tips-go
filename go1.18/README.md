# go1.18

## ジェネリクス
https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md

```
func F[T any](p T) { ... }

type M[T any] []T.

func F[T Constraint](p T) { ... }
```

`any` や `Constraint` は、インタフェース型。

```
T ... T 型そのものに制限。
~T ... 近似要素。基になる型が T である（継承したものを含む）すべての型に制限。
T1 | T2 | ... ... ユニオン要素。リストアップされた要素のどれかに制限。
```


