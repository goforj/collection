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
func New[T any](items []T) *Collection[T] {
	return &Collection[T]{items: items}
}

// NumericCollection is a Collection specialized for numeric types.
type NumericCollection[T Number] struct {
	*Collection[T]
}

// NewNumeric wraps a slice of numeric types in a NumericCollection.
// A shallow copy is made so that further operations don't mutate the original slice.
func NewNumeric[T Number](items []T) *NumericCollection[T] {
	return &NumericCollection[T]{
		Collection: &Collection[T]{items: items},
	}
}

// Items returns the underlying slice of items.
func (c *Collection[T]) Items() []T {
	return c.items
}
