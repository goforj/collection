//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {
	// Shuffle returns a shuffled copy of the collection.

	// Example: integers
	c := collection.New([]int{1, 2, 3, 4, 5})
	out1 := c.Shuffle()
	collection.Dump(out1.Items())

	// Example: strings – chaining
	out2 := collection.New([]string{"a", "b", "c"}).
		Shuffle().
		Append("d").
		Items()

	collection.Dump(out2)

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

	out3 := users.Shuffle()
	collection.Dump(out3.Items())
}
