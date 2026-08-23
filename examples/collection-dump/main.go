//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {
	// Dump prints items with godump and returns the same collection.
	// This is a no-op on the collection itself and never panics.

	// Example: integers
	c := collection.New([]int{1, 2, 3})
	c.Dump()
	// #[]int [
	//   0 => 1 #int
	//   1 => 2 #int
	//   2 => 3 #int
	// ]

	// Example: integers - chaining
	collection.New([]int{1, 2, 3}).
		Filter(func(v int) bool { return v > 1 }).
		Dump()
	// #[]int [
	//   0 => 2 #int
	//   1 => 3 #int
	// ]
}
