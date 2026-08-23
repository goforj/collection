package collection

import (
	"reflect"
	"testing"
)

func TestSkipLast(t *testing.T) {
	tests := []struct {
		n    int
		want Slice[int]
	}{
		{2, []int{1, 2, 3}}, {0, []int{1, 2, 3, 4, 5}}, {-1, []int{1, 2, 3, 4, 5}}, {5, []int{}}, {10, []int{}},
	}
	for _, test := range tests {
		got := New([]int{1, 2, 3, 4, 5}).SkipLast(test.n)
		if !reflect.DeepEqual(got, test.want) {
			t.Fatalf("SkipLast(%d) = %v, want %v", test.n, got, test.want)
		}
	}
}

func TestSkipLastStructsAndCappedView(t *testing.T) {
	type user struct{ id int }
	if got := New([]user{{1}, {2}, {3}}).SkipLast(1); !reflect.DeepEqual(got, Slice[user]{{1}, {2}}) {
		t.Fatalf("SkipLast() = %v", got)
	}
	items := []int{1, 2, 3, 4}
	got := New(items).SkipLast(1)
	got[0] = 99
	if items[0] != 99 || cap(got) != len(got) {
		t.Fatalf("SkipLast() did not return a capped view: %v cap %d", got, cap(got))
	}
}
