package collection

// TakeLast returns a new collection containing the last n items.
// If n is less than or equal to zero, TakeLast returns an empty collection.
// If n is greater than or equal to the collection length, TakeLast returns
// the full collection.
// @chainable true
//
// This operation performs no element allocations; it re-slices the
// underlying slice.
// @group Slicing
// @behavior immutable
// Example: integers
//
//	c := collection.New([]int{1, 2, 3, 4, 5})
//	out := c.TakeLast(2)
//	collection.Dump(out.Items())
//	// #[]int [
//	//   0 => 4 #int
//	//   1 => 5 #int
//	// ]
//
// Example: take none
//
//	out2 := c.TakeLast(0)
//	collection.Dump(out2.Items())
//	// #[]int []
//
// Example: take all
//
//	out3 := c.TakeLast(10)
//	collection.Dump(out3.Items())
//	// #[]int [
//	//   0 => 1 #int
//	//   1 => 2 #int
//	//   2 => 3 #int
//	//   3 => 4 #int
//	//   4 => 5 #int
//	// ]
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
//	out4 := users.TakeLast(1)
//	collection.Dump(out4.Items())
//	// #[]collection.User [
//	//   0 => {ID:3} #collection.User
//	// ]
func (c *Collection[T]) TakeLast(n int) *Collection[T] {
	items := c.items
	l := len(items)

	if n <= 0 {
		return &Collection[T]{items: items[:0]}
	}

	if n >= l {
		return &Collection[T]{items: items}
	}

	return &Collection[T]{items: items[l-n:]}
}
