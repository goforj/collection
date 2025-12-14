package collection

// Skip returns a new collection with the first n items skipped.
// If n is less than or equal to zero, Skip returns the full collection.
// If n is greater than or equal to the collection length, Skip returns
// an empty collection.
// @group Slicing
// @behavior immutable
// @chainable true
//
// This operation performs no element allocations; it re-slices the
// underlying slice.
//
// Example: integers
//
//	c := collection.New([]int{1, 2, 3, 4, 5})
//	out := c.Skip(2)
//	collection.Dump(out.Items())
//	// #[]int [
//	//   0 => 3 #int
//	//   1 => 4 #int
//	//   2 => 5 #int
//	// ]
//
// Example: skip none
//
//	out2 := c.Skip(0)
//	collection.Dump(out2.Items())
//	// #[]int [
//	//   0 => 1 #int
//	//   1 => 2 #int
//	//   2 => 3 #int
//	//   3 => 4 #int
//	//   4 => 5 #int
//	// ]
//
// Example: skip all
//
//	out3 := c.Skip(10)
//	collection.Dump(out3.Items())
//	// #[]int []
//
// Example: structs
//
//	type User struct {
//		ID int
//	}
//
//	users := collection.New([]User{
//		{ID: 1},
//		{ID: 2},
//		{ID: 3},
//	})
//
//	out4 := users.Skip(1)
//	collection.Dump(out4.Items())
//	// []main.User [
//	//  0 => #main.User {
//	//    +ID => 2 #int
//	//  }
//	//  1 => #main.User {
//	//    +ID => 3 #int
//	//  }
//	// ]
func (c *Collection[T]) Skip(n int) *Collection[T] {
	items := c.items
	l := len(items)

	if n <= 0 {
		return &Collection[T]{items: items}
	}

	if n >= l {
		return &Collection[T]{items: items[:0]}
	}

	return &Collection[T]{items: items[n:]}
}
