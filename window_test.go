package collection

import "testing"

func TestWindow_IntsStep1(t *testing.T) {
	c := New([]int{1, 2, 3, 4, 5})

	out := Window(c, 3, 1)

	exp := [][]int{
		{1, 2, 3},
		{2, 3, 4},
		{3, 4, 5},
	}

	if got := out.Items(); !doubleSlicesEqual(got, exp) {
		t.Fatalf("expected %v, got %v", exp, got)
	}
}

func TestWindow_StringsStep2(t *testing.T) {
	c := New([]string{"a", "b", "c", "d", "e"})

	out := Window(c, 2, 2)

	exp := [][]string{
		{"a", "b"},
		{"c", "d"},
	}

	if got := out.Items(); !doubleSlicesEqual(got, exp) {
		t.Fatalf("expected %v, got %v", exp, got)
	}
}

func TestWindow_Structs(t *testing.T) {
	type point struct {
		x int
		y int
	}

	c := New([]point{
		{x: 0, y: 0},
		{x: 1, y: 1},
		{x: 2, y: 4},
	})

	out := Window(c, 2, 1)

	exp := [][]point{
		{{x: 0, y: 0}, {x: 1, y: 1}},
		{{x: 1, y: 1}, {x: 2, y: 4}},
	}

	if got := out.Items(); !doubleSlicesEqual(got, exp) {
		t.Fatalf("expected %v, got %v", exp, got)
	}
}

func TestWindow_StepDefaultsToOne(t *testing.T) {
	c := New([]int{1, 2, 3})

	out := Window(c, 2, 0)

	exp := [][]int{{1, 2}, {2, 3}}
	if got := out.Items(); !doubleSlicesEqual(got, exp) {
		t.Fatalf("expected default step windows %v, got %v", exp, got)
	}
}

func TestWindow_SizeTooLargeOrZero(t *testing.T) {
	c := New([]int{1, 2})

	if got := Window(c, 3, 1).Items(); len(got) != 0 {
		t.Fatalf("expected empty for size > len, got %v", got)
	}

	if got := Window(c, 0, 1).Items(); len(got) != 0 {
		t.Fatalf("expected empty for size <= 0, got %v", got)
	}
}

func TestWindow_EmptyCollection(t *testing.T) {
	c := New([]int{})

	if got := Window(c, 2, 1).Items(); len(got) != 0 {
		t.Fatalf("expected empty for empty collection, got %v", got)
	}
}

func doubleSlicesEqual[T comparable](a, b [][]T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
