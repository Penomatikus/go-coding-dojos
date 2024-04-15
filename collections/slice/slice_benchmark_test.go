package slice

import (
	"fmt"
	"math/rand/v2"
	"testing"
)

func randomIntTestdata(b *testing.B, n int) []int {
	testdata := make([]int, 0, n)
	for i := 0; i <= n; i++ {
		testdata = append(testdata, rand.IntN(n))
	}

	b.ResetTimer()
	return testdata
}

func randomStringTestdata(b *testing.B, n int) []string {
	testdata := make([]string, 0, n)
	for i := 0; i <= n; i++ {
		testdata = append(testdata, fmt.Sprintf("%d", rand.IntN(n)))
	}

	b.ResetTimer()
	return testdata
}

func rangeIntTestdata(b *testing.B, n int64) []int64 {
	testdata := make([]int64, 0, n)
	for i := int64(0); i <= n; i++ {
		testdata = append(testdata, i)
	}

	b.ResetTimer()
	return testdata
}

func Benchmark_Filter_Range_Even_Numbers(b *testing.B) {
	testdata := rangeIntTestdata(b, 10)
	for i := 0; i < b.N; i++ {
		slice := New(testdata)
		slice.Filter(func(j int64) bool { return j%2 == 0 }).Filter(func(j int64) bool { return j > 500 && j < 1000 })
	}
}

func Benchmark_Filter_Random_Even_Numbers(b *testing.B) {
	testdata := randomIntTestdata(b, 1000)
	slice := New(testdata)
	for i := 0; i < b.N; i++ {
		slice.Filter(func(j int) bool { return j%2 == 0 }).Filter(func(j int) bool { return j > 500 && j < 1000 })
	}
}

func Benchmark_Filter_Random_String_Len(b *testing.B) {
	testdata := randomStringTestdata(b, 1000)
	slice := New(testdata)
	for i := 0; i < b.N; i++ {
		slice.Filter(func(s string) bool { return len(s) == 3 })
	}
}
