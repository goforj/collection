package main

import (
	"fmt"
	"github.com/goforj/collection"
)

func main() {
	// integers
	c := collection.New([]int{10, 20})
	s := c.DumpStr()
	fmt.Println(s)
	// #[]int [
	//   0 => 10 #int
	//   1 => 20 #int
	// ]
}
