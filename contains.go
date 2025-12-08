package collection

// Contains returns true if any item satisfies the predicate.
// Example:
//
//	c := collection.New([]int{1, 2, 3, 4})
//	hasEven := c.Contains(func(v int) bool { return v%2 == 0 }) // true
// // hasEven is true
func (c *Collection[T]) Contains(pred func(T) bool) bool {
	for _, v := range c.items {
		if pred(v) {
			return true
		}
	}
	return false
}
