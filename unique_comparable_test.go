package collection

import "testing"

func TestUniqueComparable_Ints(t *testing.T) {
	c := New([]int{1, 2, 2, 3, 4, 4, 5})

	out := UniqueComparable(c)

	exp := []int{1, 2, 3, 4, 5}
	if got := out.Items(); !slicesEqual(got, exp) {
		t.Fatalf("expected %v, got %v", exp, got)
	}
}

func TestUniqueComparable_Strings(t *testing.T) {
	c := New([]string{"a", "b", "a", "c", "b"})

	out := UniqueComparable(c)

	exp := []string{"a", "b", "c"}
	if got := out.Items(); !slicesEqual(got, exp) {
		t.Fatalf("expected %v, got %v", exp, got)
	}
}

func TestUniqueComparable_Empty(t *testing.T) {
	c := New([]int{})

	out := UniqueComparable(c)

	if len(out.Items()) != 0 {
		t.Fatalf("expected empty result, got %v", out.Items())
	}
}
