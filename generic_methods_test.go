package collection

import (
	"reflect"
	"strconv"
	"testing"
)

// TestGenericMethodsMatchCompatibilityFunctions verifies the new fluent API preserves every legacy result.
func TestGenericMethodsMatchCompatibilityFunctions(t *testing.T) {
	t.Parallel()

	values := New([]int{3, 1, 2, 1})
	other := New([]string{"a", "b", "c"})
	keyFn := func(value int) int { return value % 2 }

	assertEqual(t, "MapTo", values.MapTo(strconv.Itoa).Items(), MapTo(values, strconv.Itoa).Items())
	methodMin, methodMinOK := values.MinBy(keyFn)
	functionMin, functionMinOK := MinBy(values, keyFn)
	assertPairEqual(t, "MinBy", methodMin, methodMinOK, functionMin, functionMinOK)
	methodMax, methodMaxOK := values.MaxBy(keyFn)
	functionMax, functionMaxOK := MaxBy(values, keyFn)
	assertPairEqual(t, "MaxBy", methodMax, methodMaxOK, functionMax, functionMaxOK)
	assertEqual(t, "ToMap", values.ToMap(strconv.Itoa, func(value int) int { return value }), ToMap(values, strconv.Itoa, func(value int) int { return value }))
	assertEqual(t, "GroupBy", collectionItems(values.GroupBy(keyFn)), collectionItems(GroupBy(values, keyFn)))
	assertEqual(t, "GroupBySlice", values.GroupBySlice(keyFn), GroupBySlice(values, keyFn))
	assertEqual(t, "CountBy", values.CountBy(keyFn), CountBy(values, keyFn))
	assertEqual(t, "UniqueBy", values.UniqueBy(keyFn).Items(), UniqueBy(values, keyFn).Items())
	assertEqual(t, "Pipe", values.Pipe(func(c *Collection[int]) int { return len(c.Items()) }), Pipe(values, func(c *Collection[int]) int { return len(c.Items()) }))
	assertEqual(t, "ZipWith", values.ZipWith(other, func(value int, label string) string { return strconv.Itoa(value) + label }).Items(), ZipWith(values, other, func(value int, label string) string { return strconv.Itoa(value) + label }).Items())
}

// TestGenericMethodValuesAndExpressions verifies both callable forms supported by Go 1.27.
func TestGenericMethodValuesAndExpressions(t *testing.T) {
	t.Parallel()

	values := New([]int{1, 2, 3})
	methodValue := values.MapTo[strconv.NumError]
	methodExpression := (*Collection[int]).MapTo[string]

	errors := methodValue(func(value int) strconv.NumError {
		return strconv.NumError{Func: strconv.Itoa(value)}
	})
	if got := len(errors.Items()); got != 3 {
		t.Fatalf("method value result length = %d, want 3", got)
	}

	strings := methodExpression(values, strconv.Itoa)
	assertEqual(t, "method expression", strings.Items(), []string{"1", "2", "3"})
}

// TestGenericMethodsPreserveNilReceiverBehavior verifies compatibility at the edge of the existing pointer-based API.
func TestGenericMethodsPreserveNilReceiverBehavior(t *testing.T) {
	t.Parallel()

	var values *Collection[int]
	assertPanics(t, "method", func() { values.MapTo(strconv.Itoa) })
	assertPanics(t, "function", func() { MapTo(values, strconv.Itoa) })

	methodSawNil := values.Pipe(func(c *Collection[int]) bool { return c == nil })
	functionSawNil := Pipe(values, func(c *Collection[int]) bool { return c == nil })
	if !methodSawNil || !functionSawNil {
		t.Fatalf("Pipe nil receiver parity = (%v, %v), want both true", methodSawNil, functionSawNil)
	}
}

// assertEqual compares values that may contain maps or slices.
func assertEqual(t *testing.T, name string, got, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("%s result = %#v, want %#v", name, got, want)
	}
}

// assertPairEqual compares two value-and-presence results.
func assertPairEqual[T comparable](t *testing.T, name string, got T, gotOK bool, want T, wantOK bool) {
	t.Helper()
	if got != want || gotOK != wantOK {
		t.Fatalf("%s result = (%v, %v), want (%v, %v)", name, got, gotOK, want, wantOK)
	}
}

// collectionItems converts grouped collections into slices for stable comparison.
func collectionItems[K comparable, V any](groups map[K]*Collection[V]) map[K][]V {
	items := make(map[K][]V, len(groups))
	for key, group := range groups {
		items[key] = group.Items()
	}
	return items
}

// assertPanics verifies an operation retains its established panic behavior.
func assertPanics(t *testing.T, name string, fn func()) {
	t.Helper()
	defer func() {
		if recover() == nil {
			t.Fatalf("%s call did not panic", name)
		}
	}()
	fn()
}
