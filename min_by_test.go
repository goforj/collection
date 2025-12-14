package collection

import "testing"

func TestMinBy_Structs(t *testing.T) {
	type user struct {
		name  string
		score int
	}

	c := New([]user{
		{name: "alice", score: 10},
		{name: "bob", score: 5},
		{name: "carol", score: 8},
	})

	minUser, ok := MinBy(c, func(u user) int {
		return u.score
	})

	if !ok {
		t.Fatalf("expected ok to be true")
	}

	if minUser.name != "bob" || minUser.score != 5 {
		t.Fatalf("expected bob with score 5, got %+v", minUser)
	}
}

func TestMinBy_StringsByLength(t *testing.T) {
	c := New([]string{"strawberry", "fig", "banana"})

	minVal, ok := MinBy(c, func(s string) int {
		return len(s)
	})

	if !ok {
		t.Fatalf("expected ok to be true")
	}

	if minVal != "fig" {
		t.Fatalf("expected fig, got %q", minVal)
	}
}

func TestMinBy_Empty(t *testing.T) {
	c := New([]int{})

	minVal, ok := MinBy(c, func(v int) int {
		return v
	})

	if ok {
		t.Fatalf("expected ok to be false for empty collection")
	}

	if minVal != 0 {
		t.Fatalf("expected zero value, got %d", minVal)
	}
}

func TestMinBy_TiesReturnFirst(t *testing.T) {
	c := New([]int{3, 1, 2, 1})

	minVal, ok := MinBy(c, func(v int) int {
		return v
	})

	if !ok {
		t.Fatalf("expected ok to be true")
	}

	if minVal != 1 {
		t.Fatalf("expected first minimal value 1, got %d", minVal)
	}
}
