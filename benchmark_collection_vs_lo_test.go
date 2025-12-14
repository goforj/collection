//go:build ignore
// +build ignore

package collection

import (
	"testing"

	"github.com/samber/lo"
)

var benchData []int

func init() {
	benchData = make([]int, 10_000)
	for i := 0; i < len(benchData); i++ {
		benchData[i] = i
	}
}

// pipelineCount controls how many items we take in the final stage
const pipelineCount = 500

func Benchmark_GoForj_Collection_Pipeline(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sum := New(benchData).
			Filter(func(v int) bool { return v%2 == 0 }).
			Map(func(v int) int { return v * v }).
			Take(pipelineCount).
			Pipe(func(c *Collection[int]) any {
				return c.Reduce(0, func(acc, v int) int { return acc + v })
			}).(int)

		_ = sum
	}
}

func Benchmark_Lo_Pipeline(b *testing.B) {
	for n := 0; n < b.N; n++ {
		out := lo.Filter(benchData, func(v int, _ int) bool { return v%2 == 0 })
		out2 := lo.Map(out, func(v int, _ int) int { return v * v })
		out3 := lo.Subset(out2, 0, pipelineCount)
		sum := lo.Reduce(out3, func(acc int, item int, index int) int {
			return acc + item
		}, 0)

		_ = sum
	}
}
