package collection

// MaxBy returns the item whose key (produced by keyFn) is the largest.
// The second return value is false if the collection is empty.
// @group Aggregation
// @behavior readonly
// @chainable false
//
// This cannot be a method because methods can't introduce a new type parameter K.
// When multiple items share the same maximal key, the first such item is returned.
//
// Example: structs - highest score
//
//	type Player struct {
//		Name  string
//		Score int
//	}
//
//	players := collection.New([]Player{
//		{Name: "Alice", Score: 10},
//		{Name: "Bob", Score: 25},
//		{Name: "Carol", Score: 18},
//	})
//
//	top, ok := collection.MaxBy(players, func(p Player) int {
//		return p.Score
//	})
//
//	collection.Dump(top, ok)
//	// #main.Player {
//	//   +Name  => "Bob" #string
//	//   +Score => 25 #int
//	// }
//	// true #bool
//
// Example: strings - longest length
//
//	words := collection.New([]string{"go", "collection", "rocks"})
//
//	longest, ok := collection.MaxBy(words, func(s string) int {
//		return len(s)
//	})
//
//	collection.Dump(longest, ok)
//	// "collection" #string
//	// true #bool
//
// Example: empty collection
//
//	empty := collection.New([]int{})
//	maxVal, ok := collection.MaxBy(empty, func(v int) int { return v })
//	collection.Dump(maxVal, ok)
//	// 0 #int
//	// false #bool
func MaxBy[T any, K Number | ~string](c *Collection[T], keyFn func(T) K) (T, bool) {
	var zero T

	if len(c.items) == 0 {
		return zero, false
	}

	maxItem := c.items[0]
	maxKey := keyFn(maxItem)

	for _, item := range c.items[1:] {
		key := keyFn(item)
		if key > maxKey {
			maxKey = key
			maxItem = item
		}
	}

	return maxItem, true
}
