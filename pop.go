package collection

// Pop removes and returns the last item in the collection.
// @group Slicing
// @behavior mutable
// @fluent false
//
// If the collection is empty, the zero value of T is returned with ok=false.
//
// Example: integers
//
//	c := collection.New([]int{1, 2, 3})
//	item, ok := c.Pop()
//	collection.Dump(item, ok, c.Items())
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
//	collection.Dump(item2, ok2, c2.Items())
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
//	collection.Dump(item3, ok3, users.Items())
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
//	collection.Dump(item4, ok4, empty.Items())
//	// 0 #int
//	// false #bool
//	// #[]int [
//	// ]
func (c *Collection[T]) Pop() (T, bool) {
	n := len(c.items)

	if n == 0 {
		var zero T
		return zero, false
	}

	item := c.items[n-1]
	c.items = c.items[:n-1]
	return item, true
}

// PopN removes and returns the last n items in original order.
// @group Slicing
// @behavior mutable
// @fluent false
//
// Example: integers – pop 2
//
//	c := collection.New([]int{1, 2, 3, 4})
//	popped := c.PopN(2)
//	collection.Dump(popped, c.Items())
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
//	collection.Dump(popped2, c2.Items())
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
//	collection.Dump(popped3, users.Items())
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
//	collection.Dump(popped4, c3.Items())
//	// <nil>
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
//	collection.Dump(popped5, c4.Items())
//	// #[]string [
//	//   0 => "x" #string
//	//   1 => "y" #string
//	// ]
//	// #[]string [
//	// ]
func (c *Collection[T]) PopN(n int) []T {
	if n <= 0 || len(c.items) == 0 {
		return nil
	}

	total := len(c.items)
	if n > total {
		n = total
	}

	popped := c.items[total-n:]
	c.items = c.items[:total-n]
	return popped
}
