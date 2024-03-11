package parse

import (
	"os"
	"testing"
)

func BenchmarkAll(b *testing.B) {
	type benchmark struct {
		name      string
		benchmark func(b *testing.B)
	}

	benchmarks := []benchmark{
		{name: "1k", benchmark: Benchmark_1k},
		{name: "10k", benchmark: Benchmark_10k},
		{name: "100k", benchmark: Benchmark_100k},
		{name: "1000k", benchmark: Benchmark_1000k},
		{name: "10000k", benchmark: Benchmark_10000k},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, benchmark.benchmark)
	}
}

func Benchmark_1k(b *testing.B) {
	equation := loadTestdata_b(b, "../testdata/src/halimath/1k")
	for i := 0; i < b.N; i++ {
		ToPostfix(equation)
	}
}

func Benchmark_10k(b *testing.B) {
	equation := loadTestdata_b(b, "../testdata/src/halimath/10k")
	for i := 0; i < b.N; i++ {
		ToPostfix(equation)
	}
}

func Benchmark_100k(b *testing.B) {
	equation := loadTestdata_b(b, "../testdata/src/halimath/100k")
	for i := 0; i < b.N; i++ {
		ToPostfix(equation)
	}
}

func Benchmark_1000k(b *testing.B) {
	equation := loadTestdata_b(b, "../testdata/src/halimath/1m")
	for i := 0; i < b.N; i++ {
		ToPostfix(equation)
	}
}

func Benchmark_10000k(b *testing.B) {
	equation := loadTestdata_b(b, "../testdata/src/halimath/10m")
	for i := 0; i < b.N; i++ {
		ToPostfix(equation)
	}
}

func loadTestdata_b(b *testing.B, file string) string {
	content, err := os.ReadFile(file)
	if err != nil {
		b.Fatalf("error loading file %s: %s", file, err)
	}
	b.ResetTimer()
	return string(content)
}
