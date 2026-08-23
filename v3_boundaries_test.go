package collection

import (
	"math"
	"reflect"
	"testing"
)

func TestViewsCapTheirCapacity(t *testing.T) {
	values := New([]int{1, 2, 3, 4})
	views := []Slice[int]{
		values.After(func(value int) bool { return value == 1 }),
		values.Take(2),
		values.TakeLast(2),
		values.Skip(1),
		values.SkipLast(1),
		values.TakeUntil(func(value int) bool { return value == 3 }),
	}
	for _, view := range views {
		if cap(view) != len(view) {
			t.Fatalf("view capacity = %d, want length %d", cap(view), len(view))
		}
	}
	for _, chunk := range values.Chunk(2) {
		if cap(chunk) != len(chunk) {
			t.Fatalf("chunk capacity = %d, want length %d", cap(chunk), len(chunk))
		}
	}
	for _, window := range values.Window(2, 1) {
		if cap(window) != len(window) {
			t.Fatalf("window capacity = %d, want length %d", cap(window), len(window))
		}
	}
}

func TestTakeAcceptsOnlyPositiveCounts(t *testing.T) {
	values := New([]int{1, 2, 3})
	if got := values.Take(-1); len(got) != 0 || cap(got) != 0 {
		t.Fatalf("Take(-1) = %v with capacity %d, want empty capped view", got, cap(got))
	}
	if got := values.TakeLast(1); !reflect.DeepEqual(got, Slice[int]{3}) {
		t.Fatalf("TakeLast(1) = %v, want [3]", got)
	}
}

func TestTerminalSliceResults(t *testing.T) {
	values := New([]int{1, 2, 3, 4})
	left, right := values.Partition(func(value int) bool { return value%2 == 0 })
	if !reflect.DeepEqual(left, []int{2, 4}) || !reflect.DeepEqual(right, []int{1, 3}) {
		t.Fatalf("Partition = (%v, %v)", left, right)
	}
	if got := values.Window(2, 2); !reflect.DeepEqual(got, [][]int{{1, 2}, {3, 4}}) {
		t.Fatalf("Window = %v", got)
	}
}

func TestConcatIsVariadicAndIndependent(t *testing.T) {
	backing := make([]int, 2, 4)
	copy(backing, []int{1, 2})
	values := New(backing)
	joined := values.Concat([]int{3}, []int{4})
	if !reflect.DeepEqual(joined, Slice[int]{1, 2, 3, 4}) {
		t.Fatalf("Concat = %v", joined)
	}
	joined[0] = 9
	if values[0] != 1 {
		t.Fatalf("Concat aliases source: %v", values)
	}
}

func TestMapAndZipUsePair(t *testing.T) {
	entries := FromMap(map[string]int{"one": 1})
	if entries[0].First != "one" || entries[0].Second != 1 {
		t.Fatalf("FromMap entry = %#v", entries[0])
	}
	pairs := New([]int{1, 2}).Zip([]string{"one"})
	if !reflect.DeepEqual(pairs, []Pair[int, string]{{First: 1, Second: "one"}}) {
		t.Fatalf("Zip = %#v", pairs)
	}
	if got := New([]int{1, 2}).ZipWith([]int{10}, func(left, right int) int {
		return left + right
	}); !reflect.DeepEqual(got, Slice[int]{11}) {
		t.Fatalf("ZipWith = %v", got)
	}
	groups := New([]int{1, 2, 3}).GroupBy(func(value int) string {
		if value%2 == 0 {
			return "even"
		}
		return "odd"
	})
	if !reflect.DeepEqual(groups, map[string][]int{"even": {2}, "odd": {1, 3}}) {
		t.Fatalf("GroupBy = %v", groups)
	}
}

func TestSetOperationsAcceptMixedSliceTypes(t *testing.T) {
	native := []int{2, 3}
	named := New([]int{1, 2})

	if got := Union(named, native); !reflect.DeepEqual(got, Slice[int]{1, 2, 3}) {
		t.Fatalf("Union mixed inputs = %v", got)
	}
	if got := Intersect(named, native); !reflect.DeepEqual(got, Slice[int]{2}) {
		t.Fatalf("Intersect mixed inputs = %v", got)
	}
	if got := Difference(named, native); !reflect.DeepEqual(got, Slice[int]{1}) {
		t.Fatalf("Difference mixed inputs = %v", got)
	}
	if got := SymmetricDifference(named, native); !reflect.DeepEqual(got, Slice[int]{1, 3}) {
		t.Fatalf("SymmetricDifference mixed inputs = %v", got)
	}
}

func TestSliceBoundariesHandleMaximumCounts(t *testing.T) {
	values := New([]int{1})
	if got := values.Chunk(math.MaxInt); !reflect.DeepEqual(got, [][]int{{1}}) {
		t.Fatalf("Chunk(MaxInt) = %v", got)
	}
	if got := values.Window(1, math.MaxInt); !reflect.DeepEqual(got, [][]int{{1}}) {
		t.Fatalf("Window(1, MaxInt) = %v", got)
	}
}

func TestPopNMigrationRecipeHandlesEveryCount(t *testing.T) {
	tests := []struct {
		name       string
		count      int
		wantValues []int
		wantPopped []int
	}{
		{name: "negative", count: -1, wantValues: []int{1, 2, 3}},
		{name: "zero", count: 0, wantValues: []int{1, 2, 3}},
		{name: "partial", count: 2, wantValues: []int{1}, wantPopped: []int{2, 3}},
		{name: "oversized", count: 10, wantValues: []int{}, wantPopped: []int{1, 2, 3}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			values := []int{1, 2, 3}
			count := test.count
			var popped []int
			if count > 0 {
				count = min(count, len(values))
				split := len(values) - count
				popped = values[split:len(values):len(values)]
				values = values[:split]
			}

			if !reflect.DeepEqual(values, test.wantValues) {
				t.Fatalf("values = %v, want %v", values, test.wantValues)
			}
			if !reflect.DeepEqual(popped, test.wantPopped) {
				t.Fatalf("popped = %v, want %v", popped, test.wantPopped)
			}
		})
	}
}
