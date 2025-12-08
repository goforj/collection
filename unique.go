package collection

// Unique returns a collection with duplicate items (according to eq) removed,
// preserving the first occurrence of each unique value.
//
// eq should return true if the two values are considered equal.
// Example usage:
//
//	c := collection.New([]int{1, 2, 2, 3, 4, 4, 5})
//	unique := c.Unique(func(a, b int) bool { return a == b })
//
//	// result: [1, 2, 3, 4, 5]
func (c Collection[T]) Unique(eq func(a, b T) bool) Collection[T] {
	out := make([]T, 0, len(c.items))

	for _, v := range c.items {
		found := false
		for _, existing := range out {
			if eq(v, existing) {
				found = true
				break
			}
		}
		if !found {
			out = append(out, v)
		}
	}

	return Collection[T]{items: out}
}
