//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {
	// integers
	c := collection.New([]int{1, 2, 3})
	c.Dump()
	// #[]int [
	//   0 => 1 #int
	//   1 => 2 #int
	//   2 => 3 #int
	// ]

	// chaining
	collection.New([]int{1, 2, 3}).
		Filter(func(v int) bool { return v > 1 }).
		Dump()
	// #[]int [
	//   0 => 2 #int
	//   1 => 3 #int
	// ]
	// integers
	c2 := collection.New([]int{1, 2, 3})
	collection.Dump(c2.Items())
	// #[]int [
	//   0 => 1 #int
	//   1 => 2 #int
	//   2 => 3 #int
	// ]
}
