package collection

import (
	"reflect"
	"testing"
)

func TestMode_SimpleCase(t *testing.T) {
	nums := New([]int{1, 2, 2, 3})

	modes := Mode(nums)

	expected := []int{2}

	if !reflect.DeepEqual(modes, expected) {
		t.Fatalf("expected %v, got %v", expected, modes)
	}
}

func TestMode_MultipleModes_FirstSeenOrder(t *testing.T) {
	// 1 appears 2 times
	// 2 appears 2 times
	// 3 appears 1 time
	nums := New([]int{1, 2, 1, 2, 3})

	modes := Mode(nums)

	// The order must follow the first time each mode value was seen.
	expected := []int{1, 2}

	if !reflect.DeepEqual(modes, expected) {
		t.Fatalf("expected %v, got %v", expected, modes)
	}
}

func TestMode_AllUnique_ReturnsAllFirstSeenOrder(t *testing.T) {
	// All appear once → all values are modes
	nums := New([]int{10, 20, 30})

	modes := Mode(nums)

	expected := []int{10, 20, 30}

	if !reflect.DeepEqual(modes, expected) {
		t.Fatalf("expected %v, got %v", expected, modes)
	}
}

func TestMode_AllSame_ReturnsSingleValue(t *testing.T) {
	nums := New([]int{5, 5, 5, 5})

	modes := Mode(nums)

	expected := []int{5}

	if !reflect.DeepEqual(modes, expected) {
		t.Fatalf("expected %v, got %v", expected, modes)
	}
}

func TestMode_Empty_ReturnsNil(t *testing.T) {
	nums := New([]int{})

	modes := Mode(nums)

	if modes != nil && len(modes) != 0 {
		t.Fatalf("expected nil or empty slice, got %#v", modes)
	}
}

func TestMode_WithStrings(t *testing.T) {
	words := New([]string{"a", "b", "a", "c", "b", "b"})

	modes := Mode(words)

	expected := []string{"b"}

	if !reflect.DeepEqual(modes, expected) {
		t.Fatalf("expected %v, got %v", expected, modes)
	}
}

func TestMode_StringTieOrderPreserved(t *testing.T) {
	words := New([]string{"x", "y", "x", "y"})

	modes := Mode(words)

	expected := []string{"x", "y"}

	if !reflect.DeepEqual(modes, expected) {
		t.Fatalf("expected %v, got %v", expected, modes)
	}
}
