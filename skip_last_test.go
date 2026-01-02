package collection

import (
	"reflect"
	"testing"
)

func TestSkipLast_Normal(t *testing.T) {
	c := New([]int{1, 2, 3, 4, 5})

	out := c.SkipLast(2).Items()

	expect := []int{1, 2, 3}
	if !reflect.DeepEqual(out, expect) {
		t.Fatalf("expected %v, got %v", expect, out)
	}
}

func TestSkipLast_Zero(t *testing.T) {
	c := New([]int{1, 2, 3})

	out := c.SkipLast(0).Items()

	expect := []int{1, 2, 3}
	if !reflect.DeepEqual(out, expect) {
		t.Fatalf("expected %v, got %v", expect, out)
	}
}

func TestSkipLast_Negative(t *testing.T) {
	c := New([]int{1, 2, 3})

	out := c.SkipLast(-1).Items()

	expect := []int{1, 2, 3}
	if !reflect.DeepEqual(out, expect) {
		t.Fatalf("expected %v, got %v", expect, out)
	}
}

func TestSkipLast_ExactLength(t *testing.T) {
	c := New([]int{1, 2, 3})

	out := c.SkipLast(3).Items()

	if len(out) != 0 {
		t.Fatalf("expected empty slice, got %v", out)
	}
}

func TestSkipLast_BeyondLength(t *testing.T) {
	c := New([]int{1, 2, 3})

	out := c.SkipLast(10).Items()

	if len(out) != 0 {
		t.Fatalf("expected empty slice, got %v", out)
	}
}

func TestSkipLast_Structs(t *testing.T) {
	type User struct {
		ID int
	}

	c := New([]User{
		{ID: 1},
		{ID: 2},
		{ID: 3},
	})

	out := c.SkipLast(1).Items()

	expect := []User{
		{ID: 1},
		{ID: 2},
	}

	if !reflect.DeepEqual(out, expect) {
		t.Fatalf("expected %+v, got %+v", expect, out)
	}
}

func TestSkipLast_DoesNotMutateOriginal(t *testing.T) {
	orig := []int{1, 2, 3, 4}
	c := New(orig)

	_ = c.SkipLast(2)

	if !reflect.DeepEqual(c.Items(), orig) {
		t.Fatalf("original collection was mutated")
	}
}

func TestSkipLast_ReusesBackingSlice(t *testing.T) {
	items := []int{1, 2, 3, 4}
	c := Attach(items)

	out := c.SkipLast(1)
	out.Items()[0] = 99

	if items[0] != 99 {
		t.Fatalf("expected backing slice to be reused")
	}
}
