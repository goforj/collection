package collection

// Contains returns true if any item satisfies the predicate.
// @group Querying
// @behavior readonly
// Example: integers
//
//	c := collection.New([]int{1, 2, 3, 4, 5})
//	hasEven := c.Contains(func(v int) bool {
//		return v%2 == 0
//	})
//	collection.Dump(hasEven)
//	// true #bool
//
// Example: strings
//
//	c2 := collection.New([]string{"apple", "banana", "cherry"})
//	hasBanana := c2.Contains(func(v string) bool {
//		return v == "banana"
//	})
//	collection.Dump(hasBanana)
//	// true #bool
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
//	hasBob := users.Contains(func(u User) bool {
//		return u.Name == "Bob"
//	})
//	collection.Dump(hasBob)
//	// true #bool
func (c *Collection[T]) Contains(pred func(T) bool) bool {
	for _, v := range c.items {
		if pred(v) {
			return true
		}
	}
	return false
}
