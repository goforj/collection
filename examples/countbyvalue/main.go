//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {
	// CountByValue returns a map where each distinct item in the collection
	// is mapped to the number of times it appears.

	// Example: strings
	c1 := collection.New([]string{"a", "b", "a"})
	counts1 := collection.CountByValue(c1)
	collection.Dump(counts1)
	// #map[string]int [
	//	"a" => 2 #int
	//	"b" => 1 #int
	// ]

	// Example: integers
	c2 := collection.New([]int{1, 2, 2, 3, 3, 3})
	counts2 := collection.CountByValue(c2)
	collection.Dump(counts2)
	// #map[int]int [
	//	1 => 1 #int
	//	2 => 2 #int
	//	3 => 3 #int
	// ]

	// Example: structs (comparable)
	type Point struct {
		X int
		Y int
	}

	c3 := collection.New([]Point{
		{X: 1, Y: 1},
		{X: 2, Y: 2},
		{X: 1, Y: 1},
	})

	counts3 := collection.CountByValue(c3)
	collection.Dump(counts3)
	// #map[collection.Point]int [
	//	{X:1 Y:1} => 2 #int
	//	{X:2 Y:2} => 1 #int
	// ]
}
