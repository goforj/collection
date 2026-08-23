package collection

// Avg returns the average of the numeric slice values as a float64.
// If the slice is empty, Avg returns 0.
// @group Aggregation
// @behavior readonly
// @chainable false
// @terminal true
//
// Example: integers
//
//	collection.Dump(collection.Avg([]int{2, 4, 6}))
//	// 4.000000 #float64
//
// Example: float
//
//	collection.Dump(collection.Avg([]float64{1.5, 2.5, 3.0}))
//	// 2.333333 #float64
func Avg[S ~[]T, T Number](s S) float64 {
	if len(s) == 0 {
		return 0
	}

	var sum float64
	for _, v := range s {
		sum += float64(v)
	}

	return sum / float64(len(s))
}
