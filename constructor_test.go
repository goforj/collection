package collection

import "testing"

func TestNew_BorrowsInputSlice(t *testing.T) {
	items := []int{1, 2, 3}
	c := New(items)

	items[0] = 9

	if c.Items()[0] != 9 {
		t.Fatalf("New should borrow input slice")
	}
}

func TestNew_PreservesNilSlice(t *testing.T) {
	var items []int
	c := New(items)

	if c.Items() != nil {
		t.Fatalf("New should preserve nil slice")
	}
}

func TestNewNumeric_PreservesNilSlice(t *testing.T) {
	var items []int
	c := NewNumeric(items)

	if c.Items() != nil {
		t.Fatalf("NewNumeric should preserve nil slice")
	}
}

func TestNewNumeric_BorrowsInputSlice(t *testing.T) {
	items := []int{1, 2, 3}
	c := NewNumeric(items)

	items[0] = 9

	if c.Items()[0] != 9 {
		t.Fatalf("NewNumeric should borrow input slice")
	}
}

func TestSelectionOps_ShareBackingSlice(t *testing.T) {
	items := []int{1, 2, 3, 4}
	c := New(items)

	view := c.Take(2)
	items[0] = 9

	if view.Items()[0] != 9 {
		t.Fatalf("selection ops should return views")
	}
}

func TestItemsCopy_ReturnsCopy(t *testing.T) {
	items := []int{1, 2, 3}
	c := New(items)

	copyItems := c.ItemsCopy()
	copyItems[0] = 9

	if c.Items()[0] == 9 {
		t.Fatalf("ItemsCopy should return a copy")
	}
}
