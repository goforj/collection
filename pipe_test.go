package collection

import (
	"reflect"
	"testing"
)

func TestPipe_ReturnsTransformedValue(t *testing.T) {
	c := New([]int{1, 2, 3})

	result := c.Pipe(func(col *Collection[int]) any {
		sum := 0
		for _, v := range col.Items() {
			sum += v
		}
		return sum
	})

	if result.(int) != 6 {
		t.Fatalf("Pipe() expected sum=6, got=%v", result)
	}
}

func TestPipe_CanReturnCollection(t *testing.T) {
	c := New([]int{1, 2, 3})

	result := c.Pipe(func(col *Collection[int]) any {
		return col.Filter(func(v int) bool { return v > 1 })
	})

	out := result.(*Collection[int]).Items()
	want := []int{2, 3}

	if !reflect.DeepEqual(out, want) {
		t.Fatalf("Pipe() returning collection failed. want=%v, got=%v", want, out)
	}
}

func TestPipe_IsNonMutating(t *testing.T) {
	c := New([]int{1, 2, 3})

	_ = c.Pipe(func(col *Collection[int]) any {
		return col.Map(func(v int) int { return v * 2 })
	})

	// original must remain unchanged
	want := []int{1, 2, 3}
	if !reflect.DeepEqual(c.Items(), want) {
		t.Fatalf("Pipe() mutated original collection, want=%v got=%v", want, c.Items())
	}
}

func TestPipe_ReceivesCorrectCollection(t *testing.T) {
	c := New([]string{"a", "b"})

	calledWith := ""

	c.Pipe(func(col *Collection[string]) any {
		calledWith = col.Items()[0]
		return nil
	})

	if calledWith != "a" {
		t.Fatalf("Pipe() did not pass correct collection to callback")
	}
}
