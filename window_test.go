package collection

import (
	"reflect"
	"testing"
)

func TestWindow(t *testing.T) {
	tests := []struct {
		name       string
		size, step int
		want       [][]int
	}{
		{"overlapping", 3, 1, [][]int{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}}},
		{"stepped", 2, 2, [][]int{{1, 2}, {3, 4}}},
		{"default step", 2, 0, [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := New([]int{1, 2, 3, 4, 5}).Window(test.size, test.step)
			if !reflect.DeepEqual(got, test.want) {
				t.Fatalf("Window(%d, %d) = %v, want %v", test.size, test.step, got, test.want)
			}
		})
	}
}

func TestWindowEdgeCasesAndViews(t *testing.T) {
	values := New([]int{1, 2})
	for _, args := range [][2]int{{3, 1}, {0, 1}} {
		if got := values.Window(args[0], args[1]); got != nil {
			t.Fatalf("Window(%d, %d) = %v, want nil", args[0], args[1], got)
		}
	}
	items := []int{1, 2, 3}
	windows := New(items).Window(2, 1)
	windows[1][0] = 99
	if items[1] != 99 || cap(windows[0]) != len(windows[0]) {
		t.Fatalf("Window() did not return capped source views: %v, source %v", windows, items)
	}
}
