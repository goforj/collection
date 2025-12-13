package collection

// LastWhere returns the last element in the collection that satisfies the predicate fn.
// If fn is nil, LastWhere returns the final element in the underlying slice.
// If the collection is empty or no element matches, ok will be false.
//
// Example: integers
//
//	c := collection.New([]int{1, 2, 3, 4})
//
//	v, ok := c.LastWhere(func(v int, i int) bool {
//		return v < 3
//	})
//	collection.Dump(v, ok)
//	// 2    #int
//	// true #bool
//
// Example:
//
//	// integers without predicate (equivalent to Last())
//	c2 := collection.New([]int{10, 20, 30, 40})
//
//	v2, ok2 := c2.LastWhere(nil)
//	collection.Dump(v2, ok2)
//	// 40   #int
//	// true #bool
//
// Example:
//
//	// strings
//	c3 := collection.New([]string{"alpha", "beta", "gamma", "delta"})
//
//	v3, ok3 := c3.LastWhere(func(s string, i int) bool {
//		return strings.HasPrefix(s, "g")
//	})
//	collection.Dump(v3, ok3)
//	// "gamma" #string
//	// true    #bool
//
// Example:
//
//	// structs
//	type User struct {
//		ID   int
//		Name string
//	}
//
//	users := collection.New([]User{
//		{ID: 1, Name: "Alice"},
//		{ID: 2, Name: "Bob"},
//		{ID: 3, Name: "Alex"},
//		{ID: 4, Name: "Brian"},
//	})
//
//	u, ok4 := users.LastWhere(func(u User, i int) bool {
//		return strings.HasPrefix(u.Name, "A")
//	})
//	collection.Dump(u, ok4)
//	// #main.User {
//	//   +ID   => 3        #int
//	//   +Name => "Alex"  #string
//	// }
//	// true #bool
//
// Example:
//
//	// no matching element
//	c4 := collection.New([]int{5, 6, 7})
//
//	v4, ok5 := c4.LastWhere(func(v int, i int) bool {
//		return v > 10
//	})
//	collection.Dump(v4, ok5)
//	// 0     #int
//	// false #bool
//
// Example:
//
//	// empty collection
//	c5 := collection.New([]int{})
//
//	v5, ok6 := c5.LastWhere(nil)
//	collection.Dump(v5, ok6)
//	// 0     #int
//	// false #bool
func (c *Collection[T]) LastWhere(fn func(T, int) bool) (value T, ok bool) {
	if len(c.items) == 0 {
		return value, false
	}
	if fn == nil {
		return c.items[len(c.items)-1], true
	}
	for i := len(c.items) - 1; i >= 0; i-- {
		if fn(c.items[i], i) {
			return c.items[i], true
		}
	}
	return value, false
}
