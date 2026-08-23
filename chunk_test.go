package collection

import (
	"reflect"
	"testing"
)

func TestChunk(t *testing.T) {
	tests := []struct {
		name string
		in   Slice[int]
		size int
		want [][]int
	}{
		{"even", New([]int{1, 2, 3, 4}), 2, [][]int{{1, 2}, {3, 4}}},
		{"uneven", New([]int{1, 2, 3, 4, 5}), 2, [][]int{{1, 2}, {3, 4}, {5}}},
		{"larger than collection", New([]int{1, 2, 3}), 10, [][]int{{1, 2, 3}}},
		{"size one", New([]int{1, 2, 3}), 1, [][]int{{1}, {2}, {3}}},
		{"empty", New([]int{}), 3, [][]int{}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := test.in.Chunk(test.size); !reflect.DeepEqual(got, test.want) {
				t.Fatalf("Chunk(%d) = %v, want %v", test.size, got, test.want)
			}
		})
	}
}

func TestChunkInvalidSize(t *testing.T) {
	values := New([]int{1, 2, 3})
	for _, size := range []int{0, -5} {
		if got := values.Chunk(size); got != nil {
			t.Fatalf("Chunk(%d) = %v, want nil", size, got)
		}
	}
}

func TestChunkReturnsCappedViews(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	chunks := New(items).Chunk(2)
	chunks[1][0] = 99
	if items[2] != 99 || cap(chunks[0]) != len(chunks[0]) || cap(chunks[2]) != len(chunks[2]) {
		t.Fatalf("Chunk() did not return capped source views: chunks %v, source %v", chunks, items)
	}
}
