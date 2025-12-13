//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {
	// Example: integers
	nums := New([]int{1, 2, 3, 4, 5})
	v, ok := nums.FirstWhere(func(n int) bool {
	    return n%2 == 0
	})
	// v = 2, ok = true

	v, ok = nums.FirstWhere(func(n int) bool {
	    return n > 10
	})
	// v = 0, ok = false
}
