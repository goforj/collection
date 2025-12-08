package collection

import "sort"

// Median returns the median as float64.
// Fractional medians handled correctly.
// False if empty.
// Example:
//   c := collection.New([]int{3,1,2})
//   median, ok := Median(c) → 2, true
func Median[T Number](c *Collection[T]) (float64, bool) {
	items := c.Items()
	n := len(items)
	if n == 0 {
		return 0, false
	}

	cp := make([]T, n)
	copy(cp, items)

	sort.Slice(cp, func(i, j int) bool { return cp[i] < cp[j] })

	mid := n / 2
	if n%2 == 1 {
		return float64(cp[mid]), true
	}

	a := float64(cp[mid-1])
	b := float64(cp[mid])
	return (a + b) / 2, true
}
