package collection

// Tap invokes fn with the Slice value for side effects such as logging,
// debugging, or inspection, then returns the Slice to allow chaining.
// @group Transformation
// @behavior mutable
// @chainable true
// @terminal false
//
// The callback receives a borrowed Slice and may mutate its elements. Use Clone
// before Tap when the original backing array must remain isolated.
//
// Example: integers - capture intermediate state during a chain
//
//	captured1 := []int{}
//	c1 := collection.New([]int{3, 1, 2}).
//		Sort(func(a, b int) bool { return a < b }). // → [1, 2, 3]
//		Tap(func(col collection.Slice[int]) {
//			captured1 = append([]int(nil), col...) // snapshot copy
//		}).
//		Filter(func(v int) bool { return v >= 2 }).
//		Dump()
//		// #[]int [
//		//  0 => 2 #int
//		//  1 => 3 #int
//		// ]
//
//	// Use BOTH variables so nothing is "declared and not used"
//	collection.Dump(c1)
//	collection.Dump(captured1)
//	// #[]int [
//	//  0 => 2 #int
//	//  1 => 3 #int
//	// ]
//	// #[]int [
//	//  0 => 1 #int
//	//  1 => 2 #int
//	//  2 => 3 #int
//	// ]
//
// Example: integers - tap for debugging without changing flow
//
//	c2 := collection.New([]int{10, 20, 30}).
//		Tap(func(col collection.Slice[int]) {
//			collection.Dump(col)
//			// #[]int [
//			//  0 => 10 #int
//			//  1 => 20 #int
//			//  2 => 30 #int
//			// ]
//		}).
//		Filter(func(v int) bool { return v > 10 })
//
//	collection.Dump(c2) // ensures c2 is used
//	// #[]int [
//	//  0 => 20 #int
//	//  1 => 30 #int
//	// ]
//
// Example: structs - Tap with struct collection
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
//	users2 := users.Tap(func(col collection.Slice[User]) {
//		collection.Dump(col)
//		// #[]main.User [
//		//  0 => #main.User {
//		//    +ID   => 1 #int
//		//    +Name => "Alice" #string
//		//  }
//		//  1 => #main.User {
//		//    +ID   => 2 #int
//		//    +Name => "Bob" #string
//		//  }
//		// ]
//	})
//
//	collection.Dump(users2) // ensures users2 is used
//	// #[]main.User [
//	//  0 => #main.User {
//	//    +ID   => 1 #int
//	//    +Name => "Alice" #string
//	//  }
//	//  1 => #main.User {
//	//    +ID   => 2 #int
//	//    +Name => "Bob" #string
//	//  }
//	// ]
func (c Slice[T]) Tap(fn func(Slice[T])) Slice[T] {
	fn(c)
	return c
}
