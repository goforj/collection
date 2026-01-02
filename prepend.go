package collection

// Prepend adds the given values to the beginning of the collection.
// @group Transformation
// @behavior mutable
// @chainable true
// @terminal false
//
// This method mutates the collection in place and returns the same instance.
// It allocates a new backing slice to insert the values at the front.
//
// Example: integers
//
//	c := collection.New([]int{3, 4})
//	c.Prepend(1, 2)
//	collection.Dump(c.Items())
//	// #[]int [
//	//   0 => 1 #int
//	//   1 => 2 #int
//	//   2 => 3 #int
//	//   3 => 4 #int
//	// ]
//
// Example: strings
//
//	letters := collection.New([]string{"c", "d"})
//	letters.Prepend("a", "b")
//	collection.Dump(letters.Items())
//	// #[]string [
//	//   0 => "a" #string
//	//   1 => "b" #string
//	//   2 => "c" #string
//	//   3 => "d" #string
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
//		{ID: 2, Name: "Bob"},
//	})
//
//	users.Prepend(User{ID: 1, Name: "Alice"})
//	collection.Dump(users.Items())
//	// #[]main.User [
//	//   0 => #main.User {
//	//     +ID   => 1 #int
//	//     +Name => "Alice" #string
//	//   }
//	//   1 => #main.User {
//	//     +ID   => 2 #int
//	//     +Name => "Bob" #string
//	//   }
//	// ]
//
// Example: integers - Prepending into an empty collection
//
//	empty := collection.New([]int{})
//	empty.Prepend(9, 8)
//	collection.Dump(empty.Items())
//	// #[]int [
//	//   0 => 9 #int
//	//   1 => 8 #int
//	// ]
//
// Example: integers - Prepending no values → no change
//
//	c2 := collection.New([]int{1, 2})
//	c2.Prepend()
//	collection.Dump(c2.Items())
//	// #[]int [
//	//   0 => 1 #int
//	//   1 => 2 #int
//	// ]
func (c *Collection[T]) Prepend(values ...T) *Collection[T] {
	out := make([]T, 0, len(c.items)+len(values))
	out = append(out, values...)
	out = append(out, c.items...)
	c.items = out
	return c
}
