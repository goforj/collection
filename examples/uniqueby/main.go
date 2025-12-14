//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/collection"
	"strings"
)

func main() {
	// Example: structs – unique by ID
	type User struct {
		ID   int
		Name string
	}

	users := collection.New([]User{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
		{ID: 1, Name: "Alice Duplicate"},
	})

	out := collection.UniqueBy(users, func(u User) int { return u.ID })
	collection.Dump(out.Items())
	// #[]collection.User [
	//   0 => {ID:1 Name:"Alice"} #collection.User
	//   1 => {ID:2 Name:"Bob"}   #collection.User
	// ]

	// Example: strings – case-insensitive uniqueness
	values := collection.New([]string{"A", "a", "B", "b", "A"})

	out2 := collection.UniqueBy(values, func(s string) string {
		return strings.ToLower(s)
	})

	collection.Dump(out2.Items())
	// #[]string [
	//   0 => "A" #string
	//   1 => "B" #string
	// ]

	// Example: integers – identity key
	nums := collection.New([]int{3, 1, 2, 1, 3})

	out3 := collection.UniqueBy(nums, func(v int) int { return v })
	collection.Dump(out3.Items())
	// #[]int [
	//   0 => 3 #int
	//   1 => 1 #int
	//   2 => 2 #int
	// ]
}
