package collection

import "testing"

func TestAt_ValidIndex(t *testing.T) {
	c := New([]int{10, 20, 30})

	v, ok := c.At(1)

	if !ok {
		t.Fatalf("expected ok=true")
	}
	if v != 20 {
		t.Fatalf("expected 20, got %d", v)
	}
}

func TestAt_FirstAndLast(t *testing.T) {
	c := New([]int{10, 20, 30})

	v1, ok1 := c.At(0)
	if !ok1 || v1 != 10 {
		t.Fatalf("expected first element 10, got %v", v1)
	}

	v2, ok2 := c.At(2)
	if !ok2 || v2 != 30 {
		t.Fatalf("expected last element 30, got %v", v2)
	}
}

func TestAt_OutOfBoundsHigh(t *testing.T) {
	c := New([]int{1, 2, 3})

	v, ok := c.At(3)

	if ok {
		t.Fatalf("expected ok=false")
	}
	if v != 0 {
		t.Fatalf("expected zero value, got %v", v)
	}
}

func TestAt_OutOfBoundsNegative(t *testing.T) {
	c := New([]int{1, 2, 3})

	v, ok := c.At(-1)

	if ok {
		t.Fatalf("expected ok=false")
	}
	if v != 0 {
		t.Fatalf("expected zero value, got %v", v)
	}
}

func TestAt_EmptyCollection(t *testing.T) {
	c := New([]int{})

	v, ok := c.At(0)

	if ok {
		t.Fatalf("expected ok=false")
	}
	if v != 0 {
		t.Fatalf("expected zero value, got %v", v)
	}
}

func TestAt_Structs(t *testing.T) {
	type User struct {
		ID   int
		Name string
	}

	c := New([]User{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
	})

	u, ok := c.At(1)

	if !ok {
		t.Fatalf("expected ok=true")
	}
	if u.Name != "Bob" {
		t.Fatalf("expected Bob, got %v", u.Name)
	}
}

func TestAt_DoesNotMutate(t *testing.T) {
	items := []int{1, 2, 3}
	c := New(items)

	_, _ = c.At(1)

	if items[0] != 1 || items[1] != 2 || items[2] != 3 {
		t.Fatalf("collection was mutated")
	}
}
