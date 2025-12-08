package collection

import "github.com/goforj/godump"

// Dump pretty-prints the collection contents using goforj/godump
// and returns the collection so it can be used mid-chain.
//
// Example:
//   users.
//     Filter(func(u User) bool { return u.Age >= 35 }).
//     Dump().
//     Sort(func(a, b User) bool { return a.Age < b.Age })
func (c Collection[T]) Dump() Collection[T] {
	godump.Dump(c.Items()) // or c.items if you don't care about copying
	return c
}

// Dd pretty-prints the collection contents using goforj/godump
// and then exits the program (just like Laravel's dd()).
//
// Example:
//   users.
//     Filter(func(u User) bool { return u.Age >= 35 }).
//     Dd()
func (c Collection[T]) Dd() {
	godump.Dd(c.Items())
}

/*
DumpStr pretty-prints the collection items using godump.DumpStr
and returns the string. Unlike Dump(), this does not print to stdout
and does not interrupt a chain.

Example:
    s := users.Filter(active).DumpStr()
*/
func (c Collection[T]) DumpStr() string {
	return godump.DumpStr(c.Items())
}

/*
DdStr pretty-prints the collection items using godump.DumpStr
and returns the string, AND then exits — just like Laravel's dd(),
except here the exit is performed via godump.exitFunc so tests can override it.

Example:
    output := users.Filter(active).DdStr()
    // program exits after printing
*/
func (c Collection[T]) DdStr() string {
	out := godump.DumpStr(c.Items())
	godump.Dd(out) // this triggers exitFunc(1)
	return out     // unreachable in real usage, but returned for test purposes
}
