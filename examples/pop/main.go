//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {
	// Pop removes and returns the last item in the collection.

	// Example: integers
	c := collection.New([]int{1, 2, 3})
	item, ok := c.Pop()
	collection.Dump(item, ok, c.Items())
	// 3 #int
	// true #bool
	// #[]int [
	//   0 => 1 #int
	//   1 => 2 #int
	// ]

	// Example: strings
	c2 := collection.New([]string{"a", "b", "c"})
	item2, ok2 := c2.Pop()
	collection.Dump(item2, ok2, c2.Items())
	// "c" #string
	// true #bool
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

	item3, ok3 := users.Pop()
	collection.Dump(item3, ok3, users.Items())
	// #main.User {
	//   +ID   => 2 #int
	//   +Name => "Bob" #string
	// }
	// true #bool
	// #[]main.User [
	//   0 => #main.User {
	//     +ID   => 1 #int
	//     +Name => "Alice" #string
	//   }
	// ]

	// Example: empty collection
	empty := collection.New([]int{})
	item4, ok4 := empty.Pop()
	collection.Dump(item4, ok4, empty.Items())
	// 0 #int
	// false #bool
	// #[]int [
	// ]
}
