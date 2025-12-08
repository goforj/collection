package collection

// Reduce collapses the collection into a single value of type T.
// The accumulator has the same type as the elements.
//
// Example:
//
//	sum := New([]int{1, 2, 3}).Reduce(0, func(acc, n int) int {
//		return acc + n
//	})
//	// 6
//
//	joined := New([]string{"a", "b", "c"}).Reduce("", func(acc, s string) string {
//		return acc + s
//	})
//	// "abc"
//
func (c *Collection[T]) Reduce(initial T, fn func(T, T) T) T {
	acc := initial
	for _, v := range c.items {
		acc = fn(acc, v)
	}
	return acc
}
