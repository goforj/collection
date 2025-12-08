package collection

import "testing"

func TestIsEmpty_EmptyCollection(t *testing.T) {
	c := New([]int{})
	if !c.IsEmpty() {
		t.Fatalf("expected IsEmpty() to return true for empty collection")
	}
}

func TestIsEmpty_NonEmptyCollection(t *testing.T) {
	c := New([]string{"a"})
	if c.IsEmpty() {
		t.Fatalf("expected IsEmpty() to return false for non-empty collection")
	}
}

func TestIsEmpty_AfterNewNumeric(t *testing.T) {
	c := NewNumeric([]int{})
	if !c.IsEmpty() {
		t.Fatalf("expected IsEmpty() to return true for empty numeric collection")
	}

	c2 := NewNumeric([]int{10})
	if c2.IsEmpty() {
		t.Fatalf("expected IsEmpty() to return false for non-empty numeric collection")
	}
}

func TestIsEmpty_DoesNotMutate(t *testing.T) {
	c := New([]int{1, 2, 3})

	_ = c.IsEmpty() // call should not mutate the collection

	if len(c.items) != 3 {
		t.Fatalf("IsEmpty() mutated the collection length, got %d", len(c.items))
	}
}
