//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {
	// Avg returns the average of the collection values as a float64.
	// If the collection is empty, Avg returns 0.

	// Example: integers
	c := collection.NewNumeric([]int{2, 4, 6})
	collection.Dump(c.Avg())
	// 4.000000 #float64

	// Example: float
	c2 := collection.NewNumeric([]float64{1.5, 2.5, 3.0})
	collection.Dump(c2.Avg())
	// 2.333333 #float64
}
