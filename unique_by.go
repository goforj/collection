package collection

// UniqueBy returns a new collection containing only the first occurrence
// of each element as determined by keyFn.
// @group Set Operations
// @behavior immutable
// @fluent true
//
// The key returned by keyFn must be comparable.
// Order is preserved.
//
// Example: structs – unique by ID
//
//	type User struct {
//		ID   int
//		Name string
//	}
//
//	users := collection.New([]User{
//		{ID: 1, Name: "Alice"},
//		{ID: 2, Name: "Bob"},
//		{ID: 1, Name: "Alice Duplicate"},
//	})
//
//	out := collection.UniqueBy(users, func(u User) int { return u.ID })
//	collection.Dump(out.Items())
//	// #[]collection.User [
//	//   0 => {ID:1 Name:"Alice"} #collection.User
//	//   1 => {ID:2 Name:"Bob"}   #collection.User
//	// ]
//
// Example: strings – case-insensitive uniqueness
//
//	values := collection.New([]string{"A", "a", "B", "b", "A"})
//
//	out2 := collection.UniqueBy(values, func(s string) string {
//		return strings.ToLower(s)
//	})
//
//	collection.Dump(out2.Items())
//	// #[]string [
//	//   0 => "A" #string
//	//   1 => "B" #string
//	// ]
//
// Example: integers – identity key
//
//	nums := collection.New([]int{3, 1, 2, 1, 3})
//
//	out3 := collection.UniqueBy(nums, func(v int) int { return v })
//	collection.Dump(out3.Items())
//	// #[]int [
//	//   0 => 3 #int
//	//   1 => 1 #int
//	//   2 => 2 #int
//	// ]
func UniqueBy[T any, K comparable](c *Collection[T], keyFn func(T) K) *Collection[T] {
	items := c.items
	n := len(items)
	if n == 0 {
		return Attach([]T{})
	}

	seen := make(map[K]struct{}, n)
	out := make([]T, 0, n)

	for _, v := range items {
		k := keyFn(v)
		if _, ok := seen[k]; ok {
			continue
		}
		seen[k] = struct{}{}
		out = append(out, v)
	}

	return Attach(out)
}
