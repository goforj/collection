package collection

import (
	"reflect"
	"testing"
)

func TestMedian_OddAndEven(t *testing.T) {
	odd := New([]int{1, 3, 2})
	med, ok := Median(odd)
	if !ok || med != 2 {
		t.Fatalf("expected median 2, got %v ok=%v", med, ok)
	}

	even := New([]int{1, 2, 3, 4})
	med, ok = Median(even)
	if !ok || med != 2.5 {
		t.Fatalf("expected median 2.5, got %v ok=%v", med, ok)
	}
}

func TestMedian_Empty(t *testing.T) {
	empty := New([]int{})
	med, ok := Median(empty)

	if ok {
		t.Fatalf("expected ok=false on empty collection, got true (med=%v)", med)
	}
	if med != 0 {
		t.Fatalf("expected median=0 for empty, got %v", med)
	}
}

func TestMedian_Negatives(t *testing.T) {
	c := New([]int{-5, -1, -10})

	med, ok := Median(c)
	if !ok || med != -5 {
		t.Fatalf("expected median -5, got %v ok=%v", med, ok)
	}
}

func TestMedian_Floats(t *testing.T) {
	c := New([]float64{1.5, 3.5, 2.5})

	med, ok := Median(c)
	if !ok || med != 2.5 {
		t.Fatalf("expected median 2.5, got %v ok=%v", med, ok)
	}
}

func TestMedian_NoMutation(t *testing.T) {
	c := New([]int{3, 1, 2})
	orig := append([]int{}, c.items...) // copy original

	_, _ = Median(c)

	if !reflect.DeepEqual(c.items, orig) {
		t.Fatalf("original collection mutated: %v", c.items)
	}
}
