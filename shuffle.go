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

// Shuffle returns a shuffled copy of the collection.
// @group Ordering
// @behavior immutable
// @chainable true
// @terminal false
//
// This operation allocates a new slice and does not mutate the receiver.
//
// The shuffle uses an internal random source. Tests may override
// this source to achieve deterministic behavior.
//
// Example: integers
//
//	c := collection.New([]int{1, 2, 3, 4, 5})
//	out1 := c.Shuffle()
//	collection.Dump(out1.Items())
//
// Example: strings – chaining
//
//	out2 := collection.New([]string{"a", "b", "c"}).
//		Shuffle().
//		Append("d").
//		Items()
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
//	out3 := users.Shuffle()
//	collection.Dump(out3.Items())
func (c *Collection[T]) Shuffle() *Collection[T] {
	items := c.items
	out := make([]T, len(items))
	copy(out, items)
	n := len(out)

	// Fisher–Yates shuffle (in place on the copy)
	for i := n - 1; i > 0; i-- {
		j := shuffleRand.Intn(i + 1)
		out[i], out[j] = out[j], out[i]
	}

	return New(out)
}
