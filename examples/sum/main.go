//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {
	// integers
	c := collection.NewNumeric([]int{1, 2, 3})
	total := c.Sum()
	collection.Dump(total)
	// 6 #int

	// floats
	c2 := collection.NewNumeric([]float64{1.5, 2.5})
	total2 := c2.Sum()
	collection.Dump(total2)
	// 4.000000 #float64

	// empty collection
	c3 := collection.NewNumeric([]int{})
	total3 := c3.Sum()
	collection.Dump(total3)
	// 0 #int
}
