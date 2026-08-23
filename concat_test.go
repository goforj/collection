package collection

import (
	"reflect"
	"testing"
)

func TestConcat(t *testing.T) {
	tests := []struct {
		name string
		in   Slice[int]
		more [][]int
		want Slice[int]
	}{
		{"one slice", New([]int{1}), [][]int{{2}}, []int{1, 2}},
		{"chained inputs", New([]int{1}), [][]int{{2}, {3}}, []int{1, 2, 3}},
		{"empty receiver", New([]int{}), [][]int{{5, 6}}, []int{5, 6}},
		{"empty values", New([]int{1, 2, 3}), [][]int{{}}, []int{1, 2, 3}},
		{"both empty", New([]int{}), [][]int{{}}, []int{}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := test.in.Concat(test.more...); !reflect.DeepEqual(got, test.want) {
				t.Fatalf("Concat() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestConcatAcceptsSliceAndIsIndependent(t *testing.T) {
	left := New([]int{1, 2})
	right := New([]int{3, 4})
	got := left.Concat(right)
	got[0], got[2] = 99, 88
	if !reflect.DeepEqual(left, Slice[int]{1, 2}) || !reflect.DeepEqual(right, Slice[int]{3, 4}) {
		t.Fatalf("Concat() mutated input: left %v right %v", left, right)
	}
}

func TestConcatSupportsStructs(t *testing.T) {
	type user struct{ name string }
	got := New([]user{{"Chris"}}).Concat([]user{{"Van"}, {"Shawn"}})
	want := Slice[user]{{"Chris"}, {"Van"}, {"Shawn"}}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Concat() = %v, want %v", got, want)
	}
}
