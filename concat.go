package collection

// Concat appends the values from the given slice onto the end of the collection,
// @group Transformation
// @behavior mutable
// @fluent true
//
// Example: strings
//
//	c := collection.New([]string{"John Doe"})
//	concatenated := c.
//		Concat([]string{"Jane Doe"}).
//		Concat([]string{"Johnny Doe"}).
//		Items()
//	collection.Dump(concatenated)
//
//	// #[]string [
//	//  0 => "John Doe" #string
//	//  1 => "Jane Doe" #string
//	//  2 => "Johnny Doe" #string
//	// ]
func (c *Collection[T]) Concat(values []T) *Collection[T] {
	// Avoid nil checks + fast return
	if len(values) == 0 {
		return c
	}

	cur := c.items
	total := len(cur) + len(values)

	// FAST PATH: enough capacity → mutate in place (0 allocs)
	if cap(cur) >= total {
		oldLen := len(cur)
		cur = cur[:total]          // extend slice in place
		copy(cur[oldLen:], values) // mutate backing array
		c.items = cur
		return c
	}

	// SLOW PATH: one exact allocation
	out := make([]T, total)
	copy(out, cur)
	copy(out[len(cur):], values)

	c.items = out
	return c
}
