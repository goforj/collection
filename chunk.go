package collection

// Chunk splits the collection into chunks of the given size.
// The final chunk may be smaller if len(items) is not divisible by size.
//
// If size <= 0, nil is returned.
// Example:
//   collection.New([]int{1,2,3,4,5}).Chunk(2)
//   // [[1,2],[3,4],[5]]
func (c *Collection[T]) Chunk(size int) [][]T {
	if size <= 0 {
		return nil
	}

	n := len(c.items)
	chunks := make([][]T, 0, (n+size-1)/size)

	for i := 0; i < n; i += size {
		end := i + size
		if end > n {
			end = n
		}
		// ZERO ALLOC — slice header only
		chunks = append(chunks, c.items[i:end])
	}

	return chunks
}
