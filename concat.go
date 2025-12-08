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
        Concat([]string{"Johnny Doe"})

    // concatenated.Items()
    // ["John Doe", "Jane Doe", "Johnny Doe"]

Notes:
  • Concat never mutates the original collection.
  • Keys/indices from the appended slice are ignored; values are simply appended.
  • To concatenate another Collection[T], use:
        c.Concat(other.Items())
*/
func (c Collection[T]) Concat(values []T) Collection[T] {
	out := make([]T, 0, len(c.items)+len(values))
	out = append(out, c.items...)
	out = append(out, values...)
	return Collection[T]{items: out}
}
