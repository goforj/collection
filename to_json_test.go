package collection

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func TestToJSON_Ints(t *testing.T) {
	c := New([]int{1, 2, 3})

	jsonStr, err := c.ToJSON()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := `[1,2,3]`
	if jsonStr != expected {
		t.Fatalf("expected %s, got %s", expected, jsonStr)
	}

	// Validate it is valid JSON
	var out []int
	if err := json.Unmarshal([]byte(jsonStr), &out); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}
	if !reflect.DeepEqual(out, []int{1, 2, 3}) {
		t.Fatalf("unmarshaled JSON mismatch: %v", out)
	}
}

func TestToPrettyJSON_Ints(t *testing.T) {
	c := New([]int{1, 2, 3})

	jsonStr, err := c.ToPrettyJSON()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := `[
  1,
  2,
  3
]`

	if jsonStr != expected {
		t.Fatalf("pretty JSON mismatch.\nExpected:\n%s\nGot:\n%s", expected, jsonStr)
	}

	// Validate JSON structure
	var out []int
	if err := json.Unmarshal([]byte(jsonStr), &out); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}
}

func TestToJSON_Structs(t *testing.T) {
	type Product struct {
		Name  string `json:"name"`
		Price int    `json:"price"`
	}

	c := New([]Product{
		{"Desk", 200},
		{"Chair", 100},
	})

	jsonStr, err := c.ToJSON()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := `[{"name":"Desk","price":200},{"name":"Chair","price":100}]`

	if jsonStr != expected {
		t.Fatalf("expected %s, got %s", expected, jsonStr)
	}

	// Ensure JSON decodes properly
	var out []Product
	if err := json.Unmarshal([]byte(jsonStr), &out); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}
	if !reflect.DeepEqual(out, c.items) {
		t.Fatalf("decoded JSON mismatch: %v", out)
	}
}

func TestToPrettyJSON_Structs(t *testing.T) {
	type Product struct {
		Name  string `json:"name"`
		Price int    `json:"price"`
	}

	c := New([]Product{
		{"Desk", 200},
		{"Chair", 100},
	})

	jsonStr, err := c.ToPrettyJSON()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := `[
  {
    "name": "Desk",
    "price": 200
  },
  {
    "name": "Chair",
    "price": 100
  }
]`

	if jsonStr != expected {
		t.Fatalf("pretty JSON mismatch.\nExpected:\n%s\nGot:\n%s", expected, jsonStr)
	}

	// Validate JSON structure
	var out []Product
	if err := json.Unmarshal([]byte(jsonStr), &out); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}
}

func TestToJSON_Empty(t *testing.T) {
	c := New([]int{})

	jsonStr, err := c.ToJSON()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := `[]`
	if jsonStr != expected {
		t.Fatalf("expected [], got %s", jsonStr)
	}

	var out []int
	if err := json.Unmarshal([]byte(jsonStr), &out); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}
	if len(out) != 0 {
		t.Fatalf("expected empty list, got %v", out)
	}
}

func TestToPrettyJSON_Empty(t *testing.T) {
	c := New([]int{})

	jsonStr, err := c.ToPrettyJSON()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := `[]`
	if jsonStr != expected {
		t.Fatalf("expected %s, got %s", expected, jsonStr)
	}

	// ensure valid JSON
	var out []int
	if err := json.Unmarshal([]byte(jsonStr), &out); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}
}

func TestToJSON_NoMutation(t *testing.T) {
	c := New([]int{1, 2, 3})
	orig := append([]int{}, c.items...) // make copy

	_, _ = c.ToJSON()

	if !reflect.DeepEqual(c.items, orig) {
		t.Fatalf("collection mutated by ToJSON")
	}
}

func TestToPrettyJSON_NoMutation(t *testing.T) {
	c := New([]int{10, 20})
	orig := append([]int{}, c.items...)

	_, _ = c.ToPrettyJSON()

	if !reflect.DeepEqual(c.items, orig) {
		t.Fatalf("collection mutated by ToPrettyJSON")
	}
}

type BadJSON struct{}

func (BadJSON) MarshalJSON() ([]byte, error) {
	return nil, fmt.Errorf("marshal failure")
}

func TestToJSON_Error(t *testing.T) {
	c := New([]BadJSON{{}, {}})

	_, err := c.ToJSON()
	if err == nil {
		t.Fatalf("expected error from JSON marshal, got nil")
	}

	if err.Error() != "marshal failure" {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestToPrettyJSON_Error(t *testing.T) {
	c := New([]BadJSON{{}})

	_, err := c.ToPrettyJSON()
	if err == nil {
		t.Fatalf("expected error from JSON marshal indent, got nil")
	}

	if err.Error() != "marshal failure" {
		t.Fatalf("unexpected error: %v", err)
	}
}
