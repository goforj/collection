//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"github.com/goforj/collection"
)

func main() {
	// ToJSON converts the collection's items into a compact JSON string.

	// Example: strings - pretty JSON
	pj1 := collection.New([]string{"a", "b"})
	out1, _ := pj1.ToJSON()
	fmt.Println(out1)
	// ["a","b"]
}
