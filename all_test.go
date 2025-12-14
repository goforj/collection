package collection

import "testing"

func TestAll_AllMatch(t *testing.T) {
	c := New([]int{2, 4, 6, 8})

	ok := c.All(func(v int) bool {
		return v%2 == 0
	})

	if !ok {
		t.Fatalf("expected All to return true when all elements match")
	}
}

func TestAll_NotAllMatch(t *testing.T) {
	c := New([]int{2, 3, 4})

	ok := c.All(func(v int) bool {
		return v%2 == 0
	})

	if ok {
		t.Fatalf("expected All to return false when at least one element does not match")
	}
}

func TestAll_EmptyCollection(t *testing.T) {
	c := New([]int{})

	ok := c.All(func(v int) bool {
		return v > 0
	})

	if !ok {
		t.Fatalf("expected All to return true for empty collection (vacuously true)")
	}
}

func TestAll_ShortCircuits(t *testing.T) {
	calls := 0
	c := New([]int{1, 2, 3, 4})

	ok := c.All(func(v int) bool {
		calls++
		return v < 3 // fails at 3
	})

	if ok {
		t.Fatalf("expected All to return false")
	}

	if calls != 3 {
		t.Fatalf("expected All to short-circuit after 3 calls, got %d", calls)
	}
}

func TestAll_WithStructs(t *testing.T) {
	type User struct {
		Active bool
	}

	users := New([]User{
		{Active: true},
		{Active: true},
		{Active: false},
	})

	ok := users.All(func(u User) bool {
		return u.Active
	})

	if ok {
		t.Fatalf("expected All to return false when not all structs satisfy predicate")
	}
}
