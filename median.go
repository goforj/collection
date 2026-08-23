package collection

import "sort"

// Median returns the statistical median of a numeric slice as float64.
// It returns (0, false) if the slice is empty.
// @group Aggregation
// @behavior readonly
// @chainable false
// @terminal true
//
// Odd count  → middle value
// Even count → average of the two middle values
//
// Example: integers - odd number of items
//
//	values := []int{3, 1, 2}
//
//	median1, ok1 := collection.Median(values)
//	collection.Dump(median1, ok1)
//	// 2.000000 #float64
//	// true #bool
//
// Example: integers - even number of items
//
//	values2 := []int{10, 2, 4, 6}
//
//	median2, ok2 := collection.Median(values2)
//	collection.Dump(median2, ok2)
//	// 5.000000 #float64
//	// true #bool
//
// Example: floats
//
//	values3 := []float64{1.1, 9.9, 3.3}
//
//	median3, ok3 := collection.Median(values3)
//	collection.Dump(median3, ok3)
//	// 3.300000 #float64
//	// true #bool
//
// Example: integers - empty numeric slice
//
//	empty := []int{}
//
//	median4, ok4 := collection.Median(empty)
//	collection.Dump(median4, ok4)
//	// 0.000000 #float64
//	// false #bool
func Median[S ~[]T, T Number](s S) (float64, bool) {
	n := len(s)
	if n == 0 {
		return 0, false
	}

	// Make a copy so sorting does not mutate the original collection
	cp := make([]T, n)
	copy(cp, s)

	sort.Slice(cp, func(i, j int) bool { return cp[i] < cp[j] })

	mid := n / 2

	// Odd
	if n%2 == 1 {
		return float64(cp[mid]), true
	}

	// Even
	a := float64(cp[mid-1])
	b := float64(cp[mid])
	return (a + b) / 2, true
}
