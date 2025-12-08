package collection

import "testing"

func TestLast_ReturnsLastElement(t *testing.T) {
	c := New([]int{1, 2, 3, 4})

	v, ok := c.Last()

	if !ok {
		t.Fatalf("expected ok == true, got false")
	}
	if v != 4 {
		t.Fatalf("expected last value 4, got %v", v)
	}
}

func TestLast_EmptyCollection(t *testing.T) {
	c := New([]int{})

	v, ok := c.Last()

	if ok {
		t.Fatalf("expected ok == false for empty collection, got true with %v", v)
	}
	if v != 0 { // zero-value for int
		t.Fatalf("expected zero-value (0), got %v", v)
	}
}

func TestLast_WithStructs(t *testing.T) {
	type item struct{ ID int }
	c := New([]item{{1}, {2}, {3}})

	v, ok := c.Last()

	if !ok {
		t.Fatalf("expected ok == true, got false")
	}
	if v.ID != 3 {
		t.Fatalf("expected struct with ID 3, got %+v", v)
	}
}

func TestLast_SingleElement(t *testing.T) {
	c := New([]string{"hello"})

	v, ok := c.Last()

	if !ok {
		t.Fatalf("expected ok == true, got false")
	}
	if v != "hello" {
		t.Fatalf("expected 'hello', got %v", v)
	}
}
