package collection

// Max returns the largest numeric item in the collection.
// The second return value is false if the collection is empty.
//
// Example:
// c := collection.NewNumeric([]int{3, 1, 2})
// max, ok := c.Max()
// // → 3, true
func (c *NumericCollection[T]) Max() (T, bool) {
	var zero T

	if len(c.items) == 0 {
		return zero, false
	}

	val := c.items[0]
	for _, v := range c.items[1:] {
		if v > val {
			val = v
		}
	}

	return val, true
}
