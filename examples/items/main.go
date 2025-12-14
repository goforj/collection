//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {
	// Items returns the underlying slice of items.

	// Example: integers
	c := collection.New([]int{1, 2, 3})
	items := c.Items()
	collection.Dump(items)
	// #[]int [
	//   0 => 1 #int
	//   1 => 2 #int
	//   2 => 3 #int
	// ]

	// Example: strings
	c2 := collection.New([]string{"apple", "banana"})
	items2 := c2.Items()
	collection.Dump(items2)
	// #[]string [
	//   0 => "apple" #string
	//   1 => "banana" #string
	// ]

	// Example: structs
	type User struct {
		ID   int
		Name string
	}

	users := collection.New([]User{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
	})

	out := users.Items()
	collection.Dump(out)
	// #[]main.User [
	//   0 => #main.User {
	//     +ID   => 1 #int
	//     +Name => "Alice" #string
	//   }
	//   1 => #main.User {
	//     +ID   => 2 #int
	//     +Name => "Bob" #string
	//   }
	// ]
}
