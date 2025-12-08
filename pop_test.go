package collection

import (
	"reflect"
	"testing"
)

func TestPop_RemovesOne(t *testing.T) {
	c := New([]int{1, 2, 3})

	v, c2 := c.Pop()

	if v != 3 {
		t.Fatalf("Pop() expected 3, got %v", v)
	}

	want := []int{1, 2}
	if !reflect.DeepEqual(c2.Items(), want) {
		t.Fatalf("Pop() expected remainder %v, got %v", want, c2.Items())
	}

	// original must remain unchanged
	if !reflect.DeepEqual(c.Items(), []int{1, 2, 3}) {
		t.Fatalf("Pop() mutated original collection")
	}
}

func TestPop_OnEmptyReturnsZero(t *testing.T) {
	c := New([]int{})

	v, c2 := c.Pop()

	if v != 0 {
		t.Fatalf("Pop() on empty should return zero-value, got %v", v)
	}

	if len(c2.Items()) != 0 {
		t.Fatalf("Pop() on empty should return empty collection")
	}
}

func TestPopN_RemovesMultiple(t *testing.T) {
	c := New([]int{1, 2, 3, 4, 5})

	popped, remain := c.PopN(3)

	wantPopped := []int{5, 4, 3}
	wantRemain := []int{1, 2}

	if !reflect.DeepEqual(popped.Items(), wantPopped) {
		t.Fatalf("PopN(3) wrong popped values. want=%v got=%v", wantPopped, popped.Items())
	}

	if !reflect.DeepEqual(remain.Items(), wantRemain) {
		t.Fatalf("PopN(3) wrong remaining values. want=%v got=%v", wantRemain, remain.Items())
	}

	// original must remain unchanged
	if !reflect.DeepEqual(c.Items(), []int{1, 2, 3, 4, 5}) {
		t.Fatalf("PopN() mutated original collection")
	}
}

func TestPopN_MoreThanLength(t *testing.T) {
	c := New([]int{1, 2})

	popped, remain := c.PopN(10)

	wantPopped := []int{2, 1}

	if !reflect.DeepEqual(popped.Items(), wantPopped) {
		t.Fatalf("PopN(>len) wrong popped. want=%v got=%v", wantPopped, popped.Items())
	}

	if len(remain.Items()) != 0 {
		t.Fatalf("PopN(>len) should return empty remainder")
	}

	// original unchanged
	if !reflect.DeepEqual(c.Items(), []int{1, 2}) {
		t.Fatalf("PopN() mutated original collection")
	}
}

func TestPopN_ZeroOrNegative(t *testing.T) {
	c := New([]int{1, 2, 3})

	popped, remain := c.PopN(0)
	if len(popped.Items()) != 0 {
		t.Fatalf("PopN(0) should return empty popped collection")
	}
	if !reflect.DeepEqual(remain.Items(), []int{1, 2, 3}) {
		t.Fatalf("PopN(0) should not modify original")
	}

	popped, remain = c.PopN(-5)
	if len(popped.Items()) != 0 {
		t.Fatalf("PopN(-n) should return empty collection")
	}
	if !reflect.DeepEqual(remain.Items(), []int{1, 2, 3}) {
		t.Fatalf("PopN(-n) should not modify original")
	}

	// original still unchanged
	if !reflect.DeepEqual(c.Items(), []int{1, 2, 3}) {
		t.Fatalf("PopN zero/negative mutated original collection")
	}
}
