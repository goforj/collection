package collection

// Slice is a named slice with fluent collection operations.
//
// Because Slice is slice-backed, Go's built-in len, index, and range operations
// work directly on it. New borrows the supplied slice; use Clone or ItemsCopy
// when subsequent mutations must not share its backing array.
type Slice[T any] []T

// Number is a constraint that permits any numeric type.
type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

// Pair represents a key/value pair, typically originating from a map.
type Pair[K comparable, V any] struct {
	Key   K
	Value V
}

// New creates a Slice from items and borrows their backing array.
// @group Construction
// @behavior immutable
// @chainable true
// @terminal false
//
// Example: native slice operations
//
//	values := collection.New([]int{10, 20, 30})
//	fmt.Println(len(values))
//	// 3
//	fmt.Println(values[1])
//	// 20
//
//	total := 0
//	for _, value := range values {
//		total += value
//	}
//	fmt.Println(total)
//	// 60
func New[T any](items []T) Slice[T] {
	return Slice[T](items)
}

// Items returns the backing slice as []T.
// @group Access
// @behavior readonly
// @chainable false
// @terminal true
//
// Items remains temporarily as a migration aid. New code can pass, index, range
// over, and call len on Slice directly without converting it first.
//
// Example: migration comparison
//
//	values := collection.New([]int{10, 20, 30})
//	fmt.Println([]int(values))
//	// [10 20 30]
//	fmt.Println(values.Items())
//	// [10 20 30]
func (s Slice[T]) Items() []T {
	return []T(s)
}

// ItemsCopy returns a copy of the slice's items.
// @group Access
// @behavior immutable
// @chainable false
// @terminal true
func (s Slice[T]) ItemsCopy() []T {
	out := make([]T, len(s))
	copy(out, s)
	return out
}
