//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {
	// Example: integers
	 nums := collection.New([]int{1, 2}).Push(3, 4)
	 // nums = [1, 2, 3, 4]

	 // Complex type (structs)
	 type User struct {
	     Name string
	     Age  int
	 }

	 users := collection.New([]User{
	     {Name: "Alice", Age: 30},
	     {Name: "Bob",   Age: 25},
	 }).Push(
	     User{Name: "Carol", Age: 40},
	     User{Name: "Dave",  Age: 20},
	 )

	 // users = [
	 //   {Alice 30},
	 //   {Bob 25},
	 //   {Carol 40},
	 //   {Dave 20},
	 // ]

}
