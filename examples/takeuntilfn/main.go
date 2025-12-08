package main

import "github.com/goforj/collection"

func main() {


	c := collection.New([]int{1, 2, 3, 4})
	out := c.TakeUntilFn(func(v int) bool { return v >= 3 }) // [1, 2]
	// result is [1, 2]
}
