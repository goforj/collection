package collection

// Last returns the last element in the collection.
// If the collection is empty, ok will be false.
// @group Querying
// @behavior readonly
//
// Example: integers
//
//	c := collection.New([]int{10, 20, 30})
//
//	v, ok := c.Last()
//	collection.Dump(v, ok)
//	// 30   #int
//	// true #bool
//
// Example: strings
//
//	c2 := collection.New([]string{"alpha", "beta", "gamma"})
//
//	v2, ok2 := c2.Last()
//	collection.Dump(v2, ok2)
//	// "gamma" #string
//	// true    #bool
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
//		{ID: 3, Name: "Charlie"},
//	})
//
//	u, ok3 := users.Last()
//	collection.Dump(u, ok3)
//	// #main.User {
//	//   +ID   => 3         #int
//	//   +Name => "Charlie" #string
//	// }
//	// true #bool
//
// Example: empty collection
//
//	c3 := collection.New([]int{})
//
//	v3, ok4 := c3.Last()
//	collection.Dump(v3, ok4)
//	// 0     #int
//	// false #bool
func (c *Collection[T]) Last() (value T, ok bool) {
	if len(c.items) == 0 {
		return value, false
	}
	return c.items[len(c.items)-1], true
}
