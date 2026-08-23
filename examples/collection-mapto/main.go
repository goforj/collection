//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {
	// MapTo maps this collection to a collection with a different element type.

	// Example: map integers to labels
	numbers := collection.New([]int{1, 2, 3, 4})
	labels := numbers.MapTo(func(number int) string {
		if number%2 == 0 {
			return "even"
		}
		return "odd"
	})
	collection.Dump(labels.Items())
	// #[]string [
	//   0 => "odd" #string
	//   1 => "even" #string
	//   2 => "odd" #string
	//   3 => "even" #string
	// ]
}
