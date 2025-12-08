package collection

// LastWhere returns the last element in the collection that satisfies the predicate fn.
// If fn is nil, LastWhere returns the final element in the underlying slice.
// If the collection is empty or no element matches, ok will be false.
//
// Example: LastWhere with predicate
//
//    c := collection.New([]int{1, 2, 3, 4})
//    v, ok := c.LastWhere(func(v int, i int) bool {
//        return v < 3
//    })
//    // v == 2, ok == true
//
// Example: LastWhere without predicate
//
//    c := collection.New([]int{1, 2, 3, 4})
//    v, ok := c.LastWhere(nil)
//    // v == 4, ok == true
//
// Example: Empty collection
//
//    c := collection.New([]int{})
//    v, ok := c.LastWhere(nil)
//    // v == 0, ok == false
//
func (c Collection[T]) LastWhere(fn func(T, int) bool) (value T, ok bool) {
	// No elements?
	if len(c.items) == 0 {
		return value, false
	}

	// If no predicate, return the last item
	if fn == nil {
		return c.items[len(c.items)-1], true
	}

	// With predicate: search backwards
	for i := len(c.items) - 1; i >= 0; i-- {
		if fn(c.items[i], i) {
			return c.items[i], true
		}
	}

	// Nothing matched
	return value, false
}
