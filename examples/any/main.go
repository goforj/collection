package main

import "github.com/goforj/collection"

func main() {

	c := collection.New([]int{1, 2, 3, 4})
	hasEven := c.Any(func(v int) bool { return v%2 == 0 }) // true
	// hasEven is true
}
