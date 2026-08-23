package collection

// Concat appends values to c, using c's spare capacity when it is available.
// @group Transformation
// @behavior mutable
// @chainable true
// @terminal false
//
// Callers must capture the returned Slice because a value receiver cannot extend
// c's slice header. When c has enough capacity, Concat reuses its backing array;
// otherwise it allocates exactly enough storage for an independent result.
// Passing no values returns c unchanged.
//
// Example: strings
//
//	c := collection.New([]string{"John Doe"})
//	concatenated := c.
//		Concat([]string{"Jane Doe"}).
//		Concat([]string{"Johnny Doe"})
//	collection.Dump(concatenated)
//
//	// #[]string [
//	//  0 => "John Doe" #string
//	//  1 => "Jane Doe" #string
//	//  2 => "Johnny Doe" #string
//	// ]
//
// Example: spare capacity
//
//	backing := make([]int, 2, 4)
//	copy(backing, []int{1, 2})
//	values := collection.New(backing)
//	values = values.Concat([]int{3, 4})
//	fmt.Println(values)
//	// [1 2 3 4]
func (c Slice[T]) Concat(values []T) Slice[T] {
	if len(values) == 0 {
		return c
	}

	total := len(c) + len(values)
	if cap(c) >= total {
		oldLen := len(c)
		out := c[:total]
		copy(out[oldLen:], values)
		return out
	}

	out := make(Slice[T], total)
	copy(out, c)
	copy(out[len(c):], values)
	return out
}
