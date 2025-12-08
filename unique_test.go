package collection

import (
	"reflect"
	"testing"
)

func TestUnique_Ints(t *testing.T) {
	c := New([]int{1, 2, 2, 3, 4, 4, 5})

	out := c.Unique(func(a, b int) bool { return a == b })

	expected := []int{1, 2, 3, 4, 5}

	if !reflect.DeepEqual(out.items, expected) {
		t.Fatalf("expected %v, got %v", expected, out.items)
	}

	// Ensure original not mutated
	if !reflect.DeepEqual(c.items, []int{1, 2, 2, 3, 4, 4, 5}) {
		t.Fatalf("original collection was mutated: %v", c.items)
	}
}

func TestUnique_NoDuplicates(t *testing.T) {
	c := New([]int{1, 2, 3})

	out := c.Unique(func(a, b int) bool { return a == b })

	expected := []int{1, 2, 3}

	if !reflect.DeepEqual(out.items, expected) {
		t.Fatalf("expected %v, got %v", expected, out.items)
	}
}

func TestUnique_Empty(t *testing.T) {
	c := New([]int{})

	out := c.Unique(func(a, b int) bool { return a == b })

	if len(out.items) != 0 {
		t.Fatalf("expected empty slice, got %v", out.items)
	}
}

func TestUnique_Structs(t *testing.T) {
	type User struct {
		ID   int
		Name string
	}

	c := New([]User{
		{1, "Chris"},
		{2, "Van"},
		{3, "Shawn"},
		{4, "Chris"}, // duplicate by Name, not by struct equality
	})

	out := c.Unique(func(a, b User) bool {
		return a.Name == b.Name
	})

	expected := []User{
		{1, "Chris"},
		{2, "Van"},
		{3, "Shawn"},
	}

	if !reflect.DeepEqual(out.items, expected) {
		t.Fatalf("expected %v, got %v", expected, out.items)
	}

	// ensure original isn't mutated
	orig := []User{
		{1, "Chris"},
		{2, "Van"},
		{3, "Shawn"},
		{4, "Chris"},
	}

	if !reflect.DeepEqual(c.items, orig) {
		t.Fatalf("original collection was mutated: %v", c.items)
	}
}

func TestUnique_FirstOccurrencePreserved(t *testing.T) {
	c := New([]int{2, 1, 2, 1, 3})

	out := c.Unique(func(a, b int) bool { return a == b })

	expected := []int{2, 1, 3}

	if !reflect.DeepEqual(out.items, expected) {
		t.Fatalf("expected %v, got %v", expected, out.items)
	}
}

func TestUnique_ReturnsCopy(t *testing.T) {
	c := New([]int{1, 2, 2})

	out := c.Unique(func(a, b int) bool { return a == b })

	// Mutate original to ensure out is a copy
	c.items[0] = 999

	if out.items[0] == 999 {
		t.Fatalf("Unique returned a slice that aliases the original; must be a copy")
	}
}

func TestUnique_Chaining(t *testing.T) {
	c := New([]int{1, 2, 2, 3, 3, 3})

	out := c.
		Unique(func(a, b int) bool { return a == b }).
		Filter(func(v int) bool { return v > 1 })

	expected := []int{2, 3}

	if !reflect.DeepEqual(out.items, expected) {
		t.Fatalf("expected %v after chain, got %v", expected, out.items)
	}
}
