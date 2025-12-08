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

// fastParseInt edge cases
func TestFastParseInt_EdgeCases(t *testing.T) {
	tests := []struct {
		in   string
		want int
		ok   bool
	}{
		{"", 0, false},
		{"abc", 0, false},
		{"12a", 0, false},
		{"-1", 0, false},     // minus not allowed
		{"00123", 123, true}, // leading zeros should still parse
		{"9", 9, true},
	}

	for _, tc := range tests {
		got, ok := fastParseInt(tc.in)
		if got != tc.want || ok != tc.ok {
			t.Fatalf("fastParseInt(%q) = (%d,%v), want (%d,%v)",
				tc.in, got, ok, tc.want, tc.ok)
		}
	}
}

// map[string]T → ONLY non-numeric keys
func TestMergeMap_OnlyStringKeys(t *testing.T) {
	c := New([]int{1, 2})

	out := c.Merge(map[string]int{
		"a": 10,
		"b": 20,
	})

	got := out.Items()

	if len(got) != 4 {
		t.Fatalf("expected appended values, got len=%d", len(got))
	}

	if !(got[0] == 1 && got[1] == 2 && (got[2] == 10 || got[3] == 20)) {
		t.Fatalf("string-only merge incorrect: %v", got)
	}
}

// numeric keys in-range only → pure overwrite, no append
func TestMergeMap_OnlyOverwrite(t *testing.T) {
	c := New([]string{"a", "b", "c"})

	out := c.Merge(map[string]string{
		"0": "X",
		"2": "Z",
	})

	got := out.Items()

	want := []string{"X", "b", "Z"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("overwrite-only merge failed\nwant=%v\ngot=%v", want, got)
	}
}

// numeric keys out-of-range only → pure append
func TestMergeMap_OutOfRangeKeys(t *testing.T) {
	c := New([]int{1, 2})

	out := c.Merge(map[string]int{
		"5": 100,
		"9": 200,
	})

	got := out.Items()

	if !reflect.DeepEqual(got, []int{1, 2, 100, 200}) {
		t.Fatalf("append on out-of-range keys failed: %v", got)
	}
}

// mixed case: overwrite + append + string keys
func TestMergeMap_MixedCases(t *testing.T) {
	c := New([]int{10, 20})

	out := c.Merge(map[string]int{
		"0":   99,  // overwrite index 0
		"5":   500, // append
		"foo": 7,   // append
	})

	got := out.Items()

	if !(got[0] == 99) {
		t.Fatalf("expected overwrite at index 0: %v", got)
	}
	if len(got) != 4 {
		t.Fatalf("expected total len 4, got %d", len(got))
	}
}

// empty map → return unchanged slice
func TestMergeMap_EmptyMap(t *testing.T) {
	c := New([]int{1, 2, 3})

	out := c.Merge(map[string]int{})

	if !reflect.DeepEqual(out.Items(), c.Items()) {
		t.Fatalf("empty map should return same items")
	}
}
