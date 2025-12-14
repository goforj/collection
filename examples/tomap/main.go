//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {
	// Example: basic usage
	users := []string{"alice", "bob", "carol"}

	out := collection.ToMap(
		collection.New(users),
		func(name string) string { return name },
		func(name string) int { return len(name) },
	)

	collection.Dump(out)

	// Example: re-keying structs
	type User struct {
		ID   int
		Name string
	}

	users2 := []User{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
	}

	byID := collection.ToMap(
		collection.New(users2),
		func(u User) int { return u.ID },
		func(u User) User { return u },
	)

	collection.Dump(byID)
}
