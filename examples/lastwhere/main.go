//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {

	   c := collection.New([]int{1, 2, 3, 4})
	   v, ok := c.LastWhere(func(v int, i int) bool {
	       return v < 3
	   })
	   // v == 2, ok == true


	   c := collection.New([]int{1, 2, 3, 4})
	   v, ok := c.LastWhere(nil)
	   // v == 4, ok == true


	   c := collection.New([]int{})
	   v, ok := c.LastWhere(nil)
	   // v == 0, ok == false

}
