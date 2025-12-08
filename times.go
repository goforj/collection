package collection

// Times creates a new collection by running fn(n) for i from 1..count.
// This mirrors Laravel's Collection::times(), which is 1-indexed.
//
// Example:
//   c := collection.Times(5, func(i int) int { return i * 2 })
//   // [2,4,6,8,10]
func Times[T any](count int, fn func(int) T) *Collection[T] {
	if count <= 0 {
		return New([]T{})
	}

	out := make([]T, count)
	for i := 1; i <= count; i++ {
		out[i-1] = fn(i)
	}

	return New(out)
}
