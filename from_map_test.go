package collection

import (
	"reflect"
	"testing"
)

func TestFromMap(t *testing.T) {
	input := map[string]int{"a": 1, "b": 2, "c": 3}
	got := FromMap(input)
	seen := make(map[string]int, len(got))
	for _, pair := range got {
		seen[pair.First] = pair.Second
	}
	if !reflect.DeepEqual(seen, input) {
		t.Fatalf("FromMap() = %v, want pairs for %v", got, input)
	}
	if cap(got) < len(input) {
		t.Fatalf("FromMap() capacity = %d, want at least %d", cap(got), len(input))
	}
}

func TestFromMapDoesNotMutateInput(t *testing.T) {
	input := map[string]int{"a": 1, "b": 2}
	_ = FromMap(input)
	if !reflect.DeepEqual(input, map[string]int{"a": 1, "b": 2}) {
		t.Fatalf("FromMap() mutated input: %v", input)
	}
}

func TestFromMapEmptyAndStructValues(t *testing.T) {
	if got := FromMap(map[string]int{}); len(got) != 0 {
		t.Fatalf("FromMap(empty) = %v, want empty", got)
	}
	type user struct{ id int }
	input := map[string]user{"alice": {1}, "bob": {2}}
	seen := map[string]user{}
	for _, pair := range FromMap(input) {
		seen[pair.First] = pair.Second
	}
	if !reflect.DeepEqual(seen, input) {
		t.Fatalf("FromMap() = %v, want %v", seen, input)
	}
}
