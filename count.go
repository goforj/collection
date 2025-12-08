package collection

// Count returns the total number of items in the collection.
// Example:
//   c := collection.New([]int{1, 2, 3, 4})
//   count := c.Count() // 4
func (c *Collection[T]) Count() int {
	return len(c.items)
}
