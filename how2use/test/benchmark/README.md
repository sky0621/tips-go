# benchmark

```
go test -benchmem -bench .

go test -count 5 -benchmem -bench .
```

```
goos: darwin
goarch: arm64
pkg: github.com/sky0621/tips-go/how2use/test/benchmark
cpu: Apple M2
BenchmarkSum-8          419672557                2.591 ns/op           0 B/op          0 allocs/op
PASS
ok      github.com/sky0621/tips-go/how2use/test/benchmark       1.670s
```

ベンチマーク関数名         試行回数                  1回の実行にかかる時間    メモリアロケーションサイズ メモリアロケーション回数

```
goos: darwin
goarch: arm64
pkg: github.com/sky0621/tips-go/how2use/test/benchmark/sample
cpu: Apple M2
BenchmarkDoSomething-8                21          50734448 ns/op        88017305 B/op         38 allocs/op
PASS
ok      github.com/sky0621/tips-go/how2use/test/benchmark/sample        2.055s
```