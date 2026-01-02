package collection

// Union returns a new collection containing the unique elements from both collections.
// Items from the first collection are kept in order, followed by items from the second
// that were not already present.
// @group Set Operations
// @behavior immutable
// @fluent true
//
// Example: integers
//
//	a := collection.New([]int{1, 2, 2, 3})
//	b := collection.New([]int{3, 4, 4, 5})
//
//	out := collection.Union(a, b)
//	collection.Dump(out.Items())
//	// #[]int [
//	//   0 => 1 #int
//	//   1 => 2 #int
//	//   2 => 3 #int
//	//   3 => 4 #int
//	//   4 => 5 #int
//	// ]
//
// Example: strings
//
//	left := collection.New([]string{"apple", "banana"})
//	right := collection.New([]string{"banana", "date"})
//
//	out2 := collection.Union(left, right)
//	collection.Dump(out2.Items())
//	// #[]string [
//	//   0 => "apple" #string
//	//   1 => "banana" #string
//	//   2 => "date" #string
//	// ]
//
// Example: structs
//
//	type User struct {
//		ID   int
//		Name string
//	}
//
//	groupA := collection.New([]User{
//		{ID: 1, Name: "Alice"},
//		{ID: 2, Name: "Bob"},
//	})
//
//	groupB := collection.New([]User{
//		{ID: 2, Name: "Bob"},
//		{ID: 3, Name: "Carol"},
//	})
//
//	out3 := collection.Union(groupA, groupB)
//	collection.Dump(out3.Items())
//	// #[]main.User [
//	//   0 => #main.User {
//	//     +ID   => 1 #int
//	//     +Name => "Alice" #string
//	//   }
//	//   1 => #main.User {
//	//     +ID   => 2 #int
//	//     +Name => "Bob" #string
//	//   }
//	//   2 => #main.User {
//	//     +ID   => 3 #int
//	//     +Name => "Carol" #string
//	//   }
//	// ]
func Union[T comparable](a, b *Collection[T]) *Collection[T] {
	out := make([]T, 0, len(a.items)+len(b.items))
	seen := make(map[T]struct{}, len(a.items)+len(b.items))

	for _, v := range a.items {
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		out = append(out, v)
	}

	for _, v := range b.items {
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		out = append(out, v)
	}

	return Attach(out)
}
