//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {
	// Example: integers - stop at value 3
	c4 := collection.New([]int{1, 2, 3, 4})
	out4 := collection.TakeUntil(c4, 3)
	collection.Dump(out4.Items())
	// #[]int [
	//	0 => 1 #int
	//	1 => 2 #int
	// ]

	// Example: strings - value never appears → full slice
	c5 := collection.New([]string{"a", "b", "c"})
	out5 := collection.TakeUntil(c5, "x")
	collection.Dump(out5.Items())
	// #[]string [
	//	0 => "a" #string
	//	1 => "b" #string
	//	2 => "c" #string
	// ]

	// Example: integers - match is first item → empty result
	c6 := collection.New([]int{9, 10, 11})
	out6 := collection.TakeUntil(c6, 9)
	collection.Dump(out6.Items())
	// #[]int [
	// ]
}
