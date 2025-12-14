package collection

import (
	"reflect"
	"testing"
)

func TestClone_Independence(t *testing.T) {
	c := New([]int{1, 2, 3})
	clone := c.Clone()

	clone = clone.Append(4)

	if !reflect.DeepEqual(c.Items(), []int{1, 2, 3}) {
		t.Fatalf("original collection mutated: %v", c.Items())
	}

	if !reflect.DeepEqual(clone.Items(), []int{1, 2, 3, 4}) {
		t.Fatalf("clone incorrect: %v", clone.Items())
	}
}

func TestClone_BackendSliceIsolated(t *testing.T) {
	c := New([]int{1, 2, 3})
	clone := c.Clone()

	clone.items[0] = 99

	if c.items[0] == 99 {
		t.Fatalf("clone shares backing slice with original")
	}
}

func TestClone_EmptyCollection(t *testing.T) {
	c := New([]int{})
	clone := c.Clone()

	if len(clone.Items()) != 0 {
		t.Fatalf("expected empty clone, got %v", clone.Items())
	}
}

func TestClone_ChainedBranching(t *testing.T) {
	base := New([]int{1, 2, 3, 4, 5})

	evens := base.Clone().Filter(func(v int) bool {
		return v%2 == 0
	})

	odds := base.Clone().Filter(func(v int) bool {
		return v%2 != 0
	})

	if !reflect.DeepEqual(base.Items(), []int{1, 2, 3, 4, 5}) {
		t.Fatalf("base collection mutated: %v", base.Items())
	}

	if !reflect.DeepEqual(evens.Items(), []int{2, 4}) {
		t.Fatalf("evens incorrect: %v", evens.Items())
	}

	if !reflect.DeepEqual(odds.Items(), []int{1, 3, 5}) {
		t.Fatalf("odds incorrect: %v", odds.Items())
	}
}
