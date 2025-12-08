package collection

// Sum returns the sum of all numeric items in the NumericCollection.
// If the collection is empty, Sum returns the zero value of T.
//
// Example:
//   c := collection.NewNumeric([]int{1, 2, 3})
//   total := c.Sum()
//   // total == 6
//
// Example (float):
//   c := collection.NewNumeric([]float64{1.5, 2.5})
//   total := c.Sum()
//   // total == 4.0
func (c *NumericCollection[T]) Sum() T {
	var sum T
	for _, v := range c.items {
		sum += v
	}
	return sum
}
