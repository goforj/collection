package collection

import (
	"reflect"
	"testing"
)

func TestTap_InvokesCallback(t *testing.T) {
	called := false

	c := New([]int{3, 1, 2})

	out := c.Tap(func(col Collection[int]) {
		called = true

		// verify the collection we received is correct
		if !reflect.DeepEqual(col.items, c.items) {
			t.Fatalf("Tap received incorrect items: %v vs %v", col.items, c.items)
		}
	})

	if !called {
		t.Fatalf("Tap callback was not invoked")
	}

	// Tap must return the original collection unchanged
	if !reflect.DeepEqual(out.items, c.items) {
		t.Fatalf("Tap returned a modified collection: %v vs %v", out.items, c.items)
	}
}

func TestTap_Chainability(t *testing.T) {
	var captured []int

	out := New([]int{3, 1, 2}).
		Sort(func(a, b int) bool { return a < b }). // → [1,2,3]
		Tap(func(col Collection[int]) {
			captured = col.items
		}).
		Filter(func(v int) bool { return v >= 2 }) // → [2,3]

	if !reflect.DeepEqual(captured, []int{1, 2, 3}) {
		t.Fatalf("Tap did not receive correct snapshot: %v", captured)
	}

	if !reflect.DeepEqual(out.items, []int{2, 3}) {
		t.Fatalf("chain after Tap incorrect: %v", out.items)
	}
}

func TestTap_NoMutation(t *testing.T) {
	orig := []int{10, 20, 30}
	c := New(orig)

	c2 := c.Tap(func(col Collection[int]) {
		// do nothing
	})

	// ensure original slice unchanged
	if !reflect.DeepEqual(c.items, orig) {
		t.Fatalf("Tap mutated original collection: %v", c.items)
	}

	// ensure returned collection equals input
	if !reflect.DeepEqual(c2.items, orig) {
		t.Fatalf("Tap returned modified collection: %v", c2.items)
	}
}

func TestTap_Empty(t *testing.T) {
	c := New([]int{})

	called := false

	out := c.Tap(func(col Collection[int]) {
		called = true

		if len(col.items) != 0 {
			t.Fatalf("expected empty slice in Tap, got %v", col.items)
		}
	})

	if !called {
		t.Fatalf("Tap callback not executed for empty collection")
	}

	if len(out.items) != 0 {
		t.Fatalf("expected empty output, got %v", out.items)
	}
}

func TestTap_WithStructs(t *testing.T) {
	type User struct {
		ID   int
		Name string
	}

	c := New([]User{
		{1, "Chris"},
		{2, "Van"},
	})

	var captured []User

	out := c.Tap(func(col Collection[User]) {
		captured = col.items
	})

	if !reflect.DeepEqual(captured, c.items) {
		t.Fatalf("Tap did not receive correct struct slice: %v", captured)
	}

	if !reflect.DeepEqual(out.items, c.items) {
		t.Fatalf("Tap returned modified struct collection: %v", out.items)
	}
}
