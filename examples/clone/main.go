//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {
	// Clone returns a copy of the collection.
	// @chainable true
	// @terminal false
	// 
	// The returned collection has its own backing slice, so subsequent mutations
	// do not affect the original collection.
	// 
	// Clone is intended to be used when branching a pipeline while preserving
	// the original collection.

	// Example: basic cloning
	c := collection.New([]int{1, 2, 3})
	clone := c.Clone()

	clone.Append(4)

	collection.Dump(c.Items())
	// #[]int [
	//   0 => 1 #int
	//   1 => 2 #int
	//   2 => 3 #int
	// ]

	collection.Dump(clone.Items())
	// #[]int [
	//   0 => 1 #int
	//   1 => 2 #int
	//   2 => 3 #int
	//   3 => 4 #int
	// ]

	// Example: branching pipelines
	base := collection.New([]int{1, 2, 3, 4, 5})

	evens := base.Clone().Filter(func(v int) bool {
		return v%2 == 0
	})

	odds := base.Clone().Filter(func(v int) bool {
		return v%2 != 0
	})

	collection.Dump(base.Items())
	// #[]int [
	//   0 => 1 #int
	//   1 => 2 #int
	//   2 => 3 #int
	//   3 => 4 #int
	//   4 => 5 #int
	// ]

	collection.Dump(evens.Items())
	// #[]int [
	//   0 => 2 #int
	//   1 => 4 #int
	// ]

	collection.Dump(odds.Items())
	// #[]int [
	//   0 => 1 #int
	//   1 => 3 #int
	//   2 => 5 #int
	// ]
}
