//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/collection"
	"strings"
)

func main() {
	// integers
	c1 := collection.New([]int{1, 2, 2, 3, 4, 4, 5})
	out1 := c1.Unique(func(a, b int) bool { return a == b })
	collection.Dump(out1.Items())
	// #[]int [
	//	0 => 1 #int
	//	1 => 2 #int
	//	2 => 3 #int
	//	3 => 4 #int
	//	4 => 5 #int
	// ]

	// strings (case-insensitive uniqueness)
	c2 := collection.New([]string{"A", "a", "B", "b", "A"})
	out2 := c2.Unique(func(a, b string) bool {
		return strings.EqualFold(a, b)
	})
	collection.Dump(out2.Items())
	// #[]string [
	//	0 => "A" #string
	//	1 => "B" #string
	// ]

	// structs (unique by ID)
	type User struct {
		ID   int
		Name string
	}

	c3 := collection.New([]User{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
		{ID: 1, Name: "Alice Duplicate"},
	})

	out3 := c3.Unique(func(a, b User) bool {
		return a.ID == b.ID
	})

	collection.Dump(out3.Items())
	// #[]collection.User [
	//	0 => {ID:1 Name:"Alice"} #collection.User
	//	1 => {ID:2 Name:"Bob"}   #collection.User
	// ]
}
