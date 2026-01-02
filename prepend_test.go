package collection

import (
	"reflect"
	"testing"
)

func TestPrepend_Basic(t *testing.T) {
	c := New([]int{3, 4})

	out := c.Prepend(1, 2)

	if out != c {
		t.Fatalf("expected Prepend to return same collection instance")
	}

	expected := []int{1, 2, 3, 4}
	if !reflect.DeepEqual(c.items, expected) {
		t.Fatalf("expected %v, got %v", expected, c.items)
	}
}

func TestPrepend_EmptyCollection(t *testing.T) {
	c := New([]int{})

	out := c.Prepend(1, 2, 3)

	if out != c {
		t.Fatalf("expected Prepend to return same collection instance")
	}

	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(c.items, expected) {
		t.Fatalf("expected %v, got %v", expected, c.items)
	}
}

func TestPrepend_NoValues(t *testing.T) {
	c := New([]int{1, 2, 3})

	out := c.Prepend() // no-op

	if out != c {
		t.Fatalf("expected Prepend to return same collection instance")
	}

	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(c.items, expected) {
		t.Fatalf("expected %v, got %v", expected, c.items)
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

	if out != c {
		t.Fatalf("expected Prepend to return same collection instance")
	}

	expected := []User{
		{1, "Chris"},
		{2, "Matt"},
		{3, "Shawn"},
		{4, "Van"},
	}

	if !reflect.DeepEqual(c.items, expected) {
		t.Fatalf("expected %v, got %v", expected, c.items)
	}
}

func TestPrepend_DoesNotMutateSourceSlice(t *testing.T) {
	orig := []int{10, 20, 30}
	c := New(orig)

	out := c.Prepend(5, 7)

	if out != c {
		t.Fatalf("expected Prepend to return same collection instance")
	}

	if !reflect.DeepEqual(orig, []int{10, 20, 30}) {
		t.Fatalf("Prepend mutated source slice: %v", orig)
	}

	expected := []int{5, 7, 10, 20, 30}
	if !reflect.DeepEqual(c.items, expected) {
		t.Fatalf("expected %v, got %v", expected, c.items)
	}
}

func TestPrepend_NilSliceWithValues(t *testing.T) {
	c := New([]int(nil))

	c.Prepend(1, 2)

	expected := []int{1, 2}
	if !reflect.DeepEqual(c.Items(), expected) {
		t.Fatalf("expected %v, got %v", expected, c.Items())
	}
}

func TestPrepend_NilSliceNoValuesBecomesEmpty(t *testing.T) {
	c := New([]int(nil))

	c.Prepend()

	if c.Items() == nil {
		t.Fatalf("expected empty slice, got nil")
	}
	if len(c.Items()) != 0 {
		t.Fatalf("expected empty slice, got %v", c.Items())
	}
}
