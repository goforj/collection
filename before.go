package collection

// Before returns all items before the first element for which pred returns true.
// If no element matches, the entire collection is returned.
func (c *Collection[T]) Before(pred func(T) bool) *Collection[T] {
	idx := len(c.items)
	for i, v := range c.items {
		if pred(v) {
			idx = i
			break
		}
	}

	out := make([]T, idx)
	copy(out, c.items[:idx])
	return &Collection[T]{items: out}
}
