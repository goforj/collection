package collection

// Max returns the largest numeric item.
// Second return is false if empty.
func Max[T Number](c *Collection[T]) (T, bool) {
	items := c.Items()
	var zero T

	if len(items) == 0 {
		return zero, false
	}

	val := items[0]
	for _, v := range items[1:] {
		if v > val {
			val = v
		}
	}
	return val, true
}
