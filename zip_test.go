package collection

import (
	"reflect"
	"testing"
)

func TestZip(t *testing.T) {
	got := New([]int{1, 2, 3}).Zip([]string{"one", "two"})
	want := []Pair[int, string]{{First: 1, Second: "one"}, {First: 2, Second: "two"}}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Zip() = %v, want %v", got, want)
	}
}

func TestZipStructsAndEmpty(t *testing.T) {
	type user struct{ id int }
	type role struct{ name string }
	got := New([]user{{1}, {2}}).Zip([]role{{"admin"}, {"user"}, {"extra"}})
	want := []Pair[user, role]{{First: user{1}, Second: role{"admin"}}, {First: user{2}, Second: role{"user"}}}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Zip() = %v, want %v", got, want)
	}
	if got := New([]int{}).Zip([]int{1, 2}); len(got) != 0 {
		t.Fatalf("Zip() = %v, want empty", got)
	}
}
