package collection

// Zip combines this collection with values element-wise into pairs.
// The resulting length is the smaller of the two inputs.
// @group Transformation
// @behavior immutable
// @chainable false
// @terminal true
//
// Example: integers and strings
//
//	nums := collection.New([]int{1, 2, 3})
//	words := []string{"one", "two"}
//
//	out := nums.Zip(words)
//	collection.Dump(out)
//	// #[]collection.Pair[int,string] [
//	//   0 => #collection.Pair[int,string] {
//	//     +First  => 1 #int
//	//     +Second => "one" #string
//	//   }
//	//   1 => #collection.Pair[int,string] {
//	//     +First  => 2 #int
//	//     +Second => "two" #string
//	//   }
//	// ]
func (c Slice[T]) Zip[U any](values []U) []Pair[T, U] {
	n := len(c)
	if len(values) < n {
		n = len(values)
	}

	out := make([]Pair[T, U], n)
	for i := 0; i < n; i++ {
		out[i] = Pair[T, U]{First: c[i], Second: values[i]}
	}

	return out
}
