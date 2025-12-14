//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"github.com/goforj/collection"
)

func main() {
	// ToJSON converts the collection's items into a compact JSON string.
	// 
	// If marshalling succeeds, a JSON-encoded string and a nil error are returned.
	// If marshalling fails, the method unwraps any json.Marshal wrapping so that
	// user-defined MarshalJSON errors surface directly.
	// 
	// Returns:
	//   - string: JSON-encoded representation of the collection
	//   - error : nil on success, or the unwrapped marshalling error

	// Example: strings - pretty JSON
	pj1 := collection.New([]string{"a", "b"})
	out1, _ := pj1.ToJSON()
	fmt.Println(out1)
	// ["a","b"]
}
