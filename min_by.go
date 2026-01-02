package collection

// MinBy returns the item whose key (produced by keyFn) is the smallest.
// The second return value is false if the collection is empty.
// @group Aggregation
// @behavior readonly
// @fluent false
// @terminal true
//
// This cannot be a method because methods can't introduce a new type parameter K.
// When multiple items share the same minimal key, the first such item is returned.
//
// Example: structs - smallest age
//
//	type User struct {
//		Name string
//		Age  int
//	}
//
//	users := collection.New([]User{
//		{Name: "Alice", Age: 30},
//		{Name: "Bob", Age: 25},
//		{Name: "Carol", Age: 40},
//	})
//
//	minUser, ok := collection.MinBy(users, func(u User) int {
//		return u.Age
//	})
//
//	collection.Dump(minUser, ok)
//	// #main.User {
//	//   +Name => "Bob" #string
//	//   +Age  => 25 #int
//	// }
//	// true #bool
//
// Example: strings - shortest length
//
//	words := collection.New([]string{"apple", "fig", "banana"})
//
//	shortest, ok := collection.MinBy(words, func(s string) int {
//		return len(s)
//	})
//
//	collection.Dump(shortest, ok)
//	// "fig" #string
//	// true #bool
//
// Example: empty collection
//
//	empty := collection.New([]int{})
//	minVal, ok := collection.MinBy(empty, func(v int) int { return v })
//	collection.Dump(minVal, ok)
//	// 0 #int
//	// false #bool
func MinBy[T any, K Number | ~string](c *Collection[T], keyFn func(T) K) (T, bool) {
	var zero T

	if len(c.items) == 0 {
		return zero, false
	}

	minItem := c.items[0]
	minKey := keyFn(minItem)

	for _, item := range c.items[1:] {
		key := keyFn(item)
		if key < minKey {
			minKey = key
			minItem = item
		}
	}

	return minItem, true
}
