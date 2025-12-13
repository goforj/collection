//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {
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
