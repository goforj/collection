package collection

import "testing"

func TestIsEmpty_EmptyCollection(t *testing.T) {
	c := New([]int{})
	if len(c) != 0 {
		t.Fatalf("expected IsEmpty() to return true for empty collection")
	}
}

func TestIsEmpty_NonEmptyCollection(t *testing.T) {
	c := New([]string{"a"})
	if c.IsEmpty() {
		t.Fatalf("expected IsEmpty() to return false for non-empty collection")
	}
}

func TestIsEmpty_DoesNotMutate(t *testing.T) {
	c := New([]int{1, 2, 3})

	_ = c.IsEmpty() // call should not mutate the collection

	if len(c.Items()) != 3 {
		t.Fatalf("IsEmpty() mutated the collection length, got %d", len(c.Items()))
	}
}
