package collection

// Map applies a same-type transformation and returns a new collection.
//
// Use this when you're transforming T -> T (e.g., enrichment, normalization).
// Example usage:
//
//	c := collection.New([]int{1, 2, 3})
//	mapped := c.Map(func(v int) int { return v * 10 }) // [10, 20, 30]
//  // expected := []int{10, 20, 30}
func (c *Collection[T]) Map(fn func(T) T) *Collection[T] {
	out := make([]T, len(c.items))
	for i, v := range c.items {
		out[i] = fn(v)
	}
	return &Collection[T]{items: out}
}
