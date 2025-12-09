package collection

// Tap invokes fn with the collection pointer for side effects (logging, debugging,
// inspection) and returns the same collection to allow chaining.
//
// Tap does NOT modify the collection itself; it simply exposes the current state
// during a fluent chain.
//
// Example:
//	// capture intermediate state during a chain
//	captured1 := []int{}
//	c1 := collection.New([]int{3, 1, 2}).
//		Sort(func(a, b int) bool { return a < b }). // → [1, 2, 3]
//		Tap(func(col *collection.Collection[int]) {
//			captured1 = append([]int(nil), col.Items()...) // snapshot copy
//		}).
//		Filter(func(v int) bool { return v >= 2 })       // → [2, 3]
//
//	// Use BOTH variables so nothing is "declared and not used"
//	collection.Dump(c1.Items())
//	collection.Dump(captured1)
//	// c1 → #[]int [2,3]
//	// captured1 → #[]int [1,2,3]
//
// Example:
//	// tap for debugging without changing flow
//	c2 := collection.New([]int{10, 20, 30}).
//		Tap(func(col *collection.Collection[int]) {
//			collection.Dump(col.Items())
//		}).
//		Filter(func(v int) bool { return v > 10 })
//
//	collection.Dump(c2.Items()) // ensures c2 is used
//
// Example:
//	// Tap with struct collection
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
//	users2 := users.Tap(func(col *collection.Collection[User]) {
//		collection.Dump(col.Items())
//	})
//
//	collection.Dump(users2.Items()) // ensures users2 is used
func (c *Collection[T]) Tap(fn func(*Collection[T])) *Collection[T] {
	fn(c)
	return c
}
