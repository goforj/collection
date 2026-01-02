package collection

// Contains returns true if the collection contains the given value.
// @group Querying
// @behavior readonly
// @fluent false
// @terminal true
//
// Similar to: Any (differs by using equality on a value).
//
// Example: integers
//
//	c := collection.New([]int{1, 2, 3, 4, 5})
//	hasTwo := collection.Contains(c, 2)
//	collection.Dump(hasTwo)
//	// true #bool
//
// Example: strings
//
//	c2 := collection.New([]string{"apple", "banana", "cherry"})
//	hasBanana := collection.Contains(c2, "banana")
//	collection.Dump(hasBanana)
//	// true #bool
func Contains[T comparable](c *Collection[T], value T) bool {
	for _, v := range c.items {
		if v == value {
			return true
		}
	}
	return false
}
