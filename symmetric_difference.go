package collection

// SymmetricDifference returns a new collection containing elements that appear
// in exactly one of the two collections. Order follows the first collection for
// its unique items, then the second for its unique items. Duplicates are removed.
// @group Set Operations
// @behavior immutable
// @chainable true
// @terminal false
//
// Example: integers
//
//	a := collection.New([]int{1, 2, 3, 3})
//	b := collection.New([]int{3, 4, 4, 5})
//
//	collection.Dump(collection.SymmetricDifference(a, b))
//	// #[]int [
//	//   0 => 1 #int
//	//   1 => 2 #int
//	//   2 => 4 #int
//	//   3 => 5 #int
//	// ]
//
// Example: strings
//
//	left := collection.New([]string{"apple", "banana"})
//	right := collection.New([]string{"banana", "date"})
//
//	collection.Dump(collection.SymmetricDifference(left, right))
//	// #[]string [
//	//   0 => "apple" #string
//	//   1 => "date" #string
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
//	collection.Dump(collection.SymmetricDifference(groupA, groupB))
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
func SymmetricDifference[S1 ~[]T, S2 ~[]T, T comparable](a S1, b S2) Slice[T] {
	out := make([]T, 0, len(a)+len(b))
	seenOut := make(map[T]struct{}, len(a)+len(b))

	setA := make(map[T]struct{}, len(a))
	for _, v := range a {
		setA[v] = struct{}{}
	}

	setB := make(map[T]struct{}, len(b))
	for _, v := range b {
		setB[v] = struct{}{}
	}

	for _, v := range a {
		if _, inB := setB[v]; inB {
			continue
		}
		if _, ok := seenOut[v]; ok {
			continue
		}
		seenOut[v] = struct{}{}
		out = append(out, v)
	}

	for _, v := range b {
		if _, inA := setA[v]; inA {
			continue
		}
		if _, ok := seenOut[v]; ok {
			continue
		}
		seenOut[v] = struct{}{}
		out = append(out, v)
	}

	return New(out)
}
