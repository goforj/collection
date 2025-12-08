package collection

// Pop returns the last item and a new collection with that item removed.
// The original collection remains unchanged.
//
// If the collection is empty, the zero value of T is returned along with
// an empty collection.
func (c *Collection[T]) Pop() (T, *Collection[T]) {
	n := len(c.items)

	if n == 0 {
		var zero T
		return zero, New([]T{})
	}

	item := c.items[n-1]
	rest := c.items[:n-1]

	return item, New(rest)
}

// PopN removes and returns the last n items as a new collection,
// and returns a second collection containing the remaining items.
func (c *Collection[T]) PopN(n int) (*Collection[T], *Collection[T]) {
	if n <= 0 || len(c.items) == 0 {
		return New([]T{}), c
	}

	total := len(c.items)

	if n >= total {
		return New(reverseCopy(c.items)), New([]T{})
	}

	remain := c.items[:total-n]
	popped := c.items[total-n:]

	return New(reverseCopy(popped)), New(remain)
}

func reverseCopy[T any](src []T) []T {
	out := make([]T, len(src))
	for i := range src {
		out[i] = src[len(src)-1-i]
	}
	return out
}
