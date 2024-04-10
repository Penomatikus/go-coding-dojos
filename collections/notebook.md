# Notebook

Lessons learned:

- caching comes at a cost
- running a benchmark will automaticaly pre call \*\_test.go files
  - this increases the ns/op of the benchmark ~0.150 ns/op
- just because a benchmark tells one's implementation is a fast boy it doesn't mean the benchmark test is faster too

Open questions:

- the benchmark is behaving strange with spikes in performance and, why:
  - Benchmark_Filter_String-16 100000000 63.49 ns/op 0 B/op 0 allocs/op
  - Benchmark_Filter_String-16 676503130 1.754 ns/op 0 B/op 0 allocs/op
  - Benchmark_Filter_String-16 100000000 13.99 ns/op 0 B/op 0 allocs/op
  - Benchmark_Filter_String-16 345857420 3.536 ns/op 0 B/op 0 allocs/op

Neat online findings:

- [go.dev - Slice Tricks](https://go.dev/wiki/SliceTricks)
- [go.dev - Intro Generics](https://go.dev/blog/intro-generics)
