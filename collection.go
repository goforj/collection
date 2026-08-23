package collection

// Slice is a named slice with fluent collection operations.
//
// Because Slice is slice-backed, Go's built-in len, index, and range operations
// work directly on it. New borrows the supplied slice; use Clone when
// subsequent mutations must not share its backing array.
type Slice[T any] []T

// Number is a constraint that permits any numeric type.
type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

// Pair is an ordered pair of values, used by FromMap and Zip.
type Pair[A any, B any] struct {
	First  A
	Second B
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
