package collection

// Mode returns the most frequent numeric value or values in a slice.
// If multiple values tie for highest frequency, all are returned
// in first-seen order.
// @group Aggregation
// @behavior readonly
// @chainable false
// @terminal true
//
// Example: integers - single mode
//
//	values := []int{1, 2, 2, 3}
//	mode := collection.Mode(values)
//	collection.Dump(mode)
//	// #[]int [
//	//   0 => 2 #int
//	// ]
//
// Example: integers - tie for mode
//
//	values2 := []int{1, 2, 1, 2}
//	mode2 := collection.Mode(values2)
//	collection.Dump(mode2)
//	// #[]int [
//	//   0 => 1 #int
//	//   1 => 2 #int
//	// ]
//
// Example: floats
//
//	values3 := []float64{1.1, 2.2, 1.1, 3.3}
//	mode3 := collection.Mode(values3)
//	collection.Dump(mode3)
//	// #[]float64 [
//	//   0 => 1.100000 #float64
//	// ]
//
// Example: integers - empty collection
//
//	empty := []int{}
//	mode4 := collection.Mode(empty)
//	collection.Dump(mode4)
//	// []int(nil)
func Mode[S ~[]T, T Number](items S) []T {
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
