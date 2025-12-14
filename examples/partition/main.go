//go:build ignore
// +build ignore

package main

import (
	"strings"

	"github.com/goforj/collection"
)

func main() {
	// Partition splits the collection into two new collections based on predicate fn.
	// The first collection contains items where fn returns true; the second contains
	// items where fn returns false. Order is preserved within each partition.

	// Example: integers - even/odd
	nums := collection.New([]int{1, 2, 3, 4, 5})
	evens, odds := nums.Partition(func(n int) bool {
		return n%2 == 0
	})
	collection.Dump(evens.Items(), odds.Items())
	// #[]int [
	//   0 => 2 #int
	//   1 => 4 #int
	// ]
	// #[]int [
	//   0 => 1 #int
	//   1 => 3 #int
	//   2 => 5 #int
	// ]

	// Example: strings - prefix match
	words := collection.New([]string{"go", "gopher", "rust", "ruby"})
	goWords, other := words.Partition(func(s string) bool {
		return strings.HasPrefix(s, "go")
	})
	collection.Dump(goWords.Items(), other.Items())
	// #[]string [
	//   0 => "go" #string
	//   1 => "gopher" #string
	// ]
	// #[]string [
	//   0 => "rust" #string
	//   1 => "ruby" #string
	// ]

	// Example: structs - active vs inactive
	type User struct {
		Name   string
		Active bool
	}

	users := collection.New([]User{
		{Name: "Alice", Active: true},
		{Name: "Bob", Active: false},
		{Name: "Carol", Active: true},
	})

	active, inactive := users.Partition(func(u User) bool {
		return u.Active
	})

	collection.Dump(active.Items(), inactive.Items())
	// #[]main.User [
	//   0 => #main.User {
	//     +Name   => "Alice" #string
	//     +Active => true #bool
	//   }
	//   1 => #main.User {
	//     +Name   => "Carol" #string
	//     +Active => true #bool
	//   }
	// ]
	// #[]main.User [
	//   0 => #main.User {
	//     +Name   => "Bob" #string
	//     +Active => false #bool
	//   }
	// ]
}
