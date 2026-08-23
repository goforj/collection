package collection

import (
	"reflect"
	"testing"
)

func TestTakeUntilFn_Basic(t *testing.T) {
	c := New([]int{1, 2, 3, 4})

	out := c.TakeUntilFn(func(v int) bool {
		return v >= 3
	})

	expected := []int{1, 2}
	if !reflect.DeepEqual(out.Items(), expected) {
		t.Fatalf("expected %v, got %v", expected, out.Items())
	}
}

func TestTakeUntilFn_NoMatch(t *testing.T) {
	c := New([]int{1, 2, 3})

	out := c.TakeUntilFn(func(v int) bool { return false })

	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(out.Items(), expected) {
		t.Fatalf("expected %v, got %v", expected, out.Items())
	}
}

func TestTakeUntilFn_FirstItemMatches(t *testing.T) {
	c := New([]int{1, 2, 3})

	out := c.TakeUntilFn(func(v int) bool { return v == 1 })

	if len(out.Items()) != 0 {
		t.Fatalf("expected empty slice, got %v", out.Items())
	}
}

func TestTakeUntilFn_EmptyCollection(t *testing.T) {
	c := New([]int{})

	out := c.TakeUntilFn(func(v int) bool { return true })

	if len(out.Items()) != 0 {
		t.Fatalf("expected empty slice, got %v", out.Items())
	}
}

func TestTakeUntil_ComparableValue(t *testing.T) {
	c := New([]int{1, 2, 3, 4})

	out := TakeUntil(c, 3)

	expected := []int{1, 2}
	if !reflect.DeepEqual(out.Items(), expected) {
		t.Fatalf("expected %v, got %v", expected, out.Items())
	}
}

func TestTakeUntil_NoValueMatch(t *testing.T) {
	c := New([]int{1, 2, 3})

	out := TakeUntil(c, 999)

	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(out.Items(), expected) {
		t.Fatalf("expected %v, got %v", expected, out.Items())
	}
}

func TestTakeUntil_ComparableStruct(t *testing.T) {
	type Point struct {
		X, Y int
	}

	c := New([]Point{
		{1, 1},
		{2, 2}, // stop here
		{3, 3},
	})

	out := TakeUntil(c, Point{2, 2})

	expected := []Point{{1, 1}}

	if !reflect.DeepEqual(out.Items(), expected) {
		t.Fatalf("expected %v, got %v", expected, out.Items())
	}
}

func TestTakeUntil_NoMutation(t *testing.T) {
	c := New([]int{1, 2, 3})
	orig := append([]int{}, c.Items()...)

	_ = TakeUntil(c, 2)
	_ = c.TakeUntilFn(func(v int) bool { return v == 1 })

	if !reflect.DeepEqual(c.Items(), orig) {
		t.Fatalf("TakeUntil mutated original slice: %v", c.Items())
	}
}

func TestTakeUntil_ReturnsView(t *testing.T) {
	c := New([]int{1, 2, 3})
	out := TakeUntil(c, 3)
	out[1] = 9
	if c[1] != 9 {
		t.Fatalf("TakeUntil mutation did not reach source: %v", c)
	}
}
