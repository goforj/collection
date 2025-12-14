package collection

// IndexWhere returns the index of the first item in the collection
// for which the provided predicate function returns true.
// If no item matches, it returns (0, false).
// @group Querying
// @behavior readonly
// @fluent true
//
// This operation performs no allocations and short-circuits
// on the first match.
//
// Example: integers
//
//	c := collection.New([]int{10, 20, 30, 40})
//	idx, ok := c.IndexWhere(func(v int) bool { return v == 30 })
//	collection.Dump(idx, ok)
//	// 2 true
//
// Example: not found
//
//	idx2, ok2 := c.IndexWhere(func(v int) bool { return v == 99 })
//	collection.Dump(idx2, ok2)
//	// 0 false
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
//		{ID: 3, Name: "Carol"},
//	})
//
//	idx3, ok3 := users.IndexWhere(func(u User) bool {
//		return u.Name == "Bob"
//	})
//
//	collection.Dump(idx3, ok3)
//	// 1 true
func (c *Collection[T]) IndexWhere(fn func(T) bool) (int, bool) {
	for i, v := range c.items {
		if fn(v) {
			return i, true
		}
	}
	return 0, false
}
