package collection

// Filter keeps only the elements for which fn returns true.
//
// Filter allocates a new Slice and leaves c and its backing storage unchanged.
// @group Slicing
// @behavior immutable
// @chainable true
// @terminal false
// Example: integers
//
//	source := collection.New([]int{1, 2, 3, 4})
//	filtered := source.Filter(func(v int) bool {
//		return v%2 == 0
//	})
//	collection.Dump(filtered)
//	// #[]int [
//	//   0 => 2 #int
//	//   1 => 4 #int
//	// ]
//	fmt.Println(source[0])
//	// 1
//
// Example: strings
//
//	c2 := collection.New([]string{"apple", "banana", "cherry", "avocado"})
//	c2 = c2.Filter(func(v string) bool {
//		return strings.HasPrefix(v, "a")
//	})
//	collection.Dump(c2)
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
//	users = users.Filter(func(u User) bool {
//		return strings.HasPrefix(u.Name, "A")
//	})
//
//	collection.Dump(users)
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
func (c Slice[T]) Filter(fn func(T) bool) Slice[T] {
	out := make([]T, 0, len(c))
	for _, item := range c {
		if fn(item) {
			out = append(out, item)
		}
	}
	return New(out)
}
