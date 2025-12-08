package collection

import (
	"reflect"
	"testing"
)

func TestFilter_Ints(t *testing.T) {
	c := New([]int{1, 2, 3, 4, 5})

	filtered := c.Filter(func(v int) bool {
		return v%2 == 0 // keep even
	})

	expected := []int{2, 4}

	if !reflect.DeepEqual(filtered.items, expected) {
		t.Fatalf("expected %v, got %v", expected, filtered.items)
	}

	// Ensure original is unchanged
	if !reflect.DeepEqual(c.items, []int{1, 2, 3, 4, 5}) {
		t.Fatalf("original collection should not be mutated: %v", c.items)
	}
}

func TestFilter_NoneMatch(t *testing.T) {
	c := New([]int{1, 2, 3})

	filtered := c.Filter(func(v int) bool { return v > 100 })

	if len(filtered.items) != 0 {
		t.Fatalf("expected empty result, got %v", filtered.items)
	}
}

func TestFilter_AllMatch(t *testing.T) {
	c := New([]int{1, 2, 3})

	filtered := c.Filter(func(v int) bool { return true })

	expected := []int{1, 2, 3}

	if !reflect.DeepEqual(filtered.items, expected) {
		t.Fatalf("expected %v, got %v", expected, filtered.items)
	}
}

func TestFilter_Structs(t *testing.T) {
	type User struct {
		ID   int
		Name string
	}

	c := New([]User{
		{1, "Chris"},
		{2, "Van"},
		{3, "Shawn"},
	})

	filtered := c.Filter(func(u User) bool {
		return u.ID >= 2
	})

	expected := []User{
		{2, "Van"},
		{3, "Shawn"},
	}

	if !reflect.DeepEqual(filtered.items, expected) {
		t.Fatalf("expected %v, got %v", expected, filtered.items)
	}

	// Ensure original data was not mutated
	orig := []User{
		{1, "Chris"},
		{2, "Van"},
		{3, "Shawn"},
	}

	if !reflect.DeepEqual(c.items, orig) {
		t.Fatalf("original collection was mutated: %v", c.items)
	}
}

func TestFilter_EmptyInput(t *testing.T) {
	c := New([]int{})

	filtered := c.Filter(func(v int) bool { return v%2 == 0 })

	if len(filtered.items) != 0 {
		t.Fatalf("expected empty slice, got %v", filtered.items)
	}
}

func TestFilter_Chaining(t *testing.T) {
	c := New([]int{1, 2, 3, 4, 5})

	result := c.
		Filter(func(v int) bool { return v > 1 }). // [2,3,4,5]
		Filter(func(v int) bool { return v%2 == 1 }) // [3,5]

	expected := []int{3, 5}

	if !reflect.DeepEqual(result.items, expected) {
		t.Fatalf("expected %v, got %v", expected, result.items)
	}
}
