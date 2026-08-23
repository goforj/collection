package collection

import (
	"reflect"
	"testing"
)

func TestConcat_Slice(t *testing.T) {
	c := New([]string{"John Doe"})

	out := c.Concat([]string{"Jane Doe"})

	expected := []string{"John Doe", "Jane Doe"}
	if !reflect.DeepEqual(out.Items(), expected) {
		t.Fatalf("expected %v, got %v", expected, out.Items())
	}
}

func TestConcat_Chained(t *testing.T) {
	c := New([]string{"John Doe"})

	out := c.
		Concat([]string{"Jane Doe"}).
		Concat([]string{"Johnny Doe"})

	expected := []string{"John Doe", "Jane Doe", "Johnny Doe"}
	if !reflect.DeepEqual(out.Items(), expected) {
		t.Fatalf("expected %v, got %v", expected, out.Items())
	}
}

func TestConcat_WithOtherCollection(t *testing.T) {
	c1 := New([]int{1, 2})
	c2 := New([]int{3, 4})

	out := c1.Concat(c2.Items())

	expected := []int{1, 2, 3, 4}
	if !reflect.DeepEqual(out.Items(), expected) {
		t.Fatalf("expected %v, got %v", expected, out.Items())
	}
}

func TestConcat_EmptyCurrent(t *testing.T) {
	c := New([]int{})

	out := c.Concat([]int{5, 6})

	expected := []int{5, 6}
	if !reflect.DeepEqual(out.Items(), expected) {
		t.Fatalf("expected %v, got %v", expected, out.Items())
	}
}

func TestConcat_EmptyValues(t *testing.T) {
	c := New([]int{1, 2, 3})

	out := c.Concat([]int{})

	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(out.Items(), expected) {
		t.Fatalf("expected %v, got %v", expected, out.Items())
	}
}

func TestConcat_EmptyBoth(t *testing.T) {
	c := New([]int{})

	out := c.Concat([]int{})

	expected := []int{}
	if !reflect.DeepEqual(out.Items(), expected) {
		t.Fatalf("expected empty slice, got %v", out.Items())
	}
}

func TestConcat_ReusesSpareCapacity(t *testing.T) {
	backing := make([]int, 2, 5)
	backing[0], backing[1] = 1, 2
	c := New(backing)
	out := c.Concat([]int{3, 4})
	if !reflect.DeepEqual(out.Items(), []int{1, 2, 3, 4}) {
		t.Fatalf("result = %v", out)
	}
	if &out[0] != &c[0] || !reflect.DeepEqual(c.Items(), []int{1, 2}) {
		t.Fatalf("spare-capacity Concat did not retain source backing/header: source=%v result=%v", c, out)
	}
	if !reflect.DeepEqual(backing[:4], []int{1, 2, 3, 4}) {
		t.Fatalf("Concat did not write appended values into spare capacity: %v", backing[:4])
	}
}

func TestConcat_DifferentTypes(t *testing.T) {
	type User struct {
		Name string
	}

	c := New([]User{{"Chris"}})

	out := c.Concat([]User{{"Van"}, {"Shawn"}})

	expected := []User{{"Chris"}, {"Van"}, {"Shawn"}}
	if !reflect.DeepEqual(out.Items(), expected) {
		t.Fatalf("expected %v, got %v", expected, out.Items())
	}
}

func TestConcat_FullCapacityReturnsIndependentSlice(t *testing.T) {
	c := Slice[int]{1, 2}
	out := c.Concat([]int{3, 4, 5})
	if !reflect.DeepEqual(out.Items(), []int{1, 2, 3, 4, 5}) || !reflect.DeepEqual(c.Items(), []int{1, 2}) {
		t.Fatalf("result=%v source=%v", out, c)
	}
	if &out[0] == &c[0] {
		t.Fatalf("full-capacity Concat reused source backing")
	}
	out[0] = 99
	if c[0] != 1 {
		t.Fatalf("Concat result aliases source")
	}
}

func TestConcat_PreservesNilSliceWhenEmptyValues(t *testing.T) {
	c := New([]int(nil))

	c.Concat([]int{})

	if c.Items() != nil {
		t.Fatalf("expected nil slice to remain nil, got %v", c.Items())
	}
}

func TestConcat_NilSliceWithValues(t *testing.T) {
	c := New([]int(nil))

	out := c.Concat([]int{1, 2})

	expected := []int{1, 2}
	if !reflect.DeepEqual(out.Items(), expected) || c != nil {
		t.Fatalf("result=%v source=%v", out, c)
	}
}
