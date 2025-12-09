//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {

	   c := New([]int{1, 2, 3, 4})
	   v, ok := c.First()
	Example (empty):
	   c := New([]int{})
	   v, ok := c.First()
	   // v == 1, ok == true


	   // v == 0, ok == false

}
