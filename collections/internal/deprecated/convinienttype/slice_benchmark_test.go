package convinienttype

import (
	"fmt"
	"math/rand/v2"
	"testing"
)

func intTestdata(b *testing.B, amount int) []int {
	testdata := make([]int, 0, amount)
	for i := 0; i <= amount; i++ {
		testdata = append(testdata, rand.IntN(amount))
	}

	b.ResetTimer()
	return testdata
}

func stringTestdata(b *testing.B, amount int) []string {
	testdata := make([]string, 0, amount)
	for i := 0; i <= amount; i++ {
		testdata = append(testdata, fmt.Sprintf("%d", rand.IntN(amount)))
	}

	b.ResetTimer()
	return testdata
}

func Benchmark_Filter_Even(b *testing.B) {
	testdata := intTestdata(b, 1000)
	for i := 0; i < b.N; i++ {
		slice := New(testdata)
		slice.Filter(func(i int) bool { return i%2 == 0 }).Collect()
	}
}

func Benchmark_Filter_String(b *testing.B) {
	testdata := stringTestdata(b, 1000)
	for i := 0; i < b.N; i++ {
		slice := New(testdata)
		slice.Filter(func(s string) bool { return len(s) == 3 }).Collect()
	}
}
