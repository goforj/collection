//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {
	  c := collection.Times(5, func(i int) int { return i * 2 })
	  // [2,4,6,8,10]
}
