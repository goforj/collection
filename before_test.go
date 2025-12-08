package collection

import (
	"testing"
)

func TestBefore_PredicateMatchSimpleTypes(t *testing.T) {
	nums := New([]int{1, 2, 3, 4, 5})

	// Before first element >= 3 → expect [1,2]
	before := nums.
		Before(func(v int) bool { return v >= 3 }).
		Items()

	expected := []int{1, 2}

	if len(before) != len(expected) {
		t.Fatalf("expected %d items, got %d: %#v", len(expected), len(before), before)
	}

	for i := range expected {
		if before[i] != expected[i] {
			t.Fatalf("expected before[%d] = %d, got %d", i, expected[i], before[i])
		}
	}
}

func TestBefore_PredicateMatchStructs(t *testing.T) {
	type User struct {
		ID   int
		Name string
		Age  int
	}

	users := New([]User{
		{1, "Chris", 34},
		{2, "Van", 42},
		{3, "Shawn", 39},
	})

	// Before first user with Age >= 40 → expect only Chris
	before := users.
		Before(func(u User) bool { return u.Age >= 40 }).
		Items()

	if len(before) != 1 {
		t.Fatalf("expected 1 user, got %d (%#v)", len(before), before)
	}
	if before[0].Name != "Chris" {
		t.Fatalf("expected Chris, got %#v", before[0])
	}
}

func TestBefore_NoPredicateMatchReturnsAll(t *testing.T) {
	nums := New([]int{10, 20, 30})

	before := nums.
		Before(func(v int) bool { return v > 100 }).
		Items()

	expected := []int{10, 20, 30}

	if len(before) != len(expected) {
		t.Fatalf("expected full slice, got %#v", before)
	}
}

func TestBefore_EmptyCollection(t *testing.T) {
	nums := New([]int{})

	before := nums.
		Before(func(v int) bool { return true }).
		Items()

	if len(before) != 0 {
		t.Fatalf("expected empty slice, got %#v", before)
	}
}

func TestBefore_SingleElement(t *testing.T) {
	nums := New([]int{5})

	// If match → resulting slice should be empty
	before := nums.
		Before(func(v int) bool { return v == 5 }).
		Items()

	if len(before) != 0 {
		t.Fatalf("expected empty slice, got %v", before)
	}

	// If no match → return full slice
	before = nums.
		Before(func(v int) bool { return v == 99 }).
		Items()

	if len(before) != 1 || before[0] != 5 {
		t.Fatalf("expected [5], got %v", before)
	}
}

func TestBefore_DoesNotMutateOriginal(t *testing.T) {
	original := []int{1, 2, 3, 4}
	nums := New(original)

	_ = nums.Before(func(v int) bool { return v >= 3 })

	// Ensure original slice untouched
	for i, v := range original {
		if nums.Items()[i] != v {
			t.Fatalf("original slice mutated: expected %v, got %v", original, nums.Items())
		}
	}
}
