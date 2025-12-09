//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {
	  squared := numbers.MapTo(func(n int) int { return n * n })
	  // squared is a Collection[int] of squared numbers
}
