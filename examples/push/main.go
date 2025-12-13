//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {
	// Example: integers
	nums := collection.New([]int{1, 2}).Push(3, 4)
	nums.Dump()
	// #[]int [
	//  0 => 1 #int
	//  1 => 2 #int
	//  2 => 3 #int
	//  3 => 4 #int
	// ]

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
	users.Dump()
	// #[]main.User [
	//  0 => #main.User {
	//    +Name => "Alice" #string
	//    +Age  => 30 #int
	//  }
	//  1 => #main.User {
	//    +Name => "Bob" #string
	//    +Age  => 25 #int
	//  }
	//  2 => #main.User {
	//    +Name => "Carol" #string
	//    +Age  => 40 #int
	//  }
	//  3 => #main.User {
	//    +Name => "Dave" #string
	//    +Age  => 20 #int
	//  }
	// ]
}
