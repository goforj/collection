package collection

import (
	"reflect"
	"testing"
)

func TestFilter_Ints(t *testing.T) {
	c := New([]int{1, 2, 3, 4, 5})

	filtered := c.Filter(func(v int) bool {
		return v%2 == 0 // keep even
	})

	expected := []int{2, 4}

	if !reflect.DeepEqual(filtered.Items(), expected) {
		t.Fatalf("expected %v, got %v", expected, filtered.Items())
	}
}

func TestFilter_NoneMatch(t *testing.T) {
	c := New([]int{1, 2, 3})

	filtered := c.Filter(func(v int) bool { return v > 100 })

	if len(filtered.Items()) != 0 {
		t.Fatalf("expected empty result, got %v", filtered.Items())
	}
}

func TestFilter_AllMatch(t *testing.T) {
	c := New([]int{1, 2, 3})

	filtered := c.Filter(func(v int) bool { return true })

	expected := []int{1, 2, 3}

	if !reflect.DeepEqual(filtered.Items(), expected) {
		t.Fatalf("expected %v, got %v", expected, filtered.Items())
	}
}

func TestFilter_Structs(t *testing.T) {
	type User struct {
		ID   int
		Name string
	}

	c := New([]User{
		{1, "Chris"},
		{2, "Van"},
		{3, "Shawn"},
	})

	filtered := c.Filter(func(u User) bool {
		return u.ID >= 2
	})

	expected := []User{
		{2, "Van"},
		{3, "Shawn"},
	}

	if !reflect.DeepEqual(filtered.Items(), expected) {
		t.Fatalf("expected %v, got %v", expected, filtered.Items())
	}
}

func TestFilter_EmptyInput(t *testing.T) {
	c := New([]int{})

	filtered := c.Filter(func(v int) bool { return v%2 == 0 })

	if len(filtered.Items()) != 0 {
		t.Fatalf("expected empty slice, got %v", filtered.Items())
	}
}

func TestFilter_Chaining(t *testing.T) {
	c := New([]int{1, 2, 3, 4, 5})

	result := c.
		Filter(func(v int) bool { return v > 1 }).   // [2,3,4,5]
		Filter(func(v int) bool { return v%2 == 1 }) // [3,5]

	expected := []int{3, 5}

	if !reflect.DeepEqual(result.Items(), expected) {
		t.Fatalf("expected %v, got %v", expected, result.Items())
	}
}

func TestFilter_PreservesNilSlice(t *testing.T) {
	c := New([]int(nil))

	c.Filter(func(v int) bool { return v%2 == 0 })

	if c.Items() != nil {
		t.Fatalf("expected nil slice to remain nil, got %v", c.Items())
	}
}

func TestFilter_CompactsSourceAndRequiresReturnedHeader(t *testing.T) {
	items := []int{1, 2, 3, 4}
	c := New(items)

	out := c.Filter(func(v int) bool { return v%2 == 0 })

	want := []int{2, 4}
	if !reflect.DeepEqual(out.Items(), want) {
		t.Fatalf("expected result %v, got %v", want, out)
	}

	if !reflect.DeepEqual(items, []int{2, 4, 0, 0}) {
		t.Fatalf("Filter did not compact and clear input storage: %v", items)
	}
	if len(c) != 4 {
		t.Fatalf("ignored result changed source header length to %d", len(c))
	}
	if len(out) != 2 {
		t.Fatalf("returned header length = %d, want 2", len(out))
	}
}

func TestFilter_ClearsDiscardedReferences(t *testing.T) {
	a, b, cval := 1, 2, 3
	items := []*int{&a, &b, &cval}
	c := New(items)

	out := c.Filter(func(v *int) bool { return *v == 2 })
	if len(out) != 1 || out[0] != &b || items[1] != nil || items[2] != nil {
		t.Fatalf("result=%v source=%v", out, items)
	}
}

func TestFilter_ZeroAllocs(t *testing.T) {
	items := []int{1, 2, 3, 4}
	c := New(items)
	allocs := testing.AllocsPerRun(1000, func() {
		copy(c, []int{1, 2, 3, 4})
		_ = c.Filter(func(v int) bool { return v%2 == 0 })
	})
	if allocs != 0 {
		t.Fatalf("Filter allocations = %v, want 0", allocs)
	}
}
