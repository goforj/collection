//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {
	// GroupBySlice partitions this collection into slices keyed by the extracted value.

	// Example: group integers by parity into slices
	numbers := collection.New([]int{1, 2, 3, 4})
	groups := numbers.GroupBySlice(func(number int) string {
		if number%2 == 0 {
			return "even"
		}
		return "odd"
	})
	collection.Dump(groups["even"], groups["odd"])
	// #[]int [
	//   0 => 2 #int
	//   1 => 4 #int
	// ]
	// #[]int [
	//   0 => 1 #int
	//   1 => 3 #int
	// ]
}
