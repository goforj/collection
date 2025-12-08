package collection

// Avg returns the average of the collection values as a float64.
// If the collection is empty, Avg returns 0.
//
// Example:
//	c := collection.New([]int{2, 4, 6})
//	avg := c.Avg()
//	// avg == 4
//
// Example (float collection):
//	c := collection.New([]float64{1.5, 2.5, 3.0})
//	avg := c.Avg()
//	// avg == 2.3333333
func (c *NumericCollection[T]) Avg() float64 {
	if len(c.items) == 0 {
		return 0
	}

	var sum float64
	for _, v := range c.items {
		sum += float64(v)
	}

	return sum / float64(len(c.items))
}
