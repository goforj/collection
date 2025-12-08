package collection

// FirstWhere returns the first item in the collection for which the provided
// predicate function returns true. If no items match, ok=false is returned
// along with the zero value of T.
//
// This method is equivalent to Laravel's collection->first(fn) and mirrors
// the behavior found in functional collections in other languages.
//
// Examples:
//
//   nums := New([]int{1, 2, 3, 4, 5})
//   v, ok := nums.FirstWhere(func(n int) bool {
//       return n%2 == 0
//   })
//   // v = 2, ok = true
//
//   v, ok = nums.FirstWhere(func(n int) bool {
//       return n > 10
//   })
//   // v = 0, ok = false
//
func (c Collection[T]) FirstWhere(fn func(T) bool) (value T, ok bool) {
	for _, v := range c.items {
		if fn(v) {
			return v, true
		}
	}
	var zero T
	return zero, false
}
