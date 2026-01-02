package collection

import (
	"reflect"
	"testing"
)

func TestConcat_Slice(t *testing.T) {
	c := New([]string{"John Doe"})

	out := c.Concat([]string{"Jane Doe"})

	expected := []string{"John Doe", "Jane Doe"}
	if !reflect.DeepEqual(out.items, expected) {
		t.Fatalf("expected %v, got %v", expected, out.items)
	}
}

func TestConcat_Chained(t *testing.T) {
	c := New([]string{"John Doe"})

	out := c.
		Concat([]string{"Jane Doe"}).
		Concat([]string{"Johnny Doe"})

	expected := []string{"John Doe", "Jane Doe", "Johnny Doe"}
	if !reflect.DeepEqual(out.items, expected) {
		t.Fatalf("expected %v, got %v", expected, out.items)
	}
}

func TestConcat_WithOtherCollection(t *testing.T) {
	c1 := New([]int{1, 2})
	c2 := New([]int{3, 4})

	out := c1.Concat(c2.Items())

	expected := []int{1, 2, 3, 4}
	if !reflect.DeepEqual(out.items, expected) {
		t.Fatalf("expected %v, got %v", expected, out.items)
	}
}

func TestConcat_EmptyCurrent(t *testing.T) {
	c := New([]int{})

	out := c.Concat([]int{5, 6})

	expected := []int{5, 6}
	if !reflect.DeepEqual(out.items, expected) {
		t.Fatalf("expected %v, got %v", expected, out.items)
	}
}

func TestConcat_EmptyValues(t *testing.T) {
	c := New([]int{1, 2, 3})

	out := c.Concat([]int{})

	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(out.items, expected) {
		t.Fatalf("expected %v, got %v", expected, out.items)
	}
}

func TestConcat_EmptyBoth(t *testing.T) {
	c := New([]int{})

	out := c.Concat([]int{})

	expected := []int{}
	if !reflect.DeepEqual(out.items, expected) {
		t.Fatalf("expected empty slice, got %v", out.items)
	}
}

func TestConcat_NoMutation(t *testing.T) {
	origData := []int{1, 2, 3}
	c := New(origData)

	// perform concat
	_ = c.Concat([]int{4, 5})
}

func TestConcat_DifferentTypes(t *testing.T) {
	type User struct {
		Name string
	}

	c := New([]User{{"Chris"}})

	out := c.Concat([]User{{"Van"}, {"Shawn"}})

	expected := []User{{"Chris"}, {"Van"}, {"Shawn"}}
	if !reflect.DeepEqual(out.items, expected) {
		t.Fatalf("expected %v, got %v", expected, out.items)
	}
}

func TestConcat_NoAllocationWhenCapacityAllows(t *testing.T) {
	c := &Collection[int]{items: make([]int, 2, 10)}
	c.items[0], c.items[1] = 1, 2

	beforePtr := &c.items[0]

	c.Concat([]int{3, 4, 5})

	afterPtr := &c.items[0]

	if beforePtr != afterPtr {
		t.Fatalf("Concat allocated new backing array when it should not have")
	}

	expected := []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(c.items, expected) {
		t.Fatalf("expected %v, got %v", expected, c.items)
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

	c.Concat([]int{1, 2})

	expected := []int{1, 2}
	if !reflect.DeepEqual(c.Items(), expected) {
		t.Fatalf("expected %v, got %v", expected, c.Items())
	}
}
