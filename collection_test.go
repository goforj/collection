package collection

import (
	"testing"
)

type User struct {
	ID   int
	Name string
	Age  int
}

func TestFluentChainWithStructs(t *testing.T) {
	users := New([]User{
		{1, "Chris", 34},
		{2, "Van", 42},
		{3, "Shawn", 39},
	})

	// Fluent chain across SAME type:
	filteredAndSorted := users.
		Filter(func(u User) bool { return u.Age >= 35 }).
		Sort(func(a, b User) bool { return a.Age < b.Age })

	// Type change happens at the edge using MapTo/Pluck.
	names := Pluck(filteredAndSorted, func(u User) string {
		return u.Name
	}).Items()

	if len(names) != 2 {
		t.Fatalf("expected 2 names, got %d (%v)", len(names), names)
	}
	if names[0] != "Shawn" || names[1] != "Van" {
		t.Fatalf("unexpected order: %#v", names)
	}
}

func TestUnique(t *testing.T) {
	nums := New([]int{1, 2, 2, 3, 3, 3, 4})

	unique := nums.Unique(func(a, b int) bool { return a == b }).Items()

	if len(unique) != 4 {
		t.Fatalf("expected 4 unique items, got %d (%v)", len(unique), unique)
	}
}

func TestFluentChainWithStructsDump(t *testing.T) {
	users := New([]User{
		{1, "Chris", 34},
		{2, "Van", 42},
		{3, "Shawn", 39},
	})

	// Fluent chain across SAME type:
	users.
		Filter(func(u User) bool { return u.Age >= 35 }).
		Sort(func(a, b User) bool { return a.Age < b.Age })
}
