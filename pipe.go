package collection

// Pipe passes the entire collection into the given function
// and returns the function's result.
//
// This is useful for inline transformations, aggregations,
// or "exiting" a chain with a non-collection value.
//
// Example:
//
//   c := New([]int{1, 2, 3})
//   sum := c.Pipe(func(col Collection[int]) any {
//       return col.Sum()
//   })
//
//   // sum == 6
//
func (c *Collection[T]) Pipe(fn func(*Collection[T]) any) any {
	return fn(c)
}
