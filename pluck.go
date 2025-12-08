package collection

// MapTo maps a Collection[T] to a Collection[R] using fn(T) R.
//
// This cannot be a method because methods can't introduce a new type parameter R.
// Example:
//   squared := numbers.MapTo(func(n int) int { return n * n })
//   // squared is a Collection[int] of squared numbers
func MapTo[T any, R any](c *Collection[T], fn func(T) R) *Collection[R] {
	items := c.Items()
	out := make([]R, len(items))
	for i, v := range items {
		out[i] = fn(v)
	}
	return &Collection[R]{items: out}
}

// Pluck is an alias for MapTo with a more semantic name when projecting fields.
// Example:
//   names := users.Pluck(func(u User) string { return u.Name })
//   // names is a Collection[string] of user names
func Pluck[T any, R any](c *Collection[T], fn func(T) R) *Collection[R] {
	return MapTo(c, fn)
}
