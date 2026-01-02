//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {
	// Contains returns true if the collection contains the given value.

	// Example: integers
	c := collection.New([]int{1, 2, 3, 4, 5})
	hasTwo := collection.Contains(c, 2)
	collection.Dump(hasTwo)
	// true #bool

	// Example: strings
	c2 := collection.New([]string{"apple", "banana", "cherry"})
	hasBanana := collection.Contains(c2, "banana")
	collection.Dump(hasBanana)
	// true #bool
}
