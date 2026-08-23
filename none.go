package collection

// None returns true if fn returns false for every item in the collection.
// If the collection is empty, None returns true.
// @group Querying
// @behavior readonly
// @chainable false
// @terminal true
//
// Example: integers - none even
//
//	collection.Dump(collection.New([]int{1, 3, 5}).None(func(v int) bool { return v%2 == 0 }))
//	// true #bool
//
// Example: integers - some even
//
//	collection.Dump(collection.New([]int{1, 2, 3}).None(func(v int) bool { return v%2 == 0 }))
//	// false #bool
//
// Example: empty collection
//
//	collection.Dump(collection.New([]int{}).None(func(v int) bool { return v > 0 }))
//	// true #bool
func (c Slice[T]) None(fn func(T) bool) bool {
	for _, v := range c {
		if fn(v) {
			return false
		}
	}
	return true
}
