//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {
	// capture intermediate state during a chain
	captured1 := []int{}
	c1 := collection.New([]int{3, 1, 2}).
		Sort(func(a, b int) bool { return a < b }). // → [1, 2, 3]
		Tap(func(col *collection.Collection[int]) {
			captured1 = append([]int(nil), col.Items()...) // snapshot copy
		}).
		Filter(func(v int) bool { return v >= 2 })       // → [2, 3]

	collection.Dump(c1.Items())
	collection.Dump(captured1)
	// Use BOTH variables so nothing is "declared and not used"
	// c1 → #[]int [2,3]
	// captured1 → #[]int [1,2,3]

	// tap for debugging without changing flow
	c2 := collection.New([]int{10, 20, 30}).
		Tap(func(col *collection.Collection[int]) {
			collection.Dump(col.Items())
		}).
		Filter(func(v int) bool { return v > 10 })

	collection.Dump(c2.Items()) // ensures c2 is used

	// Tap with struct collection
	type User struct {
		ID   int
		Name string
	}

	users := collection.New([]User{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
	})

	users2 := users.Tap(func(col *collection.Collection[User]) {
		collection.Dump(col.Items())
	})

	collection.Dump(users2.Items()) // ensures users2 is used
}
