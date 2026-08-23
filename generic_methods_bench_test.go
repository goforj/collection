package collection

import (
	"strconv"
	"testing"
)

var genericMethodBenchmarkResult any

// BenchmarkGenericMethodParity compares direct compatibility-function and method calls without blocking caller inlining.
func BenchmarkGenericMethodParity(b *testing.B) {
	values := New(makeBenchmarkValues(1_000))
	other := New(makeBenchmarkValues(1_000))
	keyFn := func(value int) int { return value % 64 }

	b.Run("MapTo/Function", func(b *testing.B) {
		b.ReportAllocs()
		for range b.N {
			genericMethodBenchmarkResult = MapTo(values, strconv.Itoa)
		}
	})
	b.Run("MapTo/Method", func(b *testing.B) {
		b.ReportAllocs()
		for range b.N {
			genericMethodBenchmarkResult = values.MapTo(strconv.Itoa)
		}
	})

	b.Run("GroupBy/Function", func(b *testing.B) {
		b.ReportAllocs()
		for range b.N {
			genericMethodBenchmarkResult = GroupBy(values, keyFn)
		}
	})
	b.Run("GroupBy/Method", func(b *testing.B) {
		b.ReportAllocs()
		for range b.N {
			genericMethodBenchmarkResult = values.GroupBy(keyFn)
		}
	})

	b.Run("CountBy/Function", func(b *testing.B) {
		b.ReportAllocs()
		for range b.N {
			genericMethodBenchmarkResult = CountBy(values, keyFn)
		}
	})
	b.Run("CountBy/Method", func(b *testing.B) {
		b.ReportAllocs()
		for range b.N {
			genericMethodBenchmarkResult = values.CountBy(keyFn)
		}
	})

	b.Run("UniqueBy/Function", func(b *testing.B) {
		b.ReportAllocs()
		for range b.N {
			genericMethodBenchmarkResult = UniqueBy(values, keyFn)
		}
	})
	b.Run("UniqueBy/Method", func(b *testing.B) {
		b.ReportAllocs()
		for range b.N {
			genericMethodBenchmarkResult = values.UniqueBy(keyFn)
		}
	})

	b.Run("ToMap/Function", func(b *testing.B) {
		b.ReportAllocs()
		for range b.N {
			genericMethodBenchmarkResult = ToMap(values, keyFn, strconv.Itoa)
		}
	})
	b.Run("ToMap/Method", func(b *testing.B) {
		b.ReportAllocs()
		for range b.N {
			genericMethodBenchmarkResult = values.ToMap(keyFn, strconv.Itoa)
		}
	})

	b.Run("ZipWith/Function", func(b *testing.B) {
		b.ReportAllocs()
		for range b.N {
			genericMethodBenchmarkResult = ZipWith(values, other, addInts)
		}
	})
	b.Run("ZipWith/Method", func(b *testing.B) {
		b.ReportAllocs()
		for range b.N {
			genericMethodBenchmarkResult = values.ZipWith(other, addInts)
		}
	})
}

// addInts provides a shared non-closure callback for ZipWith benchmarks.
func addInts(a, b int) int {
	return a + b
}

// makeBenchmarkValues creates repeatable inputs with enough cardinality to exercise map allocation paths.
func makeBenchmarkValues(count int) []int {
	values := make([]int, count)
	for i := range values {
		values[i] = i
	}
	return values
}
