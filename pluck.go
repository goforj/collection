package collection

// MapTo maps a Collection[T] to a Collection[R] using fn(T) R.
// @group Transformation
// @behavior immutable
// @fluent true
//
// This cannot be a method because methods can't introduce a new type parameter R.
//
// Example: integers - extract parity label
//
//	nums := collection.New([]int{1, 2, 3, 4})
//	parity := collection.MapTo(nums, func(n int) string {
//		if n%2 == 0 {
//			return "even"
//		}
//		return "odd"
//	})
//	collection.Dump(parity.Items())
//	// #[]string [
//	//   0 => "odd" #string
//	//   1 => "even" #string
//	//   2 => "odd" #string
//	//   3 => "even" #string
//	// ]
//
// Example: strings - length of each value
//
//	words := collection.New([]string{"go", "forj", "rocks"})
//	lengths := collection.MapTo(words, func(s string) int {
//		return len(s)
//	})
//	collection.Dump(lengths.Items())
//	// #[]int [
//	//   0 => 2 #int
//	//   1 => 4 #int
//	//   2 => 5 #int
//	// ]
//
// Example: structs - MapTo a field
//
//	type User struct {
//		ID   int
//		Name string
//	}
//
//	users := collection.New([]User{
//		{ID: 1, Name: "Alice"},
//		{ID: 2, Name: "Bob"},
//	})
//
//	names := collection.MapTo(users, func(u User) string {
//		return u.Name
//	})
//
//	collection.Dump(names.Items())
//	// #[]string [
//	//   0 => "Alice" #string
//	//   1 => "Bob" #string
//	// ]
func MapTo[T any, R any](c *Collection[T], fn func(T) R) *Collection[R] {
	items := c.Items()
	out := make([]R, len(items))
	for i, v := range items {
		out[i] = fn(v)
	}
	return Attach(out)
}
