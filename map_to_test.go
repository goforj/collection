package collection

import (
	"reflect"
	"testing"
)

func TestMapTo_Ints(t *testing.T) {
	c := New([]int{1, 2, 3})

	out := MapTo(c, func(v int) string {
		if v%2 == 0 {
			return "even"
		}
		return "odd"
	})

	expected := []string{"odd", "even", "odd"}
	if !reflect.DeepEqual(out.Items(), expected) {
		t.Fatalf("expected %v, got %v", expected, out.Items())
	}
}

func TestMapTo_Strings(t *testing.T) {
	c := New([]string{"go", "forj", "rocks"})

	out := MapTo(c, func(s string) int { return len(s) })

	expected := []int{2, 4, 5}
	if !reflect.DeepEqual(out.Items(), expected) {
		t.Fatalf("expected %v, got %v", expected, out.Items())
	}
}

func TestMapTo_Structs(t *testing.T) {
	type User struct {
		ID   int
		Name string
	}

	users := New([]User{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
	})

	out := MapTo(users, func(u User) string { return u.Name })

	expected := []string{"Alice", "Bob"}
	if !reflect.DeepEqual(out.Items(), expected) {
		t.Fatalf("expected %v, got %v", expected, out.Items())
	}
}

func TestMapTo_Empty(t *testing.T) {
	c := New([]int{})

	out := MapTo(c, func(v int) int { return v * 2 })

	if len(out.Items()) != 0 {
		t.Fatalf("expected empty slice, got %v", out.Items())
	}
}

func TestMapTo_NoMutation(t *testing.T) {
	c := New([]int{1, 2, 3})

	_ = MapTo(c, func(v int) int { return v * 10 })

	if !reflect.DeepEqual(c.Items(), []int{1, 2, 3}) {
		t.Fatalf("MapTo should not mutate original collection")
	}
}
