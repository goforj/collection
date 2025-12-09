package collection

import (
	"github.com/goforj/godump"
)

// exitFunc allows tests to override the exit behavior.
var exitFunc = func(v interface{}) { godump.Dd(v) }

// Dump prints items with godump and returns the same collection.
// This is a no-op on the collection itself and never panics.
//
// Example:
//	// integers
//	c := collection.New([]int{1, 2, 3})
//	c.Dump()
//	// #[]int [
//	//   0 => 1 #int
//	//   1 => 2 #int
//	//   2 => 3 #int
//	// ]
//
// Example:
//	// chaining
//	collection.New([]int{1, 2, 3}).
//		Filter(func(v int) bool { return v > 1 }).
//		Dump()
//	// #[]int [
//	//   0 => 2 #int
//	//   1 => 3 #int
//	// ]
func (c *Collection[T]) Dump() *Collection[T] {
	godump.Dump(c.Items())
	return c
}

// Dd prints items then terminates execution.
// Like Laravel's dd(), this is intended for debugging and
// should not be used in production control flow.
//
// This method never returns.
//
// Example:
//	// strings
//	c := collection.New([]string{"a", "b"})
//	c.Dd()
//	// #[]string [
//	//   0 => "a" #string
//	//   1 => "b" #string
//	// ]
//	// Process finished with the exit code 1
func (c *Collection[T]) Dd() {
	exitFunc(c.Items())
}

// DumpStr returns the pretty-printed dump of the items as a string,
// without printing or exiting.
// Useful for logging, snapshot testing, and non-interactive debugging.
//
// Example:
//	// integers
//	c := collection.New([]int{10, 20})
//	s := c.DumpStr()
//	fmt.Println(s)
//	// #[]int [
//	//   0 => 10 #int
//	//   1 => 20 #int
//	// ]
func (c *Collection[T]) DumpStr() string {
	return godump.DumpStr(c.Items())
}

// Dump is a convenience function that calls godump.Dump.
//
// Example:
//
//	// integers
//	c2 := collection.New([]int{1, 2, 3})
//	collection.Dump(c2.Items())
//	// #[]int [
//	//   0 => 1 #int
//	//   1 => 2 #int
//	//   2 => 3 #int
//	// ]
func Dump(vs ...any) {
	godump.Dump(vs...)
}
