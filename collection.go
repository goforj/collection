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

// Pair represents a key/value pair, typically originating from a map.
//
// Pair is used to explicitly materialize unordered map data into an
// ordered collection workflow.
type Pair[K comparable, V any] struct {
	Key   K
	Value V
}

// New creates a new Collection from the provided slice.
// @group Construction
// @behavior immutable
//
// The returned Collection is a lightweight, strongly-typed wrapper
// around the slice, enabling fluent, chainable operations such as
// filtering, mapping, reducing, sorting, and more.
//
// The underlying slice is stored as-is (no copy is made), allowing
// New to be both fast and allocation-friendly. Callers should clone
// the input beforehand if they need to prevent shared mutation.
func New[T any](items []T) *Collection[T] {
	return &Collection[T]{items: items}
}

// NumericCollection is a Collection specialized for numeric types.
type NumericCollection[T Number] struct {
	*Collection[T]
}

// NewNumeric wraps a slice of numeric types in a NumericCollection.
// A shallow copy is made so that further operations don't mutate the original slice.
// @group Construction
// @behavior immutable
func NewNumeric[T Number](items []T) *NumericCollection[T] {
	return &NumericCollection[T]{
		Collection: &Collection[T]{items: items},
	}
}

// Items returns the underlying slice of items.
// @group Access
// @behavior readonly
//
// Example: integers
//
//	c := collection.New([]int{1, 2, 3})
//	items := c.Items()
//	collection.Dump(items)
//	// #[]int [
//	//   0 => 1 #int
//	//   1 => 2 #int
//	//   2 => 3 #int
//	// ]
//
// Example: strings
//
//	c2 := collection.New([]string{"apple", "banana"})
//	items2 := c2.Items()
//	collection.Dump(items2)
//	// #[]string [
//	//   0 => "apple" #string
//	//   1 => "banana" #string
//	// ]
//
// Example: structs
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
//	out := users.Items()
//	collection.Dump(out)
//	// #[]main.User [
//	//   0 => #main.User {
//	//     +ID   => 1 #int
//	//     +Name => "Alice" #string
//	//   }
//	//   1 => #main.User {
//	//     +ID   => 2 #int
//	//     +Name => "Bob" #string
//	//   }
//	// ]
func (c *Collection[T]) Items() []T {
	return c.items
}
