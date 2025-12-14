package collection

// FindWhere returns the first item in the collection for which the provided
// predicate function returns true. This is an alias for FirstWhere(fn) and
// exists for ergonomic parity with functional languages (JavaScript, Rust,
// C#, Python) where developers expect a “find” helper.
// @group Querying
//
// Example: integers
//
//	nums := collection.New([]int{1, 2, 3, 4, 5})
//
//	v1, ok1 := nums.FindWhere(func(n int) bool {
//		return n == 3
//	})
//	collection.Dump(v1, ok1)
//	// 3    #int
//	// true #bool
//
// Example: no match
//
//	v2, ok2 := nums.FindWhere(func(n int) bool {
//		return n > 10
//	})
//	collection.Dump(v2, ok2)
//	// 0     #int
//	// false #bool
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
//	u, ok3 := users.FindWhere(func(u User) bool {
//		return u.ID == 2
//	})
//	collection.Dump(u, ok3)
//	// #collection.User {
//	//   +ID    => 2   #int
//	//   +Name  => "Bob" #string
//	// }
//	// true #bool
//
// Example: integers - empty collection
//
//	empty := collection.New([]int{})
//
//	v4, ok4 := empty.FindWhere(func(n int) bool { return n == 1 })
//	collection.Dump(v4, ok4)
//	// 0     #int
//	// false #bool
func (c *Collection[T]) FindWhere(fn func(T) bool) (T, bool) {
	return c.FirstWhere(fn)
}
