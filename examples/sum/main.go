//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {
	  c := collection.NewNumeric([]int{1, 2, 3})
	  total := c.Sum()
	Example (float):
	  c := collection.NewNumeric([]float64{1.5, 2.5})
	  total := c.Sum()
	  // total == 6

	  // total == 4.0
}
