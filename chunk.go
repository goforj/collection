package collection

// Chunk splits the collection into chunks of the given size.
// The final chunk may be smaller if len(items) is not divisible by size.
// @group Slicing
// @behavior readonly
// @chainable false
// @terminal true
//
// If size <= 0, nil is returned.
//
// Chunk allocates the outer result slice. Each chunk is a capacity-capped view
// that shares the backing array with the source collection.
// Example: integers
//
//	collection.Dump(collection.New([]int{1, 2, 3, 4, 5}).Chunk(2))
//	// #[][]int [
//	//  0 => #[]int [
//	//    0 => 1 #int
//	//    1 => 2 #int
//	//  ]
//	//  1 => #[]int [
//	//    0 => 3 #int
//	//    1 => 4 #int
//	//  ]
//	//  2 => #[]int [
//	//    0 => 5 #int
//	//  ]
//	//]
//
// Example: structs
//
//	type User struct {
//		ID   int
//		Name string
//	}
//
//	users := []User{
//		{ID: 1, Name: "Alice"},
//		{ID: 2, Name: "Bob"},
//		{ID: 3, Name: "Carol"},
//		{ID: 4, Name: "Dave"},
//	}
//
//	userChunks := collection.New(users).Chunk(2)
//	collection.Dump(userChunks)
//	// #[][]main.User [
//	//  0 => #[]main.User [
//	//    0 => #main.User {
//	//      +ID   => 1 #int
//	//      +Name => "Alice" #string
//	//    }
//	//    1 => #main.User {
//	//      +ID   => 2 #int
//	//      +Name => "Bob" #string
//	//    }
//	//  ]
//	//  1 => #[]main.User [
//	//    0 => #main.User {
//	//      +ID   => 3 #int
//	//      +Name => "Carol" #string
//	//    }
//	//    1 => #main.User {
//	//      +ID   => 4 #int
//	//      +Name => "Dave" #string
//	//    }
//	//  ]
//	//]
func (c Slice[T]) Chunk(size int) [][]T {
	if size <= 0 {
		return nil
	}

	n := len(c)
	count := n / size
	if n%size != 0 {
		count++
	}
	chunks := make([][]T, 0, count)

	for i := 0; i < n; i += size {
		end := i + size
		if end > n {
			end = n
		}
		chunks = append(chunks, c[i:end:end])
	}

	return chunks
}
