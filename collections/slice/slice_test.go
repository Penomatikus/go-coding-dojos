package slice

import (
	"fmt"
	"math/rand/v2"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Slice_Filter_SHORT(t *testing.T) {
	data := []int{2, 4, 13, 99, -1, 21, 100}
	mod := func(i int) bool { return i%2 == 0 }
	area := func(i int) bool { return i > 5 && i < 101 }

	slice := New(data)
	got := slice.Filter(mod).Filter(area).Collect()
	want := []int{100}
	assert.Equal(t, want, got)
}

func Test_Slice_Take(t *testing.T) {
	data := []int{78, 4, 13, 99, -1, 21, 100, 2, 48}

	slice := New(data)
	got := slice.Take(3).Collect()
	want := []int{78, 4, 13}
	assert.Equal(t, want, got)
}

func Test_Slice_Take_And_Filter(t *testing.T) {
	data := []int{9, 4, 99, 13, 99, -1, 21, 100, 2, 48}

	slice := New(data)
	s1 := slice.Filter(func(i int) bool { return i > 10 })
	s2 := s1.Take(3)
	s3 := s2.Filter(func(i int) bool { return i == 13 })
	got := s3.Collect()

	want := []int{13}
	assert.Equal(t, want, got)
}

func Test_Slice_Take_More_Than_Possible(t *testing.T) {
	data := []int{78, 4, 13, 99, -1, 21, 100, 2, 48}

	slice := New(data)
	got := slice.Take(10).Collect()
	assert.Equal(t, data, got)
}

func Test_Slice_Skip(t *testing.T) {
	data := []int{78, 4, 13, 99, -1, 21, 100, 2, 48}

	slice := New(data)
	got := slice.Skip(3).Collect()
	want := []int{99, -1, 21, 100, 2, 48}
	assert.Equal(t, want, got)
}

func Test_Slice_Skip_More_Than_Possible(t *testing.T) {
	data := []int{78, 4, 13, 99, -1, 21, 100, 2, 48}

	slice := New(data)
	got := slice.Skip(10).Collect()
	assert.Equal(t, []int{}, got)
}

func Test_Slice_Reduce(t *testing.T) {
	data := []int{78, 4, 13, 99, -1, 21, 100, 2, 48}

	slice := New(data)
	got := slice.Reduce(func(i, j int) int { return i + j })
	want := 364
	assert.Equal(t, want, got)
}

func Test_Slice_Filter_BigData(t *testing.T) {
	data := intTestdataT(10, 10)
	result := New(data).Filter(func(i int) bool { return i%2 == 0 }).Filter(func(i int) bool { return i > 5 && i < 11 }).Collect()
	fmt.Print(result)
}

func intTestdataT(amount int, max int) []int {
	testdata := make([]int, 0, amount)
	for i := 0; i <= amount; i++ {
		testdata = append(testdata, rand.IntN(max))
	}

	return testdata
}
