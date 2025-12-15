//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {
	// UniqueComparable returns a new collection with duplicate comparable items removed.
	// The first occurrence of each value is kept, and order is preserved.
	// This is a faster, allocation-friendly path for comparable types.

	// Example: integers
	c := collection.New([]int{1, 2, 2, 3, 4, 4, 5})
	out := collection.UniqueComparable(c)
	collection.Dump(out.Items())
	// #[]int [
	//   0 => 1 #int
	//   1 => 2 #int
	//   2 => 3 #int
	//   3 => 4 #int
	//   4 => 5 #int
	// ]

	// Example: strings
	c2 := collection.New([]string{"A", "a", "B", "B"})
	out2 := collection.UniqueComparable(c2)
	collection.Dump(out2.Items())
	// #[]string [
	//   0 => "A" #string
	//   1 => "a" #string
	//   2 => "B" #string
	// ]
}
