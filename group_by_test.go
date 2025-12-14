package collection

import (
	"reflect"
	"testing"
)

func TestGroupBy_Basic(t *testing.T) {
	values := []int{1, 2, 3, 4, 5}

	groups := GroupBy(
		New(values),
		func(v int) string {
			if v%2 == 0 {
				return "even"
			}
			return "odd"
		},
	)

	if !reflect.DeepEqual(groups["even"].Items(), []int{2, 4}) {
		t.Fatalf("even group incorrect: %v", groups["even"].Items())
	}

	if !reflect.DeepEqual(groups["odd"].Items(), []int{1, 3, 5}) {
		t.Fatalf("odd group incorrect: %v", groups["odd"].Items())
	}
}

func TestGroupBy_Structs(t *testing.T) {
	type User struct {
		ID   int
		Role string
	}

	users := []User{
		{ID: 1, Role: "admin"},
		{ID: 2, Role: "user"},
		{ID: 3, Role: "admin"},
	}

	groups := GroupBy(
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

	if !reflect.DeepEqual(groups["admin"].Items(), expectAdmin) {
		t.Fatalf("admin group incorrect: %v", groups["admin"].Items())
	}

	if !reflect.DeepEqual(groups["user"].Items(), expectUser) {
		t.Fatalf("user group incorrect: %v", groups["user"].Items())
	}
}

func TestGroupBy_EmptyCollection(t *testing.T) {
	groups := GroupBy(
		New([]int{}),
		func(v int) int { return v },
	)

	if len(groups) != 0 {
		t.Fatalf("expected empty groups, got %v", groups)
	}
}

func TestGroupBy_DoesNotMutateSource(t *testing.T) {
	items := []int{1, 2, 3}
	c := New(items)

	_ = GroupBy(
		c,
		func(v int) int { return v },
	)

	if !reflect.DeepEqual(c.Items(), items) {
		t.Fatalf("source collection was mutated")
	}
}
