package collection

import (
	"reflect"
	"testing"
)

func TestShufflePreservesAllElements(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	shuffled := New(items).Shuffle()
	seen := make(map[int]int, len(shuffled))
	for _, value := range shuffled {
		seen[value]++
	}
	for _, value := range []int{1, 2, 3, 4, 5} {
		if seen[value] != 1 {
			t.Fatalf("element %d missing or duplicated", value)
		}
	}
}

func TestShuffleMutatesInPlace(t *testing.T) {
	items := []int{1, 2, 3, 4}
	shuffled := New(items).Shuffle()
	if !reflect.DeepEqual(shuffled, Slice[int](items)) {
		t.Fatalf("Shuffle result = %v, want %v", shuffled, items)
	}
}

func TestShufflePreservesNilSlice(t *testing.T) {
	var values Slice[int]
	if values.Shuffle() != nil {
		t.Fatal("Shuffle should preserve a nil slice")
	}
}

func TestShuffleDoesNotAllocate(t *testing.T) {
	values := New(make([]int, 1_000))
	allocs := testing.AllocsPerRun(1_000, func() {
		values.Shuffle()
	})
	if allocs != 0 {
		t.Fatalf("Shuffle allocations = %v, want 0", allocs)
	}
}

func TestRandomIndexStaysWithinBounds(t *testing.T) {
	state := uint64(0)
	for range 100 {
		if index := randomIndex(&state, 8); index >= 8 {
			t.Fatalf("randomIndex(_, 8) = %d", index)
		}
		limit := uint64(1<<63 + 1)
		if index := randomIndex(&state, limit); index >= limit {
			t.Fatalf("randomIndex(_, %d) = %d", limit, index)
		}
	}
}
