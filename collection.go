package collection

// Collection is a strongly-typed, fluent wrapper around a slice of T.
type Collection[T any] struct {
	items []T
}

// Number is a constraint that permits any numeric type.
type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
	~float32 | ~float64
}

// New wraps a slice in a Collection.
// A shallow copy is made so that further operations don't mutate the original slice.
func New[T any](items []T) Collection[T] {
	out := make([]T, len(items))
	copy(out, items)
	return Collection[T]{items: out}
}

// Items returns a copy of the underlying slice.
// This avoids callers mutating internal state accidentally.
func (c Collection[T]) Items() []T {
	out := make([]T, len(c.items))
	copy(out, c.items)
	return out
}

// IsEmpty returns true if the collection has no items.
func (c Collection[T]) IsEmpty() bool {
	return len(c.items) == 0
}

//
// ─── SAME-TYPE FLUENT OPERATIONS (METHODS) ─────────────────────────────────────
//

// All returns the underlying slice of items.
func (c Collection[T]) All() []T {
	out := make([]T, len(c.items))
	copy(out, c.items)
	return out
}

//
// ─── TYPE-CHANGING OPERATIONS (FREE FUNCTIONS) ─────────────────────────────────
//

// Before returns all items before the first element for which pred returns true.
// If no element matches, the entire collection is returned.
func (c Collection[T]) Before(pred func(T) bool) Collection[T] {
	idx := len(c.items)
	for i, v := range c.items {
		if pred(v) {
			idx = i
			break
		}
	}

	out := make([]T, idx)
	copy(out, c.items[:idx])
	return Collection[T]{items: out}
}

// After returns all items after the first element for which pred returns true.
// If no element matches, an empty collection is returned.
//
// Example:
//   c := collection.New([]int{1,2,3,4,5})
//   c.After(func(v int) bool { return v == 3 }) → [4,5]
func (c Collection[T]) After(pred func(T) bool) Collection[T] {
	idx := -1
	for i, v := range c.items {
		if pred(v) {
			idx = i
			break
		}
	}

	// If no match found → empty collection
	if idx == -1 || idx+1 >= len(c.items) {
		return Collection[T]{items: []T{}}
	}

	out := make([]T, len(c.items)-(idx+1))
	copy(out, c.items[idx+1:])
	return Collection[T]{items: out}
}

// AvgBy calculates the average of values extracted by fn from the collection items.
//
// Example:
//   avgAge := AvgBy(users, func(u User) float64 { return float64(u.Age) })
func AvgBy[T any](c Collection[T], fn func(T) float64) float64 {
	items := c.Items()

	if len(items) == 0 {
		return 0
	}

	var sum float64
	for _, v := range items {
		sum += fn(v)
	}

	return sum / float64(len(items))
}

// SumBy returns the sum of a numeric projection from each item.
//
// Example (structs):
//   type Row struct{ Foo int }
//   rows := New([]Row{{10}, {20}})
//   total := SumBy(rows, func(r Row) int { return r.Foo }) // 30
func SumBy[T any, N Number](c Collection[T], fn func(T) N) N {
	items := c.Items()
	var sum N
	for _, v := range items {
		sum += fn(v)
	}
	return sum
}
