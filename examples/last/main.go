package main

import "github.com/goforj/collection"

func main() {

	   c := collection.New([]int{1, 2, 3, 4})
	   v, ok := c.Last()
	Example (empty):
	   c := collection.New([]int{})
	   v, ok := c.Last()
	   // v == 4, ok == true


	   // v == 0, ok == false

}
