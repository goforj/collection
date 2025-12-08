package collection

import "testing"

func TestAvgBy_StructField(t *testing.T) {
	type Row struct {
		Foo int
	}

	rows := New([]Row{
		{Foo: 10},
		{Foo: 10},
		{Foo: 20},
		{Foo: 40},
	})

	avg := AvgBy(rows, func(r Row) float64 {
		return float64(r.Foo)
	})

	if avg != 20 {
		t.Fatalf("expected 20, got %v", avg)
	}
}

func TestAvgBy_Empty(t *testing.T) {
	type Row struct{ Foo int }

	rows := New([]Row{})

	avg := AvgBy(rows, func(r Row) float64 {
		return float64(r.Foo)
	})

	if avg != 0 {
		t.Fatalf("expected 0 for empty slice, got %v", avg)
	}
}

func TestAvgBy_Single(t *testing.T) {
	type Row struct{ Foo float64 }

	rows := New([]Row{{Foo: 99.9}})

	avg := AvgBy(rows, func(r Row) float64 {
		return r.Foo
	})

	if avg != 99.9 {
		t.Fatalf("expected 99.9, got %v", avg)
	}
}

func TestAvgBy_Negatives(t *testing.T) {
	type Row struct{ Foo float64 }

	rows := New([]Row{
		{Foo: -10},
		{Foo: 0},
		{Foo: 10},
	})

	avg := AvgBy(rows, func(r Row) float64 {
		return r.Foo
	})

	if avg != 0 {
		t.Fatalf("expected 0, got %v", avg)
	}
}
