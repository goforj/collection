//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {
	// ZipWith combines this collection with another collection using fn up to the shorter length.

	// Example: add corresponding integers
	left := collection.New([]int{1, 2, 3})
	right := collection.New([]int{10, 20})
	sums := left.ZipWith(right, func(a, b int) int {
		return a + b
	})
	collection.Dump(sums.Items())
	// #[]int [
	//   0 => 11 #int
	//   1 => 22 #int
	// ]
}
