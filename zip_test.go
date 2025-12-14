package collection

import "testing"

func TestZip_IntsAndStrings(t *testing.T) {
	a := New([]int{1, 2, 3})
	b := New([]string{"one", "two"})

	out := Zip(a, b)

	exp := []Tuple[int, string]{
		{First: 1, Second: "one"},
		{First: 2, Second: "two"},
	}

	if got := out.Items(); !pairsEqual(got, exp) {
		t.Fatalf("expected %v, got %v", exp, got)
	}
}

func TestZip_Structs(t *testing.T) {
	type user struct {
		id int
	}
	type role struct {
		name string
	}

	users := New([]user{{id: 1}, {id: 2}})
	roles := New([]role{{name: "admin"}, {name: "user"}, {name: "extra"}})

	out := Zip(users, roles)

	exp := []Tuple[user, role]{
		{First: user{id: 1}, Second: role{name: "admin"}},
		{First: user{id: 2}, Second: role{name: "user"}},
	}

	if got := out.Items(); !pairsEqual(got, exp) {
		t.Fatalf("expected %v, got %v", exp, got)
	}
}

func TestZip_Empty(t *testing.T) {
	a := New([]int{})
	b := New([]int{1, 2})

	out := Zip(a, b)

	if len(out.Items()) != 0 {
		t.Fatalf("expected empty zip, got %v", out.Items())
	}
}

func TestZipWith_Sum(t *testing.T) {
	a := New([]int{1, 2, 3})
	b := New([]int{10, 20})

	out := ZipWith(a, b, func(x, y int) int {
		return x + y
	})

	exp := []int{11, 22}
	if got := out.Items(); !slicesEqual(got, exp) {
		t.Fatalf("expected %v, got %v", exp, got)
	}
}

func TestZipWith_Structs(t *testing.T) {
	type user struct {
		name string
	}
	type role struct {
		title string
	}

	users := New([]user{{name: "alice"}, {name: "bob"}})
	roles := New([]role{{title: "admin"}})

	out := ZipWith(users, roles, func(u user, r role) string {
		return u.name + ":" + r.title
	})

	exp := []string{"alice:admin"}
	if got := out.Items(); !slicesEqual(got, exp) {
		t.Fatalf("expected %v, got %v", exp, got)
	}
}

func TestZipWith_Empty(t *testing.T) {
	a := New([]int{})
	b := New([]int{1})

	out := ZipWith(a, b, func(x, y int) int {
		return x + y
	})

	if len(out.Items()) != 0 {
		t.Fatalf("expected empty zipwith, got %v", out.Items())
	}
}

func pairsEqual[A comparable, B comparable](a, b []Tuple[A, B]) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i].First != b[i].First || a[i].Second != b[i].Second {
			return false
		}
	}
	return true
}
