package collection

// Any returns true if at least one item satisfies fn.
// Example: integers
//	c := collection.New([]int{1, 2, 3, 4})
//	has := c.Any(func(v int) bool { return v%2 == 0 }) // true
//	collection.Dump(has)
//	// true #bool
func (c *Collection[T]) Any(fn func(T) bool) bool {
	for _, v := range c.items {
		if fn(v) {
			return true
		}
	}
	return false
}
