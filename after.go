package collection

// After returns all items after the first element for which pred returns true.
// If no element matches, an empty collection is returned.
// @group Ordering
// @behavior immutable
// @fluent true
//
// NOTE: returns a view (shares backing array). Use Clone() to detach.
//
// Example: integers
//
//	c := collection.New([]int{1, 2, 3, 4, 5})
//	c.After(func(v int) bool { return v == 3 }).Dump()
//	// #[]int [
//	//  0 => 4 #int
//	//  1 => 5 #int
//	// ]
func (c *Collection[T]) After(pred func(T) bool) *Collection[T] {
	idx := -1
	for i, v := range c.items {
		if pred(v) {
			idx = i
			break
		}
	}

	// If no match found → empty collection
	if idx == -1 || idx+1 >= len(c.items) {
		return New(c.items[:0])
	}

	return New(c.items[idx+1:])
}
