//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {
	 collection.New([]int{1,2,3,4}).
		Filter(func(v int) bool { return v%2 == 0 }).
		Items()
	// []int{2,4}
}
