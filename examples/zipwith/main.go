//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {
	// ZipWith combines two collections element-wise using combiner fn.
	// The resulting length is the smaller of the two inputs.

	// Example: sum ints
	a := collection.New([]int{1, 2, 3})
	b := collection.New([]int{10, 20})

	out := collection.ZipWith(a, b, func(x, y int) int {
		return x + y
	})

	collection.Dump(out.Items())
	// #[]int [
	//   0 => 11 #int
	//   1 => 22 #int
	// ]

	// Example: format strings
	names := collection.New([]string{"alice", "bob"})
	roles := collection.New([]string{"admin", "user", "extra"})

	out2 := collection.ZipWith(names, roles, func(name, role string) string {
		return name + ":" + role
	})

	collection.Dump(out2.Items())
	// #[]string [
	//   0 => "alice:admin" #string
	//   1 => "bob:user" #string
	// ]

	// Example: structs
	type User struct {
		Name string
	}

	type Role struct {
		Title string
	}

	users := collection.New([]User{{Name: "Alice"}, {Name: "Bob"}})
	roles2 := collection.New([]Role{{Title: "admin"}})

	out3 := collection.ZipWith(users, roles2, func(u User, r Role) string {
		return u.Name + " -> " + r.Title
	})

	collection.Dump(out3.Items())
	// #[]string [
	//   0 => "Alice -> admin" #string
	// ]
}
