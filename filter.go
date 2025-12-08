package collection

// Filter keeps only the elements for which fn returns true.
// This method mutates the collection in place and returns the same instance.
//
// Example:
//  collection.New([]int{1,2,3,4}).
//		Filter(func(v int) bool { return v%2 == 0 }).
// 		Items()
//	// []int{2,4}
func (c *Collection[T]) Filter(fn func(T) bool) *Collection[T] {
	j := 0
	for i := 0; i < len(c.items); i++ {
		if fn(c.items[i]) {
			c.items[j] = c.items[i] // compact in place
			j++
		}
	}

	// Optional but ensures ZERO allocations + no GC retention of removed values.
	// Only needed when T contains pointers.
	var zero T
	for k := j; k < len(c.items); k++ {
		c.items[k] = zero // release references
	}

	c.items = c.items[:j] // shrink to new length
	return c
}
