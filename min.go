package collection

// Min returns the smallest numeric item.
// Second return is false if empty.
// Example:
//   c := collection.New([]int{3,1,2})
//   min, ok := Min(c) → 1, true
func Min[T Number](c *Collection[T]) (T, bool) {
	items := c.Items()
	var zero T

	if len(items) == 0 {
		return zero, false
	}

	val := items[0]
	for _, v := range items[1:] {
		if v < val {
			val = v
		}
	}
	return val, true
}
