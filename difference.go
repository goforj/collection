package collection

// Difference returns a new collection containing elements from the first collection
// that are not present in the second. Order follows the first collection, and
// duplicates are removed.
// @group Set Operations
// @behavior immutable
// @chainable true
// @terminal false
//
// Example: integers
//
//	a := collection.New([]int{1, 2, 2, 3, 4})
//	b := collection.New([]int{2, 4})
//
//	out := collection.Difference(a, b)
//	collection.Dump(out.Items())
//	// #[]int [
//	//   0 => 1 #int
//	//   1 => 3 #int
//	// ]
//
// Example: strings
//
//	left := collection.New([]string{"apple", "banana", "cherry"})
//	right := collection.New([]string{"banana"})
//
//	out2 := collection.Difference(left, right)
//	collection.Dump(out2.Items())
//	// #[]string [
//	//   0 => "apple" #string
//	//   1 => "cherry" #string
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
//		{ID: 3, Name: "Carol"},
//	})
//
//	groupB := collection.New([]User{
//		{ID: 2, Name: "Bob"},
//	})
//
//	out3 := collection.Difference(groupA, groupB)
//	collection.Dump(out3.Items())
//	// #[]main.User [
//	//   0 => #main.User {
//	//     +ID   => 1 #int
//	//     +Name => "Alice" #string
//	//   }
//	//   1 => #main.User {
//	//     +ID   => 3 #int
//	//     +Name => "Carol" #string
//	//   }
//	// ]
func Difference[T comparable](a, b *Collection[T]) *Collection[T] {
	if len(a.items) == 0 {
		return New([]T{})
	}

	lookup := make(map[T]struct{}, len(b.items))
	for _, v := range b.items {
		lookup[v] = struct{}{}
	}

	out := make([]T, 0, len(a.items))
	seen := make(map[T]struct{}, len(a.items))

	for _, v := range a.items {
		if _, inB := lookup[v]; inB {
			continue
		}
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		out = append(out, v)
	}

	return New(out)
}
