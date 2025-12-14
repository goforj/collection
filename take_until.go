package collection

// TakeUntilFn returns items until the predicate function returns true.
// The matching item is NOT included.
// @group Slicing
// @behavior immutable
// @chainable true
// Example: integers - stop when value >= 3
//
//	c1 := collection.New([]int{1, 2, 3, 4})
//	out1 := c1.TakeUntilFn(func(v int) bool { return v >= 3 })
//	collection.Dump(out1.Items())
//	// #[]int [
//	//	0 => 1 #int
//	//	1 => 2 #int
//	// ]
//
// Example: integers - predicate immediately true → empty result
//
//	c2 := collection.New([]int{10, 20, 30})
//	out2 := c2.TakeUntilFn(func(v int) bool { return v < 50 })
//	collection.Dump(out2.Items())
//	// #[]int [
//	// ]
//
// Example: integers - no match → full list returned
//
//	c3 := collection.New([]int{1, 2, 3})
//	out3 := c3.TakeUntilFn(func(v int) bool { return v == 99 })
//	collection.Dump(out3.Items())
//	// #[]int [
//	//	0 => 1 #int
//	//	1 => 2 #int
//	//	2 => 3 #int
//	// ]
func (c *Collection[T]) TakeUntilFn(pred func(T) bool) *Collection[T] {
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
// @chainable true
//
// Uses == comparison, so T must be comparable.
// @group Slicing
// @behavior immutable
// Example: integers - stop at value 3
//
//	c4 := collection.New([]int{1, 2, 3, 4})
//	out4 := collection.TakeUntil(c4, 3)
//	collection.Dump(out4.Items())
//	// #[]int [
//	//	0 => 1 #int
//	//	1 => 2 #int
//	// ]
//
// Example: strings - value never appears → full slice
//
//	c5 := collection.New([]string{"a", "b", "c"})
//	out5 := collection.TakeUntil(c5, "x")
//	collection.Dump(out5.Items())
//	// #[]string [
//	//	0 => "a" #string
//	//	1 => "b" #string
//	//	2 => "c" #string
//	// ]
//
// Example: integers - match is first item → empty result
//
//	c6 := collection.New([]int{9, 10, 11})
//	out6 := collection.TakeUntil(c6, 9)
//	collection.Dump(out6.Items())
//	// #[]int [
//	// ]
func TakeUntil[T comparable](c *Collection[T], value T) *Collection[T] {
	out := make([]T, 0, len(c.items))

	for _, v := range c.items {
		if v == value {
			break
		}
		out = append(out, v)
	}

	return New(out)
}
