package collection

// Concat returns an independent collection containing c followed by values.
// @group Transformation
// @behavior immutable
// @chainable true
// @terminal false
//
// Callers must capture the returned Slice because a value receiver cannot extend
// c's slice header. The returned collection never shares backing storage with c.
//
// Example: strings
//
//	c := collection.New([]string{"John Doe"})
//	concatenated := c.
//		Concat([]string{"Jane Doe"}).
//		Concat([]string{"Johnny Doe"})
//	collection.Dump(concatenated)
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
func (c Slice[T]) Concat(values ...[]T) Slice[T] {
	total := len(c)
	for _, value := range values {
		total += len(value)
	}

	out := make(Slice[T], 0, total)
	out = append(out, c...)
	for _, value := range values {
		out = append(out, value...)
	}
	return out
}
