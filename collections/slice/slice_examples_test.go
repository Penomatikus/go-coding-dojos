package slice

import (
	"fmt"
)

func ExampleSlice_Filter() {
	data := []int{2, 4, 13, 99, -1, 21, 100}
	mod := func(i int) bool { return i%2 == 0 }
	area := func(i int) bool { return i > 5 && i < 101 }

	slice := New(data)
	got := slice.Filter(mod).Filter(area).Collect()
	fmt.Print(got)

	// Output: [100]
}

func ExampleSlice_Take() {
	data := []int{78, 4, 13, 99, -1, 21, 100, 2, 48}
	slice := New(data)
	got := slice.Take(3).Collect()
	fmt.Print(got)

	// Output: [78 4 13]
}

func ExampleSlice_Skip() {
	data := []int{78, 4, 13, 99, -1, 21, 100, 2, 48}

	slice := New(data)
	got := slice.Skip(3).Collect()
	fmt.Print(got)

	// Output: [99 -1 21 100 2 48]
}

func ExampleSlice_Reduce() {
	data := []int{78, 4, 13, 99, -1, 21, 100, 2, 48, 1}

	slice := New(data)
	got := slice.Reduce(func(i, j int) int { return i + j })
	fmt.Print(got)

	// Output: 365
}

func ExampleSlice_First() {
	data := []int{78, 4, 4, 99, -1, 21, 100, 2, 48}

	slice := New(data)
	got := slice.First(func(i int) bool { return i == 4 }).Collect()
	fmt.Println(got)
	got = slice.CollectIndices()
	fmt.Print(got)

	// Output: [4]
	// [1]
}

func ExampleSlice_Last() {
	data := []int{78, 4, 4, 99, -1, 21, 100, 4, 4}

	slice := New(data)
	got := slice.Last(func(i int) bool { return i == 4 }).Collect()
	fmt.Println(got)
	got = slice.CollectIndices()
	fmt.Print(got)

	// Output: [4]
	// [8]
}

func ExampleSlice_Collect() {
	data := []int{78, 4, 13, 99, -1, 21, 100, 2, 48, 1}

	slice := New(data)
	got := slice.Collect()
	fmt.Print(got)

	// Output: [78 4 13 99 -1 21 100 2 48 1]
}

func ExampleSlice_CollectIndices() {
	data := []int{78, 4, 13, 99, -1, 21, 100, 2, 48, 1}

	slice := New(data)
	got := slice.Take(5).Skip(3).CollectIndices()
	fmt.Print(got)

	// Output: [3 4]
}
