package collection

import "github.com/goforj/godump"

// exitFunc allows tests to override the exit behavior.
var exitFunc = func(v interface{}) { godump.Dd(v) }

// Dump prints items with godump and returns the same collection.
//
// Example:
//
//   c := collection.New([]int{1, 2, 3})
//   out := c.Dump()
//   // Prints a pretty debug dump of [1, 2, 3]
//   // out == c
//
// Dump is typically used while chaining:
//
//   collection.New([]int{1, 2, 3}).
//       Filter(func(v int) bool { return v > 1 }).
//       Dump()
//
// This is a no-op on the collection itself and never panics.
func (c *Collection[T]) Dump() *Collection[T] {
	godump.Dump(c.Items())
	return c
}

// Dd prints items then terminates execution.
//
// Example:
//
//   c := collection.New([]string{"a", "b"})
//   c.Dd()    // Prints the dump and exits the program
//
// Like Laravel's dd(), this is intended for debugging and
// should not be used in production control flow.
//
// This method never returns.
func (c *Collection[T]) Dd() {
	exitFunc(c.Items())
}

// DumpStr returns the pretty-printed dump of the items as a string,
// without printing or exiting.
//
// Example:
//
//   c := collection.New([]int{10, 20})
//   s := c.DumpStr()
//   fmt.Println(s)
//   // Produces a multi-line formatted representation of [10, 20]
//
// Useful for logging, snapshot testing, and non-interactive debugging.
func (c *Collection[T]) DumpStr() string {
	return godump.DumpStr(c.Items())
}

// DdStr behaves like Dd() but also returns the formatted dump string.
//
// Because Dd() exits immediately, DdStr is helpful primarily in tests
// where exit behavior has been overridden.
//
// Example (non-fatal debug):
//
//   c := collection.New([]int{1})
//   s := c.DdStr()
//   // Prints the formatted dump, triggers exitFunc, and returns the output.
//
// The return value is mostly useful in testing environments where exitFunc
// has been replaced with a non-terminating stub.
func (c *Collection[T]) DdStr() string {
	out := godump.DumpStr(c.Items())
	exitFunc(out)
	return out
}

// Dump is a convenience function that calls the Dump method on the collection.
//
// Example:
//   c := collection.New([]int{1, 2, 3})
//   c.Dump() // Pretty-prints [1, 2, 3]
//
// This function is provided for symmetry with godump.Dump.
func Dump(vs ...any) {
	godump.Dump(vs...)
}

// Dd is a convenience function that calls the Dd method on the collection.
//
// Example:
//   c := collection.New([]string{"x", "y"})
//   c.Dd() // Pretty-prints ["x", "y"] and exits
//
// This function is provided for symmetry with godump.Dd.
func Dd(vs ...any) {
	exitFunc(vs)
}
