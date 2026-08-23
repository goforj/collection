package collection

// Sum returns the sum of all items in a numeric slice.
// If the slice is empty, Sum returns the zero value of T.
// @group Aggregation
// @behavior readonly
// @chainable false
// @terminal true
//
// Example: integers
//
//	values := []int{1, 2, 3}
//	total := collection.Sum(values)
//	collection.Dump(total)
//	// 6 #int
//
// Example: floats
//
//	values2 := []float64{1.5, 2.5}
//	total2 := collection.Sum(values2)
//	collection.Dump(total2)
//	// 4.000000 #float64
//
// Example: integers - empty collection
//
//	empty := []int{}
//	total3 := collection.Sum(empty)
//	collection.Dump(total3)
//	// 0 #int
func Sum[S ~[]T, T Number](s S) T {
	var sum T
	for _, v := range s {
		sum += v
	}
	return sum
}
