package collection

import "testing"

func TestSum_Ints(t *testing.T) {
	nums := New([]int{1, 2, 3, 4, 5})
	if got := Sum(nums); got != 15 {
		t.Fatalf("expected 15, got %v", got)
	}
}
