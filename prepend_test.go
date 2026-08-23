package collection

import (
	"reflect"
	"testing"
)

func TestPrepend_Basic(t *testing.T) {
	c := New([]int{3, 4})

	out := c.Prepend(1, 2)

	expected := []int{1, 2, 3, 4}
	if !reflect.DeepEqual(out.Items(), expected) || !reflect.DeepEqual(c.Items(), []int{3, 4}) {
		t.Fatalf("result=%v source=%v", out, c)
	}
}

func TestPrepend_EmptyCollection(t *testing.T) {
	c := New([]int{})

	out := c.Prepend(1, 2, 3)

	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(out.Items(), expected) || len(c) != 0 {
		t.Fatalf("result=%v source=%v", out, c)
	}
}

func TestPrepend_NoValues(t *testing.T) {
	c := New([]int{1, 2, 3})

	out := c.Prepend() // no-op

	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(out.Items(), expected) || !reflect.DeepEqual(c.Items(), expected) {
		t.Fatalf("result=%v source=%v", out, c)
	}
}

func TestPrepend_Structs(t *testing.T) {
	type User struct {
		ID   int
		Name string
	}

	c := New([]User{
		{3, "Shawn"},
		{4, "Van"},
	})

	out := c.Prepend(User{1, "Chris"}, User{2, "Matt"})

	expected := []User{
		{1, "Chris"},
		{2, "Matt"},
		{3, "Shawn"},
		{4, "Van"},
	}

	if !reflect.DeepEqual(out.Items(), expected) || len(c) != 2 {
		t.Fatalf("result=%v source=%v", out, c)
	}
}

func TestPrepend_DoesNotMutateSourceSlice(t *testing.T) {
	orig := []int{10, 20, 30}
	c := New(orig)

	out := c.Prepend(5, 7)

	if !reflect.DeepEqual(orig, []int{10, 20, 30}) {
		t.Fatalf("Prepend mutated source slice: %v", orig)
	}

	expected := []int{5, 7, 10, 20, 30}
	if !reflect.DeepEqual(out.Items(), expected) || !reflect.DeepEqual(c.Items(), orig) {
		t.Fatalf("result=%v source=%v", out, c)
	}
}

func TestPrepend_NilSliceWithValues(t *testing.T) {
	c := New([]int(nil))

	out := c.Prepend(1, 2)

	expected := []int{1, 2}
	if !reflect.DeepEqual(out.Items(), expected) || c != nil {
		t.Fatalf("result=%v source=%v", out, c)
	}
}

func TestPrepend_NilSliceNoValuesBecomesEmpty(t *testing.T) {
	c := New([]int(nil))

	out := c.Prepend()

	if out == nil || len(out) != 0 || c != nil {
		t.Fatalf("result=%v source=%v", out, c)
	}
}

func TestPrepend_ResultAndSourceAreIndependentWithSpareCapacity(t *testing.T) {
	backing := make([]int, 2, 8)
	copy(backing, []int{3, 4})
	c := New(backing)
	out := c.Prepend(1, 2)
	out[2] = 9
	if c[0] != 3 {
		t.Fatalf("result mutation changed source: %v", c)
	}
	c[1] = 8
	if out[3] != 4 {
		t.Fatalf("source mutation changed result: %v", out)
	}
}
