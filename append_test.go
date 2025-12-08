package collection

import (
	"reflect"
	"testing"
)

func TestAppendAndPush(t *testing.T) {
	ops := []struct {
		name string
		fn   func(Collection[int], ...int) Collection[int]
	}{
		{"Append", Collection[int].Append},
		{"Push", Collection[int].Push}, // alias
	}

	for _, op := range ops {
		t.Run(op.name+"_Basic", func(t *testing.T) {
			c := New([]int{1, 2})

			out := op.fn(c, 3, 4)
			expected := []int{1, 2, 3, 4}

			if !reflect.DeepEqual(out.items, expected) {
				t.Fatalf("%s basic expected %v, got %v", op.name, expected, out.items)
			}
		})

		t.Run(op.name+"_EmptyCollection", func(t *testing.T) {
			c := New([]int{})

			out := op.fn(c, 5, 6)
			expected := []int{5, 6}

			if !reflect.DeepEqual(out.items, expected) {
				t.Fatalf("%s empty expected %v, got %v", op.name, expected, out.items)
			}
		})

		t.Run(op.name+"_NoValues", func(t *testing.T) {
			c := New([]int{10, 20, 30})

			out := op.fn(c) // no-op
			expected := []int{10, 20, 30}

			if !reflect.DeepEqual(out.items, expected) {
				t.Fatalf("%s no-values expected %v, got %v", op.name, expected, out.items)
			}
		})

		t.Run(op.name+"_NoMutation", func(t *testing.T) {
			orig := []int{1, 2, 3}
			c := New(orig)

			_ = op.fn(c, 4, 5)

			if !reflect.DeepEqual(c.items, orig) {
				t.Fatalf("%s mutated original %v", op.name, c.items)
			}
		})
	}
}

func TestAppendAndPush_Structs(t *testing.T) {
	type User struct {
		ID   int
		Name string
	}

	ops := []struct {
		name string
		fn   func(Collection[User], ...User) Collection[User]
	}{
		{"Append", Collection[User].Append},
		{"Push", Collection[User].Push}, // alias
	}

	for _, op := range ops {
		t.Run(op.name+"_Structs", func(t *testing.T) {
			c := New([]User{
				{1, "Chris"},
				{2, "Van"},
			})

			out := op.fn(c,
				User{3, "Shawn"},
				User{4, "Matt"},
			)

			expected := []User{
				{1, "Chris"},
				{2, "Van"},
				{3, "Shawn"},
				{4, "Matt"},
			}

			if !reflect.DeepEqual(out.items, expected) {
				t.Fatalf("%s structs expected %v, got %v", op.name, expected, out.items)
			}
		})
	}
}
