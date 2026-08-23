package collection

import (
	"reflect"
	"testing"
)

func TestTap_InvokesCallback(t *testing.T) {
	called := false

	c := New([]int{3, 1, 2})

	out := c.Tap(func(col Slice[int]) {
		called = true

		// verify the collection we received is correct
		if !reflect.DeepEqual(col.Items(), c.Items()) {
			t.Fatalf("Tap received incorrect items: %v vs %v", col.Items(), c.Items())
		}
	})

	if !called {
		t.Fatalf("Tap callback was not invoked")
	}

	// Tap must return the original collection unchanged
	if !reflect.DeepEqual(out.Items(), c.Items()) {
		t.Fatalf("Tap returned a modified collection: %v vs %v", out.Items(), c.Items())
	}
}

func TestTap_Chainability(t *testing.T) {
	var captured []int

	out := New([]int{3, 1, 2}).
		Sort(func(a, b int) bool { return a < b }). // → [1,2,3]
		Tap(func(col Slice[int]) {
			captured = append([]int(nil), col.Items()...) // snapshot
		}).
		Filter(func(v int) bool { return v >= 2 }) // → [2,3]

	if !reflect.DeepEqual(captured, []int{1, 2, 3}) {
		t.Fatalf("Tap did not receive correct snapshot: %v", captured)
	}

	if !reflect.DeepEqual(out.Items(), []int{2, 3}) {
		t.Fatalf("chain after Tap incorrect: %v", out.Items())
	}
}

func TestTap_NoMutation(t *testing.T) {
	orig := []int{10, 20, 30}
	c := New(orig)

	c2 := c.Tap(func(col Slice[int]) {
		// do nothing
	})

	// ensure original slice unchanged
	if !reflect.DeepEqual(c.Items(), orig) {
		t.Fatalf("Tap mutated original collection: %v", c.Items())
	}

	// ensure returned collection equals input
	if !reflect.DeepEqual(c2.Items(), orig) {
		t.Fatalf("Tap returned modified collection: %v", c2.Items())
	}
}

func TestTap_Empty(t *testing.T) {
	c := New([]int{})

	called := false

	out := c.Tap(func(col Slice[int]) {
		called = true

		if len(col.Items()) != 0 {
			t.Fatalf("expected empty slice in Tap, got %v", col.Items())
		}
	})

	if !called {
		t.Fatalf("Tap callback not executed for empty collection")
	}

	if len(out.Items()) != 0 {
		t.Fatalf("expected empty output, got %v", out.Items())
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

	out := c.Tap(func(col Slice[User]) {
		captured = col.Items()
	})

	if !reflect.DeepEqual(captured, c.Items()) {
		t.Fatalf("Tap did not receive correct struct slice: %v", captured)
	}

	if !reflect.DeepEqual(out.Items(), c.Items()) {
		t.Fatalf("Tap returned modified struct collection: %v", out.Items())
	}
}
