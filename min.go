package collection

// Min returns the smallest numeric item in the collection.
// The second return value is false if the collection is empty.
//
// Example:
//   c := collection.NewNumeric([]int{3, 1, 2})
//   min, ok := c.Min()  // → 1, true
func (c *NumericCollection[T]) Min() (T, bool) {
	var zero T

	if len(c.items) == 0 {
		return zero, false
	}

	val := c.items[0]
	for _, v := range c.items[1:] {
		if v < val {
			val = v
		}
	}

	return val, true
}
