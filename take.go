package collection

// Take returns a new collection containing the first `n` items when n > 0,
// or the last `|n|` items when n < 0.
//
// If n exceeds the collection length, the entire collection is returned.
// If n == 0, an empty collection is returned.
//
// Mirrors Laravel's take() semantics.
//
// Example:
//	// take first 3
//	c1 := collection.New([]int{0, 1, 2, 3, 4, 5})
//	out1 := c1.Take(3)
//	collection.Dump(out1.Items())
//	// #[]int [
//	//	0 => 0 #int
//	//	1 => 1 #int
//	//	2 => 2 #int
//	// ]
//
// Example:
//	// take last 2 (negative n)
//	c2 := collection.New([]int{0, 1, 2, 3, 4, 5})
//	out2 := c2.Take(-2)
//	collection.Dump(out2.Items())
//	// #[]int [
//	//	0 => 4 #int
//	//	1 => 5 #int
//	// ]
//
// Example:
//	// n exceeds length → whole collection
//	c3 := collection.New([]int{10, 20})
//	out3 := c3.Take(10)
//	collection.Dump(out3.Items())
//	// #[]int [
//	//	0 => 10 #int
//	//	1 => 20 #int
//	// ]
//
// Example:
//	// zero → empty
//	c4 := collection.New([]int{1, 2, 3})
//	out4 := c4.Take(0)
//	collection.Dump(out4.Items())
//	// #[]int [
//	// ]
func (c *Collection[T]) Take(n int) *Collection[T] {
	length := len(c.items)

	// Zero or empty → empty collection
	if n == 0 || length == 0 {
		return New([]T{})
	}

	// n > 0 → take from start
	if n > 0 {
		if n >= length {
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
		out := make([]T, length)
		copy(out, c.items)
		return New(out)
	}

	start := length - n
	out := make([]T, n)
	copy(out, c.items[start:])

	return New(out)
}
