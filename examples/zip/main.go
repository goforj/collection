//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {
	// Zip combines two collections element-wise into a collection of tuples.
	// The resulting length is the smaller of the two inputs.

	// Example: integers and strings
	nums := collection.New([]int{1, 2, 3})
	words := collection.New([]string{"one", "two"})

	out := collection.Zip(nums, words)
	collection.Dump(out.Items())
	// #[]collection.Tuple[int,string] [
	//   0 => #collection.Tuple[int,string] {
	//     +First  => 1 #int
	//     +Second => "one" #string
	//   }
	//   1 => #collection.Tuple[int,string] {
	//     +First  => 2 #int
	//     +Second => "two" #string
	//   }
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

	roles := collection.New([]string{"admin", "user", "extra"})

	out2 := collection.Zip(users, roles)
	collection.Dump(out2.Items())
	// #[]collection.Tuple[main.User,string] [
	//   0 => #collection.Tuple[main.User,string] {
	//     +First  => #main.User {
	//       +ID   => 1 #int
	//       +Name => "Alice" #string
	//     }
	//     +Second => "admin" #string
	//   }
	//   1 => #collection.Tuple[main.User,string] {
	//     +First  => #main.User {
	//       +ID   => 2 #int
	//       +Name => "Bob" #string
	//     }
	//     +Second => "user" #string
	//   }
	// ]
}
