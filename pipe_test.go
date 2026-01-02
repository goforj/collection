package collection

import (
	"reflect"
	"testing"
)

func TestPipe_ReturnsTransformedValue(t *testing.T) {
	c := New([]int{1, 2, 3})

	result := Pipe(c, func(col *Collection[int]) int {
		sum := 0
		for _, v := range col.Items() {
			sum += v
		}
		return sum
	})

	if result != 6 {
		t.Fatalf("Pipe() expected sum=6, got=%v", result)
	}
}

func TestPipe_CanReturnCollection(t *testing.T) {
	c := New([]int{1, 2, 3})

	result := Pipe(c, func(col *Collection[int]) *Collection[int] {
		return col.Filter(func(v int) bool { return v > 1 })
	})

	out := result.Items()
	want := []int{2, 3}

	if !reflect.DeepEqual(out, want) {
		t.Fatalf("Pipe() returning collection failed. want=%v, got=%v", want, out)
	}
}

func TestPipe_IsNonMutating(t *testing.T) {
	c := New([]int{1, 2, 3})

	_ = Pipe(c, func(col *Collection[int]) *Collection[int] {
		return col
	})

	if !reflect.DeepEqual(c.Items(), []int{1, 2, 3}) {
		t.Fatalf("Pipe() should not mutate original collection")
	}
}

func TestPipe_ReceivesCorrectCollection(t *testing.T) {
	c := New([]string{"a", "b"})

	calledWith := ""

	Pipe(c, func(col *Collection[string]) string {
		calledWith = col.Items()[0]
		return ""
	})

	if calledWith != "a" {
		t.Fatalf("Pipe() did not pass correct collection to callback")
	}
}
