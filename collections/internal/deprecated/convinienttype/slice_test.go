package convinienttype

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Slice_Filter(t *testing.T) {
	data := []int{12, 4, 13, 99, -1, 21, 100}
	predicate := func(i int) bool { return i > 40 }

	slice := New(data)
	got := slice.Filter(predicate).Collect()
	want := []int{99, 100}

	assert.Equal(t, want, got)

}
