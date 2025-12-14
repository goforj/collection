package collection

import (
	"reflect"
	"testing"
)

func TestToMap_Basic(t *testing.T) {
	users := []string{"alice", "bob", "carol"}

	out := ToMap(
		New(users),
		func(name string) string { return name },
		func(name string) int { return len(name) },
	)

	expect := map[string]int{
		"alice": 5,
		"bob":   3,
		"carol": 5,
	}

	if !reflect.DeepEqual(out, expect) {
		t.Fatalf("expected %v, got %v", expect, out)
	}
}

func TestToMap_RekeyStructs(t *testing.T) {
	type User struct {
		ID   int
		Name string
	}

	users := []User{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
	}

	out := ToMap(
		New(users),
		func(u User) int { return u.ID },
		func(u User) User { return u },
	)

	expect := map[int]User{
		1: {ID: 1, Name: "Alice"},
		2: {ID: 2, Name: "Bob"},
	}

	if !reflect.DeepEqual(out, expect) {
		t.Fatalf("expected %v, got %v", expect, out)
	}
}

func TestToMap_KeyCollision_LastWins(t *testing.T) {
	values := []int{1, 2, 3, 4}

	out := ToMap(
		New(values),
		func(v int) string { return "key" },
		func(v int) int { return v },
	)

	expect := map[string]int{
		"key": 4,
	}

	if !reflect.DeepEqual(out, expect) {
		t.Fatalf("expected %v, got %v", expect, out)
	}
}

func TestToMap_EmptyCollection(t *testing.T) {
	out := ToMap(
		New([]int{}),
		func(v int) int { return v },
		func(v int) int { return v },
	)

	if len(out) != 0 {
		t.Fatalf("expected empty map, got %v", out)
	}
}

func TestToMap_DoesNotMutateCollection(t *testing.T) {
	items := []int{1, 2, 3}
	c := New(items)

	_ = ToMap(
		c,
		func(v int) int { return v },
		func(v int) int { return v },
	)

	if !reflect.DeepEqual(c.Items(), items) {
		t.Fatalf("collection was mutated: expected %v, got %v", items, c.Items())
	}
}

func TestToMap_MapLengthMatchesUniqueKeys(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}

	out := ToMap(
		New(items),
		func(v int) int { return v % 2 }, // keys: 1,0
		func(v int) int { return v },
	)

	if len(out) != 2 {
		t.Fatalf("expected 2 keys, got %d", len(out))
	}
}
