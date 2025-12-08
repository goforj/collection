package collection

import (
	"reflect"
	"testing"
)

func TestFindWhere_FindsMatchingValue(t *testing.T) {
	c := New([]int{1, 2, 3, 4, 5})

	v, ok := c.FindWhere(func(n int) bool { return n%2 == 0 })

	if !ok {
		t.Fatalf("FindWhere should return ok=true when match exists")
	}

	if v != 2 {
		t.Fatalf("FindWhere returned wrong value. expected 2, got %v", v)
	}
}

func TestFindWhere_NoMatchReturnsZeroValue(t *testing.T) {
	c := New([]int{1, 3, 5})

	v, ok := c.FindWhere(func(n int) bool { return n%2 == 0 })

	if ok {
		t.Fatalf("FindWhere should return ok=false when no match exists")
	}

	if v != 0 {
		t.Fatalf("FindWhere should return zero-value when no match exists; got %v", v)
	}
}

func TestFindWhere_AliasToFirst(t *testing.T) {
	c := New([]int{1, 2, 3, 4, 5})

	v1, ok1 := c.FirstWhere(func(n int) bool { return n > 3 })
	v2, ok2 := c.FindWhere(func(n int) bool { return n > 3 })

	if v1 != v2 || ok1 != ok2 {
		t.Fatalf("FindWhere should behave exactly the same as FirstWhere(fn)")
	}
}

func TestFindWhere_WorksWithStructs(t *testing.T) {
	type User struct {
		ID   int
		Name string
	}

	c := New([]User{
		{ID: 1, Name: "A"},
		{ID: 2, Name: "B"},
		{ID: 3, Name: "C"},
	})

	expected := User{ID: 2, Name: "B"}

	v, ok := c.FindWhere(func(u User) bool {
		return u.ID == 2
	})

	if !ok {
		t.Fatalf("FindWhere should locate struct matching predicate")
	}

	if !reflect.DeepEqual(v, expected) {
		t.Fatalf("FindWhere returned wrong struct. expected %#v, got %#v", expected, v)
	}
}

func TestFindWhere_NoStructMatch(t *testing.T) {
	type User struct {
		ID   int
		Name string
	}

	c := New([]User{
		{ID: 1, Name: "A"},
		{ID: 2, Name: "B"},
	})

	v, ok := c.FindWhere(func(u User) bool {
		return u.ID == 999
	})

	if ok {
		t.Fatalf("FindWhere should return ok=false when no struct matches")
	}

	if (v != User{}) {
		t.Fatalf("FindWhere should return zero-value struct when no match")
	}
}
