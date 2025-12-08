package collection

import (
	"testing"
)

func TestAny_BasicMatch(t *testing.T) {
	c := New([]int{1, 2, 3, 4})

	out := c.Any(func(v int) bool {
		return v%2 == 0 // even numbers
	})

	if !out {
		t.Fatalf("expected true, got false")
	}
}

func TestAny_NoMatch(t *testing.T) {
	c := New([]int{1, 3, 5})

	out := c.Any(func(v int) bool {
		return v%2 == 0
	})

	if out {
		t.Fatalf("expected false, got true")
	}
}

func TestAny_FirstElementMatches(t *testing.T) {
	c := New([]int{10, 2, 3})

	out := c.Any(func(v int) bool {
		return v == 10
	})

	if !out {
		t.Fatalf("expected true, got false")
	}
}

func TestAny_LastElementMatches(t *testing.T) {
	c := New([]int{1, 2, 3, 99})

	out := c.Any(func(v int) bool {
		return v == 99
	})

	if !out {
		t.Fatalf("expected true, got false")
	}
}

func TestAny_EmptyCollection(t *testing.T) {
	c := New([]int{})

	out := c.Any(func(v int) bool {
		return true
	})

	if out {
		t.Fatalf("expected false on empty collection, got true")
	}
}

func TestAny_Structs(t *testing.T) {
	type User struct {
		ID   int
		Name string
	}

	c := New([]User{
		{1, "Chris"},
		{2, "Van"},
		{3, "Shawn"},
	})

	out := c.Any(func(u User) bool {
		return u.Name == "Van"
	})

	if !out {
		t.Fatalf("expected true, got false")
	}
}

func TestAny_NoMutation(t *testing.T) {
	c := New([]int{1, 2, 3})
	original := append([]int{}, c.items...)

	_ = c.Any(func(v int) bool { return v == 2 })

	for i := range c.items {
		if c.items[i] != original[i] {
			t.Fatalf("Any mutated the underlying slice: %v vs %v", c.items, original)
		}
	}
}
