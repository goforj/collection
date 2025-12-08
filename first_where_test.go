package collection

import (
	"reflect"
	"testing"
)

func TestFirstWhere_Ints(t *testing.T) {
	c := New([]int{1, 2, 3, 4, 5})

	value, ok := c.FirstWhere(func(v int) bool {
		return v%2 == 0 // first even
	})

	if !ok {
		t.Fatalf("expected ok=true, got false")
	}
	if value != 2 {
		t.Fatalf("expected 2, got %d", value)
	}
}

func TestFirstWhere_NoMatch(t *testing.T) {
	c := New([]int{1, 3, 5})

	value, ok := c.FirstWhere(func(v int) bool {
		return v%2 == 0 // no evens
	})

	if ok {
		t.Fatalf("expected ok=false, got true (value=%v)", value)
	}

	// ensure returned value is zero for the type
	if value != 0 {
		t.Fatalf("expected zero value, got %v", value)
	}
}

func TestFirstWhere_EmptyCollection(t *testing.T) {
	c := New([]int{})

	value, ok := c.FirstWhere(func(v int) bool { return true })

	if ok {
		t.Fatalf("expected ok=false from empty collection, got true")
	}
	if value != 0 {
		t.Fatalf("expected zero value, got %v", value)
	}
}

func TestFirstWhere_Structs(t *testing.T) {
	type User struct {
		ID   int
		Name string
	}

	c := New([]User{
		{1, "Chris"},
		{2, "Van"},
		{3, "Shawn"},
	})

	value, ok := c.FirstWhere(func(u User) bool {
		return u.Name == "Van"
	})

	expected := User{2, "Van"}

	if !ok {
		t.Fatalf("expected ok=true, got false")
	}
	if !reflect.DeepEqual(value, expected) {
		t.Fatalf("expected %v, got %v", expected, value)
	}
}

func TestFirstWhere_OrderMatters(t *testing.T) {
	c := New([]int{4, 2, 2, 1})

	value, ok := c.FirstWhere(func(v int) bool {
		return v == 2
	})

	if !ok {
		t.Fatalf("expected ok=true, got false")
	}
	if value != 2 {
		t.Fatalf("expected first occurrence (2), got %v", value)
	}
}

func TestFirstWhere_NoMutation(t *testing.T) {
	c := New([]int{1, 2, 3})

	_, _ = c.FirstWhere(func(v int) bool { return v == 2 })

	// ensure original was not mutated
	if !reflect.DeepEqual(c.items, []int{1, 2, 3}) {
		t.Fatalf("original collection was mutated: %v", c.items)
	}
}

func TestFirstWhere_ChainingBehavior(t *testing.T) {
	c := New([]int{1, 2, 3, 4, 5})

	// typical Laravel-style chain:
	// first even number after filtering evens >= 4 → should be 4
	value, ok := c.
		Filter(func(v int) bool { return v%2 == 0 }). // [2, 4]
		FirstWhere(func(v int) bool { return v >= 4 }) // → 4

	if !ok {
		t.Fatalf("expected ok=true, got false")
	}
	if value != 4 {
		t.Fatalf("expected 4, got %v", value)
	}
}
