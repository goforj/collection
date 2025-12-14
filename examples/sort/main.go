//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {
	// Sort returns a new collection sorted using the provided comparison function.

	// Example: integers
	c := collection.New([]int{5, 1, 4, 2})
	sorted := c.Sort(func(a, b int) bool { return a < b })
	collection.Dump(sorted.Items())
	// #[]int [
	//   0 => 1 #int
	//   1 => 2 #int
	//   2 => 4 #int
	//   3 => 5 #int
	// ]

	// Example: strings (descending)
	c2 := collection.New([]string{"apple", "banana", "cherry"})
	sorted2 := c2.Sort(func(a, b string) bool { return a > b })
	collection.Dump(sorted2.Items())
	// #[]string [
	//   0 => "cherry" #string
	//   1 => "banana" #string
	//   2 => "apple" #string
	// ]

	// Example: structs
	type User struct {
		Name string
		Age  int
	}

	users := collection.New([]User{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
		{Name: "Carol", Age: 40},
	})

	// Sort by age ascending
	sortedUsers := users.Sort(func(a, b User) bool {
		return a.Age < b.Age
	})
	collection.Dump(sortedUsers.Items())
	// #[]main.User [
	//   0 => #main.User {
	//     +Name => "Bob" #string
	//     +Age  => 25 #int
	//   }
	//   1 => #main.User {
	//     +Name => "Alice" #string
	//     +Age  => 30 #int
	//   }
	//   2 => #main.User {
	//     +Name => "Carol" #string
	//     +Age  => 40 #int
	//   }
	// ]
}
