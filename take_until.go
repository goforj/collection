package collection

// TakeUntil returns items until the predicate function returns true.
// The matching item is NOT included.
// @group Slicing
// @behavior immutable
// @chainable true
// @terminal false
//
// NOTE: returns a view (shares backing array). Use Clone() to detach.
// Example: integers - stop when value >= 3
//
//	c1 := collection.New([]int{1, 2, 3, 4})
//	out1 := c1.TakeUntil(func(v int) bool { return v >= 3 })
//	collection.Dump(out1)
//	// #[]int [
//	//	0 => 1 #int
//	//	1 => 2 #int
//	// ]
//
// Example: integers - predicate immediately true → empty result
//
//	c2 := collection.New([]int{10, 20, 30})
//	out2 := c2.TakeUntil(func(v int) bool { return v < 50 })
//	collection.Dump(out2)
//	// #[]int [
//	// ]
//
// Example: integers - no match → full list returned
//
//	c3 := collection.New([]int{1, 2, 3})
//	out3 := c3.TakeUntil(func(v int) bool { return v == 99 })
//	collection.Dump(out3)
//	// #[]int [
//	//	0 => 1 #int
//	//	1 => 2 #int
//	//	2 => 3 #int
//	// ]
func (c Slice[T]) TakeUntil(pred func(T) bool) Slice[T] {
	idx := len(c)
	for i, v := range c {
		if pred(v) {
			idx = i
			break
		}
	}

	return New(c[:idx:idx])
}
