package main

import "github.com/goforj/collection"

func main() {


	  c := collection.New([]string{"a", "b"})
	  c.Dd()    // Prints the dump and exits the program

	Like Laravel's dd(), this is intended for debugging and
	should not be used in production control flow.

	This method never returns.
	  c := collection.New([]string{"x", "y"})
	  c.Dd() // Pretty-prints ["x", "y"] and exits

	This function is provided for symmetry with godump.Dd.
}
