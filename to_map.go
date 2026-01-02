package collection

// ToMap reduces a collection into a map using the provided key and value
// selector functions.
// @group Maps
// @behavior readonly
// @fluent false
// @terminal true
//
// If multiple items produce the same key, the last value wins.
//
// This operation allocates a map sized to the collection length.
//
// Example: basic usage
//
//	users := []string{"alice", "bob", "carol"}
//
//	out := collection.ToMap(
//		collection.New(users),
//		func(name string) string { return name },
//		func(name string) int { return len(name) },
//	)
//
//	collection.Dump(out)
//
// Example: re-keying structs
//
//	type User struct {
//		ID   int
//		Name string
//	}
//
//	users2 := []User{
//		{ID: 1, Name: "Alice"},
//		{ID: 2, Name: "Bob"},
//	}
//
//	byID := collection.ToMap(
//		collection.New(users2),
//		func(u User) int { return u.ID },
//		func(u User) User { return u },
//	)
//
//	collection.Dump(byID)
func ToMap[T any, K comparable, V any](
	c *Collection[T],
	keyFn func(T) K,
	valueFn func(T) V,
) map[K]V {
	out := make(map[K]V, len(c.items))
	for _, item := range c.items {
		out[keyFn(item)] = valueFn(item)
	}
	return out
}
