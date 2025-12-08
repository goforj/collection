package main

import "github.com/goforj/collection"

func main() {

	  collection.NewNumeric([]int{1, 2, 2, 3}).Mode() // → []int{2}

	Example (tie):
	  collection.NewNumeric([]int{1, 2, 1, 2}).Mode() // → []int{1, 2}
}
