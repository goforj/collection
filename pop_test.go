package collection

import (
	"reflect"
	"testing"
)

func TestPop_RemovesOne(t *testing.T) {
	c := New([]int{1, 2, 3})

	v, ok := c.Pop()

	if v != 3 {
		t.Fatalf("Pop() expected 3, got %v", v)
	}

	if !ok {
		t.Fatalf("Pop() expected ok=true")
	}

	want := []int{1, 2}
	if !reflect.DeepEqual(c.Items(), want) {
		t.Fatalf("Pop() expected remainder %v, got %v", want, c.Items())
	}
}

func TestPop_OnEmptyReturnsZero(t *testing.T) {
	c := New([]int{})

	v, ok := c.Pop()

	if v != 0 {
		t.Fatalf("Pop() on empty should return zero-value, got %v", v)
	}

	if ok {
		t.Fatalf("Pop() on empty should return ok=false")
	}

	if len(c.Items()) != 0 {
		t.Fatalf("Pop() on empty should keep collection empty")
	}
}

func TestPopN_RemovesMultiple(t *testing.T) {
	c := New([]int{1, 2, 3, 4, 5})

	popped := c.PopN(3)

	wantPopped := []int{3, 4, 5}
	wantRemain := []int{1, 2}

	if !reflect.DeepEqual(popped, wantPopped) {
		t.Fatalf("PopN(3) wrong popped values. want=%v got=%v", wantPopped, popped)
	}

	if !reflect.DeepEqual(c.Items(), wantRemain) {
		t.Fatalf("PopN(3) wrong remaining values. want=%v got=%v", wantRemain, c.Items())
	}
}

func TestPopN_MoreThanLength(t *testing.T) {
	c := New([]int{1, 2})

	popped := c.PopN(10)

	wantPopped := []int{1, 2}

	if !reflect.DeepEqual(popped, wantPopped) {
		t.Fatalf("PopN(>len) wrong popped. want=%v got=%v", wantPopped, popped)
	}

	if len(c.Items()) != 0 {
		t.Fatalf("PopN(>len) should leave empty remainder")
	}
}

func TestPopN_ZeroOrNegative(t *testing.T) {
	c := New([]int{1, 2, 3})

	popped := c.PopN(0)
	if popped != nil {
		t.Fatalf("PopN(0) should return nil")
	}
	if !reflect.DeepEqual(c.Items(), []int{1, 2, 3}) {
		t.Fatalf("PopN(0) should not modify original")
	}

	popped = c.PopN(-5)
	if popped != nil {
		t.Fatalf("PopN(-n) should return nil")
	}
	if !reflect.DeepEqual(c.Items(), []int{1, 2, 3}) {
		t.Fatalf("PopN(-n) should not modify original")
	}
}
