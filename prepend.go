package collection

// Prepend returns a new collection with the given values added
// to the *beginning* of the collection.
// @group Transformation
// @behavior mutable
// @fluent true
//
// The original collection is not modified.
//
// Example: integers
//
//	c := collection.New([]int{3, 4})
//	newC := c.Prepend(1, 2)
//	collection.Dump(newC.Items())
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
//	out := letters.Prepend("a", "b")
//	collection.Dump(out.Items())
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
//	out2 := users.Prepend(User{ID: 1, Name: "Alice"})
//	collection.Dump(out2.Items())
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
//	out3 := empty.Prepend(9, 8)
//	collection.Dump(out3.Items())
//	// #[]int [
//	//   0 => 9 #int
//	//   1 => 8 #int
//	// ]
//
// Example: integers - Prepending no values → returns a copy of original
//
//	c2 := collection.New([]int{1, 2})
//	out4 := c2.Prepend()
//	collection.Dump(out4.Items())
//	// #[]int [
//	//   0 => 1 #int
//	//   1 => 2 #int
//	// ]
func (c *Collection[T]) Prepend(values ...T) *Collection[T] {
	out := make([]T, 0, len(c.items)+len(values))
	out = append(out, values...)
	out = append(out, c.items...)
	return &Collection[T]{items: out}
}
