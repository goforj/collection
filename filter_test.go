package collection

import (
	"reflect"
	"testing"
)

func TestFilter_AllocatesWithoutMutatingSource(t *testing.T) {
	items := []int{1, 2, 3, 4}
	values := New(items)
	filtered := values.Filter(func(value int) bool { return value%2 == 0 })
	want := Slice[int]{2, 4}
	if !reflect.DeepEqual(filtered, want) {
		t.Fatalf("Filter result = %v, want %v", filtered, want)
	}
	wantSource := []int{1, 2, 3, 4}
	if !reflect.DeepEqual(items, wantSource) {
		t.Fatalf("Filter mutated source = %v, want %v", items, wantSource)
	}
}

func TestFilter_IgnoredResultLeavesSourceUnchanged(t *testing.T) {
	items := []int{1, 2, 3, 4}
	values := New(items)
	_ = values.Filter(func(value int) bool { return value%2 == 0 })
	want := []int{1, 2, 3, 4}
	if !reflect.DeepEqual(items, want) {
		t.Fatalf("ignored Filter mutated source = %v, want %v", items, want)
	}
}

func TestFilter_NilSliceProducesAllocatedEmptyResult(t *testing.T) {
	var values Slice[int]
	filtered := values.Filter(func(int) bool { return true })
	if filtered == nil || len(filtered) != 0 {
		t.Fatalf("Filter(nil) = %#v, want allocated empty Slice", filtered)
	}
}
