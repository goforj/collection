//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {
	// Dd prints items then terminates execution.
	// Like Laravel's dd(), this is intended for debugging and
	// should not be used in production control flow.

	// Example: strings
	c := collection.New([]string{"a", "b"})
	c.Dd()
	// #[]string [
	//   0 => "a" #string
	//   1 => "b" #string
	// ]
	// Process finished with the exit code 1
}
