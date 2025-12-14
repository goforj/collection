package collection

// Mode returns the most frequent numeric value(s) in the collection.
// If multiple values tie for highest frequency, all are returned
// in first-seen order.
// @group Aggregation
//
// Example: integers – single mode
//
//	c := collection.NewNumeric([]int{1, 2, 2, 3})
//	mode := c.Mode()
//	collection.Dump(mode)
//	// #[]int [
//	//   0 => 2 #int
//	// ]
//
// Example: integers – tie for mode
//
//	c2 := collection.NewNumeric([]int{1, 2, 1, 2})
//	mode2 := c2.Mode()
//	collection.Dump(mode2)
//	// #[]int [
//	//   0 => 1 #int
//	//   1 => 2 #int
//	// ]
//
// Example: floats
//
//	c3 := collection.NewNumeric([]float64{1.1, 2.2, 1.1, 3.3})
//	mode3 := c3.Mode()
//	collection.Dump(mode3)
//	// #[]float64 [
//	//   0 => 1.100000 #float64
//	// ]
//
// Example: integers - empty collection
//
//	empty := collection.NewNumeric([]int{})
//	mode4 := empty.Mode()
//	collection.Dump(mode4)
//	// <nil>
func (c *NumericCollection[T]) Mode() []T {
	items := c.items
	n := len(items)
	if n == 0 {
		return nil
	}

	counts := make(map[T]int, n)
	order := make([]T, 0, n)
	maxCount := 0

	for _, v := range items {
		if _, seen := counts[v]; !seen {
			order = append(order, v)
		}
		counts[v]++

		if counts[v] > maxCount {
			maxCount = counts[v]
		}
	}

	result := make([]T, 0, len(order))
	for _, v := range order {
		if counts[v] == maxCount {
			result = append(result, v)
		}
	}

	return result
}
