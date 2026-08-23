package collection

import (
	"reflect"
	"testing"
)

func TestRetain_CompactsExistingBackingArray(t *testing.T) {
	items := []int{1, 2, 3, 4}
	values := New(items)
	first := &values[0]
	retained := values.Retain(func(value int) bool { return value%2 == 0 })
	if &retained[0] != first {
		t.Fatal("Retain did not retain the source backing array")
	}
	want := Slice[int]{2, 4}
	if !reflect.DeepEqual(retained, want) {
		t.Fatalf("Retain result = %v, want %v", retained, want)
	}
	wantSource := []int{2, 4, 0, 0}
	if !reflect.DeepEqual(items, wantSource) {
		t.Fatalf("Retain source storage = %v, want %v", items, wantSource)
	}
}

func TestRetain_ClearsDiscardedReferences(t *testing.T) {
	one, two := 1, 2
	items := []*int{&one, &two}
	retained := New(items).Retain(func(value *int) bool { return value == &two })
	if len(retained) != 1 || retained[0] != &two || items[1] != nil {
		t.Fatalf("Retain result=%v storage=%v", retained, items)
	}
}
