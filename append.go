package collection

// Append returns a new collection with the given values appended.
// Example:
//  c := collection.New([]int{1, 2})
//  newC := c.Append(3, 4) // Collection with items [1, 2, 3, 4]
//  // newC.Items() == []int{1, 2, 3, 4}
func (c Collection[T]) Append(values ...T) Collection[T] {
	out := make([]T, 0, len(c.items)+len(values))
	out = append(out, c.items...)
	out = append(out, values...)
	return Collection[T]{items: out}
}

// Push returns a new collection with the given values appended.
// Example:
//  c := collection.New([]int{1, 2})
//  newC := c.Push(3, 4) // Collection with items [1, 2, 3, 4]
//  // newC.Items() == []int{1, 2, 3, 4}
func (c Collection[T]) Push(values ...T) Collection[T] {
	return c.Append(values...)
}
