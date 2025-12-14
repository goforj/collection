//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {
	// Dump is a convenience function that calls godump.Dump.

	// Example: integers
	c2 := collection.New([]int{1, 2, 3})
	collection.Dump(c2.Items())
	// #[]int [
	//   0 => 1 #int
	//   1 => 2 #int
	//   2 => 3 #int
	// ]
}
