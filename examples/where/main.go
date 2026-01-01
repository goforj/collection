//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {
	// Where keeps only the elements for which fn returns true.
	// This is an alias for Filter(fn) for SQL-style ergonomics.
	// This method mutates the collection in place and returns the same instance.

	// Example: integers
	nums := collection.New([]int{1, 2, 3, 4})
	nums.Where(func(v int) bool {
		return v%2 == 0
	})
	collection.Dump(nums.Items())
	// #[]int [
	//   0 => 2 #int
	//   1 => 4 #int
	// ]

	// Example: structs
	type User struct {
		ID   int
		Name string
	}

	users := collection.New([]User{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
		{ID: 3, Name: "Carol"},
	})

	users.Where(func(u User) bool {
		return u.ID >= 2
	})

	collection.Dump(users.Items())
	// #[]main.User [
	//   0 => #main.User {
	//     +ID   => 2 #int
	//     +Name => "Bob" #string
	//   }
	//   1 => #main.User {
	//     +ID   => 3 #int
	//     +Name => "Carol" #string
	//   }
	// ]
}
