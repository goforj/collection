package collection

import (
	"reflect"
	"testing"
)

func TestMerge_AppendsSlice(t *testing.T) {
	c := New([]string{"Desk", "Chair"})
	merged := c.Merge([]string{"Bookcase", "Door"})

	want := []string{"Desk", "Chair", "Bookcase", "Door"}
	got := merged.Items()

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Merge([]T) append failed.\nwant=%v\ngot=%v", want, got)
	}
}

func TestMerge_AppendsCollection(t *testing.T) {
	c1 := New([]int{1, 2})
	c2 := New([]int{3, 4})

	out := c1.Merge(c2)
	want := []int{1, 2, 3, 4}

	if !reflect.DeepEqual(out.Items(), want) {
		t.Fatalf("Merge(Collection) failed.\nwant=%v\ngot=%v", want, out.Items())
	}
}

func TestMerge_AssociativeOverwrite(t *testing.T) {
	type Product struct {
		ID    int
		Price int
	}

	// emulate Laravel example:
	// ['product_id' => 1, 'price' => 100]
	c := New([]Product{
		{ID: 1, Price: 100},
	})

	// associative merge overwrites existing keys
	merged := c.Merge(map[string]Product{
		"0":        {ID: 1, Price: 200}, // overwrites index 0
		"discount": {ID: 0, Price: 0},   // new key
	})

	got := merged.Items()

	// order is not guaranteed for associative maps — so we verify membership
	foundPrice200 := false
	foundDiscount := false

	for _, v := range got {
		if v.Price == 200 {
			foundPrice200 = true
		}
		if v.ID == 0 && v.Price == 0 {
			foundDiscount = true
		}
	}

	if !foundPrice200 {
		t.Fatalf("Merge(map[string]T) did not overwrite existing key with new value")
	}
	if !foundDiscount {
		t.Fatalf("Merge(map[string]T) did not include associative new key")
	}
}

func TestMerge_UnsupportedTypeReturnsOriginal(t *testing.T) {
	c := New([]int{1, 2, 3})

	// Unsupported types:
	out1 := c.Merge(123)
	out2 := c.Merge("not valid")
	out3 := c.Merge(struct{}{})

	want := []int{1, 2, 3}

	// out1
	if !reflect.DeepEqual(out1.Items(), want) {
		t.Fatalf("Merge(unsupported int) should return original collection.\nwant=%v\ngot=%v",
			want, out1.Items())
	}

	// out2
	if !reflect.DeepEqual(out2.Items(), want) {
		t.Fatalf("Merge(unsupported string) should return original collection.\nwant=%v\ngot=%v",
			want, out2.Items())
	}

	// out3
	if !reflect.DeepEqual(out3.Items(), want) {
		t.Fatalf("Merge(unsupported struct) should return original collection.\nwant=%v\ngot=%v",
			want, out3.Items())
	}
}
