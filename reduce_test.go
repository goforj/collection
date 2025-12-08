package collection

import "testing"

func TestReduce_IntSum(t *testing.T) {
	c := New([]int{1, 2, 3, 4})

	sum := c.Reduce(0, func(acc, n int) int {
		return acc + n
	})

	if sum != 10 {
		t.Fatalf("expected 10, got %d", sum)
	}
}

func TestReduce_StringConcat(t *testing.T) {
	c := New([]string{"go", "forj", "rocks"})

	out := c.Reduce("", func(acc, s string) string {
		if acc == "" {
			return s
		}
		return acc + " " + s
	})

	if out != "go forj rocks" {
		t.Fatalf(`expected "go forj rocks", got %q`, out)
	}
}

func TestReduce_EmptyCollectionReturnsInitial(t *testing.T) {
	c := New([]int{})

	out := c.Reduce(42, func(acc, n int) int {
		return acc + n
	})

	if out != 42 {
		t.Fatalf("expected initial value 42 for empty collection, got %d", out)
	}
}

func TestReduce_OrderIsLeftToRight(t *testing.T) {
	c := New([]string{"a", "b", "c"})

	out := c.Reduce("", func(acc, s string) string {
		return acc + s
	})

	if out != "abc" {
		t.Fatalf(`expected "abc" (left-to-right accumulation), got %q`, out)
	}
}
