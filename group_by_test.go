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
	var _ map[string]Slice[int] = groups

	if !reflect.DeepEqual(groups["even"], Slice[int]{2, 4}) {
		t.Fatalf("even group incorrect: %v", groups["even"])
	}

	if !reflect.DeepEqual(groups["odd"], Slice[int]{1, 3, 5}) {
		t.Fatalf("odd group incorrect: %v", groups["odd"])
	}
	filtered := groups["odd"].Filter(func(value int) bool { return value > 1 })
	if !reflect.DeepEqual(filtered, Slice[int]{3, 5}) {
		t.Fatalf("fluent group Filter result = %v", filtered)
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

	expectAdmin := Slice[User]{
		{ID: 1, Role: "admin"},
		{ID: 3, Role: "admin"},
	}

	expectUser := Slice[User]{
		{ID: 2, Role: "user"},
	}

	if !reflect.DeepEqual(groups["admin"], expectAdmin) {
		t.Fatalf("admin group incorrect: %v", groups["admin"])
	}

	if !reflect.DeepEqual(groups["user"], expectUser) {
		t.Fatalf("user group incorrect: %v", groups["user"])
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
