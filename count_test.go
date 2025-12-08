package collection

import "testing"

func TestCount_Basic(t *testing.T) {
	nums := New([]int{1, 2, 3, 4})

	if nums.Count() != 4 {
		t.Fatalf("expected 4, got %d", nums.Count())
	}
}

func TestCount_Empty(t *testing.T) {
	nums := New([]int{})

	if nums.Count() != 0 {
		t.Fatalf("expected 0, got %d", nums.Count())
	}
}

func TestCount_Structs(t *testing.T) {
	type User struct {
		ID   int
		Name string
	}

	users := New([]User{
		{1, "Chris"},
		{2, "Van"},
		{3, "Shawn"},
	})

	if users.Count() != 3 {
		t.Fatalf("expected 3, got %d", users.Count())
	}
}
