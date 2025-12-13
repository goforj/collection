//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {
	// Example: integers – computing a sum
	c := collection.New([]int{1, 2, 3})
	sum := c.Pipe(func(col *collection.Collection[int]) any {
		total := 0
		for _, v := range col.Items() {
			total += v
		}
		return total
	})
	collection.Dump(sum)
	// 6 #int

	// strings – joining values
	c2 := collection.New([]string{"a", "b", "c"})
	joined := c2.Pipe(func(col *collection.Collection[string]) any {
		out := ""
		for _, v := range col.Items() {
			out += v
		}
		return out
	})
	collection.Dump(joined)
	// "abc" #string

	// structs – extracting just the names
	type User struct {
		ID   int
		Name string
	}

	users := collection.New([]User{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
	})

	names := users.Pipe(func(col *collection.Collection[User]) any {
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
