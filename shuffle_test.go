package collection

import (
	"math/rand"
	"reflect"
	"testing"
)

func withDeterministicShuffle(t *testing.T, seed int64, fn func()) {
	t.Helper()

	orig := shuffleRand
	setShuffleRand(rand.New(rand.NewSource(seed)))
	defer setShuffleRand(orig)

	fn()
}

func TestShuffle_Deterministic(t *testing.T) {
	withDeterministicShuffle(t, 42, func() {
		c := New([]int{1, 2, 3, 4, 5})

		c.Shuffle()
		first := append([]int(nil), c.Items()...) // snapshot

		// Shuffle again with the same deterministic RNG reset
		setShuffleRand(rand.New(rand.NewSource(42)))
		c2 := New([]int{1, 2, 3, 4, 5})
		c2.Shuffle()
		second := c2.Items()

		if !reflect.DeepEqual(first, second) {
			t.Fatalf("expected deterministic shuffle within same RNG, got %v vs %v", first, second)
		}
	})
}

func TestShuffle_PreservesAllElements(t *testing.T) {
	withDeterministicShuffle(t, 1, func() {
		orig := []int{1, 2, 3, 4, 5}
		c := New(orig)

		c.Shuffle()
		out := c.Items()

		if len(out) != len(orig) {
			t.Fatalf("expected length %d, got %d", len(orig), len(out))
		}

		seen := make(map[int]int)
		for _, v := range out {
			seen[v]++
		}

		for _, v := range orig {
			if seen[v] != 1 {
				t.Fatalf("element %d missing or duplicated", v)
			}
		}
	})
}

func TestShuffle_Empty(t *testing.T) {
	c := New([]int{})
	c.Shuffle()

	if len(c.Items()) != 0 {
		t.Fatalf("expected empty slice, got %v", c.Items())
	}
}

func TestShuffle_SingleElement(t *testing.T) {
	c := New([]int{42})
	c.Shuffle()

	expect := []int{42}
	if !reflect.DeepEqual(c.Items(), expect) {
		t.Fatalf("expected %v, got %v", expect, c.Items())
	}
}

func TestShuffle_Structs(t *testing.T) {
	withDeterministicShuffle(t, 7, func() {
		type User struct {
			ID int
		}

		c := New([]User{
			{ID: 1},
			{ID: 2},
			{ID: 3},
		})

		c.Shuffle()

		ids := map[int]bool{}
		for _, u := range c.Items() {
			ids[u.ID] = true
		}

		for i := 1; i <= 3; i++ {
			if !ids[i] {
				t.Fatalf("missing user ID %d", i)
			}
		}
	})
}

func TestShuffle_MutatesInPlace(t *testing.T) {
	withDeterministicShuffle(t, 99, func() {
		items := []int{1, 2, 3, 4}
		c := New(items)

		c.Shuffle()

		if reflect.DeepEqual(items, []int{1, 2, 3, 4}) {
			t.Fatalf("expected underlying slice to be shuffled in place")
		}
	})
}
