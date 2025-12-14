package collection

// FirstWhere returns the first item in the collection for which the provided
// predicate function returns true. If no items match, ok=false is returned
// along with the zero value of T.
// @group Querying
// @behavior readonly
// @chainable false
//
// This method is equivalent to Laravel's collection->first(fn) and mirrors
// the behavior found in functional collections in other languages.
//
// Example: integers
//
//	nums := collection.New([]int{1, 2, 3, 4, 5})
//	v, ok := nums.FirstWhere(func(n int) bool {
//		return n%2 == 0
//	})
//	collection.Dump(v, ok)
//	// 2 #int
//	// true #bool
//
//	v, ok = nums.FirstWhere(func(n int) bool {
//		return n > 10
//	})
//	collection.Dump(v, ok)
//	// 0 #int
//	// false #bool
func (c *Collection[T]) FirstWhere(fn func(T) bool) (value T, ok bool) {
	for _, v := range c.items {
		if fn(v) {
			return v, true
		}
	}
	var zero T
	return zero, false
}
