//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {
	// SymmetricDifference returns a new collection containing elements that appear
	// in exactly one of the two collections. Order follows the first collection for
	// its unique items, then the second for its unique items. Duplicates are removed.

	// Example: integers
	a := collection.New([]int{1, 2, 3, 3})
	b := collection.New([]int{3, 4, 4, 5})

	out := collection.SymmetricDifference(a, b)
	collection.Dump(out.Items())
	// #[]int [
	//   0 => 1 #int
	//   1 => 2 #int
	//   2 => 4 #int
	//   3 => 5 #int
	// ]

	// Example: strings
	left := collection.New([]string{"apple", "banana"})
	right := collection.New([]string{"banana", "date"})

	out2 := collection.SymmetricDifference(left, right)
	collection.Dump(out2.Items())
	// #[]string [
	//   0 => "apple" #string
	//   1 => "date" #string
	// ]

	// Example: structs
	type User struct {
		ID   int
		Name string
	}

	groupA := collection.New([]User{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
	})

	groupB := collection.New([]User{
		{ID: 2, Name: "Bob"},
		{ID: 3, Name: "Carol"},
	})

	out3 := collection.SymmetricDifference(groupA, groupB)
	collection.Dump(out3.Items())
	// #[]main.User [
	//   0 => #main.User {
	//     +ID   => 1 #int
	//     +Name => "Alice" #string
	//   }
	//   1 => #main.User {
	//     +ID   => 3 #int
	//     +Name => "Carol" #string
	//   }
	// ]
}
