package collection

import (
	"reflect"
	"testing"
)

func TestWhere_AliasToFilter(t *testing.T) {
	fn := func(v int) bool { return v%2 == 0 }

	c1 := New([]int{1, 2, 3, 4, 5})
	c2 := New([]int{1, 2, 3, 4, 5})

	out1 := c1.Filter(fn)
	out2 := c2.Where(fn)

	if out2 != c2 {
		t.Fatalf("Where should return the same collection instance")
	}

	if !reflect.DeepEqual(out1.items, out2.items) {
		t.Fatalf("Where should behave exactly the same as Filter(fn)")
	}
}
