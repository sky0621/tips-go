
## 新旧ベンチマーク結果の比較用ツールをインストール

```
go install golang.org/x/perf/cmd/benchstat@latest
```

## 新旧ベンチマーク結果の比較

```
benchstat old.log new.log
```
