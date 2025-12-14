//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {
	// After returns all items after the first element for which pred returns true.
	// If no element matches, an empty collection is returned.

	// Example: integers
	c := collection.New([]int{1, 2, 3, 4, 5})
	c.After(func(v int) bool { return v == 3 }).Dump()
	// #[]int [
	//  0 => 4 #int
	//  1 => 5 #int
	// ]
}
