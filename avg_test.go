package collection

import "testing"

func TestAvg_Ints(t *testing.T) {
	nums := New([]int{1, 1, 2, 4})

	avg := Avg(nums)

	if avg != 2 {
		t.Fatalf("expected 2, got %v", avg)
	}
}

func TestAvg_Floats(t *testing.T) {
	nums := New([]float64{1.5, 2.5, 3.0})

	avg := Avg(nums)

	if avg != 2.3333333333333335 {
		t.Fatalf("expected ~2.3333, got %v", avg)
	}
}

func TestAvg_EmptySlice(t *testing.T) {
	nums := New([]int{})

	avg := Avg(nums)

	if avg != 0 {
		t.Fatalf("expected 0 for empty slice, got %v", avg)
	}
}

func TestAvg_SingleValue(t *testing.T) {
	nums := New([]int{10})

	avg := Avg(nums)

	if avg != 10 {
		t.Fatalf("expected 10, got %v", avg)
	}
}
