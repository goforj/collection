//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {
	// ItemsCopy returns a copy of the collection's items.

	// Example: integers
	c := collection.New([]int{1, 2, 3})
	items := c.ItemsCopy()
	collection.Dump(items)
	// #[]int [
	//   0 => 1 #int
	//   1 => 2 #int
	//   2 => 3 #int
	// ]
}
