package collection

// Sum returns the sum of all numeric items.
func Sum[T Number](c Collection[T]) T {
	items := c.Items()
	var sum T
	for _, v := range items {
		sum += v
	}
	return sum
}
