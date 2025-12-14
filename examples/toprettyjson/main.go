//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"github.com/goforj/collection"
)

func main() {
	// ToPrettyJSON converts the collection's items into a human-readable,
	// indented JSON string.
	// 
	// If marshalling succeeds, a formatted JSON string and nil error are returned.
	// If marshalling fails, the underlying error is unwrapped so user-defined
	// MarshalJSON failures surface directly.
	// 
	// Returns:
	//   - string: the pretty-printed JSON representation
	//   - error : nil on success, or the unwrapped marshalling error

	// Example: strings - pretty JSON
	pj1 := collection.New([]string{"a", "b"})
	out1, _ := pj1.ToPrettyJSON()
	fmt.Println(out1)
	// [
	//  "a",
	//  "b"
	// ]
}
