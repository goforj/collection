package collection

// Prepend returns a new collection with the given values prepended.
// Example:
//  c := collection.New([]int{3, 4})
//  newC := c.Prepend(1, 2) // Collection with items [1, 2, 3, 4]
//  // newC.Items() == []int{1, 2, 3, 4}
func (c Collection[T]) Prepend(values ...T) Collection[T] {
	out := make([]T, 0, len(c.items)+len(values))
	out = append(out, values...)
	out = append(out, c.items...)
	return Collection[T]{items: out}
}
