package collection

// Merge merges the given data into the current collection.
// @group Transformation
// @behavior mutable
// @chainable true
//
// Example: integers - merging slices
//
//	ints := collection.New([]int{1, 2})
//	extra := []int{3, 4}
//	// Merge the extra slice into the ints collection
//	merged1 := ints.Merge(extra)
//	collection.Dump(merged1.Items())
//	// #[]int [
//	//   0 => 1 #int
//	//   1 => 2 #int
//	//   2 => 3 #int
//	//   3 => 4 #int
//	// ]
//
// Example: strings - merging another collection
//
//	strs := collection.New([]string{"a", "b"})
//	more := collection.New([]string{"c", "d"})
//
//	merged2 := strs.Merge(more)
//	collection.Dump(merged2.Items())
//	// #[]string [
//	//   0 => "a" #string
//	//   1 => "b" #string
//	//   2 => "c" #string
//	//   3 => "d" #string
//	// ]
//
// Example: structs - merging struct slices
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
//	moreUsers := []User{
//		{ID: 3, Name: "Carol"},
//		{ID: 4, Name: "Dave"},
//	}
//
//	merged3 := users.Merge(moreUsers)
//	collection.Dump(merged3.Items())
//	// #[]main.User [
//	//   0 => #main.User {
//	//     +ID   => 1 #int
//	//     +Name => "Alice" #string
//	//   }
//	//   1 => #main.User {
//	//     +ID   => 2 #int
//	//     +Name => "Bob" #string
//	//   }
//	//   2 => #main.User {
//	//     +ID   => 3 #int
//	//     +Name => "Carol" #string
//	//   }
//	//   3 => #main.User {
//	//     +ID   => 4 #int
//	//     +Name => "Dave" #string
//	//   }
//	// ]
func (c *Collection[T]) Merge(other any) *Collection[T] {
	switch v := other.(type) {
	case []T:
		return c.mergeSlice(v)

	case *Collection[T]:
		return c.mergeSlice(v.items)

	case map[string]T:
		return c.mergeMap(v)

	default:
		return c
	}
}

/*
mergeSlice handles Laravel-style numeric merges.

Given a slice ([]T), values are appended to the end of the current items.

This function is immutable and returns a new collection.
*/
func (c *Collection[T]) mergeSlice(values []T) *Collection[T] {
	out := make([]T, len(c.items)+len(values))
	copy(out, c.items)
	copy(out[len(c.items):], values)
	return &Collection[T]{items: out}
}

/*
mergeMap handles Laravel-style associative merges.

Steps:

 1. Convert the current slice into a map[string]T with numeric keys.
 2. Apply associative merge rules (overwrite or add).
 3. Convert the map back into a slice.

Map iteration order is not guaranteed — this mirrors Laravel's
behavior when working with associative arrays.

This function is immutable.
*/
func (c *Collection[T]) mergeMap(values map[string]T) *Collection[T] {
	// Precalculate how many values will be appended.
	// Numeric keys <= len(out) overwrite, others append.
	appendCount := 0
	for k := range values {
		if idx, ok := fastParseInt(k); ok {
			if idx < 0 || idx >= len(c.items) {
				appendCount++
			}
		} else {
			appendCount++
		}
	}

	// Pre-size output slice for NO reallocs.
	outLen := len(c.items)
	out := make([]T, outLen, outLen+appendCount)
	copy(out, c.items)

	// Apply merge semantics.
	for k, v := range values {
		if idx, ok := fastParseInt(k); ok {
			if idx >= 0 && idx < len(out) {
				out[idx] = v
				continue
			}
			// numeric but out of range → append
			out = append(out, v)
			continue
		}

		// string key → append value (Laravel)
		out = append(out, v)
	}

	// IMPORTANT: return without copying out again
	return &Collection[T]{items: out}
}

// fastParseInt is much lighter than strconv.Atoi
// Returns (value, ok)
func fastParseInt(s string) (int, bool) {
	if len(s) == 0 {
		return 0, false
	}
	n := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c < '0' || c > '9' {
			return 0, false
		}
		n = n*10 + int(c-'0')
	}
	return n, true
}
