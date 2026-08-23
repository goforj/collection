//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {
	// GroupBy partitions this collection into collections keyed by the extracted value.

	// Example: group integers by parity
	numbers := collection.New([]int{1, 2, 3, 4})
	groups := numbers.GroupBy(func(number int) string {
		if number%2 == 0 {
			return "even"
		}
		return "odd"
	})
	collection.Dump(groups["even"].Items(), groups["odd"].Items())
	// #[]int [
	//   0 => 2 #int
	//   1 => 4 #int
	// ]
	// #[]int [
	//   0 => 1 #int
	//   1 => 3 #int
	// ]
}
