package collection

import "sort"

// Sort returns a new collection sorted using the given comparison function.
//
// less should return true if a should come before b.
//
// Example:
//   sorted := users.Sort(func(a, b User) bool { return a.Age < b.Age })
//  // sorted by Age ascending
func (c Collection[T]) Sort(less func(a, b T) bool) Collection[T] {
	out := c.Items()
	sort.Slice(out, func(i, j int) bool {
		return less(out[i], out[j])
	})
	return Collection[T]{items: out}
}
