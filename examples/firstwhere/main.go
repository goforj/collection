//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {
	// FirstWhere returns the first item in the collection for which the provided
	// predicate function returns true. If no items match, ok=false is returned
	// along with the zero value of T.
	// 
	// This method is equivalent to Laravel's collection->first(fn) and mirrors
	// the behavior found in functional collections in other languages.

	// Example: integers
	nums := collection.New([]int{1, 2, 3, 4, 5})
	v, ok := nums.FirstWhere(func(n int) bool {
		return n%2 == 0
	})
	collection.Dump(v, ok)
	// 2 #int
	// true #bool

	v, ok = nums.FirstWhere(func(n int) bool {
		return n > 10
	})
	collection.Dump(v, ok)
	// 0 #int
	// false #bool
}
