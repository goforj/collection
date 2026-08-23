package collection

import (
	"reflect"
	"testing"
)

func TestTakeLast(t *testing.T) {
	tests := []struct {
		n    int
		want Slice[int]
	}{
		{2, []int{4, 5}}, {0, []int{}}, {-1, []int{}}, {5, []int{1, 2, 3, 4, 5}}, {10, []int{1, 2, 3, 4, 5}},
	}
	for _, test := range tests {
		got := New([]int{1, 2, 3, 4, 5}).TakeLast(test.n)
		if !reflect.DeepEqual(got, test.want) {
			t.Fatalf("TakeLast(%d) = %v, want %v", test.n, got, test.want)
		}
	}
}

func TestTakeLastStructsAndCappedView(t *testing.T) {
	type user struct{ id int }
	if got := New([]user{{1}, {2}, {3}}).TakeLast(1); !reflect.DeepEqual(got, Slice[user]{{3}}) {
		t.Fatalf("TakeLast() = %v", got)
	}
	items := []int{1, 2, 3, 4}
	got := New(items).TakeLast(2)
	got[0] = 99
	if items[2] != 99 || cap(got) != len(got) {
		t.Fatalf("TakeLast() did not return a capped view: %v cap %d", got, cap(got))
	}
}
