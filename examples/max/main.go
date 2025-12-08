package main

import "github.com/goforj/collection"

func main() {

	c := collection.NewNumeric([]int{3, 1, 2})
	max, ok := c.Max()
	// → 3, true
}
