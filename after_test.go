package collection

import "testing"

func TestAfter_PredicateMatchSimpleTypes(t *testing.T) {
	nums := New([]int{1, 2, 3, 4, 5})

	after := nums.
		After(func(v int) bool { return v >= 3 }).
		Items()

	expected := []int{4, 5}

	if len(after) != len(expected) {
		t.Fatalf("expected %d items, got %d: %#v", len(expected), len(after), after)
	}

	for i := range expected {
		if after[i] != expected[i] {
			t.Fatalf("expected after[%d] = %d, got %d", i, expected[i], after[i])
		}
	}
}

func TestAfter_PredicateMatchStructs(t *testing.T) {
	type User struct {
		ID   int
		Name string
		Age  int
	}

	users := New([]User{
		{1, "Chris", 34},
		{2, "Van", 42},
		{3, "Shawn", 39},
	})

	after := users.
		After(func(u User) bool { return u.Age >= 40 }).
		Items()

	if len(after) != 1 {
		t.Fatalf("expected 1 user, got %d (%#v)", len(after), after)
	}
	if after[0].Name != "Shawn" {
		t.Fatalf("expected Shawn, got %#v", after[0])
	}
}

func TestAfter_LastElementMatchReturnsEmpty(t *testing.T) {
	nums := New([]int{10, 20, 30})

	after := nums.
		After(func(v int) bool { return v == 30 }).
		Items()

	if len(after) != 0 {
		t.Fatalf("expected empty, got %#v", after)
	}
}

func TestAfter_NoPredicateMatchReturnsEmpty(t *testing.T) {
	nums := New([]int{1, 2, 3})

	after := nums.
		After(func(v int) bool { return v > 100 }).
		Items()

	if len(after) != 0 {
		t.Fatalf("expected empty slice, got %#v", after)
	}
}

func TestAfter_EmptyCollection(t *testing.T) {
	nums := New([]int{})

	after := nums.
		After(func(v int) bool { return true }).
		Items()

	if len(after) != 0 {
		t.Fatalf("expected empty slice, got %#v", after)
	}
}

func TestAfter_SingleElement(t *testing.T) {
	nums := New([]int{5})

	// If match → resulting slice should be empty
	after := nums.
		After(func(v int) bool { return v == 5 }).
		Items()

	if len(after) != 0 {
		t.Fatalf("expected empty slice, got %v", after)
	}

	// If no match → still empty
	after = nums.
		After(func(v int) bool { return v == 99 }).
		Items()

	if len(after) != 0 {
		t.Fatalf("expected empty slice, got %v", after)
	}
}

func TestAfter_DoesNotMutateOriginal(t *testing.T) {
	original := []int{1, 2, 3, 4}
	nums := New(original)

	_ = nums.After(func(v int) bool { return v >= 3 })

	items := nums.Items()
	for i := range original {
		if items[i] != original[i] {
			t.Fatalf("original slice mutated: expected %v, got %v", original, items)
		}
	}
}
