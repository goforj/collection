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
//	collection.Dump(collection.Sum([]int{1, 2, 3}))
//	// 6 #int
//
// Example: floats
//
//	collection.Dump(collection.Sum([]float64{1.5, 2.5}))
//	// 4.000000 #float64
//
// Example: integers - empty collection
//
//	collection.Dump(collection.Sum([]int{}))
//	// 0 #int
func Sum[S ~[]T, T Number](s S) T {
	var sum T
	for _, v := range s {
		sum += v
	}
	return sum
}
