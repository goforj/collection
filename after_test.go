package collection

import "testing"

func TestAfterPredicateMatch(t *testing.T) {
	values := New([]int{1, 2, 3, 4, 5})
	got := values.After(func(value int) bool { return value >= 3 })
	want := []int{4, 5}
	if !slicesEqual(got, want) {
		t.Fatalf("After() = %v, want %v", got, want)
	}
}

func TestAfterStructs(t *testing.T) {
	type user struct{ age int }
	values := New([]user{{34}, {42}, {39}})
	got := values.After(func(value user) bool { return value.age >= 40 })
	want := []user{{39}}
	if len(got) != len(want) || got[0] != want[0] {
		t.Fatalf("After() = %v, want %v", got, want)
	}
}

func TestAfterNoFollowingItems(t *testing.T) {
	for _, values := range []Slice[int]{New([]int{10, 20, 30}), New([]int{1, 2, 3}), New([]int{}), New([]int{5})} {
		got := values.After(func(value int) bool { return value == 30 || value == 5 })
		if len(got) != 0 || cap(got) != 0 {
			t.Fatalf("After() = len %d cap %d, want capped empty", len(got), cap(got))
		}
	}
}

func TestAfterReturnsCappedViewWithoutMutatingSource(t *testing.T) {
	items := []int{1, 2, 3, 4}
	got := New(items).After(func(value int) bool { return value == 2 })
	got[0] = 99
	if items[2] != 99 || cap(got) != len(got) {
		t.Fatalf("After() did not return a capped view: got %v cap %d, source %v", got, cap(got), items)
	}
}
