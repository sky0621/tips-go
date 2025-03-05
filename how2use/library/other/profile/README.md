# profile

## exec

```
go build ~~~

2025/03/05 22:29:51 profile: memory profiling enabled (rate 4096), mem.pprof
1.041434秒
2025/03/05 22:29:52 profile: memory profiling disabled, mem.pprof
```

## profiling

```
go tool pprof memory mem.pprof
```

top

list main

