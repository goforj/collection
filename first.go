package collection

// First returns the first element in the collection.
// If the collection is empty, ok will be false.
//
// Example:
//
//    c := New([]int{1, 2, 3, 4})
//    v, ok := c.First()
//    // v == 1, ok == true
//
// Example (empty):
//
//    c := New([]int{})
//    v, ok := c.First()
//    // v == 0, ok == false
//
func (c Collection[T]) First() (value T, ok bool) {
	if len(c.items) == 0 {
		return value, false
	}
	return c.items[0], true
}
