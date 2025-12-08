package collection

/*
Merge merges the given data into the current collection using
Laravel-style semantics.

Behavior depends on the type of `other`:

  • []T (numeric merges)
      Values are appended to the end of the collection.
  • Collection[T]
      Values are appended, same as merging a slice.
  • map[string]T (associative merges)
      Keys that already exist overwrite the original values;
      new keys are added.

Unsupported merge types are ignored. This method
never panics and always returns a new Collection.
*/
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
