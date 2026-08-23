package collection

// CountByValue returns the number of occurrences of each distinct item in c.
// @group Aggregation
// @behavior readonly
// @chainable false
// @terminal true
//
// T must be comparable.
//
// Example: strings
//
//	values := []string{"go", "forj", "go"}
//	counts := collection.CountByValue(values)
//	collection.Dump(counts)
//	// #map[string]int {
//	//   forj => 1 #int
//	//   go => 2 #int
//	// }
func CountByValue[S ~[]T, T comparable](c S) map[T]int {
	result := make(map[T]int)
	for _, value := range c {
		result[value]++
	}
	return result
}
