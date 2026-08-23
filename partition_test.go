package collection

import "testing"

func TestPartition(t *testing.T) {
	values := New([]int{1, 2, 3, 4, 5})
	left, right := values.Partition(func(value int) bool { return value%2 == 0 })
	if !slicesEqual(left, []int{2, 4}) || !slicesEqual(right, []int{1, 3, 5}) {
		t.Fatalf("Partition() = %v, %v", left, right)
	}
}

func TestPartitionStringsAndStructs(t *testing.T) {
	words := New([]string{"go", "gopher", "rust", "ruby"})
	long, short := words.Partition(func(word string) bool { return len(word) >= 4 })
	if !slicesEqual(long, []string{"gopher", "rust", "ruby"}) || !slicesEqual(short, []string{"go"}) {
		t.Fatalf("Partition() = %v, %v", long, short)
	}
	type user struct{ active bool }
	active, inactive := New([]user{{true}, {false}, {true}}).Partition(func(value user) bool { return value.active })
	if len(active) != 2 || len(inactive) != 1 || !active[0].active || inactive[0].active {
		t.Fatalf("Partition() = %v, %v", active, inactive)
	}
}

func TestPartitionEmptyAndIndependent(t *testing.T) {
	left, right := New([]int{}).Partition(func(int) bool { return true })
	if left == nil || right == nil || len(left) != 0 || len(right) != 0 {
		t.Fatalf("Partition(empty) = %v, %v, want non-nil empties", left, right)
	}
	items := []int{1, 2}
	left, right = New(items).Partition(func(value int) bool { return value%2 == 0 })
	left[0], right[0] = 20, 10
	if items[0] != 1 || items[1] != 2 {
		t.Fatalf("Partition() shared backing storage with source: %v", items)
	}
}
