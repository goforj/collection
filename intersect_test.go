package collection

import "testing"

func TestIntersect_Ints(t *testing.T) {
	a := New([]int{1, 2, 2, 3, 4})
	b := New([]int{2, 4, 4, 5})

	out := Intersect(a, b)

	exp := []int{2, 4}
	if got := out.Items(); !slicesEqual(got, exp) {
		t.Fatalf("expected %v, got %v", exp, got)
	}
}

func TestIntersect_Strings(t *testing.T) {
	left := New([]string{"apple", "banana", "cherry"})
	right := New([]string{"banana", "date", "cherry", "banana"})

	out := Intersect(left, right)

	exp := []string{"banana", "cherry"}
	if got := out.Items(); !slicesEqual(got, exp) {
		t.Fatalf("expected %v, got %v", exp, got)
	}
}

func TestIntersect_Empty(t *testing.T) {
	a := New([]int{})
	b := New([]int{1, 2})

	out := Intersect(a, b)

	if len(out.Items()) != 0 {
		t.Fatalf("expected empty collection, got %v", out.Items())
	}
}

func TestIntersect_DuplicateInFirstOnlyOnce(t *testing.T) {
	a := New([]int{1, 1, 1, 2})
	b := New([]int{1, 2})

	out := Intersect(a, b)

	exp := []int{1, 2}
	if got := out.Items(); !slicesEqual(got, exp) {
		t.Fatalf("expected %v, got %v", exp, got)
	}
}

func slicesEqual[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
