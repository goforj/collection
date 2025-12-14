//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {
	// Pluck is an alias for MapTo with a more semantic name when projecting fields.
	// It extracts a single field or computed value from every element and returns a
	// new typed collection.

	// Example: integers - extract parity label
	nums := collection.New([]int{1, 2, 3, 4})
	parity := collection.Pluck(nums, func(n int) string {
		if n%2 == 0 {
			return "even"
		}
		return "odd"
	})
	collection.Dump(parity.Items())
	// #[]string [
	//   0 => "odd" #string
	//   1 => "even" #string
	//   2 => "odd" #string
	//   3 => "even" #string
	// ]

	// Example: strings - length of each value
	words := collection.New([]string{"go", "forj", "rocks"})
	lengths := collection.Pluck(words, func(s string) int {
		return len(s)
	})
	collection.Dump(lengths.Items())
	// #[]int [
	//   0 => 2 #int
	//   1 => 4 #int
	//   2 => 5 #int
	// ]

	// Example: structs - pluck a field
	type User struct {
		ID   int
		Name string
	}

	users := collection.New([]User{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
	})

	names := collection.Pluck(users, func(u User) string {
		return u.Name
	})

	collection.Dump(names.Items())
	// #[]string [
	//   0 => "Alice" #string
	//   1 => "Bob" #string
	// ]
}
