
```
$ echo "GET http://localhost:8080/customers" | vegeta attack -rate=1500 -duration=10s | tee result.bin | vegeta report
Requests      [total, rate, throughput]         15000, 1500.11, 35.83
Duration      [total, attack, wait]             10.998s, 9.999s, 998.481ms
Latencies     [min, mean, 50, 90, 95, 99, max]  102.179µs, 7.306ms, 209.715µs, 353.06µs, 1.69ms, 134.746ms, 1.032s
Bytes In      [total, mean]                     255312, 17.02
Bytes Out     [total, mean]                     0, 0.00
Success       [ratio]                           2.63%
Status Codes  [code:count]                      0:14606  200:394  
Error Set:
```