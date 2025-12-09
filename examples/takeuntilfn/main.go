//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {
	// stop when value >= 3
	c1 := collection.New([]int{1, 2, 3, 4})
	out1 := c1.TakeUntilFn(func(v int) bool { return v >= 3 })
	collection.Dump(out1.Items())
	// #[]int [
	//	0 => 1 #int
	//	1 => 2 #int
	// ]

	// predicate immediately true → empty result
	c2 := collection.New([]int{10, 20, 30})
	out2 := c2.TakeUntilFn(func(v int) bool { return v < 50 })
	collection.Dump(out2.Items())
	// #[]int [
	// ]

	// no match → full list returned
	c3 := collection.New([]int{1, 2, 3})
	out3 := c3.TakeUntilFn(func(v int) bool { return v == 99 })
	collection.Dump(out3.Items())
	// #[]int [
	//	0 => 1 #int
	//	1 => 2 #int
	//	2 => 3 #int
	// ]
}
