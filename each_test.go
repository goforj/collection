package collection

import (
	"reflect"
	"testing"
)

func TestEach_SideEffects(t *testing.T) {
	c := New([]int{1, 2, 3})

	var sum int
	c.Each(func(v int) {
		sum += v
	})

	if sum != 6 {
		t.Fatalf("expected sum=6, got %d", sum)
	}
}

func TestEach_ReturnsSameCollection(t *testing.T) {
	c := New([]int{1, 2, 3})

	out := c.Each(func(v int) {})

	if out != c {
		t.Fatalf("Each should return the same collection instance")
	}
}

func TestEach_Chaining(t *testing.T) {
	var seen []int

	out := New([]int{1, 2, 3}).
		Each(func(v int) {
			seen = append(seen, v)
		}).
		Map(func(v int) int {
			return v * 2
		})

	if !reflect.DeepEqual(seen, []int{1, 2, 3}) {
		t.Fatalf("Each did not see all items: %v", seen)
	}

	if !reflect.DeepEqual(out.items, []int{2, 4, 6}) {
		t.Fatalf("Map after Each returned wrong result: %v", out.items)
	}
}
