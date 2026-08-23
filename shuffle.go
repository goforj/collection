package collection

import "math/rand/v2"

// Shuffle shuffles the collection in place and returns the same collection.
// @group Ordering
// @behavior mutable
// @chainable true
// @terminal false
//
// This operation mutates the receiver's backing slice.
//
// Example: integers
//
//	c := collection.New([]int{1, 2, 3, 4, 5})
//	c.Shuffle()
//	fmt.Println(len(c), collection.Sum(c))
//	// 5 15
//
// Example: strings - chaining
//
//	out2 := collection.New([]string{"a", "b", "c"}).
//		Shuffle().
//		Concat([]string{"d"})
//
//	fmt.Println(len(out2))
//	// 4
//
// Example: structs
//
//	type User struct {
//		ID int
//	}
//
//	users := collection.New([]User{
//		{ID: 1},
//		{ID: 2},
//		{ID: 3},
//		{ID: 4},
//	})
//
//	users.Shuffle()
//	fmt.Println(len(users))
//	// 4
func (c Slice[T]) Shuffle() Slice[T] {
	n := len(c)

	for i := n - 1; i > 0; i-- {
		j := rand.IntN(i + 1)
		c[i], c[j] = c[j], c[i]
	}

	return c
}
