package collection

// UniqueComparable returns a new collection with duplicate comparable items removed.
// The first occurrence of each value is kept, and order is preserved.
// This is a faster, allocation-friendly path for comparable types.
// @group Set Operations
// @behavior immutable
// @fluent true
//
// Example: integers
//
//	c := collection.New([]int{1, 2, 2, 3, 4, 4, 5})
//	out := collection.UniqueComparable(c)
//	collection.Dump(out.Items())
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
//	c2 := collection.New([]string{"A", "a", "B", "B"})
//	out2 := collection.UniqueComparable(c2)
//	collection.Dump(out2.Items())
//	// #[]string [
//	//   0 => "A" #string
//	//   1 => "a" #string
//	//   2 => "B" #string
//	// ]
func UniqueComparable[T comparable](c *Collection[T]) *Collection[T] {
	n := len(c.items)
	if n == 0 {
		return New([]T{})
	}

	seen := make(map[T]struct{}, n)

	out := make([]T, 0, n)

	for _, v := range c.items {
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		out = append(out, v)
	}

	return &Collection[T]{items: out}
}
