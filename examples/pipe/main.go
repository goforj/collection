//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {
	// Pipe passes the entire collection into the given function
	// and returns the function's result.

	// Example: integers – computing a sum
	c := collection.New([]int{1, 2, 3})
	sum := collection.Pipe(c, func(col *collection.Collection[int]) int {
		total := 0
		for _, v := range col.Items() {
			total += v
		}
		return total
	})
	collection.Dump(sum)
	// 6 #int

	// Example: strings – joining values
	c2 := collection.New([]string{"a", "b", "c"})
	joined := collection.Pipe(c2, func(col *collection.Collection[string]) string {
		out := ""
		for _, v := range col.Items() {
			out += v
		}
		return out
	})
	collection.Dump(joined)
	// "abc" #string

	// Example: structs – extracting just the names
	type User struct {
		ID   int
		Name string
	}

	users := collection.New([]User{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
	})

	names := collection.Pipe(users, func(col *collection.Collection[User]) []string {
		result := make([]string, 0, len(col.Items()))
		for _, u := range col.Items() {
			result = append(result, u.Name)
		}
		return result
	})

	collection.Dump(names)
	// #[]string [
	//   0 => "Alice" #string
	//   1 => "Bob" #string
	// ]
}
