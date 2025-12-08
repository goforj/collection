package collection

// Transform applies fn to every item *in place* and replaces the values
// with the returned values. This matches Laravel's transform(), which mutates
// the collection instead of returning a new one.
// Example:
//   c := collection.New([]int{1,2,3})
//   c.Transform(func(v int) int { return v * 2 })
//   // c is now [2,4,6]
func (c *Collection[T]) Transform(fn func(T) T) {
	for i, v := range c.items {
		c.items[i] = fn(v)
	}
}
