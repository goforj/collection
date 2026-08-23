package collection

import (
	"reflect"
	"testing"
)

func TestMode_SingleMode(t *testing.T) {
	c := []int{1, 2, 2, 3}
	m := Mode(c)

	exp := []int{2}
	if !reflect.DeepEqual(m, exp) {
		t.Fatalf("expected %v, got %v", exp, m)
	}
}

func TestMode_MultipleModes(t *testing.T) {
	c := []int{1, 2, 1, 2}
	m := Mode(c)

	exp := []int{1, 2} // first-seen order preserved
	if !reflect.DeepEqual(m, exp) {
		t.Fatalf("expected %v, got %v", exp, m)
	}
}

func TestMode_Empty(t *testing.T) {
	c := []float64{}
	m := Mode(c)

	if m != nil {
		t.Fatalf("expected nil for empty collection, got %v", m)
	}
}

func TestMode_SingleValue(t *testing.T) {
	c := []int{42}
	m := Mode(c)

	exp := []int{42}
	if !reflect.DeepEqual(m, exp) {
		t.Fatalf("expected %v, got %v", exp, m)
	}
}

func TestMode_OrderPreserved(t *testing.T) {
	c := []int{3, 1, 3, 2, 1}
	// counts: 3→2, 1→2, 2→1 → modes are [3, 1]
	m := Mode(c)

	exp := []int{3, 1}
	if !reflect.DeepEqual(m, exp) {
		t.Fatalf("expected %v, got %v", exp, m)
	}
}
