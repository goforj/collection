package collection

// Max returns the largest numeric item in the collection.
// The second return value is false if the collection is empty.
//
// Example: integers
//
//	c := collection.NewNumeric([]int{3, 1, 2})
//
//	max1, ok1 := c.Max()
//	collection.Dump(max1, ok1)
//	// 3    #int
//	// true #bool
//
// Example: floats
//
//	c2 := collection.NewNumeric([]float64{1.5, 9.2, 4.4})
//
//	max2, ok2 := c2.Max()
//	collection.Dump(max2, ok2)
//	// 9.200000 #float64
//	// true     #bool
//
// Example: empty numeric collection
//
//	c3 := collection.NewNumeric([]int{})
//
//	max3, ok3 := c3.Max()
//	collection.Dump(max3, ok3)
//	// 0     #int
//	// false #bool
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
