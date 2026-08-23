package regression_test

import (
	"runtime"
	"testing"

	collectionv2 "github.com/goforj/collection"
	collectionv4 "github.com/goforj/collection/v4"
)

const benchmarkSize = 1_000

// BenchmarkMutableFilterV2 measures the v2 in-place filtering path with fresh input.
func BenchmarkMutableFilterV2(b *testing.B) {
	original, work := benchmarkInputs()
	b.ReportAllocs()
	b.ResetTimer()
	for b.Loop() {
		copy(work, original)
		result := collectionv2.New(work).Filter(keepEven)
		runtime.KeepAlive(result)
	}
}

// BenchmarkMutableFilterV4 measures the v4 in-place filtering path with fresh input.
func BenchmarkMutableFilterV4(b *testing.B) {
	original, work := benchmarkInputs()
	b.ReportAllocs()
	b.ResetTimer()
	for b.Loop() {
		copy(work, original)
		result := collectionv4.New(work).Retain(keepEven)
		runtime.KeepAlive(result)
	}
}

// BenchmarkCopiedFilterV2 measures v2 filtering while preserving the input.
func BenchmarkCopiedFilterV2(b *testing.B) {
	original, _ := benchmarkInputs()
	b.ReportAllocs()
	b.ResetTimer()
	for b.Loop() {
		result := collectionv2.New(original).Clone().Filter(keepEven)
		runtime.KeepAlive(result)
	}
}

// BenchmarkCopiedFilterV4 measures v4 filtering while preserving the input.
func BenchmarkCopiedFilterV4(b *testing.B) {
	original, _ := benchmarkInputs()
	b.ReportAllocs()
	b.ResetTimer()
	for b.Loop() {
		result := collectionv4.New(original).Filter(keepEven)
		runtime.KeepAlive(result)
	}
}

// BenchmarkMutableTransformV2 measures the v2 in-place transformation path with fresh input.
func BenchmarkMutableTransformV2(b *testing.B) {
	original, work := benchmarkInputs()
	b.ReportAllocs()
	b.ResetTimer()
	for b.Loop() {
		copy(work, original)
		result := collectionv2.New(work).Map(double)
		runtime.KeepAlive(result)
	}
}

// BenchmarkMutableTransformV4 measures the v4 in-place transformation path with fresh input.
func BenchmarkMutableTransformV4(b *testing.B) {
	original, work := benchmarkInputs()
	b.ReportAllocs()
	b.ResetTimer()
	for b.Loop() {
		copy(work, original)
		result := collectionv4.New(work).Transform(double)
		runtime.KeepAlive(result)
	}
}

// BenchmarkMutablePipelineV2 measures the v2 in-place fluent pipeline with fresh input.
func BenchmarkMutablePipelineV2(b *testing.B) {
	original, work := benchmarkInputs()
	b.ReportAllocs()
	b.ResetTimer()
	for b.Loop() {
		copy(work, original)
		result := collectionv2.New(work).Filter(keepEven).Map(double).Take(40).Reduce(0, sum)
		runtime.KeepAlive(result)
	}
}

// BenchmarkMutablePipelineV4 measures the v4 in-place fluent pipeline with fresh input.
func BenchmarkMutablePipelineV4(b *testing.B) {
	original, work := benchmarkInputs()
	b.ReportAllocs()
	b.ResetTimer()
	for b.Loop() {
		copy(work, original)
		result := collectionv4.New(work).Retain(keepEven).Transform(double).Take(40).Reduce(0, sum)
		runtime.KeepAlive(result)
	}
}

// BenchmarkCopiedPipelineV2 measures a v2 pipeline that preserves its input.
func BenchmarkCopiedPipelineV2(b *testing.B) {
	original, _ := benchmarkInputs()
	b.ReportAllocs()
	b.ResetTimer()
	for b.Loop() {
		result := collectionv2.New(original).Clone().Filter(keepEven).Map(double).Take(40).Reduce(0, sum)
		runtime.KeepAlive(result)
	}
}

// BenchmarkCopiedPipelineV4 measures a v4 pipeline that preserves its input.
func BenchmarkCopiedPipelineV4(b *testing.B) {
	original, _ := benchmarkInputs()
	b.ReportAllocs()
	b.ResetTimer()
	for b.Loop() {
		result := collectionv4.New(original).Filter(keepEven).Map(double).Take(40).Reduce(0, sum)
		runtime.KeepAlive(result)
	}
}

// BenchmarkShuffleV2 measures v2 Shuffle with fresh input.
func BenchmarkShuffleV2(b *testing.B) {
	original, work := benchmarkInputs()
	b.ReportAllocs()
	b.ResetTimer()
	for b.Loop() {
		copy(work, original)
		result := collectionv2.New(work).Shuffle()
		runtime.KeepAlive(result)
	}
}

// BenchmarkShuffleV4 measures v4 Shuffle with fresh input.
func BenchmarkShuffleV4(b *testing.B) {
	original, work := benchmarkInputs()
	b.ReportAllocs()
	b.ResetTimer()
	for b.Loop() {
		copy(work, original)
		result := collectionv4.New(work).Shuffle()
		runtime.KeepAlive(result)
	}
}

// benchmarkInputs returns immutable source data and same-sized scratch storage.
func benchmarkInputs() ([]int, []int) {
	original := make([]int, benchmarkSize)
	for index := range original {
		original[index] = index
	}
	return original, make([]int, len(original))
}

// keepEven selects even benchmark values.
func keepEven(value int) bool {
	return value%2 == 0
}

// double transforms a benchmark value without changing its type.
func double(value int) int {
	return value * 2
}

// sum adds a benchmark value to its accumulator.
func sum(total, value int) int {
	return total + value
}
