package collection

import (
	"math/bits"
	"math/rand/v2"
)

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
	state := rand.Uint64()
	for i := len(c) - 1; i > 0; i-- {
		j := int(randomIndex(&state, uint64(i+1)))
		c[i], c[j] = c[j], c[i]
	}

	return c
}

// randomIndex reduces a random value into [0, n) without modulo bias.
func randomIndex(state *uint64, n uint64) uint64 {
	if n&(n-1) == 0 {
		return nextRandom(state) & (n - 1)
	}

	high, low := bits.Mul64(nextRandom(state), n)
	if low < n {
		threshold := -n % n
		for low < threshold {
			high, low = bits.Mul64(nextRandom(state), n)
		}
	}
	return high
}

// nextRandom advances a local SplitMix64 stream so Shuffle avoids shared-source contention.
func nextRandom(state *uint64) uint64 {
	*state += 0x9e3779b97f4a7c15
	value := *state
	value = (value ^ (value >> 30)) * 0xbf58476d1ce4e5b9
	value = (value ^ (value >> 27)) * 0x94d049bb133111eb
	return value ^ (value >> 31)
}
