package collection

import (
	"reflect"
	"testing"
)

func TestMultiply_RepeatsItemsCorrectly(t *testing.T) {
	type User struct {
		Name  string
		Email string
	}

	users := New([]User{
		{"User #1", "user1@example.com"},
		{"User #2", "user2@example.com"},
	})

	out := users.Multiply(3)

	want := []User{
		{"User #1", "user1@example.com"},
		{"User #2", "user2@example.com"},
		{"User #1", "user1@example.com"},
		{"User #2", "user2@example.com"},
		{"User #1", "user1@example.com"},
		{"User #2", "user2@example.com"},
	}

	if !reflect.DeepEqual(out.Items(), want) {
		t.Fatalf("Multiply(3) failed.\nwant=%v\ngot=%v", want, out.Items())
	}
}

func TestMultiply_One_ReturnsSameItems(t *testing.T) {
	c := New([]int{1, 2, 3})
	out := c.Multiply(1)

	if !reflect.DeepEqual(out.Items(), []int{1, 2, 3}) {
		t.Fatalf("Multiply(1) should return original items")
	}
}

func TestMultiply_Zero_ReturnsEmpty(t *testing.T) {
	c := New([]int{1, 2, 3})
	out := c.Multiply(0)

	if len(out.Items()) != 0 {
		t.Fatalf("Multiply(0) should return empty collection")
	}
}

func TestMultiply_Negative_ReturnsEmpty(t *testing.T) {
	c := New([]int{1, 2, 3})
	out := c.Multiply(-2)

	if len(out.Items()) != 0 {
		t.Fatalf("Multiply(-2) should return empty collection")
	}
}

func TestMultiply_IsNonMutating(t *testing.T) {
	c := New([]int{1, 2})
	_ = c.Multiply(3)

	// original should remain unchanged
	if !reflect.DeepEqual(c.Items(), []int{1, 2}) {
		t.Fatalf("Multiply should not mutate the original collection")
	}
}
