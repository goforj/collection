//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {
	// Pipe passes this collection to fn and returns fn's result.

	// Example: sum a collection
	numbers := collection.New([]int{1, 2, 3})
	total := numbers.Pipe(func(values *collection.Collection[int]) int {
		sum := 0
		for _, value := range values.Items() {
			sum += value
		}
		return sum
	})
	collection.Dump(total)
	// 6 #int
}
