//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {
	// Pop returns the last item and a new collection with that item removed.
	// The original collection remains unchanged.

	// Example: integers
	c := collection.New([]int{1, 2, 3})
	item, rest := c.Pop()
	collection.Dump(item, rest.Items())
	// 3 #int
	// #[]int [
	//   0 => 1 #int
	//   1 => 2 #int
	// ]

	// Example: strings
	c2 := collection.New([]string{"a", "b", "c"})
	item2, rest2 := c2.Pop()
	collection.Dump(item2, rest2.Items())
	// "c" #string
	// #[]string [
	//   0 => "a" #string
	//   1 => "b" #string
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

	item3, rest3 := users.Pop()
	collection.Dump(item3, rest3.Items())
	// #main.User {
	//   +ID   => 2 #int
	//   +Name => "Bob" #string
	// }
	// #[]main.User [
	//   0 => #main.User {
	//     +ID   => 1 #int
	//     +Name => "Alice" #string
	//   }
	// ]

	// Example: empty collection
	empty := collection.New([]int{})
	item4, rest4 := empty.Pop()
	collection.Dump(item4, rest4.Items())
	// 0 #int
	// #[]int [
	// ]
}
