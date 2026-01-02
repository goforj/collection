//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {
	// TakeLast returns a new collection containing the last n items.
	// If n is less than or equal to zero, TakeLast returns an empty collection.
	// If n is greater than or equal to the collection length, TakeLast returns
	// the full collection.
	// @chainable true
	// @terminal false
	// 
	// This operation performs no element allocations; it re-slices the
	// underlying slice.
	// 
	// NOTE: returns a view (shares backing array). Use Clone() to detach.

	// Example: integers
	c := collection.New([]int{1, 2, 3, 4, 5})
	out := c.TakeLast(2)
	collection.Dump(out.Items())
	// #[]int [
	//   0 => 4 #int
	//   1 => 5 #int
	// ]

	// Example: take none
	out2 := c.TakeLast(0)
	collection.Dump(out2.Items())
	// #[]int [
	// ]

	// Example: take all
	out3 := c.TakeLast(10)
	collection.Dump(out3.Items())
	// #[]int [
	//   0 => 1 #int
	//   1 => 2 #int
	//   2 => 3 #int
	//   3 => 4 #int
	//   4 => 5 #int
	// ]

	// Example: structs
	type User struct {
		ID int
	}

	users := collection.New([]User{
		{ID: 1},
		{ID: 2},
		{ID: 3},
	})

	out4 := users.TakeLast(1)
	collection.Dump(out4.Items())
	// #[]main.User [
	//  0 => #main.User {
	//    +ID => 3 #int
	//  }
	// ]
}
