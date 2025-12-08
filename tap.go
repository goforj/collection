package collection

// Tap invokes fn with the collection pointer for side effects (logging, debugging,
// inspection) and returns the same collection to allow chaining.
//
// Tap does NOT modify the collection itself; it simply exposes the current state
// during a fluent chain.
//
// Example:
//
//   captured := []int{}
//   c := New([]int{3,1,2}).
//       Sort(func(a,b int) bool { return a < b }).  // → [1,2,3]
//       Tap(func(col *Collection[int]) {
//           captured = append([]int(nil), col.items...) // snapshot
//       }).
//       Filter(func(v int) bool { return v >= 2 })     // → [2,3]
//
// After Tap, 'captured' contains the sorted state: []int{1,2,3}
// and the chain continues unaffected.
func (c *Collection[T]) Tap(fn func(*Collection[T])) *Collection[T] {
	fn(c)
	return c
}
