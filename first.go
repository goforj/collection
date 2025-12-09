package collection

// First returns the first element in the collection.
// If the collection is empty, ok will be false.
//
// Example:
//	// integers
//	c := collection.New([]int{10, 20, 30})
//
//	v, ok := c.First()
//	collection.Dump(v, ok)
//	// 10   #int
//	// true #bool
//
// Example:
//	// strings
//	c2 := collection.New([]string{"alpha", "beta", "gamma"})
//
//	v2, ok2 := c2.First()
//	collection.Dump(v2, ok2)
//	// "alpha" #string
//	// true    #bool
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
//	u, ok3 := users.First()
//	collection.Dump(u, ok3)
//	// #main.User {
//	//   +ID   => 1      #int
//	//   +Name => "Alice" #string
//	// }
//	// true #bool
//
// Example:
//	// empty collection
//	c3 := collection.New([]int{})
//	v3, ok4 := c3.First()
//	collection.Dump(v3, ok4)
//	// 0    #int
//	// false #bool
func (c *Collection[T]) First() (value T, ok bool) {
	if len(c.items) == 0 {
		return value, false
	}
	return c.items[0], true
}
