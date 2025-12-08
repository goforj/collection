package collection

// Any returns true if at least one item satisfies fn.
// Example:
// c := collection.New([]int{1, 2, 3, 4})
// hasEven := c.Any(func(v int) bool { return v%2 == 0 }) // true
// // hasEven is true
func (c *Collection[T]) Any(fn func(T) bool) bool {
	for _, v := range c.items {
		if fn(v) {
			return true
		}
	}
	return false
}
