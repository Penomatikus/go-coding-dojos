# Notebook

Lessons learned:

- caching comes at a cost
- running a benchmark will automaticaly pre call \*\_test.go files
  - this increases the ns/op of the benchmark ~0.150 ns/op
- just because a benchmark tells one's implementation is a fast boy it doesn't mean the benchmark test is faster too


Neat online findings:

- [go.dev - Slice Tricks](https://go.dev/wiki/SliceTricks)
- [go.dev - Intro Generics](https://go.dev/blog/intro-generics)
