package collection

import "testing"

func TestIndexWhere_Found(t *testing.T) {
	c := New([]int{10, 20, 30, 40})

	idx, ok := c.IndexWhere(func(v int) bool {
		return v == 30
	})

	if !ok {
		t.Fatalf("expected ok=true")
	}
	if idx != 2 {
		t.Fatalf("expected index 2, got %d", idx)
	}
}

func TestIndexWhere_NotFound(t *testing.T) {
	c := New([]int{1, 2, 3})

	idx, ok := c.IndexWhere(func(v int) bool {
		return v == 99
	})

	if ok {
		t.Fatalf("expected ok=false")
	}
	if idx != 0 {
		t.Fatalf("expected index 0 when not found, got %d", idx)
	}
}

func TestIndexWhere_Empty(t *testing.T) {
	c := New([]int{})

	idx, ok := c.IndexWhere(func(v int) bool {
		return true
	})

	if ok {
		t.Fatalf("expected ok=false for empty collection")
	}
	if idx != 0 {
		t.Fatalf("expected index 0, got %d", idx)
	}
}

func TestIndexWhere_Structs(t *testing.T) {
	type User struct {
		ID   int
		Name string
	}

	c := New([]User{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
		{ID: 3, Name: "Carol"},
	})

	idx, ok := c.IndexWhere(func(u User) bool {
		return u.ID == 2
	})

	if !ok {
		t.Fatalf("expected ok=true")
	}
	if idx != 1 {
		t.Fatalf("expected index 1, got %d", idx)
	}
}

func TestIndexWhere_ShortCircuits(t *testing.T) {
	calls := 0
	c := New([]int{1, 2, 3, 4, 5})

	_, _ = c.IndexWhere(func(v int) bool {
		calls++
		return v == 3
	})

	if calls != 3 {
		t.Fatalf("expected predicate to be called 3 times, got %d", calls)
	}
}

func TestIndexWhere_DoesNotMutate(t *testing.T) {
	items := []int{1, 2, 3}
	c := New(items)

	_, _ = c.IndexWhere(func(v int) bool {
		return v == 2
	})

	// Ensure original slice unchanged
	if items[0] != 1 || items[1] != 2 || items[2] != 3 {
		t.Fatalf("collection was mutated")
	}
}
