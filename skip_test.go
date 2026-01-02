package collection

import (
	"reflect"
	"testing"
)

func TestSkip_Normal(t *testing.T) {
	c := New([]int{1, 2, 3, 4, 5})

	out := c.Skip(2).Items()

	expect := []int{3, 4, 5}
	if !reflect.DeepEqual(out, expect) {
		t.Fatalf("expected %v, got %v", expect, out)
	}
}

func TestSkip_Zero(t *testing.T) {
	c := New([]int{1, 2, 3})

	out := c.Skip(0).Items()

	expect := []int{1, 2, 3}
	if !reflect.DeepEqual(out, expect) {
		t.Fatalf("expected %v, got %v", expect, out)
	}
}

func TestSkip_Negative(t *testing.T) {
	c := New([]int{1, 2, 3})

	out := c.Skip(-1).Items()

	expect := []int{1, 2, 3}
	if !reflect.DeepEqual(out, expect) {
		t.Fatalf("expected %v, got %v", expect, out)
	}
}

func TestSkip_ExactLength(t *testing.T) {
	c := New([]int{1, 2, 3})

	out := c.Skip(3).Items()

	if len(out) != 0 {
		t.Fatalf("expected empty slice, got %v", out)
	}
}

func TestSkip_BeyondLength(t *testing.T) {
	c := New([]int{1, 2, 3})

	out := c.Skip(10).Items()

	if len(out) != 0 {
		t.Fatalf("expected empty slice, got %v", out)
	}
}

func TestSkip_Structs(t *testing.T) {
	type User struct {
		ID int
	}

	c := New([]User{
		{ID: 1},
		{ID: 2},
		{ID: 3},
	})

	out := c.Skip(1).Items()

	expect := []User{
		{ID: 2},
		{ID: 3},
	}

	if !reflect.DeepEqual(out, expect) {
		t.Fatalf("expected %+v, got %+v", expect, out)
	}
}

func TestSkip_DoesNotMutateOriginal(t *testing.T) {
	orig := []int{1, 2, 3, 4}
	c := New(orig)

	_ = c.Skip(2)

	if !reflect.DeepEqual(c.Items(), orig) {
		t.Fatalf("original collection was mutated")
	}
}

func TestSkip_ReusesBackingSlice(t *testing.T) {
	items := []int{1, 2, 3, 4}
	c := New(items)

	out := c.Skip(1)

	// Mutating the skipped collection should reflect in the original backing slice
	out.Items()[0] = 99

	if items[1] != 99 {
		t.Fatalf("expected backing slice to be reused")
	}
}
