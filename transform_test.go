package collection

import (
	"reflect"
	"testing"
)

func TestTransform_MutatesInPlaceAndChains(t *testing.T) {
	items := []int{1, 2, 3}
	values := New(items)
	first := &values[0]
	result := values.Transform(func(value int) int { return value * 2 })
	if &result[0] != first {
		t.Fatal("Transform did not retain the receiver backing array")
	}
	wantSource := []int{2, 4, 6}
	if !reflect.DeepEqual(items, wantSource) {
		t.Fatalf("source = %v, want %v", items, wantSource)
	}
	want := Slice[int]{2, 4, 6}
	if !reflect.DeepEqual(result, want) {
		t.Fatalf("result = %v, want %v", result, want)
	}
}

func TestTransform_PreservesNilSlice(t *testing.T) {
	var values Slice[int]
	if values.Transform(func(value int) int { return value }) != nil {
		t.Fatal("Transform(nil) should remain nil")
	}
}
