package collection

// Sum returns the sum of all numeric items in the NumericCollection.
// If the collection is empty, Sum returns the zero value of T.
// @group Aggregation
// @behavior readonly
// @chainable false
// @terminal true
//
// Example: integers
//
//	c := collection.NewNumeric([]int{1, 2, 3})
//	total := c.Sum()
//	collection.Dump(total)
//	// 6 #int
//
// Example: floats
//
//	c2 := collection.NewNumeric([]float64{1.5, 2.5})
//	total2 := c2.Sum()
//	collection.Dump(total2)
//	// 4.000000 #float64
//
// Example: integers - empty collection
//
//	c3 := collection.NewNumeric([]int{})
//	total3 := c3.Sum()
//	collection.Dump(total3)
//	// 0 #int
func (c *NumericCollection[T]) Sum() T {
	var sum T
	for _, v := range c.items {
		sum += v
	}
	return sum
}
