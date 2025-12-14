package collection

import (
	"reflect"
	"testing"
)

func TestToMapKV_Basic(t *testing.T) {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	c := FromMap(m)
	out := ToMapKV(c)

	if !reflect.DeepEqual(out, m) {
		t.Fatalf("expected %v, got %v", m, out)
	}
}

func TestToMapKV_AfterFilter(t *testing.T) {
	type Config struct {
		Enabled bool
		Timeout int
	}

	configs := map[string]Config{
		"a": {Enabled: true, Timeout: 10},
		"b": {Enabled: false, Timeout: 20},
		"c": {Enabled: true, Timeout: 30},
	}

	c := FromMap(configs).
		Filter(func(p Pair[string, Config]) bool {
			return p.Value.Enabled
		})

	out := ToMapKV(c)

	expect := map[string]Config{
		"a": {Enabled: true, Timeout: 10},
		"c": {Enabled: true, Timeout: 30},
	}

	if !reflect.DeepEqual(out, expect) {
		t.Fatalf("expected %v, got %v", expect, out)
	}
}

func TestToMapKV_KeyCollision_LastWins(t *testing.T) {
	c := New([]Pair[string, int]{
		{Key: "x", Value: 1},
		{Key: "x", Value: 2},
		{Key: "x", Value: 3},
	})

	out := ToMapKV(c)

	expect := map[string]int{"x": 3}
	if !reflect.DeepEqual(out, expect) {
		t.Fatalf("expected %v, got %v", expect, out)
	}
}

func TestToMapKV_EmptyCollection(t *testing.T) {
	c := New([]Pair[string, int]{})
	out := ToMapKV(c)

	if len(out) != 0 {
		t.Fatalf("expected empty map, got %v", out)
	}
}

func TestToMapKV_DoesNotMutateCollection(t *testing.T) {
	items := []Pair[string, int]{
		{Key: "a", Value: 1},
		{Key: "b", Value: 2},
	}

	c := New(items)
	_ = ToMapKV(c)

	if !reflect.DeepEqual(c.Items(), items) {
		t.Fatalf("collection was mutated")
	}
}

func TestToMapKV_MapCapacity(t *testing.T) {
	items := []Pair[string, int]{
		{Key: "a", Value: 1},
		{Key: "b", Value: 2},
		{Key: "c", Value: 3},
	}

	c := New(items)
	out := ToMapKV(c)

	if len(out) != len(items) {
		t.Fatalf("expected map length %d, got %d", len(items), len(out))
	}
}
