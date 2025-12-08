package collection

// IsEmpty returns true if the collection has no items.
func (c *Collection[T]) IsEmpty() bool {
	return len(c.items) == 0
}
