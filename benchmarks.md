| goos | goarch | pkg                                             | cpu                                       |
|------|--------|-------------------------------------------------|-------------------------------------------|
| linux| amd64  | github.com/Penomatikus/calculator/internal/rpn/parse | AMD Ryzen 7 5700U with Radeon Graphics |

| Benchmark              | Iterations | Duration          |
|------------------------|------------|-------------------|
| All/1k-16              | 43,706     | 27,203 ns/op      |
| All/10k-16             | 3,782      | 304,756 ns/op     |
| All/100k-16            | 345        | 3,504,545 ns/op   |
| All/1000k-16           | 28         | 39,396,625 ns/op  |
| All/10000k-16          | 3          | 384,424,661 ns/op |


| Benchmark              | Iterations | Duration          |
|------------------------|------------|-------------------|
| All/1k-16              | 28,231     | 41,980 ns/op      |
| All/10k-16             | 2,473      | 448,812 ns/op     |
| All/100k-16            | 258        | 4,699,024 ns/op   |
| All/1000k-16           | 21         | 52,132,953 ns/op  |
| All/10000k-16          | 2          | 501,802,299 ns/op |