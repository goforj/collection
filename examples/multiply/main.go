//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {
	 users := collection.New([]User{{Name: "A"}, {Name: "B"}})
	 out := users.Multiply(3)
	If n <= 0, the method returns an empty collection.
	// [A, B, A, B, A, B]

}
