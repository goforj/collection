package collection

import "testing"

func TestSymmetricDifference_Ints(t *testing.T) {
	a := New([]int{1, 2, 3, 3})
	b := New([]int{3, 4, 4, 5})

	out := SymmetricDifference(a, b)

	exp := []int{1, 2, 4, 5}
	if got := out.Items(); !slicesEqual(got, exp) {
		t.Fatalf("expected %v, got %v", exp, got)
	}
}

func TestSymmetricDifference_Strings(t *testing.T) {
	left := New([]string{"apple", "banana"})
	right := New([]string{"banana", "date"})

	out := SymmetricDifference(left, right)

	exp := []string{"apple", "date"}
	if got := out.Items(); !slicesEqual(got, exp) {
		t.Fatalf("expected %v, got %v", exp, got)
	}
}

func TestSymmetricDifference_NoOverlap(t *testing.T) {
	left := New([]int{1, 2})
	right := New([]int{3, 4})

	out := SymmetricDifference(left, right)

	exp := []int{1, 2, 3, 4}
	if got := out.Items(); !slicesEqual(got, exp) {
		t.Fatalf("expected %v, got %v", exp, got)
	}
}

func TestSymmetricDifference_EmptyLeft(t *testing.T) {
	left := New([]int{})
	right := New([]int{1, 1, 2})

	out := SymmetricDifference(left, right)

	exp := []int{1, 2}
	if got := out.Items(); !slicesEqual(got, exp) {
		t.Fatalf("expected %v, got %v", exp, got)
	}
}

func TestSymmetricDifference_DedupFromLeftOnly(t *testing.T) {
	left := New([]int{1, 1, 2})
	right := New([]int{})

	out := SymmetricDifference(left, right)

	exp := []int{1, 2}
	if got := out.Items(); !slicesEqual(got, exp) {
		t.Fatalf("expected %v, got %v", exp, got)
	}
}
