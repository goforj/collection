package collection

// Prepend returns an independently backed Slice containing values followed by c.
// @group Transformation
// @behavior immutable
// @chainable true
// @terminal false
//
// It allocates exactly enough storage for the result and leaves c unchanged.
//
// Example: integers
//
//	c := collection.New([]int{3, 4})
//	result := c.Prepend(1, 2)
//	collection.Dump(result)
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
//	result2 := letters.Prepend("a", "b")
//	collection.Dump(result2)
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
//	result3 := users.Prepend(User{ID: 1, Name: "Alice"})
//	collection.Dump(result3)
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
//	result4 := empty.Prepend(9, 8)
//	collection.Dump(result4)
//	// #[]int [
//	//   0 => 9 #int
//	//   1 => 8 #int
//	// ]
//
// Example: integers - Prepending no values → no change
//
//	c2 := collection.New([]int{1, 2})
//	result5 := c2.Prepend()
//	collection.Dump(result5)
//	// #[]int [
//	//   0 => 1 #int
//	//   1 => 2 #int
//	// ]
func (c Slice[T]) Prepend(values ...T) Slice[T] {
	out := make([]T, 0, len(c)+len(values))
	out = append(out, values...)
	out = append(out, c...)
	return out
}
