package collection

// Multiply creates `n` copies of all items in the collection
// and returns a new collection.
//
// Example: integers
//
//	ints := collection.New([]int{1, 2})
//	out := ints.Multiply(3)
//	collection.Dump(out.Items())
//	// #[]int [
//	//   0 => 1 #int
//	//   1 => 2 #int
//	//   2 => 1 #int
//	//   3 => 2 #int
//	//   4 => 1 #int
//	//   5 => 2 #int
//	// ]
//
// Example:
//
//	// strings
//	strs := collection.New([]string{"a", "b"})
//	out2 := strs.Multiply(2)
//	collection.Dump(out2.Items())
//	// #[]string [
//	//   0 => "a" #string
//	//   1 => "b" #string
//	//   2 => "a" #string
//	//   3 => "b" #string
//	// ]
//
// Example:
//
//	// structs
//	type User struct {
//		Name string
//	}
//
//	users := collection.New([]User{{Name: "Alice"}, {Name: "Bob"}})
//	out3 := users.Multiply(2)
//	collection.Dump(out3.Items())
//	// #[]main.User [
//	//   0 => #main.User {
//	//     +Name => "Alice" #string
//	//   }
//	//   1 => #main.User {
//	//     +Name => "Bob" #string
//	//   }
//	//   2 => #main.User {
//	//     +Name => "Alice" #string
//	//   }
//	//   3 => #main.User {
//	//     +Name => "Bob" #string
//	//   }
//	// ]
//
// Example:
//
//	// multiplying by zero or negative returns empty
//	none := ints.Multiply(0)
//	collection.Dump(none.Items())
//	// #[]int [
//	// ]
func (c *Collection[T]) Multiply(n int) *Collection[T] {
	if n <= 0 {
		return New([]T{})
	}

	orig := c.items
	out := make([]T, 0, len(orig)*n)

	for i := 0; i < n; i++ {
		out = append(out, orig...)
	}

	return New(out)
}
