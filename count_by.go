package collection

// CountBy returns a map of keys extracted by fn to their occurrence counts.
// K must be comparable.
// Example:
//   counts := CountBy(users, func(u User) string { return u.Role })
//  // counts == map[string]int{"admin": 3, "user": 5}
func CountBy[T any, K comparable](c Collection[T], fn func(T) K) map[K]int {
	items := c.Items()
	result := make(map[K]int, len(items))

	for _, v := range items {
		key := fn(v)
		result[key]++
	}

	return result
}

// CountByValue returns a map of item values to their occurrence counts.
// T must be comparable.
// Example:
//   counts := CountByValue(collection.New([]string{"a", "b", "a"}))
//  // counts == map[string]int{"a": 2, "b": 1}
func CountByValue[T comparable](c Collection[T]) map[T]int {
	items := c.Items()
	result := make(map[T]int, len(items))

	for _, v := range items {
		result[v]++
	}

	return result
}
