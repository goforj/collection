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

// Shuffle randomly shuffles the items in the collection in place
// and returns the same collection for chaining.
// @group Ordering
// @behavior immutable
// @chainable true
//
// This operation performs no allocations.
//
// The shuffle uses an internal random source. Tests may override
// this source to achieve deterministic behavior.
//
// Example: integers
//
//	c := collection.New([]int{1, 2, 3, 4, 5})
//	c.Shuffle()
//	collection.Dump(c.Items())
//
// Example: strings – chaining
//
//	out := collection.New([]string{"a", "b", "c"}).
//		Shuffle().
//		Append("d").
//		Items()
//
//	collection.Dump(out)
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
//	collection.Dump(users.Items())
func (c *Collection[T]) Shuffle() *Collection[T] {
	items := c.items
	n := len(items)

	// Fisher–Yates shuffle (in place, zero alloc)
	for i := n - 1; i > 0; i-- {
		j := shuffleRand.Intn(i + 1)
		items[i], items[j] = items[j], items[i]
	}

	return c
}
