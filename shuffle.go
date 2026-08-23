package collection

import (
	"math/rand"
	"time"
)

// shuffleRand is the RNG used by Shuffle.
// It is overridden in tests for deterministic behavior.
var shuffleRand = rand.New(rand.NewSource(time.Now().UnixNano()))

// setShuffleRand allows tests to inject a deterministic RNG.
// Not exported — production code should not touch this.
func setShuffleRand(r *rand.Rand) {
	shuffleRand = r
}

// Shuffle shuffles the collection in place and returns the same collection.
// @group Ordering
// @behavior mutable
// @chainable true
// @terminal false
//
// This operation mutates the receiver's backing slice.
//
// The shuffle uses an internal random source. Tests may override
// this source to achieve deterministic behavior.
//
// Example: integers
//
//	c := collection.New([]int{1, 2, 3, 4, 5})
//	c.Shuffle()
//	collection.Dump(c)
//
// Example: strings – chaining
//
//	out2 := collection.New([]string{"a", "b", "c"}).
//		Shuffle().
//		Append("d")
//
//	collection.Dump(out2)
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
//	collection.Dump(users)
func (c Slice[T]) Shuffle() Slice[T] {
	n := len(c)

	// Fisher–Yates shuffle (in place)
	for i := n - 1; i > 0; i-- {
		j := shuffleRand.Intn(i + 1)
		c[i], c[j] = c[j], c[i]
	}

	return c
}
