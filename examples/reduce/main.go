//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {
	// Reduce collapses the collection into a single accumulated value.
	// The accumulator has the same type T as the collection's elements.

	// Example: integers - sum
	sum := collection.New([]int{1, 2, 3}).Reduce(0, func(acc, n int) int {
		return acc + n
	})
	collection.Dump(sum)
	// 6 #int

	// Example: strings
	joined := collection.New([]string{"a", "b", "c"}).Reduce("", func(acc, s string) string {
		return acc + s
	})
	collection.Dump(joined)
	// "abc" #string

	// Example: structs
	type Stats struct {
		Count int
		Sum   int
	}

	stats := collection.New([]Stats{
		{Count: 1, Sum: 10},
		{Count: 1, Sum: 20},
		{Count: 1, Sum: 30},
	})

	total := stats.Reduce(Stats{}, func(acc, s Stats) Stats {
		acc.Count += s.Count
		acc.Sum += s.Sum
		return acc
	})

	collection.Dump(total)
	// #main.Stats [
	//   +Count => 3 #int
	//   +Sum   => 60 #int
	// ]
}
