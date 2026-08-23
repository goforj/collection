package collection

import (
	"reflect"
	"strconv"
	"testing"
)

func TestMap_MapsToDifferentElementTypeWithoutMutatingSource(t *testing.T) {
	items := []int{1, 2, 3}
	values := New(items)
	mapped := values.Map(strconv.Itoa)
	want := Slice[string]{"1", "2", "3"}
	if !reflect.DeepEqual(mapped, want) {
		t.Fatalf("Map result = %v, want %v", mapped, want)
	}
	wantSource := []int{1, 2, 3}
	if !reflect.DeepEqual(items, wantSource) {
		t.Fatalf("Map mutated source = %v, want %v", items, wantSource)
	}
}

func TestMap_AllocatesForSameElementType(t *testing.T) {
	values := New([]int{1, 2, 3})
	mapped := values.Map(func(value int) int { return value * 2 })
	mapped[0] = 99
	want := Slice[int]{1, 2, 3}
	if !reflect.DeepEqual(values, want) {
		t.Fatalf("source = %v, want %v", values, want)
	}
}

func TestMap_EmptySliceAllocatesAnEmptyResult(t *testing.T) {
	values := New([]int(nil))
	mapped := values.Map(strconv.Itoa)
	if mapped == nil || len(mapped) != 0 {
		t.Fatalf("Map(nil) = %#v, want allocated empty Slice", mapped)
	}
}
