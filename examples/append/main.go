package main

import "github.com/goforj/collection"

func main() {

	 c := collection.New([]int{1, 2})
	 newC := c.Append(3, 4) // Collection with items [1, 2, 3, 4]
	 // newC.Items() == []int{1, 2, 3, 4}
}
