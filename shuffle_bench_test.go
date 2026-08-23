package collection

import (
	"runtime"
	"testing"
)

// BenchmarkShuffleFresh measures Shuffle with a restored 1,000-element input.
func BenchmarkShuffleFresh(b *testing.B) {
	original := make([]int, 1_000)
	for i := range original {
		original[i] = i
	}
	work := make([]int, len(original))

	b.ReportAllocs()
	b.ResetTimer()
	for b.Loop() {
		copy(work, original)
		result := New(work).Shuffle()
		runtime.KeepAlive(result)
	}
}
