//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {
	// Tap invokes fn with the collection pointer for side effects (logging, debugging,
	// inspection) and returns the same collection to allow chaining.

	// Example: integers - capture intermediate state during a chain
	captured1 := []int{}
	c1 := collection.New([]int{3, 1, 2}).
		Sort(func(a, b int) bool { return a < b }). // → [1, 2, 3]
		Tap(func(col *collection.Collection[int]) {
			captured1 = append([]int(nil), col.Items()...) // snapshot copy
		}).
		Filter(func(v int) bool { return v >= 2 }).
		Dump()
		// #[]int [
		//  0 => 2 #int
		//  1 => 3 #int
		// ]

	// Use BOTH variables so nothing is "declared and not used"
	collection.Dump(c1.Items())
	collection.Dump(captured1)
	// #[]int [
	//  0 => 2 #int
	//  1 => 3 #int
	// ]
	// #[]int [
	//  0 => 1 #int
	//  1 => 2 #int
	//  2 => 3 #int
	// ]

	// Example: integers - tap for debugging without changing flow
	c2 := collection.New([]int{10, 20, 30}).
		Tap(func(col *collection.Collection[int]) {
			collection.Dump(col.Items())
			// #[]int [
			//  0 => 10 #int
			//  1 => 20 #int
			//  2 => 30 #int
			// ]
		}).
		Filter(func(v int) bool { return v > 10 })

	collection.Dump(c2.Items()) // ensures c2 is used
	// #[]int [
	//  0 => 20 #int
	//  1 => 30 #int
	// ]

	// Example: structs - Tap with struct collection
	type User struct {
		ID   int
		Name string
	}

	users := collection.New([]User{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
	})

	users2 := users.Tap(func(col *collection.Collection[User]) {
		collection.Dump(col.Items())
		// #[]main.User [
		//  0 => #main.User {
		//    +ID   => 1 #int
		//    +Name => "Alice" #string
		//  }
		//  1 => #main.User {
		//    +ID   => 2 #int
		//    +Name => "Bob" #string
		//  }
		// ]
	})

	collection.Dump(users2.Items()) // ensures users2 is used
	// #[]main.User [
	//  0 => #main.User {
	//    +ID   => 1 #int
	//    +Name => "Alice" #string
	//  }
	//  1 => #main.User {
	//    +ID   => 2 #int
	//    +Name => "Bob" #string
	//  }
	// ]
}
