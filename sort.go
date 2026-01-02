package collection

import "sort"

// Sort sorts the collection in place using the provided comparison function and
// returns the same collection for chaining.
// @group Ordering
// @behavior mutable
// @chainable true
// @terminal false
//
// The comparison function `less(a, b)` should return true if `a` should come
// before `b` in the sorted order.
//
// This operation mutates the underlying slice (no allocation).
//
// Example: integers
//
//	c := collection.New([]int{5, 1, 4, 2})
//	c.Sort(func(a, b int) bool { return a < b })
//	collection.Dump(c.Items())
//	// #[]int [
//	//   0 => 1 #int
//	//   1 => 2 #int
//	//   2 => 4 #int
//	//   3 => 5 #int
//	// ]
//
// Example: strings (descending)
//
//	c2 := collection.New([]string{"apple", "banana", "cherry"})
//	c2.Sort(func(a, b string) bool { return a > b })
//	collection.Dump(c2.Items())
//	// #[]string [
//	//   0 => "cherry" #string
//	//   1 => "banana" #string
//	//   2 => "apple" #string
//	// ]
//
// Example: structs
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
//	// Sort by age ascending
//	users.Sort(func(a, b User) bool {
//		return a.Age < b.Age
//	})
//	collection.Dump(users.Items())
//	// #[]main.User [
//	//   0 => #main.User {
//	//     +Name => "Bob" #string
//	//     +Age  => 25 #int
//	//   }
//	//   1 => #main.User {
//	//     +Name => "Alice" #string
//	//     +Age  => 30 #int
//	//   }
//	//   2 => #main.User {
//	//     +Name => "Carol" #string
//	//     +Age  => 40 #int
//	//   }
//	// ]
func (c *Collection[T]) Sort(less func(a, b T) bool) *Collection[T] {
	sort.Slice(c.items, func(i, j int) bool {
		return less(c.items[i], c.items[j])
	})
	return c
}
