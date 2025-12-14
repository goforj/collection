//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {
	// SkipLast returns a new collection with the last n items skipped.
	// If n is less than or equal to zero, SkipLast returns the full collection.
	// If n is greater than or equal to the collection length, SkipLast returns
	// an empty collection.
	// 
	// This operation performs no element allocations; it re-slices the
	// underlying slice.

	// Example: integers
	c := collection.New([]int{1, 2, 3, 4, 5})
	out := c.SkipLast(2)
	collection.Dump(out.Items())
	// #[]int [
	//   0 => 1 #int
	//   1 => 2 #int
	//   2 => 3 #int
	// ]

	// Example: skip none
	out2 := c.SkipLast(0)
	collection.Dump(out2.Items())
	// #[]int [
	//   0 => 1 #int
	//   1 => 2 #int
	//   2 => 3 #int
	//   3 => 4 #int
	//   4 => 5 #int
	// ]

	// Example: skip all
	out3 := c.SkipLast(10)
	collection.Dump(out3.Items())
	// #[]int []

	// Example: structs
	type User struct {
		ID int
	}

	users := collection.New([]User{
		{ID: 1},
		{ID: 2},
		{ID: 3},
	})

	out4 := users.SkipLast(1)
	collection.Dump(out4.Items())
	// #[]collection.User [
	//   0 => {ID:1} #collection.User
	//   1 => {ID:2} #collection.User
	// ]
}
