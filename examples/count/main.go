//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {
	count := collection.New([]int{1, 2, 3, 4}).Count()
	collection.Dump(count)
	// 4 #int
}
