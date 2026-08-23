package collection

import (
	"reflect"
	"strconv"
	"testing"
)

// TestGenericMethodValuesAndExpressions verifies the callable forms supported by Go 1.27.
func TestGenericMethodValuesAndExpressions(t *testing.T) {
	values := New([]int{1, 2, 3})
	methodValue := values.Map[strconv.NumError]
	methodExpression := (Slice[int]).Map[string]

	errors := methodValue(func(value int) strconv.NumError {
		return strconv.NumError{Func: strconv.Itoa(value)}
	})
	if len(errors) != 3 {
		t.Fatalf("method value result length = %d, want 3", len(errors))
	}

	strings := methodExpression(values, strconv.Itoa)
	want := Slice[string]{"1", "2", "3"}
	if !reflect.DeepEqual(strings, want) {
		t.Fatalf("method expression result = %v, want %v", strings, want)
	}
}

func TestGroupByReturnsIndependentBuiltInSlices(t *testing.T) {
	values := New([]int{1, 2, 3, 4})
	groups := values.GroupBy(func(value int) int { return value % 2 })
	groups[0][0] = 20

	want := Slice[int]{1, 2, 3, 4}
	if !reflect.DeepEqual(values, want) {
		t.Fatalf("GroupBy mutated source = %v, want %v", values, want)
	}
	wantEven := []int{20, 4}
	if !reflect.DeepEqual(groups[0], wantEven) {
		t.Fatalf("even group = %v, want %v", groups[0], wantEven)
	}
}

func TestZipWithAcceptsBuiltInSlice(t *testing.T) {
	values := New([]int{1, 2, 3})
	result := values.ZipWith([]string{"a", "b"}, func(value int, label string) string {
		return strconv.Itoa(value) + label
	})
	want := Slice[string]{"1a", "2b"}
	if !reflect.DeepEqual(result, want) {
		t.Fatalf("ZipWith result = %v, want %v", result, want)
	}
}
