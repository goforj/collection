//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {
	// IsEmpty returns true if the collection has no items.

	// Example: integers (non-empty)
	c := collection.New([]int{1, 2, 3})

	empty := c.IsEmpty()
	collection.Dump(empty)
	// false #bool

	// Example: strings (empty)
	c2 := collection.New([]string{})

	empty2 := c2.IsEmpty()
	collection.Dump(empty2)
	// true #bool

	// Example: structs (non-empty)
	type User struct {
		ID   int
		Name string
	}

	users := collection.New([]User{
		{ID: 1, Name: "Alice"},
	})

	empty3 := users.IsEmpty()
	collection.Dump(empty3)
	// false #bool

	// Example: structs (empty)
	none := collection.New([]User{})

	empty4 := none.IsEmpty()
	collection.Dump(empty4)
	// true #bool
}
