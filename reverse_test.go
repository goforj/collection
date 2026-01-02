package collection

import (
	"reflect"
	"testing"
)

func TestReverse_Integers(t *testing.T) {
	c := New([]int{1, 2, 3, 4})

	c.Reverse()

	expect := []int{4, 3, 2, 1}
	if !reflect.DeepEqual(c.Items(), expect) {
		t.Fatalf("expected %v, got %v", expect, c.Items())
	}
}

func TestReverse_OddLength(t *testing.T) {
	c := New([]int{1, 2, 3})

	c.Reverse()

	expect := []int{3, 2, 1}
	if !reflect.DeepEqual(c.Items(), expect) {
		t.Fatalf("expected %v, got %v", expect, c.Items())
	}
}

func TestReverse_Empty(t *testing.T) {
	c := New([]int{})

	c.Reverse()

	if len(c.Items()) != 0 {
		t.Fatalf("expected empty slice, got %v", c.Items())
	}
}

func TestReverse_SingleElement(t *testing.T) {
	c := New([]int{42})

	c.Reverse()

	expect := []int{42}
	if !reflect.DeepEqual(c.Items(), expect) {
		t.Fatalf("expected %v, got %v", expect, c.Items())
	}
}

func TestReverse_Structs(t *testing.T) {
	type User struct {
		ID int
	}

	c := New([]User{
		{ID: 1},
		{ID: 2},
		{ID: 3},
	})

	c.Reverse()

	expect := []User{
		{ID: 3},
		{ID: 2},
		{ID: 1},
	}

	if !reflect.DeepEqual(c.Items(), expect) {
		t.Fatalf("expected %+v, got %+v", expect, c.Items())
	}
}

func TestReverse_Chainable(t *testing.T) {
	out := New([]int{1, 2, 3}).
		Reverse().
		Append(4).
		Items()

	expect := []int{3, 2, 1, 4}
	if !reflect.DeepEqual(out, expect) {
		t.Fatalf("expected %v, got %v", expect, out)
	}
}

func TestReverse_MutatesInPlace(t *testing.T) {
	items := []int{1, 2, 3}
	c := New(items)

	c.Reverse()

	// underlying slice should be reversed
	expect := []int{3, 2, 1}
	if !reflect.DeepEqual(items, expect) {
		t.Fatalf("expected underlying slice to be mutated, got %v", items)
	}
}

func TestReverse_PreservesNilSlice(t *testing.T) {
	c := New([]int(nil))

	c.Reverse()

	if c.Items() != nil {
		t.Fatalf("expected nil slice to remain nil, got %v", c.Items())
	}
}
