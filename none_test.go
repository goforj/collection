package collection

import "testing"

func TestNone_NoMatches(t *testing.T) {
	c := New([]int{1, 3, 5})

	ok := c.None(func(v int) bool {
		return v%2 == 0
	})

	if !ok {
		t.Fatalf("expected None to return true when no elements match")
	}
}

func TestNone_WithMatch(t *testing.T) {
	c := New([]int{1, 2, 3})

	ok := c.None(func(v int) bool {
		return v%2 == 0
	})

	if ok {
		t.Fatalf("expected None to return false when an element matches")
	}
}

func TestNone_EmptyCollection(t *testing.T) {
	c := New([]int{})

	ok := c.None(func(v int) bool {
		return v > 0
	})

	if !ok {
		t.Fatalf("expected None to return true for empty collection")
	}
}

func TestNone_ShortCircuits(t *testing.T) {
	calls := 0
	c := New([]int{1, 2, 3, 4})

	ok := c.None(func(v int) bool {
		calls++
		return v == 2
	})

	if ok {
		t.Fatalf("expected None to return false")
	}

	if calls != 2 {
		t.Fatalf("expected None to short-circuit after 2 calls, got %d", calls)
	}
}

func TestNone_WithStructs(t *testing.T) {
	type User struct {
		Admin bool
	}

	users := New([]User{
		{Admin: false},
		{Admin: false},
		{Admin: true},
	})

	ok := users.None(func(u User) bool {
		return u.Admin
	})

	if ok {
		t.Fatalf("expected None to return false when a struct matches predicate")
	}
}
