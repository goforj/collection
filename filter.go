package collection

// Filter keeps only the elements for which fn returns true.
// This method mutates the collection in place and returns the same instance.
// @group Slicing
// @behavior mutable
// @fluent true
// Example: integers
//
//	c := collection.New([]int{1, 2, 3, 4})
//	c.Filter(func(v int) bool {
//		return v%2 == 0
//	})
//	collection.Dump(c.Items())
//	// #[]int [
//	//   0 => 2 #int
//	//   1 => 4 #int
//	// ]
//
// Example: strings
//
//	c2 := collection.New([]string{"apple", "banana", "cherry", "avocado"})
//	c2.Filter(func(v string) bool {
//		return strings.HasPrefix(v, "a")
//	})
//	collection.Dump(c2.Items())
//	// #[]string [
//	//   0 => "apple" #string
//	//   1 => "avocado" #string
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
//		{ID: 3, Name: "Andrew"},
//		{ID: 4, Name: "Carol"},
//	})
//
//	users.Filter(func(u User) bool {
//		return strings.HasPrefix(u.Name, "A")
//	})
//
//	collection.Dump(users.Items())
//	// #[]main.User [
//	//   0 => #main.User {
//	//     +ID   => 1 #int
//	//     +Name => "Alice" #string
//	//   }
//	//   1 => #main.User {
//	//     +ID   => 3 #int
//	//     +Name => "Andrew" #string
//	//   }
//	// ]
func (c *Collection[T]) Filter(fn func(T) bool) *Collection[T] {
	j := 0
	for i := 0; i < len(c.items); i++ {
		if fn(c.items[i]) {
			c.items[j] = c.items[i] // compact in place
			j++
		}
	}

	c.items = c.items[:j] // shrink to new length
	return c
}
