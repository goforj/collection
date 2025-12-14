//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"github.com/goforj/collection"
)

func main() {
	// DumpStr returns the pretty-printed dump of the items as a string,
	// without printing or exiting.
	// Useful for logging, snapshot testing, and non-interactive debugging.

	// Example: integers
	c := collection.New([]int{10, 20})
	s := c.DumpStr()
	fmt.Println(s)
	// #[]int [
	//   0 => 10 #int
	//   1 => 20 #int
	// ]
}
