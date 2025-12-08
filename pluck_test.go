package collection

import (
	"reflect"
	"testing"
)

func TestMapTo_Ints(t *testing.T) {
	c := New([]int{1, 2, 3})

	out := MapTo(c, func(v int) int {
		return v * v
	})

	expected := []int{1, 4, 9}

	if !reflect.DeepEqual(out.items, expected) {
		t.Fatalf("expected %v, got %v", expected, out.items)
	}
}

func TestMapTo_ChangeType(t *testing.T) {
	c := New([]int{1, 2, 3})

	out := MapTo(c, func(v int) string {
		return "num:" + string(rune('0'+v))
	})

	expected := []string{"num:1", "num:2", "num:3"}

	if !reflect.DeepEqual(out.items, expected) {
		t.Fatalf("expected %v, got %v", expected, out.items)
	}
}

func TestMapTo_StructToInt(t *testing.T) {
	type User struct {
		ID   int
		Name string
	}

	c := New([]User{
		{1, "Chris"},
		{2, "Van"},
		{3, "Shawn"},
	})

	out := MapTo(c, func(u User) int {
		return u.ID
	})

	expected := []int{1, 2, 3}

	if !reflect.DeepEqual(out.items, expected) {
		t.Fatalf("expected %v, got %v", expected, out.items)
	}
}

func TestPluck_StructField(t *testing.T) {
	type User struct {
		ID   int
		Name string
	}

	c := New([]User{
		{1, "Chris"},
		{2, "Van"},
		{3, "Shawn"},
	})

	out := Pluck(c, func(u User) string {
		return u.Name
	})

	expected := []string{"Chris", "Van", "Shawn"}

	if !reflect.DeepEqual(out.items, expected) {
		t.Fatalf("expected %v, got %v", expected, out.items)
	}
}

func TestMapTo_Empty(t *testing.T) {
	c := New([]int{})

	out := MapTo(c, func(v int) int {
		return v * 10
	})

	if len(out.items) != 0 {
		t.Fatalf("expected empty result, got %v", out.items)
	}
}

func TestPluck_Empty(t *testing.T) {
	type User struct{ Name string }

	c := New([]User{})

	out := Pluck(c, func(u User) string {
		return u.Name
	})

	if len(out.items) != 0 {
		t.Fatalf("expected empty result, got %v", out.items)
	}
}

func TestMapTo_NoMutation(t *testing.T) {
	c := New([]int{1, 2, 3})
	orig := append([]int{}, c.items...)

	_ = MapTo(c, func(v int) int {
		return v * 100
	})

	if !reflect.DeepEqual(c.items, orig) {
		t.Fatalf("MapTo mutated original collection: %v vs %v", c.items, orig)
	}
}

func TestPluck_NoMutation(t *testing.T) {
	type User struct {
		ID   int
		Name string
	}

	c := New([]User{
		{1, "Chris"},
		{2, "Van"},
	})

	orig := append([]User{}, c.items...)

	_ = Pluck(c, func(u User) string { return u.Name })

	if !reflect.DeepEqual(c.items, orig) {
		t.Fatalf("Pluck mutated original collection: %v vs %v", c.items, orig)
	}
}
