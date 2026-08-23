package collection

import (
	"reflect"
	"testing"
)

func TestMapTo_Ints(t *testing.T) {
	c := New([]int{1, 2, 3})

	out := c.Map(func(v int) string {
		if v%2 == 0 {
			return "even"
		}
		return "odd"
	})

	expected := Slice[string]{"odd", "even", "odd"}
	if !reflect.DeepEqual(out, expected) {
		t.Fatalf("expected %v, got %v", expected, out)
	}
}

func TestMapTo_Strings(t *testing.T) {
	c := New([]string{"go", "forj", "rocks"})

	out := c.Map(func(s string) int { return len(s) })

	expected := Slice[int]{2, 4, 5}
	if !reflect.DeepEqual(out, expected) {
		t.Fatalf("expected %v, got %v", expected, out)
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

	out := users.Map(func(u User) string { return u.Name })

	expected := Slice[string]{"Alice", "Bob"}
	if !reflect.DeepEqual(out, expected) {
		t.Fatalf("expected %v, got %v", expected, out)
	}
}

func TestMapTo_Empty(t *testing.T) {
	c := New([]int{})

	out := c.Map(func(v int) int { return v * 2 })

	if len(out) != 0 {
		t.Fatalf("expected empty slice, got %v", out)
	}
}

func TestMapTo_NoMutation(t *testing.T) {
	c := New([]int{1, 2, 3})

	_ = c.Map(func(v int) int { return v * 10 })

	if !reflect.DeepEqual(c, Slice[int]{1, 2, 3}) {
		t.Fatalf("Map should not mutate original collection")
	}
}
