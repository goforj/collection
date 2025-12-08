package collection

import (
	"reflect"
	"testing"
)

func TestReduce_IntSum(t *testing.T) {
	c := New([]int{1, 2, 3, 4})

	sum := Reduce(c, 0, func(acc, v int) int {
		return acc + v
	})

	if sum != 10 {
		t.Fatalf("expected 10, got %v", sum)
	}
}

func TestReduce_StringConcat(t *testing.T) {
	c := New([]string{"a", "b", "c"})

	out := Reduce(c, "", func(acc, v string) string {
		return acc + v
	})

	if out != "abc" {
		t.Fatalf("expected \"abc\", got %q", out)
	}
}

func TestReduce_EmptyCollection(t *testing.T) {
	c := New([]int{})

	sum := Reduce(c, 10, func(acc, v int) int {
		return acc + v
	})

	// Should return the initial value unchanged
	if sum != 10 {
		t.Fatalf("expected 10, got %v", sum)
	}
}

func TestReduce_StructAccumulator(t *testing.T) {
	type User struct {
		ID   int
		Name string
	}

	c := New([]User{
		{1, "Chris"},
		{2, "Van"},
		{3, "Shawn"},
	})

	acc := Reduce(c, []string{}, func(acc []string, u User) []string {
		return append(acc, u.Name)
	})

	expected := []string{"Chris", "Van", "Shawn"}

	if !reflect.DeepEqual(acc, expected) {
		t.Fatalf("expected %v, got %v", expected, acc)
	}
}

func TestReduce_MapAccumulator(t *testing.T) {
	c := New([]int{2, 4, 6})

	acc := Reduce(c, map[int]int{}, func(acc map[int]int, v int) map[int]int {
		acc[v] = v * 10
		return acc
	})

	expected := map[int]int{2: 20, 4: 40, 6: 60}

	if !reflect.DeepEqual(acc, expected) {
		t.Fatalf("expected %v, got %v", expected, acc)
	}
}

func TestReduce_NoMutation(t *testing.T) {
	c := New([]int{1, 2, 3})
	orig := append([]int{}, c.items...)

	_ = Reduce(c, 0, func(acc, v int) int {
		return acc + v
	})

	if !reflect.DeepEqual(c.items, orig) {
		t.Fatalf("Reduce mutated the original collection: %v vs %v", c.items, orig)
	}
}
