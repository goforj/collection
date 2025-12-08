package collection

// Multiply creates `n` copies of all items in the collection
// and returns a new collection.
//
// Example:
//   users := New([]User{{Name: "A"}, {Name: "B"}})
//   out := users.Multiply(3)
//
// Resulting items:
//   [A, B, A, B, A, B]
//
// If n <= 0, the method returns an empty collection.
func (c *Collection[T]) Multiply(n int) *Collection[T] {
	if n <= 0 {
		return New([]T{})
	}

	orig := c.items
	out := make([]T, 0, len(orig)*n)

	for i := 0; i < n; i++ {
		out = append(out, orig...)
	}

	return New(out)
}
