package collection

// Map applies a same-type transformation and returns a new collection.
//
// Use this when you're transforming T -> T (e.g., enrichment, normalization).
//
// Example:
//	// integers
//	c := collection.New([]int{1, 2, 3})
//
//	mapped := c.Map(func(v int) int {
//		return v * 10
//	})
//
//	collection.Dump(mapped.Items())
//	// #[]int [
//	//   0 => 10 #int
//	//   1 => 20 #int
//	//   2 => 30 #int
//	// ]
//
// Example:
//	// strings
//	c2 := collection.New([]string{"apple", "banana", "cherry"})
//
//	upper := c2.Map(func(s string) string {
//		return strings.ToUpper(s)
//	})
//
//	collection.Dump(upper.Items())
//	// #[]string [
//	//   0 => "APPLE"  #string
//	//   1 => "BANANA" #string
//	//   2 => "CHERRY" #string
//	// ]
//
// Example:
//	// structs
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
//	updated := users.Map(func(u User) User {
//		u.Name = strings.ToUpper(u.Name)
//		return u
//	})
//
//	collection.Dump(updated.Items())
//	// #[]main.User [
//	//   0 => #main.User {
//	//     +ID   => 1        #int
//	//     +Name => "ALICE"  #string
//	//   }
//	//   1 => #main.User {
//	//     +ID   => 2        #int
//	//     +Name => "BOB"    #string
//	//   }
//	// ]
func (c *Collection[T]) Map(fn func(T) T) *Collection[T] {
	for i := range c.items {
		c.items[i] = fn(c.items[i])
	}
	return c
}
