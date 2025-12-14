package collection

import "testing"

func TestPartition_Ints(t *testing.T) {
	c := New([]int{1, 2, 3, 4, 5})

	evens, odds := c.Partition(func(n int) bool {
		return n%2 == 0
	})

	if got := evens.Items(); !slicesEqual(got, []int{2, 4}) {
		t.Fatalf("evens mismatch, got %v", got)
	}
	if got := odds.Items(); !slicesEqual(got, []int{1, 3, 5}) {
		t.Fatalf("odds mismatch, got %v", got)
	}
}

func TestPartition_Strings(t *testing.T) {
	c := New([]string{"go", "gopher", "rust", "ruby"})

	goWords, other := c.Partition(func(s string) bool {
		return len(s) >= 4
	})

	if got := goWords.Items(); !slicesEqual(got, []string{"gopher", "rust", "ruby"}) {
		t.Fatalf("first partition mismatch, got %v", got)
	}
	if got := other.Items(); !slicesEqual(got, []string{"go"}) {
		t.Fatalf("second partition mismatch, got %v", got)
	}
}

func TestPartition_Structs(t *testing.T) {
	type user struct {
		name   string
		active bool
	}

	c := New([]user{
		{name: "alice", active: true},
		{name: "bob", active: false},
		{name: "carol", active: true},
	})

	active, inactive := c.Partition(func(u user) bool {
		return u.active
	})

	expActive := []user{{name: "alice", active: true}, {name: "carol", active: true}}
	expInactive := []user{{name: "bob", active: false}}

	if got := active.Items(); !slicesEqual(got, expActive) {
		t.Fatalf("active mismatch, got %v", got)
	}
	if got := inactive.Items(); !slicesEqual(got, expInactive) {
		t.Fatalf("inactive mismatch, got %v", got)
	}
}

func TestPartition_Empty(t *testing.T) {
	c := New([]int{})

	left, right := c.Partition(func(n int) bool { return n > 0 })

	if len(left.Items()) != 0 || len(right.Items()) != 0 {
		t.Fatalf("expected both empty, got %v and %v", left.Items(), right.Items())
	}
}
