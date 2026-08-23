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
