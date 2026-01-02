package collection

// Avg returns the average of the collection values as a float64.
// If the collection is empty, Avg returns 0.
// @group Aggregation
// @behavior readonly
// @fluent false
// @terminal true
//
// Example: integers
//
//	c := collection.NewNumeric([]int{2, 4, 6})
//	collection.Dump(c.Avg())
//	// 4.000000 #float64
//
// Example: float
//
//	c2 := collection.NewNumeric([]float64{1.5, 2.5, 3.0})
//	collection.Dump(c2.Avg())
//	// 2.333333 #float64
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
