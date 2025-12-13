//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {
	// Example: integers - merging slices
	ints := collection.New([]int{1, 2})
	extra := []int{3, 4}
	// Merge the extra slice into the ints collection
	merged1 := ints.Merge(extra)
	collection.Dump(merged1.Items())
	// #[]int [
	//   0 => 1 #int
	//   1 => 2 #int
	//   2 => 3 #int
	//   3 => 4 #int
	// ]

	// Example: strings - merging another collection
	strs := collection.New([]string{"a", "b"})
	more := collection.New([]string{"c", "d"})

	merged2 := strs.Merge(more)
	collection.Dump(merged2.Items())
	// #[]string [
	//   0 => "a" #string
	//   1 => "b" #string
	//   2 => "c" #string
	//   3 => "d" #string
	// ]

	// Example: structs - merging struct slices
	type User struct {
		ID   int
		Name string
	}

	users := collection.New([]User{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
	})

	moreUsers := []User{
		{ID: 3, Name: "Carol"},
		{ID: 4, Name: "Dave"},
	}

	merged3 := users.Merge(moreUsers)
	collection.Dump(merged3.Items())
	// #[]main.User [
	//   0 => #main.User {
	//     +ID   => 1 #int
	//     +Name => "Alice" #string
	//   }
	//   1 => #main.User {
	//     +ID   => 2 #int
	//     +Name => "Bob" #string
	//   }
	//   2 => #main.User {
	//     +ID   => 3 #int
	//     +Name => "Carol" #string
	//   }
	//   3 => #main.User {
	//     +ID   => 4 #int
	//     +Name => "Dave" #string
	//   }
	// ]
}
