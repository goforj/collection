package collection

import "testing"

func TestUnion_Ints(t *testing.T) {
	a := New([]int{1, 2, 2, 3})
	b := New([]int{3, 4, 4, 5})

	out := Union(a, b)

	exp := []int{1, 2, 3, 4, 5}
	if got := out.Items(); !slicesEqual(got, exp) {
		t.Fatalf("expected %v, got %v", exp, got)
	}
}

func TestUnion_Strings(t *testing.T) {
	left := New([]string{"apple", "banana"})
	right := New([]string{"banana", "date"})

	out := Union(left, right)

	exp := []string{"apple", "banana", "date"}
	if got := out.Items(); !slicesEqual(got, exp) {
		t.Fatalf("expected %v, got %v", exp, got)
	}
}

func TestUnion_FirstEmpty(t *testing.T) {
	left := New([]int{})
	right := New([]int{1, 2})

	out := Union(left, right)

	exp := []int{1, 2}
	if got := out.Items(); !slicesEqual(got, exp) {
		t.Fatalf("expected %v, got %v", exp, got)
	}
}
