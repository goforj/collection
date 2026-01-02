package collection

import (
	"reflect"
	"testing"
)

func TestSort_IntsAscending(t *testing.T) {
	c := New([]int{5, 3, 1, 4, 2})

	sorted := c.Sort(func(a, b int) bool {
		return a < b
	})

	expected := []int{1, 2, 3, 4, 5}

	if !reflect.DeepEqual(sorted.items, expected) {
		t.Fatalf("expected %v, got %v", expected, sorted.items)
	}

}

func TestSort_IntsDescending(t *testing.T) {
	c := New([]int{1, 4, 2, 5, 3})

	sorted := c.Sort(func(a, b int) bool {
		return a > b
	})

	expected := []int{5, 4, 3, 2, 1}

	if !reflect.DeepEqual(sorted.items, expected) {
		t.Fatalf("expected %v, got %v", expected, sorted.items)
	}
}

func TestSort_StructsByField(t *testing.T) {
	type User struct {
		ID   int
		Age  int
		Name string
	}

	c := New([]User{
		{1, 40, "Shawn"},
		{2, 25, "Chris"},
		{3, 30, "Van"},
	})

	sorted := c.Sort(func(a, b User) bool {
		return a.Age < b.Age
	})

	expected := []User{
		{2, 25, "Chris"},
		{3, 30, "Van"},
		{1, 40, "Shawn"},
	}

	if !reflect.DeepEqual(sorted.items, expected) {
		t.Fatalf("expected %v, got %v", expected, sorted.items)
	}
}

func TestSort_StableWhenEqual(t *testing.T) {
	type Item struct {
		ID    int
		Value int
	}

	// Value ties: 1, 1, 1 — original order must be preserved (Go sort.Slice is stable)
	c := New([]Item{
		{1, 10},
		{2, 10},
		{3, 10},
	})

	sorted := c.Sort(func(a, b Item) bool {
		return a.Value < b.Value
	})

	expected := []Item{
		{1, 10},
		{2, 10},
		{3, 10},
	}

	if !reflect.DeepEqual(sorted.items, expected) {
		t.Fatalf("expected stable ordering %v, got %v", expected, sorted.items)
	}
}

func TestSort_EmptyCollection(t *testing.T) {
	c := New([]int{})

	sorted := c.Sort(func(a, b int) bool { return a < b })

	if len(sorted.items) != 0 {
		t.Fatalf("expected empty slice, got %v", sorted.items)
	}
}

func TestSort_SingleElement(t *testing.T) {
	c := New([]int{42})

	sorted := c.Sort(func(a, b int) bool { return a < b })

	expected := []int{42}

	if !reflect.DeepEqual(sorted.items, expected) {
		t.Fatalf("expected %v, got %v", expected, sorted.items)
	}
}

func TestSort_PreservesNilSlice(t *testing.T) {
	c := New([]int(nil))

	c.Sort(func(a, b int) bool { return a < b })

	if c.Items() != nil {
		t.Fatalf("expected nil slice to remain nil, got %v", c.Items())
	}
}

func TestSort_WritesThroughSourceSlice(t *testing.T) {
	items := []int{3, 1, 2}
	c := New(items)

	c.Sort(func(a, b int) bool { return a < b })

	want := []int{1, 2, 3}
	if !reflect.DeepEqual(items, want) {
		t.Fatalf("expected source slice %v, got %v", want, items)
	}
}

func TestSort_LengthUnchanged(t *testing.T) {
	c := New([]int{3, 1, 2})

	c.Sort(func(a, b int) bool { return a < b })

	if len(c.Items()) != 3 {
		t.Fatalf("expected length 3, got %d", len(c.Items()))
	}
}
