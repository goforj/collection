package collection

// Retain keeps items for which fn returns true in c's existing backing array.
// @group Slicing
// @behavior mutable
// @chainable true
// @terminal false
//
// Retain returns a capacity-capped, shortened slice header, so callers should
// retain its result when subsequent operations must observe the new length.
//
// Example: keep even integers without allocating another backing array
//
//	values := collection.New([]int{1, 2, 3, 4})
//	evens := values.Retain(func(value int) bool { return value%2 == 0 })
//	collection.Dump(evens)
//	// #[]int [
//	//   0 => 2 #int
//	//   1 => 4 #int
//	// ]
//	fmt.Println(values)
//	// [2 4 0 0]
func (c Slice[T]) Retain(fn func(T) bool) Slice[T] {
	retained := 0
	for _, item := range c {
		if fn(item) {
			c[retained] = item
			retained++
		}
	}
	clear(c[retained:])
	return c[:retained:retained]
}
