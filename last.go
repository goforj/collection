package collection

// Last returns the last element in the collection.
// If the collection is empty, ok will be false.
//
// Example:
//
//    c := collection.New([]int{1, 2, 3, 4})
//    v, ok := c.Last()
//    // v == 4, ok == true
//
// Example (empty):
//
//    c := collection.New([]int{})
//    v, ok := c.Last()
//    // v == 0, ok == false
//
func (c *Collection[T]) Last() (value T, ok bool) {
	if len(c.items) == 0 {
		return value, false
	}
	return c.items[len(c.items)-1], true
}
