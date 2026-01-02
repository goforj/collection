//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {
	// At returns the item at the given index and a boolean indicating
	// whether the index was within bounds.

	// Example: integers
	c := collection.New([]int{10, 20, 30})
	v, ok := c.At(1)
	collection.Dump(v, ok)
	// 20 #int
	// true #bool

	// Example: out of bounds
	v2, ok2 := c.At(10)
	collection.Dump(v2, ok2)
	// 0 #int
	// false #bool

	// Example: structs
	type User struct {
		ID   int
		Name string
	}

	users := collection.New([]User{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
	})

	u, ok3 := users.At(0)
	collection.Dump(u, ok3)
	// #main.User {
	//   +ID   => 1 #int
	//   +Name => "Alice" #string
	// }
	// true #bool
}
