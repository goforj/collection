package collection

// Count returns the total number of items in the collection.
// @group Aggregation
// @behavior readonly
// @fluent true
// Example: integers
//
//	count := collection.New([]int{1, 2, 3, 4}).Count()
//	collection.Dump(count)
//	// 4 #int
func (c *Collection[T]) Count() int {
	return len(c.items)
}
