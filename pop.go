package collection

// Pop removes and returns the last item from an addressable Slice variable.
// @group Slicing
// @behavior mutable
// @chainable false
// @terminal true
//
// If the collection is empty, the zero value of T is returned with ok=false.
// Pop clears the removed element and shortens the receiver's slice header.
//
// Example: integers
//
//	c := collection.New([]int{1, 2, 3})
//	item, ok := c.Pop()
//	collection.Dump(item, ok, c)
//	// 3 #int
//	// true #bool
//	// #[]int [
//	//   0 => 1 #int
//	//   1 => 2 #int
//	// ]
//
// Example: strings
//
//	c2 := collection.New([]string{"a", "b", "c"})
//	item2, ok2 := c2.Pop()
//	collection.Dump(item2, ok2, c2)
//	// "c" #string
//	// true #bool
//	// #[]string [
//	//   0 => "a" #string
//	//   1 => "b" #string
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
//	item3, ok3 := users.Pop()
//	collection.Dump(item3, ok3, users)
//	// #main.User {
//	//   +ID   => 2 #int
//	//   +Name => "Bob" #string
//	// }
//	// true #bool
//	// #[]main.User [
//	//   0 => #main.User {
//	//     +ID   => 1 #int
//	//     +Name => "Alice" #string
//	//   }
//	// ]
//
// Example: empty collection
//
//	empty := collection.New([]int{})
//	item4, ok4 := empty.Pop()
//	collection.Dump(item4, ok4, empty)
//	// 0 #int
//	// false #bool
//	// #[]int [
//	// ]
func (c *Slice[T]) Pop() (T, bool) {
	n := len(*c)

	if n == 0 {
		var zero T
		return zero, false
	}

	item := (*c)[n-1]
	var zero T
	(*c)[n-1] = zero
	*c = (*c)[:n-1]
	return item, true
}

// PopN removes and returns the last n items in original order from an
// addressable Slice variable. It shortens the receiver's slice header.
//
// The returned slice is a view into the receiver's backing array, not a copy.
// Other copied Slice headers keep their original length, and later appending or
// concatenating into the shortened receiver's spare capacity can overwrite the
// values visible through the returned view.
// @group Slicing
// @behavior mutable
// @chainable false
// @terminal true
//
// Example: integers – pop 2
//
//	c := collection.New([]int{1, 2, 3, 4})
//	popped := c.PopN(2)
//	collection.Dump(popped, c)
//	// #[]int [
//	//   0 => 3 #int
//	//   1 => 4 #int
//	// ]
//	// #[]int [
//	//   0 => 1 #int
//	//   1 => 2 #int
//	// ]
//
// Example: strings – pop 1
//
//	c2 := collection.New([]string{"a", "b", "c"})
//	popped2 := c2.PopN(1)
//	collection.Dump(popped2, c2)
//	// #[]string [
//	//   0 => "c" #string
//	// ]
//	// #[]string [
//	//   0 => "a" #string
//	//   1 => "b" #string
//	// ]
//
// Example: structs – pop 2
//
//	type User struct {
//		ID   int
//		Name string
//	}
//
//	users := collection.New([]User{
//		{ID: 1, Name: "Alice"},
//		{ID: 2, Name: "Bob"},
//		{ID: 3, Name: "Carol"},
//	})
//
//	popped3 := users.PopN(2)
//	collection.Dump(popped3, users)
//	// #[]main.User [
//	//   0 => #main.User {
//	//     +ID   => 2 #int
//	//     +Name => "Bob" #string
//	//   }
//	//   1 => #main.User {
//	//     +ID   => 3 #int
//	//     +Name => "Carol" #string
//	//   }
//	// ]
//	// #[]main.User [
//	//   0 => #main.User {
//	//     +ID   => 1 #int
//	//     +Name => "Alice" #string
//	//   }
//	// ]
//
// Example: integers - n <= 0 → returns nil, no change
//
//	c3 := collection.New([]int{1, 2, 3})
//	popped4 := c3.PopN(0)
//	collection.Dump(popped4, c3)
//	// []int(nil)
//	// #[]int [
//	//   0 => 1 #int
//	//   1 => 2 #int
//	//   2 => 3 #int
//	// ]
//
// Example: strings - n exceeds length → all items popped, rest empty
//
//	c4 := collection.New([]string{"x", "y"})
//	popped5 := c4.PopN(10)
//	collection.Dump(popped5, c4)
//	// #[]string [
//	//   0 => "x" #string
//	//   1 => "y" #string
//	// ]
//	// #[]string [
//	// ]
func (c *Slice[T]) PopN(n int) []T {
	if n <= 0 || len(*c) == 0 {
		return nil
	}

	total := len(*c)
	if n > total {
		n = total
	}

	popped := (*c)[total-n:]
	*c = (*c)[:total-n]
	return popped
}
