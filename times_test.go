package collection

import (
	"reflect"
	"testing"
)

func TestTimes_Basic(t *testing.T) {
	c := Times(10, func(n int) int {
		return n * 9
	})

	expected := []int{9, 18, 27, 36, 45, 54, 63, 72, 81, 90}

	if !reflect.DeepEqual(c.items, expected) {
		t.Fatalf("expected %v, got %v", expected, c.items)
	}
}

func TestTimes_Zero(t *testing.T) {
	c := Times[int](0, func(n int) int { return n })

	if len(c.items) != 0 {
		t.Fatalf("expected empty collection, got %v", c.items)
	}
}

func TestTimes_Negative(t *testing.T) {
	c := Times[int](-5, func(n int) int { return n })

	if len(c.items) != 0 {
		t.Fatalf("expected empty collection, got %v", c.items)
	}
}

func TestTimes_Structs(t *testing.T) {
	type User struct {
		ID   int
		Name string
	}

	c := Times(3, func(n int) User {
		return User{ID: n, Name: "User"}
	})

	expected := []User{
		{1, "User"},
		{2, "User"},
		{3, "User"},
	}

	if !reflect.DeepEqual(c.items, expected) {
		t.Fatalf("expected %v, got %v", expected, c.items)
	}
}

func TestTimes_Large(t *testing.T) {
	c := Times(1000, func(n int) int {
		return n
	})

	if len(c.items) != 1000 {
		t.Fatalf("expected 1000 items, got %d", len(c.items))
	}

	// spot check
	if c.items[0] != 1 || c.items[999] != 1000 {
		t.Fatalf("indexing mismatch: first=%v last=%v", c.items[0], c.items[999])
	}
}
