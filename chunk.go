package collection

// Chunk splits the collection into chunks of the given size.
// The final chunk may be smaller if len(items) is not divisible by size.
//
// If size <= 0, nil is returned.
// Example:
//   c := collection.New([]int{1,2,3,4,5})
//   chunks := c.Chunk(2) → [[1,2],[3,4],[5]]
func (c Collection[T]) Chunk(size int) [][]T {
	if size <= 0 {
		return nil
	}

	chunks := make([][]T, 0, (len(c.items)+size-1)/size)
	for i := 0; i < len(c.items); i += size {
		end := i + size
		if end > len(c.items) {
			end = len(c.items)
		}
		chunk := make([]T, end-i)
		copy(chunk, c.items[i:end])
		chunks = append(chunks, chunk)
	}
	return chunks
}
