//go:build benchlo
// +build benchlo

package collection

import (
	"testing"

	"github.com/samber/lo"
)

const (
	benchLoSize        = 20_000
	benchLoPipelineLen = 1_000
	benchLoChunkSize   = 50
)

var (
	benchLoInts    []int
	benchLoIntsDup []int
)

func init() {
	benchLoInts = make([]int, benchLoSize)
	for i := 0; i < benchLoSize; i++ {
		benchLoInts[i] = i
	}

	benchLoIntsDup = make([]int, benchLoSize)
	for i := 0; i < benchLoSize; i++ {
		benchLoIntsDup[i] = i % 128 // intentional dupes for Unique
	}
}

func Benchmark_Lo_Comparison(b *testing.B) {
	b.Run("Pipeline_Filter/Map/Take/Reduce", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			sum := New(benchLoInts).
				Filter(func(v int) bool { return v%2 == 0 }).
				Map(func(v int) int { return v * v }).
				Take(benchLoPipelineLen).
				Reduce(0, func(acc, v int) int { return acc + v })
			_ = sum
		}
	})

	b.Run("Pipeline_Filter/Map/Take/Reduce_lo", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			out := lo.Filter(benchLoInts, func(v int, _ int) bool { return v%2 == 0 })
			out2 := lo.Map(out, func(v int, _ int) int { return v * v })
			out3 := lo.Subset(out2, 0, benchLoPipelineLen)
			sum := lo.Reduce(out3, func(acc int, v int, _ int) int {
				return acc + v
			}, 0)
			_ = sum
		}
	})

	b.Run("Map", func(b *testing.B) {
		b.ReportAllocs()
		c := New(benchLoInts)
		for i := 0; i < b.N; i++ {
			_ = c.Map(func(v int) int { return v * 3 })
		}
	})

	b.Run("Map_lo", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = lo.Map(benchLoInts, func(v int, _ int) int { return v * 3 })
		}
	})

	b.Run("Filter", func(b *testing.B) {
		b.ReportAllocs()
		c := New(benchLoInts)
		for i := 0; i < b.N; i++ {
			_ = c.Filter(func(v int) bool { return v%3 == 0 })
		}
	})

	b.Run("Filter_lo", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = lo.Filter(benchLoInts, func(v int, _ int) bool { return v%3 == 0 })
		}
	})

	b.Run("Unique", func(b *testing.B) {
		b.ReportAllocs()
		c := New(benchLoIntsDup)
		for i := 0; i < b.N; i++ {
			_ = UniqueBy(c, func(v int) int { return v })
		}
	})

	b.Run("Unique_lo", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = lo.Uniq(benchLoIntsDup)
		}
	})

	b.Run("Chunk", func(b *testing.B) {
		b.ReportAllocs()
		c := New(benchLoInts)
		for i := 0; i < b.N; i++ {
			_ = c.Chunk(benchLoChunkSize)
		}
	})

	b.Run("Chunk_lo", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = lo.Chunk(benchLoInts, benchLoChunkSize)
		}
	})
}
