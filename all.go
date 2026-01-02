package collection

// All returns true if fn returns true for every item in the collection.
// If the collection is empty, All returns true (vacuously true).
// @group Querying
// @behavior readonly
// @chainable false
// @terminal true
//
// Example: integers – all even
//
//	c := collection.New([]int{2, 4, 6})
//	allEven := c.All(func(v int) bool { return v%2 == 0 })
//	collection.Dump(allEven)
//	// true #bool
//
// Example: integers – not all even
//
//	c2 := collection.New([]int{2, 3, 4})
//	allEven2 := c2.All(func(v int) bool { return v%2 == 0 })
//	collection.Dump(allEven2)
//	// false #bool
//
// Example: strings – all non-empty
//
//	c3 := collection.New([]string{"a", "b", "c"})
//	allNonEmpty := c3.All(func(s string) bool { return s != "" })
//	collection.Dump(allNonEmpty)
//	// true #bool
//
// Example: empty collection (vacuously true)
//
//	empty := collection.New([]int{})
//	all := empty.All(func(v int) bool { return v > 0 })
//	collection.Dump(all)
//	// true #bool
func (c *Collection[T]) All(fn func(T) bool) bool {
	for _, v := range c.items {
		if !fn(v) {
			return false
		}
	}
	return true
}
