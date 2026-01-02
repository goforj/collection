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

	if mapped != c {
		t.Fatalf("Map should return the same collection")
	}

	if !reflect.DeepEqual(c.Items(), expected) {
		t.Fatalf("Map should mutate original collection")
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

	if mapped != c {
		t.Fatalf("Map should return the same collection")
	}

	if !reflect.DeepEqual(c.Items(), expected) {
		t.Fatalf("Map should mutate original collection")
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

	if mapped != c {
		t.Fatalf("Map should return the same collection")
	}
}

func TestMap_PreservesNilSlice(t *testing.T) {
	c := New([]int(nil))

	c.Map(func(v int) int { return v * 2 })

	if c.Items() != nil {
		t.Fatalf("expected nil slice to remain nil, got %v", c.Items())
	}
}

func TestMap_WritesThroughSourceSlice(t *testing.T) {
	items := []int{1, 2, 3}
	c := New(items)

	c.Map(func(v int) int { return v * 2 })

	want := []int{2, 4, 6}
	if !reflect.DeepEqual(items, want) {
		t.Fatalf("expected source slice %v, got %v", want, items)
	}
}

func TestMap_LengthUnchanged(t *testing.T) {
	c := New([]int{1, 2, 3})

	c.Map(func(v int) int { return v + 1 })

	if len(c.Items()) != 3 {
		t.Fatalf("expected length 3, got %d", len(c.Items()))
	}
}
