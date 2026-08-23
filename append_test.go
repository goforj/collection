package collection

import (
	"reflect"
	"testing"
)

func TestAppend(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		c := New([]int{1, 2})

		out := c.Append(3, 4)
		expected := []int{1, 2, 3, 4}

		if !reflect.DeepEqual(out.Items(), expected) {
			t.Fatalf("Append basic expected %v, got %v", expected, out.Items())
		}
	})

	t.Run("EmptyCollection", func(t *testing.T) {
		c := New([]int{})

		out := c.Append(5, 6)
		expected := []int{5, 6}

		if !reflect.DeepEqual(out.Items(), expected) {
			t.Fatalf("Append empty expected %v, got %v", expected, out.Items())
		}
	})

	t.Run("NoValues", func(t *testing.T) {
		c := New([]int{10, 20, 30})

		out := c.Append() // no-op
		expected := []int{10, 20, 30}

		if !reflect.DeepEqual(out.Items(), expected) {
			t.Fatalf("Append no-values expected %v, got %v", expected, out.Items())
		}
	})

	t.Run("NoMutation", func(t *testing.T) {
		orig := []int{1, 2, 3}
		c := New(orig)

		_ = c.Append(4, 5)

		if !reflect.DeepEqual(c.Items(), orig) {
			t.Fatalf("Append mutated original %v", c.Items())
		}
	})
}

func TestAppend_Structs(t *testing.T) {
	type User struct {
		ID   int
		Name string
	}

	c := New([]User{
		{1, "Chris"},
		{2, "Van"},
	})

	out := c.Append(
		User{3, "Shawn"},
		User{4, "Matt"},
	)

	expected := []User{
		{1, "Chris"},
		{2, "Van"},
		{3, "Shawn"},
		{4, "Matt"},
	}

	if !reflect.DeepEqual(out.Items(), expected) {
		t.Fatalf("Append structs expected %v, got %v", expected, out.Items())
	}
}

func TestAppend_ResultAndSourceAreIndependentWithSpareCapacity(t *testing.T) {
	backing := make([]int, 2, 8)
	copy(backing, []int{1, 2})
	c := New(backing)
	out := c.Append(3)
	out[0] = 9
	if c[0] != 1 {
		t.Fatalf("result mutation changed source: %v", c)
	}
	c[1] = 8
	if out[1] != 2 {
		t.Fatalf("source mutation changed result: %v", out)
	}
}
