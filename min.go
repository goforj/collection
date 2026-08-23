package collection

// Min returns the smallest item in a numeric slice.
// The second return value is false if the slice is empty.
// @group Aggregation
// @behavior readonly
// @chainable false
// @terminal true
//
// Example: integers
//
//	values := []int{3, 1, 2}
//	min, ok := collection.Min(values)
//	collection.Dump(min, ok)
//	// 1 #int
//	// true #bool
//
// Example: floats
//
//	values2 := []float64{2.5, 9.1, 1.2}
//	min2, ok2 := collection.Min(values2)
//	collection.Dump(min2, ok2)
//	// 1.200000 #float64
//	// true #bool
//
// Example: integers - empty collection
//
//	empty := []int{}
//	min3, ok3 := collection.Min(empty)
//	collection.Dump(min3, ok3)
//	// 0 #int
//	// false #bool
func Min[S ~[]T, T Number](s S) (T, bool) {
	var zero T

	if len(s) == 0 {
		return zero, false
	}

	val := s[0]
	for _, v := range s[1:] {
		if v < val {
			val = v
		}
	}

	return val, true
}
