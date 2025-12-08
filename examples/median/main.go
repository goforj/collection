package main

import "github.com/goforj/collection"

func main() {

	  c := collection.NewNumeric([]int{3, 1, 2})
	  median, ok := c.Median()   // → 2, true
}
