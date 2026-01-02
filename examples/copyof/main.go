//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {
	// CopyOf creates a new Collection by copying the provided slice.

	// Example: copying input slice
	items := []int{1, 2, 3}
	c := collection.CopyOf(items)

	items[0] = 9
	collection.Dump(c.Items())
	// #[]int [
	//   0 => 1 #int
	//   1 => 2 #int
	//   2 => 3 #int
	// ]
}
