package collection

// None returns true if fn returns false for every item in the collection.
// If the collection is empty, None returns true.
//
// Example: integers – none even
//
//	c := collection.New([]int{1, 3, 5})
//	noneEven := c.None(func(v int) bool { return v%2 == 0 })
//	collection.Dump(noneEven)
//	// true #bool
//
// Example: integers – some even
//
//	c2 := collection.New([]int{1, 2, 3})
//	noneEven2 := c2.None(func(v int) bool { return v%2 == 0 })
//	collection.Dump(noneEven2)
//	// false #bool
//
// Example: empty collection
//
//	empty := collection.New([]int{})
//	none := empty.None(func(v int) bool { return v > 0 })
//	collection.Dump(none)
//	// true #bool
func (c *Collection[T]) None(fn func(T) bool) bool {
	for _, v := range c.items {
		if fn(v) {
			return false
		}
	}
	return true
}
