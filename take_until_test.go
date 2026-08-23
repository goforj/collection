package collection

import (
	"reflect"
	"testing"
)

func TestTakeUntil(t *testing.T) {
	tests := []struct {
		name string
		pred func(int) bool
		want Slice[int]
	}{
		{"match", func(value int) bool { return value >= 3 }, []int{1, 2}},
		{"no match", func(int) bool { return false }, []int{1, 2, 3, 4}},
		{"first matches", func(value int) bool { return value == 1 }, []int{}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := New([]int{1, 2, 3, 4}).TakeUntil(test.pred)
			if !reflect.DeepEqual(got, test.want) {
				t.Fatalf("TakeUntil() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestTakeUntilEmptyAndCappedView(t *testing.T) {
	if got := New([]int{}).TakeUntil(func(int) bool { return true }); len(got) != 0 || cap(got) != 0 {
		t.Fatalf("TakeUntil() = len %d cap %d, want capped empty", len(got), cap(got))
	}
	items := []int{1, 2, 3, 4}
	got := New(items).TakeUntil(func(value int) bool { return value == 3 })
	got[0] = 99
	if items[0] != 99 || cap(got) != len(got) {
		t.Fatalf("TakeUntil() did not return a capped view: %v cap %d", got, cap(got))
	}
}
