package collection

// TakeUntilFn returns items until the predicate function returns true.
// The matching item is NOT included.
// Example:
//
//	c := collection.New([]int{1, 2, 3, 4})
//	out := c.TakeUntilFn(func(v int) bool { return v >= 3 }) // [1, 2]
// // result is [1, 2]
func (c Collection[T]) TakeUntilFn(pred func(T) bool) Collection[T] {
	out := make([]T, 0, len(c.items))

	for _, v := range c.items {
		if pred(v) {
			break
		}
		out = append(out, v)
	}

	return New(out)
}

// TakeUntil returns items until the first element equals `value`.
// The matching item is NOT included.
//
// Uses == comparison, so T must be comparable.
func TakeUntil[T comparable](c Collection[T], value T) Collection[T] {
	out := make([]T, 0, len(c.items))

	for _, v := range c.items {
		if v == value {
			break
		}
		out = append(out, v)
	}

	return New(out)
}
