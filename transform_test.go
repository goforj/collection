package collection

import (
	"reflect"
	"testing"
)

func TestTransform_Ints(t *testing.T) {
	c := New([]int{1, 2, 3, 4, 5})

	c.Transform(func(v int) int {
		return v * 2
	})

	expected := []int{2, 4, 6, 8, 10}

	if !reflect.DeepEqual(c.items, expected) {
		t.Fatalf("expected %v, got %v", expected, c.items)
	}
}

func TestTransform_Structs(t *testing.T) {
	type User struct {
		ID   int
		Name string
	}

	c := New([]User{
		{1, "Chris"},
		{2, "Van"},
		{3, "Shawn"},
	})

	c.Transform(func(u User) User {
		u.Name = u.Name + "!"
		return u
	})

	expected := []User{
		{1, "Chris!"},
		{2, "Van!"},
		{3, "Shawn!"},
	}

	if !reflect.DeepEqual(c.items, expected) {
		t.Fatalf("expected %v, got %v", expected, c.items)
	}
}

func TestTransform_Empty(t *testing.T) {
	c := New([]int{})

	c.Transform(func(v int) int {
		return v * 10
	})

	if len(c.items) != 0 {
		t.Fatalf("expected empty slice, got %v", c.items)
	}
}

func TestTransform_MutationInPlace(t *testing.T) {
	c := New([]int{1, 2, 3})

	originalPtr := &c.items[0] // memory address check

	c.Transform(func(v int) int {
		return v + 1
	})

	// If transform is truly in-place, slice capacity and identity stay same
	if originalPtr != &c.items[0] {
		t.Fatalf("Transform should mutate in place, but underlying slice changed")
	}

	expected := []int{2, 3, 4}

	if !reflect.DeepEqual(c.items, expected) {
		t.Fatalf("expected %v, got %v", expected, c.items)
	}
}

func TestTransform_ChainingCompatibility(t *testing.T) {
	c := New([]int{1, 2, 3})

	c.Transform(func(v int) int { return v * 2 }) // -> [2,4,6]
	out := c.Filter(func(v int) bool { return v > 3 })

	expected := []int{4, 6}

	if !reflect.DeepEqual(out.items, expected) {
		t.Fatalf("expected %v, got %v", expected, out.items)
	}
}

func TestTransform_PreservesNilSlice(t *testing.T) {
	c := New([]int(nil))

	c.Transform(func(v int) int { return v + 1 })

	if c.Items() != nil {
		t.Fatalf("expected nil slice to remain nil, got %v", c.Items())
	}
}

func TestTransform_WritesThroughSourceSlice(t *testing.T) {
	items := []int{1, 2, 3}
	c := New(items)

	c.Transform(func(v int) int { return v + 10 })

	want := []int{11, 12, 13}
	if !reflect.DeepEqual(items, want) {
		t.Fatalf("expected source slice %v, got %v", want, items)
	}
}

func TestTransform_LengthUnchanged(t *testing.T) {
	c := New([]int{1, 2, 3})

	c.Transform(func(v int) int { return v + 1 })

	if len(c.Items()) != 3 {
		t.Fatalf("expected length 3, got %d", len(c.Items()))
	}
}
