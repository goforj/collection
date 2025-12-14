package collection

import (
	"reflect"
	"testing"
)

func TestFromMap_Basic(t *testing.T) {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	c := FromMap(m)
	items := c.Items()

	if len(items) != 3 {
		t.Fatalf("expected 3 items, got %d", len(items))
	}

	seen := make(map[string]int)
	for _, p := range items {
		seen[p.Key] = p.Value
	}

	if !reflect.DeepEqual(seen, m) {
		t.Fatalf("expected %v, got %v", m, seen)
	}
}

func TestFromMap_DoesNotMutateInput(t *testing.T) {
	m := map[string]int{
		"a": 1,
		"b": 2,
	}

	_ = FromMap(m)

	expect := map[string]int{
		"a": 1,
		"b": 2,
	}

	if !reflect.DeepEqual(m, expect) {
		t.Fatalf("input map was mutated")
	}
}

func TestFromMap_EmptyMap(t *testing.T) {
	m := map[string]int{}

	c := FromMap(m)

	if len(c.Items()) != 0 {
		t.Fatalf("expected empty collection, got %v", c.Items())
	}
}

func TestFromMap_StructValues(t *testing.T) {
	type User struct {
		ID   int
		Name string
	}

	m := map[string]User{
		"alice": {ID: 1, Name: "Alice"},
		"bob":   {ID: 2, Name: "Bob"},
	}

	c := FromMap(m)
	items := c.Items()

	if len(items) != 2 {
		t.Fatalf("expected 2 items, got %d", len(items))
	}

	seen := make(map[string]User)
	for _, p := range items {
		seen[p.Key] = p.Value
	}

	if !reflect.DeepEqual(seen, m) {
		t.Fatalf("expected %v, got %v", m, seen)
	}
}

func TestFromMap_Capacity(t *testing.T) {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	c := FromMap(m)
	items := c.Items()

	if cap(items) < len(m) {
		t.Fatalf("expected capacity >= %d, got %d", len(m), cap(items))
	}
}
