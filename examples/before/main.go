//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {
	// Before returns a new collection containing all items that appear
	// *before* the first element for which pred returns true.
	// 
	// If no element matches the predicate, the entire collection is returned.

	// Example: integers
	c1 := collection.New([]int{1, 2, 3, 4, 5})
	out1 := c1.Before(func(v int) bool { return v >= 3 })
	collection.Dump(out1.Items())
	// #[]int [
	//	0 => 1 #int
	//	1 => 2 #int
	// ]

	// Example: predicate never matches → whole collection returned
	c2 := collection.New([]int{10, 20, 30})
	out2 := c2.Before(func(v int) bool { return v == 99 })
	collection.Dump(out2.Items())
	// #[]int [
	//	0 => 10 #int
	//	1 => 20 #int
	//	2 => 30 #int
	// ]

	// Example: structs: get all users before the first admin
	type User struct {
		Name  string
		Admin bool
	}

	c3 := collection.New([]User{
		{Name: "Alice", Admin: false},
		{Name: "Bob", Admin: false},
		{Name: "Eve", Admin: true},
		{Name: "Mallory", Admin: false},
	})

	out3 := c3.Before(func(u User) bool { return u.Admin })
	collection.Dump(out3.Items())
	// #[]collection.User [
	//	0 => {Name:"Alice" Admin:false}  #collection.User
	//	1 => {Name:"Bob"   Admin:false}  #collection.User
	// ]
}
