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
// @fluent true
//
// The returned Collection is a lightweight, strongly-typed wrapper
// around the slice, enabling fluent, chainable operations such as
// filtering, mapping, reducing, sorting, and more.
//
// New copies the input slice to avoid shared backing.
func New[T any](items []T) *Collection[T] {
	var out []T
	if items == nil {
		out = nil
	} else {
		out = make([]T, len(items))
		copy(out, items)
	}
	return &Collection[T]{items: out}
}

// NumericCollection is a Collection specialized for numeric types.
type NumericCollection[T Number] struct {
	*Collection[T]
}

// NewNumeric wraps a slice of numeric types in a NumericCollection.
// @group Construction
// @behavior immutable
// @fluent true
//
// NewNumeric copies the input slice to avoid shared backing.
func NewNumeric[T Number](items []T) *NumericCollection[T] {
	var out []T
	if items == nil {
		out = nil
	} else {
		out = make([]T, len(items))
		copy(out, items)
	}
	return &NumericCollection[T]{
		Collection: &Collection[T]{items: out},
	}
}

// Attach wraps a slice without copying.
// @group Construction
// @behavior immutable
// @fluent true
//
// Attach shares the backing array with the caller. Mutating either side
// will affect the other. Use New to copy the input.
//
// Example: sharing backing slice
//
//	items := []int{1, 2, 3}
//	c := collection.Attach(items)
//
//	items[0] = 9
//	collection.Dump(c.Items())
//	// #[]int [
//	//   0 => 9 #int
//	//   1 => 2 #int
//	//   2 => 3 #int
//	// ]
func Attach[T any](items []T) *Collection[T] {
	return &Collection[T]{items: items}
}

// AttachNumeric wraps a slice of numeric types without copying.
// @group Construction
// @behavior immutable
// @fluent true
//
// AttachNumeric shares the backing array with the caller. Mutating either side
// will affect the other. Use NewNumeric to copy the input.
//
// Example: sharing backing slice
//
//	items := []int{1, 2, 3}
//	c := collection.AttachNumeric(items)
//
//	items[0] = 9
//	collection.Dump(c.Items())
//	// #[]int [
//	//   0 => 9 #int
//	//   1 => 2 #int
//	//   2 => 3 #int
//	// ]
func AttachNumeric[T Number](items []T) *NumericCollection[T] {
	return &NumericCollection[T]{Collection: &Collection[T]{items: items}}
}

// Items returns the backing slice of items.
// @group Access
// @behavior readonly
// @fluent false
// @terminal true
//
// Items shares the backing array with the collection. Mutating the returned
// slice will mutate the collection.
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

// ItemsCopy returns a copy of the collection's items.
// @group Access
// @behavior readonly
// @fluent false
// @terminal true
//
// ItemsCopy allocates a new slice.
//
// Example: integers
//
//	c := collection.New([]int{1, 2, 3})
//	items := c.ItemsCopy()
//	collection.Dump(items)
//	// #[]int [
//	//   0 => 1 #int
//	//   1 => 2 #int
//	//   2 => 3 #int
//	// ]
func (c *Collection[T]) ItemsCopy() []T {
	out := make([]T, len(c.items))
	copy(out, c.items)
	return out
}
