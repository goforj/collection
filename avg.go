package collection

// Avg returns the average as float64.
// Even integer averages may be fractional.
func Avg[T Number](c *Collection[T]) float64 {
	items := c.Items()
	if len(items) == 0 {
		return 0
	}

	var sum float64
	for _, v := range items {
		sum += float64(v)
	}
	return sum / float64(len(items))
}
