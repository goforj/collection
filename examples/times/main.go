//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"github.com/goforj/collection"
)

func main() {
	// Times creates a new collection by calling fn(i) for i = 1..count.
	// This mirrors Laravel's Collection::times(), which is 1-indexed.

	// Example: integers - double each index
	cTimes1 := collection.Times(5, func(i int) int {
		return i * 2
	})
	collection.Dump(cTimes1.Items())
	// #[]int [
	//	0 => 2  #int
	//	1 => 4  #int
	//	2 => 6  #int
	//	3 => 8  #int
	//	4 => 10 #int
	// ]

	// Example: strings
	cTimes2 := collection.Times(3, func(i int) string {
		return fmt.Sprintf("item-%d", i)
	})
	collection.Dump(cTimes2.Items())
	// #[]string [
	//	0 => "item-1" #string
	//	1 => "item-2" #string
	//	2 => "item-3" #string
	// ]

	// Example: structs
	type Point struct {
		X int
		Y int
	}

	cTimes3 := collection.Times(4, func(i int) Point {
		return Point{X: i, Y: i * i}
	})
	collection.Dump(cTimes3.Items())
	// #[]main.Point [
	//	0 => #main.Point {
	//		+X => 1 #int
	//		+Y => 1 #int
	//	}
	//	1 => #main.Point {
	//		+X => 2 #int
	//		+Y => 4 #int
	//	}
	//	2 => #main.Point {
	//		+X => 3 #int
	//		+Y => 9 #int
	//	}
	//	3 => #main.Point {
	//		+X => 4 #int
	//		+Y => 16 #int
	//	}
	// ]
}
