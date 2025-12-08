package collection

import "sort"

// Median returns the statistical median of the numeric collection as float64.
// Returns (0, false) if the collection is empty.
//
// Odd count  → middle value
// Even count → average of the two middle values
//
// Example:
//   c := collection.NewNumeric([]int{3, 1, 2})
//   median, ok := c.Median()   // → 2, true
func (c *NumericCollection[T]) Median() (float64, bool) {
	n := len(c.items)
	if n == 0 {
		return 0, false
	}

	// Make a copy so sorting does not mutate the original collection
	cp := make([]T, n)
	copy(cp, c.items)

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
