package collection

// Min returns the smallest numeric item in the collection.
// The second return value is false if the collection is empty.
//
// Example: integers
//
//	c := collection.NewNumeric([]int{3, 1, 2})
//	min, ok := c.Min()
//	collection.Dump(min, ok)
//	// 1 #int
//	// true #bool
//
// Example: floats
//
//	c2 := collection.NewNumeric([]float64{2.5, 9.1, 1.2})
//	min2, ok2 := c2.Min()
//	collection.Dump(min2, ok2)
//	// 1.200000 #float64
//	// true #bool
//
// Example: integers - empty collection
//
//	empty := collection.NewNumeric([]int{})
//	min3, ok3 := empty.Min()
//	collection.Dump(min3, ok3)
//	// 0 #int
//	// false #bool
func (c *NumericCollection[T]) Min() (T, bool) {
	var zero T

	if len(c.items) == 0 {
		return zero, false
	}

	val := c.items[0]
	for _, v := range c.items[1:] {
		if v < val {
			val = v
		}
	}

	return val, true
}
