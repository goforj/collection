package collection

// Mode returns the most frequent value(s).
// If tie, returns all values with max freq in first-seen order.
func Mode[T comparable](c *Collection[T]) []T {
	items := c.Items()
	if len(items) == 0 {
		return nil
	}

	counts := make(map[T]int)
	order := make([]T, 0, len(items))
	maxCount := 0

	for _, v := range items {
		if _, exists := counts[v]; !exists {
			order = append(order, v)
		}
		counts[v]++

		if counts[v] > maxCount {
			maxCount = counts[v]
		}
	}

	result := make([]T, 0)
	for _, v := range order {
		if counts[v] == maxCount {
			result = append(result, v)
		}
	}

	return result
}
