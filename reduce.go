package collection

// Reduce collapses the collection into a single value of type T.
// The accumulator has the same type as the elements.
//
// Example:
//
//	sum := collection.New([]int{1, 2, 3}).Reduce(0, func(acc, n int) int {
//		return acc + n
//	})
//	// 6
//
//	joined := collection.New([]string{"a", "b", "c"}).Reduce("", func(acc, s string) string {
//		return acc + s
//	})
//	// "abc"
//
// type Stats struct {
//     Count int
//     Sum   int
// }
//
// c := collection.New([]Stats{
//     {Count: 1, Sum: 10},
//     {Count: 1, Sum: 20},
//     {Count: 1, Sum: 30},
// })
//
// total := c.Reduce(Stats{}, func(acc, s Stats) Stats {
//     acc.Count += s.Count
//     acc.Sum += s.Sum
//     return acc
// })
//
// // Stats{Count: 3, Sum: 60}
func (c *Collection[T]) Reduce(initial T, fn func(T, T) T) T {
	acc := initial
	for _, v := range c.items {
		acc = fn(acc, v)
	}
	return acc
}
