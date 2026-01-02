package collection

// Zip combines two collections element-wise into a collection of tuples.
// The resulting length is the smaller of the two inputs.
// @group Transformation
// @behavior immutable
// @fluent true
//
// Example: integers and strings
//
//	nums := collection.New([]int{1, 2, 3})
//	words := collection.New([]string{"one", "two"})
//
//	out := collection.Zip(nums, words)
//	collection.Dump(out.Items())
//	// #[]collection.Tuple[int,string] [
//	//   0 => #collection.Tuple[int,string] {
//	//     +First  => 1 #int
//	//     +Second => "one" #string
//	//   }
//	//   1 => #collection.Tuple[int,string] {
//	//     +First  => 2 #int
//	//     +Second => "two" #string
//	//   }
//	// ]
//
// Example: structs
//
//	type User struct {
//		ID   int
//		Name string
//	}
//
//	users := collection.New([]User{
//		{ID: 1, Name: "Alice"},
//		{ID: 2, Name: "Bob"},
//	})
//
//	roles := collection.New([]string{"admin", "user", "extra"})
//
//	out2 := collection.Zip(users, roles)
//	collection.Dump(out2.Items())
//	// #[]collection.Tuple[main.User,string] [
//	//   0 => #collection.Tuple[main.User,string] {
//	//     +First  => #main.User {
//	//       +ID   => 1 #int
//	//       +Name => "Alice" #string
//	//     }
//	//     +Second => "admin" #string
//	//   }
//	//   1 => #collection.Tuple[main.User,string] {
//	//     +First  => #main.User {
//	//       +ID   => 2 #int
//	//       +Name => "Bob" #string
//	//     }
//	//     +Second => "user" #string
//	//   }
//	// ]
func Zip[A any, B any](a *Collection[A], b *Collection[B]) *Collection[Tuple[A, B]] {
	n := len(a.items)
	if len(b.items) < n {
		n = len(b.items)
	}

	out := make([]Tuple[A, B], n)
	for i := 0; i < n; i++ {
		out[i] = Tuple[A, B]{First: a.items[i], Second: b.items[i]}
	}

	return New(out)
}

// ZipWith combines two collections element-wise using combiner fn.
// The resulting length is the smaller of the two inputs.
// @group Transformation
// @behavior immutable
// @fluent true
//
// Example: sum ints
//
//	a := collection.New([]int{1, 2, 3})
//	b := collection.New([]int{10, 20})
//
//	out := collection.ZipWith(a, b, func(x, y int) int {
//		return x + y
//	})
//
//	collection.Dump(out.Items())
//	// #[]int [
//	//   0 => 11 #int
//	//   1 => 22 #int
//	// ]
//
// Example: format strings
//
//	names := collection.New([]string{"alice", "bob"})
//	roles := collection.New([]string{"admin", "user", "extra"})
//
//	out2 := collection.ZipWith(names, roles, func(name, role string) string {
//		return name + ":" + role
//	})
//
//	collection.Dump(out2.Items())
//	// #[]string [
//	//   0 => "alice:admin" #string
//	//   1 => "bob:user" #string
//	// ]
//
// Example: structs
//
//	type User struct {
//		Name string
//	}
//
//	type Role struct {
//		Title string
//	}
//
//	users := collection.New([]User{{Name: "Alice"}, {Name: "Bob"}})
//	roles2 := collection.New([]Role{{Title: "admin"}})
//
//	out3 := collection.ZipWith(users, roles2, func(u User, r Role) string {
//		return u.Name + " -> " + r.Title
//	})
//
//	collection.Dump(out3.Items())
//	// #[]string [
//	//   0 => "Alice -> admin" #string
//	// ]
func ZipWith[A any, B any, R any](a *Collection[A], b *Collection[B], fn func(A, B) R) *Collection[R] {
	n := len(a.items)
	if len(b.items) < n {
		n = len(b.items)
	}

	out := make([]R, n)
	for i := 0; i < n; i++ {
		out[i] = fn(a.items[i], b.items[i])
	}

	return New(out)
}
