package collection

// Intersect returns a new collection containing elements present in both collections.
// Order follows the first collection, and duplicates are removed.
// @group Set Operations
// @behavior immutable
// @chainable true
//
// Example: integers
//
//	a := collection.New([]int{1, 2, 2, 3, 4})
//	b := collection.New([]int{2, 4, 4, 5})
//
//	out := collection.Intersect(a, b)
//	collection.Dump(out.Items())
//	// #[]int [
//	//   0 => 2 #int
//	//   1 => 4 #int
//	// ]
//
// Example: strings
//
//	left := collection.New([]string{"apple", "banana", "cherry"})
//	right := collection.New([]string{"banana", "date", "cherry", "banana"})
//
//	out2 := collection.Intersect(left, right)
//	collection.Dump(out2.Items())
//	// #[]string [
//	//   0 => "banana" #string
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
//		{ID: 3, Name: "Carol"},
//		{ID: 4, Name: "Dave"},
//	})
//
//	out3 := collection.Intersect(groupA, groupB)
//	collection.Dump(out3.Items())
//	// #[]main.User [
//	//   0 => #main.User {
//	//     +ID   => 2 #int
//	//     +Name => "Bob" #string
//	//   }
//	//   1 => #main.User {
//	//     +ID   => 3 #int
//	//     +Name => "Carol" #string
//	//   }
//	// ]
func Intersect[T comparable](a, b *Collection[T]) *Collection[T] {
	if len(a.items) == 0 || len(b.items) == 0 {
		return New([]T{})
	}

	lookup := make(map[T]struct{}, len(b.items))
	for _, v := range b.items {
		lookup[v] = struct{}{}
	}

	out := make([]T, 0, len(a.items))
	seen := make(map[T]struct{}, len(a.items))

	for _, v := range a.items {
		if _, ok := lookup[v]; !ok {
			continue
		}
		if _, already := seen[v]; already {
			continue
		}
		seen[v] = struct{}{}
		out = append(out, v)
	}

	return &Collection[T]{items: out}
}
