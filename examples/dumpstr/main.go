package main

import "github.com/goforj/collection"

func main() {

	  c := collection.New([]int{10, 20})
	  s := c.DumpStr()
	  fmt.Println(s)
	Useful for logging, snapshot testing, and non-interactive debugging.
	  // Produces a multi-line formatted representation of [10, 20]

}
