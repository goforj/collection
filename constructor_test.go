package collection

import "testing"

func TestNew_BorrowsInputSlice(t *testing.T) {
	items := []int{1, 2, 3}
	c := New(items)

	items[0] = 9

	if c[0] != 9 {
		t.Fatalf("New should borrow input slice")
	}
}

func TestNew_PreservesNilSlice(t *testing.T) {
	var items []int
	c := New(items)

	if c != nil {
		t.Fatalf("New should preserve nil slice")
	}
}

func TestSelectionOps_ShareBackingSlice(t *testing.T) {
	items := []int{1, 2, 3, 4}
	c := New(items)

	view := c.Take(2)
	items[0] = 9

	if view[0] != 9 {
		t.Fatalf("selection ops should return views")
	}
}

func TestClone_ReturnsCopy(t *testing.T) {
	items := []int{1, 2, 3}
	c := New(items)

	copyItems := c.Clone()
	copyItems[0] = 9

	if c[0] == 9 {
		t.Fatalf("Clone should return a copy")
	}
}
