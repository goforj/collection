//go:build ignore
// +build ignore

package main

import (
	"github.com/goforj/collection"
	"strings"
)

func main() {
	// integers
	c := collection.New([]int{1, 2, 3})

	sum := 0
	c.Each(func(v int) {
		sum += v
	})

	collection.Dump(sum)
	// 6 #int

	// strings
	c2 := collection.New([]string{"apple", "banana", "cherry"})

	var out []string
	c2.Each(func(s string) {
		out = append(out, strings.ToUpper(s))
	})

	collection.Dump(out)
	// #[]string [
	//   0 => "APPLE"  #string
	//   1 => "BANANA" #string
	//   2 => "CHERRY" #string
	// ]

	// structs
	type User struct {
		ID   int
		Name string
	}

	users := collection.New([]User{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
		{ID: 3, Name: "Charlie"},
	})

	var names []string
	users.Each(func(u User) {
		names = append(names, u.Name)
	})

	collection.Dump(names)
	// #[]string [
	//   0 => "Alice"   #string
	//   1 => "Bob"     #string
	//   2 => "Charlie" #string
	// ]
}
