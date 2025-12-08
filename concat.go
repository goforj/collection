package collection

/*
Concat appends the values from the given slice onto the end of the collection,
returning a new collection. The original collection is never modified.

This mirrors Laravel's concat(): the appended values are numerically reindexed
in the resulting collection, regardless of their original keys or positions.

Example:
    c := collection.New([]string{"John Doe"})

    concatenated := c.
        Concat([]string{"Jane Doe"}).
        Concat([]string{"Johnny Doe"}).
		Items()

    // ["John Doe", "Jane Doe", "Johnny Doe"]

Notes:
  • Concat never mutates the original collection.
  • Keys/indices from the appended slice are ignored; values are simply appended.
  • To concatenate another Collection[T], use:
        c.Concat(other.Items())
*/
func (c *Collection[T]) Concat(values []T) *Collection[T] {
	total := len(c.items) + len(values)

	// CASE 1: Enough capacity → NO allocation
	if cap(c.items) >= total {
		// Extend length, then copy in-place.
		oldLen := len(c.items)
		c.items = c.items[:total]
		copy(c.items[oldLen:], values)
		return c
	}

	// CASE 2: Need to grow → ONE allocation, exact final size
	out := make([]T, total)
	copy(out, c.items)
	copy(out[len(c.items):], values)
	c.items = out
	return c
}
