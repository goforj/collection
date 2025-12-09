package main

import "github.com/goforj/collection"

func main() {

	  c := New([]int{1, 2, 3})
	  sum := c.Pipe(func(col Collection[int]) any {
	      return col.Sum()
	  })

	  // sum == 6

}
