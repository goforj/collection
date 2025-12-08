package collection

// Take returns a new collection containing the first `n` items when n > 0,
// or the last `|n|` items when n < 0. If n exceeds the collection length,
// the entire slice (or nothing) is returned.
//
// Mirrors Laravel's take() semantics.
//
// Examples:
//   New([]int{0,1,2,3,4,5}).Take(3)  → [0,1,2]
//   New([]int{0,1,2,3,4,5}).Take(-2) → [4,5]
func (c Collection[T]) Take(n int) Collection[T] {
	length := len(c.items)

	// Zero or empty → empty collection
	if n == 0 || length == 0 {
		return New([]T{})
	}

	// n > 0 → take from start
	if n > 0 {
		if n >= length {
			// return whole collection
			out := make([]T, length)
			copy(out, c.items)
			return New(out)
		}
		out := make([]T, n)
		copy(out, c.items[:n])
		return New(out)
	}

	// n < 0 → take from end
	n = -n // absolute count
	if n >= length {
		// return whole collection
		out := make([]T, length)
		copy(out, c.items)
		return New(out)
	}

	start := length - n
	out := make([]T, n)
	copy(out, c.items[start:])

	return New(out)
}
