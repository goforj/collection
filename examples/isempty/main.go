//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {
	// integers (non-empty)
	c := collection.New([]int{1, 2, 3})

	empty := c.IsEmpty()
	collection.Dump(empty)
	// false #bool

	// strings (empty)
	c2 := collection.New([]string{})

	empty2 := c2.IsEmpty()
	collection.Dump(empty2)
	// true #bool

	// structs (non-empty)
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

	// structs (empty)
	none := collection.New([]User{})

	empty4 := none.IsEmpty()
	collection.Dump(empty4)
	// true #bool
}
