| goos  | goarch | pkg                                                      | cpu                                    |
| ----- | ------ | -------------------------------------------------------- | -------------------------------------- |
| linux | amd64  | github.com/Penomatikus/calculator/internal/rpn/**parse** | AMD Ryzen 7 5700U with Radeon Graphics |

| Benchmark     | Iterations | Duration          |
| ------------- | ---------- | ----------------- |
| All/1k-16     | 43,706     | 27,203 ns/op      |
| All/10k-16    | 3,782      | 304,756 ns/op     |
| All/100k-16   | 345        | 3,504,545 ns/op   |
| All/1000k-16  | 28         | 39,396,625 ns/op  |
| All/10000k-16 | 3          | 384,424,661 ns/op |

| goos  | goarch | pkg                                                           | cpu                                    |
| ----- | ------ | ------------------------------------------------------------- | -------------------------------------- |
| linux | amd64  | github.com/Penomatikus/calculator/internal/rpn/**calculator** | AMD Ryzen 7 5700U with Radeon Graphics |

| Benchmark     | Iterations | Duration          |
| ------------- | ---------- | ----------------- |
| All/1k-16     | 28,231     | 41,980 ns/op      |
| All/10k-16    | 2,473      | 448,812 ns/op     |
| All/100k-16   | 258        | 4,699,024 ns/op   |
| All/1000k-16  | 21         | 52,132,953 ns/op  |
| All/10000k-16 | 2          | 501,802,299 ns/op |

| goos    | goarch | pkg                                                      | cpu                                |
| ------- | ------ | -------------------------------------------------------- | ---------------------------------- |
| windows | amd64  | github.com/Penomatikus/calculator/internal/rpn/**parse** | AMD Ryzen 7 5800X 8-Core Processor |

goos: windows
goarch: amd64
pkg: github.com/Penomatikus/calculator/internal/rpn/parse
cpu: AMD Ryzen 7 5800X 8-Core Processor  
BenchmarkAll/1k-16 61802 24057 ns/op
BenchmarkAll/10k-16 4206 265088 ns/op
BenchmarkAll/100k-16 379 2673800 ns/op
BenchmarkAll/1000k-16 34 31409353 ns/op
BenchmarkAll/10000k-16 3 356474100 ns/op

BenchmarkAll/1k-16 42896 28652 ns/op
BenchmarkAll/10k-16 3805 304861 ns/op
BenchmarkAll/100k-16 369 3102938 ns/op
BenchmarkAll/1000k-16 28 36066704 ns/op
BenchmarkAll/10000k-16 3 374655467 ns/op

ok github.com/Penomatikus/calculator/internal/rpn/calculator 0.589s
