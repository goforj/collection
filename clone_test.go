package collection

import (
	"reflect"
	"testing"
)

func TestClone_Independence(t *testing.T) {
	c := New([]int{1, 2, 3})
	clone := c.Clone()

	clone.Transform(func(value int) int { return value * 10 })

	if !reflect.DeepEqual(c, Slice[int]{1, 2, 3}) {
		t.Fatalf("original collection mutated: %v", c)
	}

	if !reflect.DeepEqual(clone, Slice[int]{10, 20, 30}) {
		t.Fatalf("clone incorrect: %v", clone)
	}
}

func TestClone_BackendSliceIsolated(t *testing.T) {
	c := New([]int{1, 2, 3})
	clone := c.Clone()

	clone[0] = 99

	if c[0] == 99 {
		t.Fatalf("clone shares backing slice with original")
	}
}

func TestClone_EmptyCollection(t *testing.T) {
	c := New([]int{})
	clone := c.Clone()

	if clone == nil {
		t.Fatal("Clone of an empty slice should be non-nil")
	}
	if len(clone) != 0 {
		t.Fatalf("expected empty clone, got %v", clone)
	}

	var nilSlice Slice[int]
	nilClone := nilSlice.Clone()
	if nilClone == nil {
		t.Fatal("Clone of a nil slice should be non-nil")
	}
}

func TestClone_ShallowCopySharesReferencedData(t *testing.T) {
	original := New([]map[string]int{{"count": 1}})
	clone := original.Clone()

	clone[0]["count"] = 2
	if original[0]["count"] != 2 {
		t.Fatal("Clone should retain references stored in elements")
	}
}

func TestClone_ChainedBranching(t *testing.T) {
	base := New([]int{1, 2, 3, 4, 5})

	evens := base.Clone().Retain(func(v int) bool {
		return v%2 == 0
	})

	odds := base.Clone().Retain(func(v int) bool {
		return v%2 != 0
	})

	if !reflect.DeepEqual(base, Slice[int]{1, 2, 3, 4, 5}) {
		t.Fatalf("base collection mutated: %v", base)
	}

	if !reflect.DeepEqual(evens, Slice[int]{2, 4}) {
		t.Fatalf("evens incorrect: %v", evens)
	}

	if !reflect.DeepEqual(odds, Slice[int]{1, 3, 5}) {
		t.Fatalf("odds incorrect: %v", odds)
	}
}
