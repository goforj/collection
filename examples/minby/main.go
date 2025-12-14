//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {
	// MinBy returns the item whose key (produced by keyFn) is the smallest.
	// The second return value is false if the collection is empty.

	// Example: structs - smallest age
	type User struct {
		Name string
		Age  int
	}

	users := collection.New([]User{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
		{Name: "Carol", Age: 40},
	})

	minUser, ok := collection.MinBy(users, func(u User) int {
		return u.Age
	})

	collection.Dump(minUser, ok)
	// #main.User {
	//   +Name => "Bob" #string
	//   +Age  => 25 #int
	// }
	// true #bool

	// Example: strings - shortest length
	words := collection.New([]string{"apple", "fig", "banana"})

	shortest, ok := collection.MinBy(words, func(s string) int {
		return len(s)
	})

	collection.Dump(shortest, ok)
	// "fig" #string
	// true #bool

	// Example: empty collection
	empty := collection.New([]int{})
	minVal, ok := collection.MinBy(empty, func(v int) int { return v })
	collection.Dump(minVal, ok)
	// 0 #int
	// false #bool
}
