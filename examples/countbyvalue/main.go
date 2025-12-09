//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {
	  counts := CountByValue(collection.New([]string{"a", "b", "a"}))
	 // counts == map[string]int{"a": 2, "b": 1}
}
