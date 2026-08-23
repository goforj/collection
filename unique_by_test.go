package collection

import (
	"reflect"
	"testing"
)

func TestUniqueBy_Basic(t *testing.T) {
	type User struct {
		ID   int
		Name string
	}

	c := New([]User{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
		{ID: 1, Name: "Alice Duplicate"},
	})

	out := c.UniqueBy(func(u User) int { return u.ID })

	expect := Slice[User]{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
	}

	if !reflect.DeepEqual(out, expect) {
		t.Fatalf("expected %+v, got %+v", expect, out)
	}
}

func TestUniqueBy_OrderPreserved(t *testing.T) {
	c := New([]int{3, 1, 2, 1, 3})

	out := c.UniqueBy(func(v int) int { return v })

	expect := Slice[int]{3, 1, 2}
	if !reflect.DeepEqual(out, expect) {
		t.Fatalf("order not preserved: expected %v, got %v", expect, out)
	}
}

func TestUniqueBy_AlreadyUnique(t *testing.T) {
	c := New([]string{"a", "b", "c"})

	out := c.UniqueBy(func(s string) string { return s })

	expect := Slice[string]{"a", "b", "c"}
	if !reflect.DeepEqual(out, expect) {
		t.Fatalf("expected %v, got %v", expect, out)
	}
}

func TestUniqueBy_EmptyCollection(t *testing.T) {
	c := New([]int{})

	out := c.UniqueBy(func(v int) int { return v })

	if len(out) != 0 {
		t.Fatalf("expected empty result, got %v", out)
	}
}

func TestUniqueBy_KeyCollisions(t *testing.T) {
	type Item struct {
		Value int
	}

	c := New([]Item{
		{Value: 1},
		{Value: -1},
		{Value: 2},
		{Value: -2},
	})

	// abs() collapses 1/-1 and 2/-2
	out := c.UniqueBy(func(i Item) int {
		if i.Value < 0 {
			return -i.Value
		}
		return i.Value
	})

	expect := Slice[Item]{
		{Value: 1},
		{Value: 2},
	}

	if !reflect.DeepEqual(out, expect) {
		t.Fatalf("expected %v, got %v", expect, out)
	}
}

func TestUniqueBy_DoesNotMutateOriginal(t *testing.T) {
	orig := []int{1, 2, 2, 3}
	c := New(orig)

	_ = c.UniqueBy(func(v int) int { return v })

	if !reflect.DeepEqual(c, Slice[int](orig)) {
		t.Fatalf("original collection was mutated")
	}
}

func TestUniqueBy_KeyFnCalledOncePerItem(t *testing.T) {
	calls := 0
	c := New([]int{1, 2, 3, 2, 1})

	_ = c.UniqueBy(func(v int) int {
		calls++
		return v
	})

	if calls != len(c) {
		t.Fatalf(
			"expected keyFn to be called %d times, got %d",
			len(c),
			calls,
		)
	}
}
