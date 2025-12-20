package collection

import (
	"reflect"
	"testing"
)

func TestGroupBySlice_Basic(t *testing.T) {
	values := []int{1, 2, 3, 4, 5}

	groups := GroupBySlice(
		New(values),
		func(v int) string {
			if v%2 == 0 {
				return "even"
			}
			return "odd"
		},
	)

	if !reflect.DeepEqual(groups["even"], []int{2, 4}) {
		t.Fatalf("even group incorrect: %v", groups["even"])
	}

	if !reflect.DeepEqual(groups["odd"], []int{1, 3, 5}) {
		t.Fatalf("odd group incorrect: %v", groups["odd"])
	}
}

func TestGroupBySlice_Structs(t *testing.T) {
	type User struct {
		ID   int
		Role string
	}

	users := []User{
		{ID: 1, Role: "admin"},
		{ID: 2, Role: "user"},
		{ID: 3, Role: "admin"},
	}

	groups := GroupBySlice(
		New(users),
		func(u User) string { return u.Role },
	)

	expectAdmin := []User{
		{ID: 1, Role: "admin"},
		{ID: 3, Role: "admin"},
	}

	expectUser := []User{
		{ID: 2, Role: "user"},
	}

	if !reflect.DeepEqual(groups["admin"], expectAdmin) {
		t.Fatalf("admin group incorrect: %v", groups["admin"])
	}

	if !reflect.DeepEqual(groups["user"], expectUser) {
		t.Fatalf("user group incorrect: %v", groups["user"])
	}
}

func TestGroupBySlice_EmptyCollection(t *testing.T) {
	groups := GroupBySlice(
		New([]int{}),
		func(v int) int { return v },
	)

	if len(groups) != 0 {
		t.Fatalf("expected empty groups, got %v", groups)
	}
}

func TestGroupBySlice_DoesNotMutateSource(t *testing.T) {
	items := []int{1, 2, 3}
	c := New(items)

	groups := GroupBySlice(
		c,
		func(v int) int { return v },
	)

	if !reflect.DeepEqual(c.Items(), items) {
		t.Fatalf("source collection was mutated")
	}

	groups[1][0] = 99
	if !reflect.DeepEqual(c.Items(), items) {
		t.Fatalf("group mutation should not affect source collection")
	}
}
