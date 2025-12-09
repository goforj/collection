package collection

import (
	"reflect"
	"testing"
)

func TestMap_Ints(t *testing.T) {
	c := New([]int{1, 2, 3})

	mapped := c.Map(func(v int) int {
		return v * 10
	})

	expected := []int{10, 20, 30}

	if !reflect.DeepEqual(mapped.items, expected) {
		t.Fatalf("expected %v, got %v", expected, mapped.items)
	}
}

func TestMap_Structs(t *testing.T) {
	type User struct {
		ID   int
		Name string
	}

	c := New([]User{
		{1, "Chris"},
		{2, "Van"},
	})

	mapped := c.Map(func(u User) User {
		u.Name = u.Name + "!"
		return u
	})

	expected := []User{
		{1, "Chris!"},
		{2, "Van!"},
	}

	if !reflect.DeepEqual(mapped.items, expected) {
		t.Fatalf("expected %v, got %v", expected, mapped.items)
	}
}

func TestMap_Empty(t *testing.T) {
	c := New([]int{})

	mapped := c.Map(func(v int) int {
		return v * 2
	})

	if len(mapped.items) != 0 {
		t.Fatalf("expected empty slice, got %v", mapped.items)
	}
}
