package main

import "github.com/goforj/collection"

func main() {

	  c := collection.New([]int{1,2,3})
	  c.Transform(func(v int) int { return v * 2 })
	  // c is now [2,4,6]
}
