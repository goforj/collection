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
// This method never panics.
//
// Example:
//
//   c := collection.New([]int{1, 2, 3})
//   out, err := c.ToJSON()
//   // out: "[1,2,3]"
//   // err: nil
//
// Example (error):
//
//   type Bad struct{}
//   func (Bad) MarshalJSON() ([]byte, error) {
//       return nil, fmt.Errorf("marshal failure")
//   }
//
//   c := collection.New([]Bad{{}})
//   out, err := c.ToJSON()
//   // out: ""
//   // err.Error(): "marshal failure"
//
// Returns:
//   - string: the JSON-encoded representation of the collection
//   - error : nil on success, or the unwrapped marshalling error
func (c Collection[T]) ToJSON() (string, error) {
	b, err := json.Marshal(c.items)
	if err != nil {
		return "", errors.Unwrap(err)
	}
	return string(b), nil
}

// ToPrettyJSON converts the collection's items into an indented,
// human-readable JSON string.
//
// If marshalling succeeds, a formatted JSON string and nil error are returned.
// If marshalling fails, the underlying error is unwrapped so that user-defined
// MarshalJSON failures surface directly (e.g., "marshal failure") rather than
// the json.MarshalIndent wrapper.
//
// This method never panics.
//
// Example:
//
//   c := collection.New([]string{"a", "b"})
//   out, err := c.ToPrettyJSON()
//   // out:
//   // [
//   //   "a",
//   //   "b"
//   // ]
//   // err: nil
//
// Example (error):
//
//   type Bad struct{}
//   func (Bad) MarshalJSON() ([]byte, error) {
//       return nil, fmt.Errorf("marshal failure")
//   }
//
//   c := collection.New([]Bad{{}})
//   out, err := c.ToPrettyJSON()
//   // out: ""
//   // err.Error(): "marshal failure"
//
// Returns:
//   - string: the pretty-printed JSON representation
//   - error : nil on success, or the unwrapped marshalling error
func (c Collection[T]) ToPrettyJSON() (string, error) {
	b, err := json.MarshalIndent(c.items, "", "  ")
	if err != nil {
		return "", errors.Unwrap(err)
	}
	return string(b), nil
}
