//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {
	// Example: strings
	c := collection.New([]string{"a", "b"})
	c.Dd()
	// #[]string [
	//   0 => "a" #string
	//   1 => "b" #string
	// ]
	// Process finished with the exit code 1
}
