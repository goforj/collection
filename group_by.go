package collection

// GroupBy partitions the collection into groups keyed by the value
// returned from keyFn.
//
// The order of items within each group is preserved.
// The order of the groups themselves is unspecified.
//
// This function does not mutate the source collection.
//
// Example: grouping integers by parity
//
//	values := []int{1, 2, 3, 4, 5}
//
//	groups := collection.GroupBy(
//		collection.New(values),
//		func(v int) string {
//			if v%2 == 0 {
//				return "even"
//			}
//			return "odd"
//		},
//	)
//
//	collection.Dump(groups["even"].Items())
//	// []int [
//	//  0 => 2 #int
//	//  1 => 4 #int
//	// ]
//	collection.Dump(groups["odd"].Items())
//	// []int [
//	//  0 => 1 #int
//	//  1 => 3 #int
//	//  2 => 5 #int
//	// ]
//
// Example: grouping structs by field
//
//	type User struct {
//		ID   int
//		Role string
//	}
//
//	users := []User{
//		{ID: 1, Role: "admin"},
//		{ID: 2, Role: "user"},
//		{ID: 3, Role: "admin"},
//	}
//
//	groups2 := collection.GroupBy(
//		collection.New(users),
//		func(u User) string { return u.Role },
//	)
//
//	collection.Dump(groups2["admin"].Items())
//	// []main.User [
//	//  0 => #main.User {
//	//    +ID   => 1 #int
//	//    +Role => "admin" #string
//	//  }
//	//  1 => #main.User {
//	//    +ID   => 3 #int
//	//    +Role => "admin" #string
//	//  }
//	// ]
//	collection.Dump(groups2["user"].Items())
//  // []main.User [
//	//  0 => #main.User {
//	//    +ID   => 2 #int
//	//    +Role => "user" #string
//	//  }
//	// ]
func GroupBy[T any, K comparable](
	c *Collection[T],
	keyFn func(T) K,
) map[K]*Collection[T] {
	out := make(map[K]*Collection[T])

	for _, item := range c.items {
		key := keyFn(item)

		group, ok := out[key]
		if !ok {
			group = &Collection[T]{items: make([]T, 0)}
			out[key] = group
		}

		group.items = append(group.items, item)
	}

	return out
}
