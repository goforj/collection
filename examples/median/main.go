//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {
	// Median returns the statistical median of the numeric collection as float64.
	// Returns (0, false) if the collection is empty.
	// 
	// Odd count  → middle value
	// Even count → average of the two middle values

	// Example: integers - odd number of items
	c := collection.NewNumeric([]int{3, 1, 2})

	median1, ok1 := c.Median()
	collection.Dump(median1, ok1)
	// 2.000000 #float64
	// true     #bool

	// Example: integers - even number of items
	c2 := collection.NewNumeric([]int{10, 2, 4, 6})

	median2, ok2 := c2.Median()
	collection.Dump(median2, ok2)
	// 5.000000 #float64
	// true     #bool

	// Example: floats
	c3 := collection.NewNumeric([]float64{1.1, 9.9, 3.3})

	median3, ok3 := c3.Median()
	collection.Dump(median3, ok3)
	// 3.300000 #float64
	// true     #bool

	// Example: integers - empty numeric collection
	c4 := collection.NewNumeric([]int{})

	median4, ok4 := c4.Median()
	collection.Dump(median4, ok4)
	// 0.000000 #float64
	// false    #bool
}
