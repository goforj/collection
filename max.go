package collection

// Max returns the largest item in a numeric slice.
// The second return value is false if the slice is empty.
// @group Aggregation
// @behavior readonly
// @chainable false
// @terminal true
//
// Example: integers
//
//	values := []int{3, 1, 2}
//
//	max1, ok1 := collection.Max(values)
//	collection.Dump(max1, ok1)
//	// 3 #int
//	// true #bool
//
// Example: floats
//
//	values2 := []float64{1.5, 9.2, 4.4}
//
//	max2, ok2 := collection.Max(values2)
//	collection.Dump(max2, ok2)
//	// 9.200000 #float64
//	// true #bool
//
// Example: empty numeric slice
//
//	empty := []int{}
//
//	max3, ok3 := collection.Max(empty)
//	collection.Dump(max3, ok3)
//	// 0 #int
//	// false #bool
func Max[S ~[]T, T Number](s S) (T, bool) {
	var zero T

	if len(s) == 0 {
		return zero, false
	}

	val := s[0]
	for _, v := range s[1:] {
		if v > val {
			val = v
		}
	}

	return val, true
}
