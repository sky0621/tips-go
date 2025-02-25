
## 新旧ベンチマーク結果の比較用ツールをインストール

```
go install golang.org/x/perf/cmd/benchstat@latest
```

## 新旧ベンチマーク結果の比較

```
benchstat old.log new.log
```

```
goos: darwin
goarch: arm64
pkg: github.com/sky0621/tips-go/how2use/test/compare_benchmark/sample
cpu: Apple M2
         │   old.log   │      new.log       │
         │   sec/op    │   sec/op     vs base   │
Before-8   380.7µ ± 2%
After-8                  357.0µ ± 1%
geomean    380.7µ        357.0µ       ? ¹ ²
¹ benchmark set differs from baseline; geomeans may not be comparable
² ratios must be >0 to compute geomean

         │   old.log    │       new.log       │
         │     B/op     │     B/op      vs base   │
Before-8   415.8Ki ± 0%
After-8                   196.8Ki ± 0%
geomean    415.8Ki        196.8Ki       ? ¹ ²
¹ benchmark set differs from baseline; geomeans may not be comparable
² ratios must be >0 to compute geomean

         │   old.log   │      new.log       │
         │  allocs/op  │  allocs/op   vs base   │
Before-8   9.761k ± 0%
After-8                  9.747k ± 0%
geomean    9.761k        9.747k       ? ¹ ²
¹ benchmark set differs from baseline; geomeans may not be comparable
² ratios must be >0 to compute geomean
```
