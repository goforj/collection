package collection

import "testing"

func TestMaxBy_Structs(t *testing.T) {
	type player struct {
		name  string
		score int
	}

	c := New([]player{
		{name: "alice", score: 10},
		{name: "bob", score: 25},
		{name: "carol", score: 18},
	})

	maxPlayer, ok := MaxBy(c, func(p player) int {
		return p.score
	})

	if !ok {
		t.Fatalf("expected ok to be true")
	}

	if maxPlayer.name != "bob" || maxPlayer.score != 25 {
		t.Fatalf("expected bob with score 25, got %+v", maxPlayer)
	}
}

func TestMaxBy_StringsByLength(t *testing.T) {
	c := New([]string{"go", "collection", "rocks"})

	maxVal, ok := MaxBy(c, func(s string) int {
		return len(s)
	})

	if !ok {
		t.Fatalf("expected ok to be true")
	}

	if maxVal != "collection" {
		t.Fatalf("expected collection, got %q", maxVal)
	}
}

func TestMaxBy_Empty(t *testing.T) {
	c := New([]int{})

	maxVal, ok := MaxBy(c, func(v int) int {
		return v
	})

	if ok {
		t.Fatalf("expected ok to be false for empty collection")
	}

	if maxVal != 0 {
		t.Fatalf("expected zero value, got %d", maxVal)
	}
}

func TestMaxBy_TiesReturnFirst(t *testing.T) {
	c := New([]int{1, 3, 2, 3})

	maxVal, ok := MaxBy(c, func(v int) int {
		return v
	})

	if !ok {
		t.Fatalf("expected ok to be true")
	}

	if maxVal != 3 {
		t.Fatalf("expected first maximal value 3, got %d", maxVal)
	}
}
