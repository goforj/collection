//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {
	// Attach wraps a slice without copying.

	// Example: sharing backing slice
	items := []int{1, 2, 3}
	c := collection.Attach(items)

	items[0] = 9
	collection.Dump(c.Items())
	// #[]int [
	//   0 => 9 #int
	//   1 => 2 #int
	//   2 => 3 #int
	// ]
}
