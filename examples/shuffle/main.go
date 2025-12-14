//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {
	// Shuffle randomly shuffles the items in the collection in place
	// and returns the same collection for chaining.
	// 
	// This operation performs no allocations.
	// 
	// The shuffle uses an internal random source. Tests may override
	// this source to achieve deterministic behavior.

	// Example: integers
	c := collection.New([]int{1, 2, 3, 4, 5})
	c.Shuffle()
	collection.Dump(c.Items())

	// Example: strings – chaining
	out := collection.New([]string{"a", "b", "c"}).
		Shuffle().
		Append("d").
		Items()

	collection.Dump(out)

	// Example: structs
	type User struct {
		ID int
	}

	users := collection.New([]User{
		{ID: 1},
		{ID: 2},
		{ID: 3},
		{ID: 4},
	})

	users.Shuffle()
	collection.Dump(users.Items())
}
