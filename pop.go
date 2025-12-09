package collection

// Pop returns the last item and a new collection with that item removed.
// The original collection remains unchanged.
//
// If the collection is empty, the zero value of T is returned along with
// an empty collection.
//
// Example:
//	// integers
//	c := collection.New([]int{1, 2, 3})
//	item, rest := c.Pop()
//	collection.Dump(item, rest.Items())
//	// 3 #int
//	// #[]int [
//	//   0 => 1 #int
//	//   1 => 2 #int
//	// ]
//
// Example:
//	// strings
//	c2 := collection.New([]string{"a", "b", "c"})
//	item2, rest2 := c2.Pop()
//	collection.Dump(item2, rest2.Items())
//	// "c" #string
//	// #[]string [
//	//   0 => "a" #string
//	//   1 => "b" #string
//	// ]
//
// Example:
//	// structs
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
//	item3, rest3 := users.Pop()
//	collection.Dump(item3, rest3.Items())
//	// #main.User {
//	//   +ID   => 2 #int
//	//   +Name => "Bob" #string
//	// }
//	// #[]main.User [
//	//   0 => #main.User {
//	//     +ID   => 1 #int
//	//     +Name => "Alice" #string
//	//   }
//	// ]
//
// Example:
//	// empty collection
//	empty := collection.New([]int{})
//	item4, rest4 := empty.Pop()
//	collection.Dump(item4, rest4.Items())
//	// 0 #int
//	// #[]int [
//	// ]
func (c *Collection[T]) Pop() (T, *Collection[T]) {
	n := len(c.items)

	if n == 0 {
		var zero T
		return zero, New([]T{})
	}

	item := c.items[n-1]
	rest := c.items[:n-1]

	return item, New(rest)
}

// PopN removes and returns the last n items as a new collection,
// and returns a second collection containing the remaining items.
func (c *Collection[T]) PopN(n int) (*Collection[T], *Collection[T]) {
	if n <= 0 || len(c.items) == 0 {
		return New([]T{}), c
	}

	total := len(c.items)

	if n >= total {
		return New(reverseCopy(c.items)), New([]T{})
	}

	remain := c.items[:total-n]
	popped := c.items[total-n:]

	return New(reverseCopy(popped)), New(remain)
}

func reverseCopy[T any](src []T) []T {
	out := make([]T, len(src))
	for i := range src {
		out[i] = src[len(src)-1-i]
	}
	return out
}
