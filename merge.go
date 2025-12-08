package collection

import (
	"strconv"
)

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
func (c Collection[T]) Merge(other any) Collection[T] {
	switch v := other.(type) {

	case []T:
		return c.mergeSlice(v)

	case Collection[T]:
		return c.mergeSlice(v.items)

	case map[string]T:
		return c.mergeMap(v)

	default:
		// Unsupported type — return original collection unchanged.
		// This matches Laravel's fail-soft behavior.
		return c
	}
}

/*
mergeSlice handles Laravel-style numeric merges.

Given a slice ([]T), values are appended to the end of the current items.

This function is immutable and returns a new collection.
*/
func (c Collection[T]) mergeSlice(values []T) Collection[T] {
	out := make([]T, len(c.items))
	copy(out, c.items)
	out = append(out, values...)
	return New(out)
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
func (c Collection[T]) mergeMap(values map[string]T) Collection[T] {
	tmp := make(map[string]T)

	for i, v := range c.items {
		tmp[strconv.Itoa(i)] = v
	}

	for k, v := range values {
		tmp[k] = v
	}

	out := make([]T, 0, len(tmp))
	for _, v := range tmp {
		out = append(out, v)
	}

	return New(out)
}
