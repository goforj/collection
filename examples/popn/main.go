//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {
	// PopN removes and returns the last n items as a new collection,
	// and returns a second collection containing the remaining items.

	// Example: integers – pop 2
	c := collection.New([]int{1, 2, 3, 4})
	popped, rest := c.PopN(2)
	collection.Dump(popped.Items(), rest.Items())
	// #[]int [
	//   0 => 4 #int
	//   1 => 3 #int
	// ]
	// #[]int [
	//   0 => 1 #int
	//   1 => 2 #int
	// ]

	// Example: strings – pop 1
	c2 := collection.New([]string{"a", "b", "c"})
	popped2, rest2 := c2.PopN(1)
	collection.Dump(popped2.Items(), rest2.Items())
	// #[]string [
	//   0 => "c" #string
	// ]
	// #[]string [
	//   0 => "a" #string
	//   1 => "b" #string
	// ]

	// Example: structs – pop 2
	type User struct {
		ID   int
		Name string
	}

	users := collection.New([]User{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
		{ID: 3, Name: "Carol"},
	})

	popped3, rest3 := users.PopN(2)
	collection.Dump(popped3.Items(), rest3.Items())
	// #[]main.User [
	//   0 => #main.User {
	//     +ID   => 3 #int
	//     +Name => "Carol" #string
	//   }
	//   1 => #main.User {
	//     +ID   => 2 #int
	//     +Name => "Bob" #string
	//   }
	// ]
	// #[]main.User [
	//   0 => #main.User {
	//     +ID   => 1 #int
	//     +Name => "Alice" #string
	//   }
	// ]

	// Example: integers - n <= 0 → returns empty popped + original collection
	c3 := collection.New([]int{1, 2, 3})
	popped4, rest4 := c3.PopN(0)
	collection.Dump(popped4.Items(), rest4.Items())
	// #[]int [
	// ]
	// #[]int [
	//   0 => 1 #int
	//   1 => 2 #int
	//   2 => 3 #int
	// ]

	// Example: strings - n exceeds length → all items popped, rest empty
	c4 := collection.New([]string{"x", "y"})
	popped5, rest5 := c4.PopN(10)
	collection.Dump(popped5.Items(), rest5.Items())
	// #[]string [
	//   0 => "y" #string
	//   1 => "x" #string
	// ]
	// #[]string [
	// ]
}
