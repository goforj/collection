package collection

import "testing"

func TestDifference_Ints(t *testing.T) {
	a := New([]int{1, 2, 2, 3, 4})
	b := New([]int{2, 4})

	out := Difference(a, b)

	exp := []int{1, 3}
	if got := out.Items(); !slicesEqual(got, exp) {
		t.Fatalf("expected %v, got %v", exp, got)
	}
}

func TestDifference_Strings(t *testing.T) {
	left := New([]string{"apple", "banana", "cherry"})
	right := New([]string{"banana"})

	out := Difference(left, right)

	exp := []string{"apple", "cherry"}
	if got := out.Items(); !slicesEqual(got, exp) {
		t.Fatalf("expected %v, got %v", exp, got)
	}
}

func TestDifference_EmptySecond(t *testing.T) {
	left := New([]int{1, 2})
	right := New([]int{})

	out := Difference(left, right)

	exp := []int{1, 2}
	if got := out.Items(); !slicesEqual(got, exp) {
		t.Fatalf("expected %v, got %v", exp, got)
	}
}
