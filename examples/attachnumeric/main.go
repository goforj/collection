//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {
	// AttachNumeric wraps a slice of numeric types without copying.

	// Example: sharing backing slice
	items := []int{1, 2, 3}
	c := collection.AttachNumeric(items)

	items[0] = 9
	collection.Dump(c.Items())
	// #[]int [
	//   0 => 9 #int
	//   1 => 2 #int
	//   2 => 3 #int
	// ]
}
