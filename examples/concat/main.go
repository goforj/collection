//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {
	// Concat appends the values from the given slice onto the end of the collection,

	// Example: strings
	c := collection.New([]string{"John Doe"})
	concatenated := c.
		Concat([]string{"Jane Doe"}).
		Concat([]string{"Johnny Doe"}).
		Items()
	collection.Dump(concatenated)

	// #[]string [
	//  0 => "John Doe" #string
	//  1 => "Jane Doe" #string
	//  2 => "Johnny Doe" #string
	// ]
}
