package collection

// Each runs fn for every item in the collection and returns the same collection,
// so it can be used in chains for side effects (logging, debugging, etc.).
func (c *Collection[T]) Each(fn func(T)) *Collection[T] {
	for _, v := range c.items {
		fn(v)
	}
	return c
}
