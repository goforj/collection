package collection

import "testing"

func TestContains_ValueMatch(t *testing.T) {
	nums := New([]int{1, 2, 3})

	if !nums.Contains(func(v int) bool { return v == 2 }) {
		t.Fatalf("expected true, got false")
	}
}

func TestContains_NoMatch(t *testing.T) {
	nums := New([]int{1, 2, 3})

	if nums.Contains(func(v int) bool { return v == 99 }) {
		t.Fatalf("expected false, got true")
	}
}

func TestContains_EmptyCollection(t *testing.T) {
	nums := New([]int{})

	if nums.Contains(func(v int) bool { return true }) {
		t.Fatalf("expected false for empty collection")
	}
}

func TestContains_Structs(t *testing.T) {
	type User struct {
		ID   int
		Name string
	}

	users := New([]User{
		{1, "Chris"},
		{2, "Van"},
		{3, "Shawn"},
	})

	if !users.Contains(func(u User) bool { return u.Name == "Van" }) {
		t.Fatalf("expected true, got false")
	}

	if users.Contains(func(u User) bool { return u.Name == "Zach" }) {
		t.Fatalf("expected false, got true")
	}
}
