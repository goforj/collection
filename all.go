package collection

// All returns true if fn returns true for every item in the collection.
// If the collection is empty, All returns true (vacuously true).
// @group Querying
// @behavior readonly
// @chainable false
// @terminal true
//
// Example: integers - all even
//
//	collection.Dump(collection.New([]int{2, 4, 6}).All(func(v int) bool { return v%2 == 0 }))
//	// true #bool
//
// Example: integers - not all even
//
//	collection.Dump(collection.New([]int{2, 3, 4}).All(func(v int) bool { return v%2 == 0 }))
//	// false #bool
//
// Example: strings - all non-empty
//
//	collection.Dump(collection.New([]string{"a", "b", "c"}).All(func(s string) bool { return s != "" }))
//	// true #bool
//
// Example: empty collection (vacuously true)
//
//	collection.Dump(collection.New([]int{}).All(func(v int) bool { return v > 0 }))
//	// true #bool
func (c Slice[T]) All(fn func(T) bool) bool {
	for _, v := range c {
		if !fn(v) {
			return false
		}
	}
	return true
}
