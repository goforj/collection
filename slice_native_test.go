package collection

import (
	"reflect"
	"testing"
)

func TestSliceNativeInterop(t *testing.T) {
	values := New([]int{1, 2, 3})
	if len(values) != 3 || values[1] != 2 {
		t.Fatalf("len/index = (%d, %d), want (3, 2)", len(values), values[1])
	}

	total := 0
	for _, value := range values {
		total += value
	}
	if total != 6 {
		t.Fatalf("range sum = %d, want 6", total)
	}

	acceptInts := func(items []int) int { return len(items) }
	if acceptInts(values) != 3 {
		t.Fatalf("Slice was not accepted as []int")
	}
}

func TestSliceZeroValueAndAliasing(t *testing.T) {
	var zero Slice[int]
	if zero != nil || len(zero) != 0 {
		t.Fatalf("zero value = %v, want nil empty slice", zero)
	}

	items := []int{1, 2}
	values := New(items)
	values[0] = 9
	if !reflect.DeepEqual(items, []int{9, 2}) {
		t.Fatalf("New did not preserve backing-slice aliasing: %v", items)
	}
}
