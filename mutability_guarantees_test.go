package collection

import (
	"reflect"
	"testing"
)

func TestMutators_PreserveNilSlice(t *testing.T) {
	cases := []struct {
		name string
		fn   func(*Collection[int])
	}{
		{
			name: "Map",
			fn: func(c *Collection[int]) {
				c.Map(func(v int) int { return v * 2 })
			},
		},
		{
			name: "Filter",
			fn: func(c *Collection[int]) {
				c.Filter(func(v int) bool { return v%2 == 0 })
			},
		},
		{
			name: "Sort",
			fn: func(c *Collection[int]) {
				c.Sort(func(a, b int) bool { return a < b })
			},
		},
		{
			name: "Reverse",
			fn: func(c *Collection[int]) {
				c.Reverse()
			},
		},
		{
			name: "Shuffle",
			fn: func(c *Collection[int]) {
				c.Shuffle()
			},
		},
		{
			name: "Transform",
			fn: func(c *Collection[int]) {
				c.Transform(func(v int) int { return v + 1 })
			},
		},
		{
			name: "Pop",
			fn: func(c *Collection[int]) {
				_, _ = c.Pop()
			},
		},
		{
			name: "PopN",
			fn: func(c *Collection[int]) {
				_ = c.PopN(2)
			},
		},
		{
			name: "ConcatEmpty",
			fn: func(c *Collection[int]) {
				c.Concat([]int{})
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			c := New([]int(nil))
			tc.fn(c)

			if c.Items() != nil {
				t.Fatalf("expected nil slice to remain nil, got %v", c.Items())
			}
		})
	}
}

func TestMap_WritesThroughSourceSlice(t *testing.T) {
	items := []int{1, 2, 3}
	c := New(items)

	c.Map(func(v int) int { return v * 2 })

	want := []int{2, 4, 6}
	if !reflect.DeepEqual(items, want) {
		t.Fatalf("expected source slice %v, got %v", want, items)
	}
}

func TestFilter_WritesThroughSourceSlice(t *testing.T) {
	items := []int{1, 2, 3, 4}
	c := New(items)

	c.Filter(func(v int) bool { return v%2 == 0 })

	want := []int{2, 4}
	if !reflect.DeepEqual(items[:len(want)], want) {
		t.Fatalf("expected source prefix %v, got %v", want, items[:len(want)])
	}

	if len(c.Items()) != len(want) {
		t.Fatalf("expected filtered length %d, got %d", len(want), len(c.Items()))
	}
}

func TestSort_WritesThroughSourceSlice(t *testing.T) {
	items := []int{3, 1, 2}
	c := New(items)

	c.Sort(func(a, b int) bool { return a < b })

	want := []int{1, 2, 3}
	if !reflect.DeepEqual(items, want) {
		t.Fatalf("expected source slice %v, got %v", want, items)
	}
}

func TestTransform_WritesThroughSourceSlice(t *testing.T) {
	items := []int{1, 2, 3}
	c := New(items)

	c.Transform(func(v int) int { return v + 10 })

	want := []int{11, 12, 13}
	if !reflect.DeepEqual(items, want) {
		t.Fatalf("expected source slice %v, got %v", want, items)
	}
}

func TestReverse_WritesThroughSourceSlice(t *testing.T) {
	items := []int{1, 2, 3}
	c := New(items)

	c.Reverse()

	want := []int{3, 2, 1}
	if !reflect.DeepEqual(items, want) {
		t.Fatalf("expected source slice %v, got %v", want, items)
	}
}

func TestMutators_LengthGuarantees(t *testing.T) {
	cases := []struct {
		name    string
		input   []int
		fn      func(*Collection[int]) any
		wantLen int
	}{
		{
			name:    "Map",
			input:   []int{1, 2, 3},
			fn:      func(c *Collection[int]) any { return c.Map(func(v int) int { return v * 2 }) },
			wantLen: 3,
		},
		{
			name:    "Filter",
			input:   []int{1, 2, 3, 4},
			fn:      func(c *Collection[int]) any { return c.Filter(func(v int) bool { return v%2 == 0 }) },
			wantLen: 2,
		},
		{
			name:    "Sort",
			input:   []int{3, 1, 2},
			fn:      func(c *Collection[int]) any { return c.Sort(func(a, b int) bool { return a < b }) },
			wantLen: 3,
		},
		{
			name:    "Transform",
			input:   []int{1, 2, 3},
			fn:      func(c *Collection[int]) any { c.Transform(func(v int) int { return v + 1 }); return nil },
			wantLen: 3,
		},
		{
			name:    "Reverse",
			input:   []int{1, 2, 3},
			fn:      func(c *Collection[int]) any { return c.Reverse() },
			wantLen: 3,
		},
		{
			name:    "Shuffle",
			input:   []int{1, 2, 3, 4},
			fn:      func(c *Collection[int]) any { return c.Shuffle() },
			wantLen: 4,
		},
		{
			name:    "Pop",
			input:   []int{1, 2, 3},
			fn:      func(c *Collection[int]) any { _, _ = c.Pop(); return nil },
			wantLen: 2,
		},
		{
			name:    "PopN",
			input:   []int{1, 2, 3, 4},
			fn:      func(c *Collection[int]) any { _ = c.PopN(2); return nil },
			wantLen: 2,
		},
		{
			name:    "Prepend",
			input:   []int{3, 4},
			fn:      func(c *Collection[int]) any { return c.Prepend(1, 2) },
			wantLen: 4,
		},
		{
			name:    "Concat",
			input:   []int{1, 2},
			fn:      func(c *Collection[int]) any { return c.Concat([]int{3, 4, 5}) },
			wantLen: 5,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			c := New(append([]int(nil), tc.input...))
			tc.fn(c)
			if len(c.items) != tc.wantLen {
				t.Fatalf("expected len %d, got %d", tc.wantLen, len(c.items))
			}
		})
	}
}
