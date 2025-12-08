package collection

// FindWhere returns the first item in the collection for which the provided
// predicate function returns true. This method is an alias for FirstWhere(fn)
// and is provided for ergonomic parity with functional libraries and
// languages such as JavaScript, Rust, C#, and Python.
//
// FindWhere improves discoverability for developers who naturally search for a
// "find" helper when retrieving an element that matches a condition.
//
// Examples:
//
//   nums := New([]int{1, 2, 3, 4, 5})
//
//   v, ok := nums.FindWhere(func(n int) bool {
//       return n == 3
//   })
//   // v = 3, ok = true
//
//   v, ok = nums.FindWhere(func(n int) bool {
//       return n > 10
//   })
//   // v = 0, ok = false
//
func (c Collection[T]) FindWhere(fn func(T) bool) (T, bool) {
	return c.FirstWhere(fn)
}
