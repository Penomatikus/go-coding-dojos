package structtype

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Slice_Filter_SHORT(t *testing.T) {
	data := []int{2, 4, 13, 99, -1, 21, 100}
	mod := func(i int) bool { return i%2 == 0 }
	area := func(i int) bool { return i > 5 && i < 101 }

	slice := NewSlice(data)
	got := slice.Filter(mod).Filter(area).Collect()
	want := []int{100}
	assert.Equal(t, want, got)
}
