package collection

import (
	"runtime"
	"testing"
)

// BenchmarkRetainFresh measures Retain with restored input on every iteration.
func BenchmarkRetainFresh(b *testing.B) {
	original := make([]int, 1_000)
	for index := range original {
		original[index] = index
	}
	work := make([]int, len(original))

	b.ReportAllocs()
	b.ResetTimer()
	for range b.N {
		copy(work, original)
		result := New(work).Retain(func(value int) bool { return value%2 == 0 })
		runtime.KeepAlive(result)
	}
}

// BenchmarkTransformFresh measures Transform with restored input on every iteration.
func BenchmarkTransformFresh(b *testing.B) {
	original := make([]int, 1_000)
	for index := range original {
		original[index] = index
	}
	work := make([]int, len(original))

	b.ReportAllocs()
	b.ResetTimer()
	for range b.N {
		copy(work, original)
		result := New(work).Transform(func(value int) int { return value + 1 })
		runtime.KeepAlive(result)
	}
}
