package collection

import (
	"strconv"
	"testing"
)

var genericMethodBenchmarkResult any

// BenchmarkGenericMethods measures the retained generic method APIs.
func BenchmarkGenericMethods(b *testing.B) {
	values := New(genericMethodBenchmarkValues(1_000))
	other := genericMethodBenchmarkValues(1_000)
	keyFn := func(value int) int { return value % 64 }

	b.Run("Map", func(b *testing.B) {
		for range b.N {
			genericMethodBenchmarkResult = values.Map(strconv.Itoa)
		}
	})
	b.Run("GroupBy", func(b *testing.B) {
		for range b.N {
			genericMethodBenchmarkResult = values.GroupBy(keyFn)
		}
	})
	b.Run("CountBy", func(b *testing.B) {
		for range b.N {
			genericMethodBenchmarkResult = values.CountBy(keyFn)
		}
	})
	b.Run("UniqueBy", func(b *testing.B) {
		for range b.N {
			genericMethodBenchmarkResult = values.UniqueBy(keyFn)
		}
	})
	b.Run("ToMap", func(b *testing.B) {
		for range b.N {
			genericMethodBenchmarkResult = values.ToMap(keyFn, strconv.Itoa)
		}
	})
	b.Run("ZipWith", func(b *testing.B) {
		for range b.N {
			genericMethodBenchmarkResult = values.ZipWith(other, func(a, b int) int { return a + b })
		}
	})
}

// genericMethodBenchmarkValues creates representative input without relying on removed API benchmarks.
func genericMethodBenchmarkValues(size int) []int {
	values := make([]int, size)
	for i := range values {
		values[i] = i
	}
	return values
}
