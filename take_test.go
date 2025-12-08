package collection

import (
	"reflect"
	"testing"
)

func TestTake_Positive(t *testing.T) {
	c := New([]int{0, 1, 2, 3, 4, 5})

	out := c.Take(3)

	expected := []int{0, 1, 2}
	if !reflect.DeepEqual(out.items, expected) {
		t.Fatalf("expected %v, got %v", expected, out.items)
	}
}

func TestTake_Negative(t *testing.T) {
	c := New([]int{0, 1, 2, 3, 4, 5})

	out := c.Take(-2)

	expected := []int{4, 5}
	if !reflect.DeepEqual(out.items, expected) {
		t.Fatalf("expected %v, got %v", expected, out.items)
	}
}

func TestTake_Zero(t *testing.T) {
	c := New([]int{1, 2, 3})

	out := c.Take(0)

	if len(out.items) != 0 {
		t.Fatalf("expected empty result, got %v", out.items)
	}
}

func TestTake_PositiveOvershoot(t *testing.T) {
	c := New([]int{1, 2, 3})

	out := c.Take(10)

	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(out.items, expected) {
		t.Fatalf("expected full slice %v, got %v", expected, out.items)
	}
}

func TestTake_NegativeOvershoot(t *testing.T) {
	c := New([]int{1, 2, 3})

	out := c.Take(-10)

	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(out.items, expected) {
		t.Fatalf("expected full slice %v, got %v", expected, out.items)
	}
}

func TestTake_EmptyCollection(t *testing.T) {
	c := New([]int{})

	out := c.Take(5)
	if len(out.items) != 0 {
		t.Fatalf("expected empty result, got %v", out.items)
	}
}

func TestTake_Structs(t *testing.T) {
	type User struct {
		ID int
	}

	c := New([]User{
		{1}, {2}, {3}, {4},
	})

	out := c.Take(-2)

	expected := []User{{3}, {4}}
	if !reflect.DeepEqual(out.items, expected) {
		t.Fatalf("expected %v, got %v", expected, out.items)
	}
}

func TestTake_NoMutation(t *testing.T) {
	c := New([]int{1, 2, 3, 4})
	orig := append([]int{}, c.items...)

	_ = c.Take(2)

	if !reflect.DeepEqual(c.items, orig) {
		t.Fatalf("Take mutated original collection: %v", c.items)
	}
}
