package collection

import "testing"

type namedInts []int

func TestNumericFunctionsInferSliceLikeInputs(t *testing.T) {
	assertNumericFunctions(t, []int{1, 2, 3})
	assertNumericFunctions(t, New([]int{1, 2, 3}))
	assertNumericFunctions(t, namedInts{1, 2, 3})
}

func assertNumericFunctions[S ~[]T, T Number](t *testing.T, values S) {
	t.Helper()
	if Sum(values) != T(6) || Avg(values) != 2 {
		t.Fatalf("Sum/Avg(%v) = (%v, %v)", values, Sum(values), Avg(values))
	}
	if min, ok := Min(values); !ok || min != T(1) {
		t.Fatalf("Min(%v) = (%v, %v)", values, min, ok)
	}
	if max, ok := Max(values); !ok || max != T(3) {
		t.Fatalf("Max(%v) = (%v, %v)", values, max, ok)
	}
	if median, ok := Median(values); !ok || median != 2 {
		t.Fatalf("Median(%v) = (%v, %v)", values, median, ok)
	}
	if modes := Mode(values); len(modes) != 3 {
		t.Fatalf("Mode(%v) = %v", values, modes)
	}
}
