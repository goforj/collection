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

	// original must be unchanged
	if !reflect.DeepEqual(c.items, origData) {
		t.Fatalf("Concat mutated the original collection: %v != %v", c.items, origData)
	}
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
