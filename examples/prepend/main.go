//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {
	 c := collection.New([]int{3, 4})
	 newC := c.Prepend(1, 2) // Collection with items [1, 2, 3, 4]
	 // newC.Items() == []int{1, 2, 3, 4}
}
