package collection

// UniqueComparable returns a new collection with duplicate comparable items removed.
// The first occurrence of each value is kept, and order is preserved.
// It uses a map to track seen values, so it has expected linear time and allocates
// storage for both the map and the result.
// @group Set Operations
// @behavior immutable
// @chainable true
// @terminal false
//
// Example: integers
//
//	collection.Dump(collection.UniqueComparable([]int{1, 2, 2, 3, 4, 4, 5}))
//	// #[]int [
//	//   0 => 1 #int
//	//   1 => 2 #int
//	//   2 => 3 #int
//	//   3 => 4 #int
//	//   4 => 5 #int
//	// ]
//
// Example: strings
//
//	collection.Dump(collection.UniqueComparable([]string{"A", "a", "B", "B"}))
//	// #[]string [
//	//   0 => "A" #string
//	//   1 => "a" #string
//	//   2 => "B" #string
//	// ]
func UniqueComparable[S ~[]T, T comparable](c S) Slice[T] {
	n := len(c)
	if n == 0 {
		return New([]T{})
	}

	seen := make(map[T]struct{}, n)

	out := make([]T, 0, n)

	for _, v := range c {
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		out = append(out, v)
	}

	return New(out)
}
