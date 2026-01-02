package collection

// Each runs fn for every item in the collection and returns the same collection,
// so it can be used in chains for side effects (logging, debugging, etc.).
// @group Transformation
// @behavior readonly
// @chainable true
// @terminal false
//
// Example: integers
//
//	c := collection.New([]int{1, 2, 3})
//
//	sum := 0
//	c.Each(func(v int) {
//		sum += v
//	})
//
//	collection.Dump(sum)
//	// 6 #int
//
// Example: strings
//
//	c2 := collection.New([]string{"apple", "banana", "cherry"})
//
//	var out []string
//	c2.Each(func(s string) {
//		out = append(out, strings.ToUpper(s))
//	})
//
//	collection.Dump(out)
//	// #[]string [
//	//   0 => "APPLE" #string
//	//   1 => "BANANA" #string
//	//   2 => "CHERRY" #string
//	// ]
//
// Example: structs
//
//	type User struct {
//		ID   int
//		Name string
//	}
//
//	users := collection.New([]User{
//		{ID: 1, Name: "Alice"},
//		{ID: 2, Name: "Bob"},
//		{ID: 3, Name: "Charlie"},
//	})
//
//	var names []string
//	users.Each(func(u User) {
//		names = append(names, u.Name)
//	})
//
//	collection.Dump(names)
//	// #[]string [
//	//   0 => "Alice" #string
//	//   1 => "Bob" #string
//	//   2 => "Charlie" #string
//	// ]
func (c *Collection[T]) Each(fn func(T)) *Collection[T] {
	for _, v := range c.items {
		fn(v)
	}
	return c
}
