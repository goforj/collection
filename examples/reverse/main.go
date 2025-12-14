//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {
	// Reverse reverses the order of items in the collection in place
	// and returns the same collection for chaining.

	// Example: integers
	c := collection.New([]int{1, 2, 3, 4})
	c.Reverse()
	collection.Dump(c.Items())
	// #[]int [
	//   0 => 4 #int
	//   1 => 3 #int
	//   2 => 2 #int
	//   3 => 1 #int
	// ]

	// Example: strings – chaining
	out := collection.New([]string{"a", "b", "c"}).
		Reverse().
		Append("d").
		Items()

	collection.Dump(out)
	// #[]string [
	//   0 => "c" #string
	//   1 => "b" #string
	//   2 => "a" #string
	//   3 => "d" #string
	// ]

	// Example: structs
	type User struct {
		ID int
	}

	users := collection.New([]User{
		{ID: 1},
		{ID: 2},
		{ID: 3},
	})

	users.Reverse()
	collection.Dump(users.Items())
	// #[]collection.User [
	//   0 => {ID:3} #collection.User
	//   1 => {ID:2} #collection.User
	//   2 => {ID:1} #collection.User
	// ]
}
