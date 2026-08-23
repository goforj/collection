package collection

import (
	"reflect"
	"testing"
)

func TestSkip(t *testing.T) {
	tests := []struct {
		n    int
		want Slice[int]
	}{
		{2, []int{3, 4, 5}}, {0, []int{1, 2, 3, 4, 5}}, {-1, []int{1, 2, 3, 4, 5}}, {5, []int{}}, {10, []int{}},
	}
	for _, test := range tests {
		got := New([]int{1, 2, 3, 4, 5}).Skip(test.n)
		if !reflect.DeepEqual(got, test.want) {
			t.Fatalf("Skip(%d) = %v, want %v", test.n, got, test.want)
		}
	}
}

func TestSkipStructsAndCappedView(t *testing.T) {
	type user struct{ id int }
	if got := New([]user{{1}, {2}, {3}}).Skip(1); !reflect.DeepEqual(got, Slice[user]{{2}, {3}}) {
		t.Fatalf("Skip() = %v", got)
	}
	items := []int{1, 2, 3, 4}
	got := New(items).Skip(1)
	got[0] = 99
	if items[1] != 99 || cap(got) != len(got) {
		t.Fatalf("Skip() did not return a capped view: %v cap %d", got, cap(got))
	}
}
