package collection

import (
	"reflect"
	"testing"
)

func TestPrepend_Basic(t *testing.T) {
	c := New([]int{3, 4})

	out := c.Prepend(1, 2)

	expected := []int{1, 2, 3, 4}
	if !reflect.DeepEqual(out.items, expected) {
		t.Fatalf("expected %v, got %v", expected, out.items)
	}
}

func TestPrepend_EmptyCollection(t *testing.T) {
	c := New([]int{})

	out := c.Prepend(1, 2, 3)

	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(out.items, expected) {
		t.Fatalf("expected %v, got %v", expected, out.items)
	}
}

func TestPrepend_NoValues(t *testing.T) {
	c := New([]int{1, 2, 3})

	out := c.Prepend() // no-op

	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(out.items, expected) {
		t.Fatalf("expected %v, got %v", expected, out.items)
	}
}

func TestPrepend_Structs(t *testing.T) {
	type User struct {
		ID   int
		Name string
	}

	c := New([]User{
		{3, "Shawn"},
		{4, "Van"},
	})

	out := c.Prepend(User{1, "Chris"}, User{2, "Matt"})

	expected := []User{
		{1, "Chris"},
		{2, "Matt"},
		{3, "Shawn"},
		{4, "Van"},
	}

	if !reflect.DeepEqual(out.items, expected) {
		t.Fatalf("expected %v, got %v", expected, out.items)
	}
}

func TestPrepend_NoMutation(t *testing.T) {
	orig := []int{10, 20, 30}
	c := New(orig)

	out := c.Prepend(5, 7)

	// original must remain unchanged
	if !reflect.DeepEqual(c.items, orig) {
		t.Fatalf("Prepend mutated original collection: %v", c.items)
	}

	// ensure new collection is different & correct
	expected := []int{5, 7, 10, 20, 30}
	if !reflect.DeepEqual(out.items, expected) {
		t.Fatalf("expected %v, got %v", expected, out.items)
	}
}
