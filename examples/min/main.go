//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {
	// integers
	c := collection.NewNumeric([]int{3, 1, 2})
	min, ok := c.Min()
	collection.Dump(min, ok)
	// 1 #int
	// true #bool

	// floats
	c2 := collection.NewNumeric([]float64{2.5, 9.1, 1.2})
	min2, ok2 := c2.Min()
	collection.Dump(min2, ok2)
	// 1.200000 #float64
	// true #bool

	// empty collection
	empty := collection.NewNumeric([]int{})
	min3, ok3 := empty.Min()
	collection.Dump(min3, ok3)
	// 0 #int
	// false #bool
}
