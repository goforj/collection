//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {
	// CountBy returns a map of keys extracted by fn to their occurrence counts.
	// K must be comparable.

	// Example: integers
	c := collection.New([]int{1, 2, 2, 3, 3, 3})
	counts := collection.CountBy(c, func(v int) int {
		return v
	})
	collection.Dump(counts)
	// #map[int]int {
	//   1 => 1 #int
	//   2 => 2 #int
	//   3 => 3 #int
	// }

	// Example: strings
	c2 := collection.New([]string{"apple", "banana", "apple", "cherry", "banana"})
	counts2 := collection.CountBy(c2, func(v string) string {
		return v
	})
	collection.Dump(counts2)
	// #map[string]int {
	//   apple => 2 #int
	//   banana => 2 #int
	//   cherry => 1 #int
	// }

	// Example: structs
	type User struct {
		Name string
		Role string
	}

	users := collection.New([]User{
		{Name: "Alice", Role: "admin"},
		{Name: "Bob", Role: "user"},
		{Name: "Carol", Role: "admin"},
		{Name: "Dave", Role: "user"},
		{Name: "Eve", Role: "admin"},
	})

	roleCounts := collection.CountBy(users, func(u User) string {
		return u.Role
	})

	collection.Dump(roleCounts)
	// #map[string]int {
	//   admin => 3 #int
	//   user => 2 #int
	// }
}
