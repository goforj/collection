package collection

// Append returns a new collection with the given values appended.
// @group Transformation
// @behavior immutable
// @fluent true
// Example: integers
//
//	c := collection.New([]int{1, 2})
//	c.Append(3, 4).Dump()
//	// #[]int [
//	//  0 => 1 #int
//	//  1 => 2 #int
//	//  2 => 3 #int
//	//  3 => 4 #int
//	// ]
//
// Example: structs
//
//	type User struct {
//		ID   int
//		Name string
//	}
//
//	users := collection.New([]User{
//		{ID: 1, Name: "Alice"},
//		{ID: 2, Name: "Bob"},
//	})
//
//	users.Append(
//		User{ID: 3, Name: "Carol"},
//		User{ID: 4, Name: "Dave"},
//	).Dump()
//
//	// #[]main.User [
//	//  0 => #main.User {
//	//    +ID   => 1 #int
//	//    +Name => "Alice" #string
//	//  }
//	//  1 => #main.User {
//	//    +ID   => 2 #int
//	//    +Name => "Bob" #string
//	//  }
//	//  2 => #main.User {
//	//    +ID   => 3 #int
//	//    +Name => "Carol" #string
//	//  }
//	//  3 => #main.User {
//	//    +ID   => 4 #int
//	//    +Name => "Dave" #string
//	//  }
//	// ]
func (c *Collection[T]) Append(values ...T) *Collection[T] {
	out := make([]T, 0, len(c.items)+len(values))
	out = append(out, c.items...)
	out = append(out, values...)
	return Attach(out)
}
