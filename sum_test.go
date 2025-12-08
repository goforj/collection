package collection

import "testing"

func TestSum_Ints(t *testing.T) {
	nums := New([]int{1, 2, 3, 4, 5})
	if got := Sum(nums); got != 15 {
		t.Fatalf("expected 15, got %v", got)
	}
}

func TestSumBy_Structs(t *testing.T) {
	type Row struct{ Foo int }
	rows := New([]Row{{10}, {20}})
	if got := SumBy(rows, func(r Row) int { return r.Foo }); got != 30 {
		t.Fatalf("expected 30, got %v", got)
	}
}
