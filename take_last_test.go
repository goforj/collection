package collection

import (
	"reflect"
	"testing"
)

func TestTakeLast_Normal(t *testing.T) {
	c := New([]int{1, 2, 3, 4, 5})

	out := c.TakeLast(2).Items()

	expect := []int{4, 5}
	if !reflect.DeepEqual(out, expect) {
		t.Fatalf("expected %v, got %v", expect, out)
	}
}

func TestTakeLast_Zero(t *testing.T) {
	c := New([]int{1, 2, 3})

	out := c.TakeLast(0).Items()

	if len(out) != 0 {
		t.Fatalf("expected empty slice, got %v", out)
	}
}

func TestTakeLast_Negative(t *testing.T) {
	c := New([]int{1, 2, 3})

	out := c.TakeLast(-1).Items()

	if len(out) != 0 {
		t.Fatalf("expected empty slice, got %v", out)
	}
}

func TestTakeLast_ExactLength(t *testing.T) {
	c := New([]int{1, 2, 3})

	out := c.TakeLast(3).Items()

	expect := []int{1, 2, 3}
	if !reflect.DeepEqual(out, expect) {
		t.Fatalf("expected %v, got %v", expect, out)
	}
}

func TestTakeLast_BeyondLength(t *testing.T) {
	c := New([]int{1, 2, 3})

	out := c.TakeLast(10).Items()

	expect := []int{1, 2, 3}
	if !reflect.DeepEqual(out, expect) {
		t.Fatalf("expected %v, got %v", expect, out)
	}
}

func TestTakeLast_Structs(t *testing.T) {
	type User struct {
		ID int
	}

	c := New([]User{
		{ID: 1},
		{ID: 2},
		{ID: 3},
	})

	out := c.TakeLast(1).Items()

	expect := []User{
		{ID: 3},
	}

	if !reflect.DeepEqual(out, expect) {
		t.Fatalf("expected %+v, got %+v", expect, out)
	}
}

func TestTakeLast_DoesNotMutateOriginal(t *testing.T) {
	orig := []int{1, 2, 3, 4}
	c := New(orig)

	_ = c.TakeLast(2)

	if !reflect.DeepEqual(c.Items(), orig) {
		t.Fatalf("original collection was mutated")
	}
}

func TestTakeLast_ReusesBackingSlice(t *testing.T) {
	items := []int{1, 2, 3, 4}
	c := New(items)

	out := c.TakeLast(2)
	out.Items()[0] = 99

	if items[2] != 99 {
		t.Fatalf("expected backing slice to be reused")
	}
}
