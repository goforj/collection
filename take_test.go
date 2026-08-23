package collection

import (
	"reflect"
	"testing"
)

func TestTake(t *testing.T) {
	tests := []struct {
		n    int
		want Slice[int]
	}{
		{3, []int{0, 1, 2}}, {-2, []int{}}, {0, []int{}}, {10, []int{0, 1, 2, 3, 4, 5}},
	}
	for _, test := range tests {
		got := New([]int{0, 1, 2, 3, 4, 5}).Take(test.n)
		if !reflect.DeepEqual(got, test.want) {
			t.Fatalf("Take(%d) = %v, want %v", test.n, got, test.want)
		}
	}
}

func TestTakeEmptyStructsAndCappedView(t *testing.T) {
	if got := New([]int{}).Take(5); len(got) != 0 {
		t.Fatalf("Take() = %v, want empty", got)
	}
	type user struct{ id int }
	if got := New([]user{{1}, {2}, {3}, {4}}).Take(2); !reflect.DeepEqual(got, Slice[user]{{1}, {2}}) {
		t.Fatalf("Take() = %v", got)
	}
	items := []int{1, 2, 3, 4}
	got := New(items).Take(2)
	got[0] = 99
	if items[0] != 99 || cap(got) != len(got) {
		t.Fatalf("Take() did not return a capped view: %v cap %d", got, cap(got))
	}
}
