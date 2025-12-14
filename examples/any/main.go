//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {
	// Any returns true if at least one item satisfies fn.

	// Example: integers
	c := collection.New([]int{1, 2, 3, 4})
	has := c.Any(func(v int) bool { return v%2 == 0 }) // true
	collection.Dump(has)
	// true #bool
}
