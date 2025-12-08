package collection

// Filter returns a new collection containing only values for which fn returns true.
//
// This keeps T the same, so it can be a method.
// Example usage:
//
//	c := collection.New([]int{1, 2, 3, 4, 5})
//	filtered := c.Filter(func(v int) bool { return v%2 == 0 }) // keeps even numbers
//
// // result: [2, 4]
func (c Collection[T]) Filter(fn func(T) bool) Collection[T] {
	out := make([]T, 0, len(c.items))
	for _, v := range c.items {
		if fn(v) {
			out = append(out, v)
		}
	}
	return Collection[T]{items: out}
}
