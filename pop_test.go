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

func TestPop_PreservesNilSlice(t *testing.T) {
	c := New([]int(nil))

	v, ok := c.Pop()

	if v != 0 || ok {
		t.Fatalf("Pop() on nil should return zero-value and ok=false")
	}

	if c.Items() != nil {
		t.Fatalf("expected nil slice to remain nil, got %v", c.Items())
	}
}

func TestPop_ClearsRemovedSlot(t *testing.T) {
	a, b := 1, 2
	items := []*int{&a, &b}
	c := New(items)

	popped, ok := c.Pop()

	if !ok || popped != &b {
		t.Fatalf("expected to pop second pointer")
	}
	if items[1] != nil {
		t.Fatalf("expected removed slot to be cleared, got %v", items[1])
	}
}

func TestPopN_PreservesNilSlice(t *testing.T) {
	c := New([]int(nil))

	popped := c.PopN(2)

	if popped != nil {
		t.Fatalf("PopN on nil should return nil")
	}

	if c.Items() != nil {
		t.Fatalf("expected nil slice to remain nil, got %v", c.Items())
	}
}

func TestPopN_ReturnsBackingViewWithIndependentHeader(t *testing.T) {
	backing := make([]int, 4, 6)
	copy(backing, []int{1, 2, 3, 4})
	c := New(backing)
	copyBeforePop := c

	popped := c.PopN(2)
	if !reflect.DeepEqual(popped, []int{3, 4}) || !reflect.DeepEqual(c.Items(), []int{1, 2}) {
		t.Fatalf("PopN result=%v remainder=%v", popped, c)
	}
	if &popped[0] != &backing[2] {
		t.Fatalf("PopN result does not view the original backing array")
	}
	if len(copyBeforePop) != 4 || !reflect.DeepEqual(copyBeforePop.Items(), []int{1, 2, 3, 4}) {
		t.Fatalf("copied Slice header changed after PopN: %v", copyBeforePop)
	}

	c = append(c, 9, 10)
	if !reflect.DeepEqual(popped, []int{9, 10}) {
		t.Fatalf("growth into spare capacity did not overwrite popped view: %v", popped)
	}
}
