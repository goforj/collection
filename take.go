package collection

// Take returns a capacity-capped view containing the first n items.
//
// If n exceeds the collection length, the entire collection is returned.
// If n == 0, an empty collection is returned.
//
// NOTE: returns a view (shares backing array). Use Clone() to detach.
//
// @group Slicing
// @behavior immutable
// @chainable true
// @terminal false
// Example: integers - take first 3
//
//	c1 := collection.New([]int{0, 1, 2, 3, 4, 5})
//	out1 := c1.Take(3)
//	collection.Dump(out1)
//	// #[]int [
//	//	0 => 0 #int
//	//	1 => 1 #int
//	//	2 => 2 #int
//	// ]
//
// Example: integers - n exceeds length → whole collection
//
//	c3 := collection.New([]int{10, 20})
//	out3 := c3.Take(10)
//	collection.Dump(out3)
//	// #[]int [
//	//	0 => 10 #int
//	//	1 => 20 #int
//	// ]
//
// Example: integers - zero → empty
//
//	c4 := collection.New([]int{1, 2, 3})
//	out4 := c4.Take(0)
//	collection.Dump(out4)
//	// #[]int [
//	// ]
func (c Slice[T]) Take(n int) Slice[T] {
	length := len(c)

	if n <= 0 || length == 0 {
		return New(c[:0:0])
	}

	if n >= length {
		return New(c[:length:length])
	}
	return New(c[:n:n])
}
