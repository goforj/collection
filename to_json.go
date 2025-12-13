package collection

import (
	"encoding/json"
	"errors"
)

// ToJSON converts the collection's items into a compact JSON string.
//
// If marshalling succeeds, a JSON-encoded string and a nil error are returned.
// If marshalling fails, the method unwraps any json.Marshal wrapping so that
// user-defined MarshalJSON errors surface directly.
//
// Returns:
//   - string: JSON-encoded representation of the collection
//   - error : nil on success, or the unwrapped marshalling error
//
// Example: strings - pretty JSON
//
//	pj1 := collection.New([]string{"a", "b"})
//	out1, _ := pj1.ToJSON()
//	fmt.Println(out1)
//	// ["a","b"]
func (c *Collection[T]) ToJSON() (string, error) {
	b, err := json.Marshal(c.items)
	if err != nil {
		return "", errors.Unwrap(err)
	}
	return string(b), nil
}

// ToPrettyJSON converts the collection's items into a human-readable,
// indented JSON string.
//
// If marshalling succeeds, a formatted JSON string and nil error are returned.
// If marshalling fails, the underlying error is unwrapped so user-defined
// MarshalJSON failures surface directly.
//
// Returns:
//   - string: the pretty-printed JSON representation
//   - error : nil on success, or the unwrapped marshalling error
//
// Example: strings - pretty JSON
//
//	pj1 := collection.New([]string{"a", "b"})
//	out1, _ := pj1.ToPrettyJSON()
//	fmt.Println(out1)
//	// [
//	//  "a",
//	//  "b"
//	// ]
func (c *Collection[T]) ToPrettyJSON() (string, error) {
	b, err := json.MarshalIndent(c.items, "", "  ")
	if err != nil {
		return "", errors.Unwrap(err)
	}
	return string(b), nil
}
