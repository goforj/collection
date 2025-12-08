package main

import "github.com/goforj/collection"

func main() {


	  c := collection.New([]int{1, 2, 3})
	  out := c.Dump()
	Dump is typically used while chaining:
	  collection.New([]int{1, 2, 3}).
	      Filter(func(v int) bool { return v > 1 }).
	      Dump()
	This is a no-op on the collection itself and never panics.
	  // Prints a pretty debug dump of [1, 2, 3]
	  // out == c



	  c := collection.New([]int{1, 2, 3})
	  c.Dump() // Pretty-prints [1, 2, 3]

	This function is provided for symmetry with godump.Dump.
}
