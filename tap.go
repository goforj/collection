package collection

// Tap passes the collection to fn for side effects (logging, inspection, debugging)
// without modifying the collection. The collection is returned so it can continue
// chaining, matching Laravel's behavior.
//
// Example:
//
//   c := New([]int{3,1,2}).Sort(func(a,b int) bool { return a < b }).Tap(func(col Collection[int]) {
//       fmt.Println("After sorting:", col.Items())
//   })
//
//   // c is still the sorted collection.
func (c Collection[T]) Tap(fn func(Collection[T])) Collection[T] {
	fn(c) // pass a copy of the collection value (safe)
	return c
}
