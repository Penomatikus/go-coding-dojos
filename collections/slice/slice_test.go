package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Slice_Take_And_Filter(t *testing.T) {
	data := []int{9, 4, 99, 13, 99, -1, 21, 100, 2, 48}

	slice := New(data)
	got := slice.Filter(func(i int) bool { return i > 10 }).
		Take(3).
		Filter(func(i int) bool { return i == 13 }).
		Collect()
	want := []int{13}
	assert.Equal(t, want, got)
}

func Test_Slice_Take_More_Than_Possible(t *testing.T) {
	data := []int{78, 4, 13, 99, -1, 21, 100, 2, 48}

	slice := New(data)
	got := slice.Take(10).Collect()
	assert.Equal(t, data, got)
}

func Test_Slice_Skip_More_Than_Possible(t *testing.T) {
	data := []int{78, 4, 13, 99, -1, 21, 100, 2, 48}

	slice := New(data)
	got := slice.Skip(10).Collect()
	assert.Equal(t, []int{}, got)
}

func Test_Slice_First_Empty(t *testing.T) {
	data := []int{78, 4, 4, 99, -1, 21, 100, 2, 48}

	slice := New(data)
	got := slice.First(func(i int) bool { return i == 1337 }).Collect()
	want := []int{}
	assert.Equal(t, want, got)

	got = slice.CollectIndices()
	want = []int{}
	assert.Equal(t, want, got)
}

func Test_Slice_Last_Empty(t *testing.T) {
	data := []int{78, 4, 4, 99, -1, 21, 100, 2, 48}

	slice := New(data)
	got := slice.Last(func(i int) bool { return i == 1337 }).Collect()
	want := []int{}
	assert.Equal(t, want, got)

	got = slice.CollectIndices()
	want = []int{}
	assert.Equal(t, want, got)
}

func Test_Slice_Dry_Operations(t *testing.T) {
	data := []int{78, 4, 4, 99, -1, 21, 100, 2, 48}

	slice := New(data)
	got := slice.Take(5).
		Filter(func(i int) bool { return i == 4 }).
		Filter(func(i int) bool { return i == 5 }).
		First(func(i int) bool { return i == 6 }).
		Last(func(i int) bool { return i == 7 }).
		Collect()
	want := []int{}
	assert.Equal(t, want, got)

	got = slice.CollectIndices()
	want = []int{}
	assert.Equal(t, want, got)

	assert.True(t, slice.dry)
}
