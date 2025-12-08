package collection

import (
	"encoding/json"
	"math/rand"
	"testing"
	"time"
)

// utility for generating big slices
func makeIntSlice(n int) []int {
	out := make([]int, n)
	for i := 0; i < n; i++ {
		out[i] = rand.Int()
	}
	return out
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

//
// ───────────────────────── BASIC OPS ─────────────────────────
//

func BenchmarkNew(b *testing.B) {
	data := makeIntSlice(100000)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = New(data)
	}
}

func BenchmarkItems(b *testing.B) {
	c := New(makeIntSlice(100000))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = c.Items()
	}
}

func BenchmarkFirst(b *testing.B) {
	c := New(makeIntSlice(100000))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		c.First()
	}
}

func BenchmarkLast(b *testing.B) {
	c := New(makeIntSlice(100000))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		c.Last()
	}
}

func BenchmarkAny(b *testing.B) {
	c := New(makeIntSlice(100000))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		c.Any(func(v int) bool { return v == -1 })
	}
}

func BenchmarkContains(b *testing.B) {
	c := New(makeIntSlice(100000))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		c.Contains(func(v int) bool { return v == -1 })
	}
}

//
// ───────────────────────── FILTER / MAP / UNIQUE ─────────────────────────
//

func BenchmarkFilter(b *testing.B) {
	c := New(makeIntSlice(200000))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = c.Filter(func(v int) bool { return v%2 == 0 })
	}
}

func BenchmarkMap(b *testing.B) {
	c := New(makeIntSlice(200000))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = c.Map(func(v int) int { return v * 2 })
	}
}

func BenchmarkTransform(b *testing.B) {
	c := New(makeIntSlice(200000))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		cc := New(c.items) // clone each loop
		cc.Transform(func(v int) int { return v * 2 })
	}
}

func BenchmarkUnique(b *testing.B) {
	c := New(makeIntSlice(50000))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = c.Unique(func(a, b int) bool { return a == b })
	}
}

//
// ───────────────────────── SORTING ─────────────────────────
//

func BenchmarkSort(b *testing.B) {
	c := New(makeIntSlice(100000))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = c.Sort(func(a, b int) bool { return a < b })
	}
}

//
// ───────────────────────── MERGE OPS ─────────────────────────
//

func BenchmarkMergeSlice(b *testing.B) {
	c := New(makeIntSlice(50000))
	other := makeIntSlice(50000)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = c.Merge(other)
	}
}

func BenchmarkMergeCollection(b *testing.B) {
	c1 := New(makeIntSlice(50000))
	c2 := New(makeIntSlice(50000))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = c1.Merge(c2)
	}
}

func BenchmarkMergeMap(b *testing.B) {
	c := New(makeIntSlice(50000))
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = c.Merge(m)
	}
}

//
// ───────────────────────── JSON OPS ─────────────────────────
//

func BenchmarkToJSON(b *testing.B) {
	c := New(makeIntSlice(50000))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = c.ToJSON()
	}
}

func BenchmarkToPrettyJSON(b *testing.B) {
	c := New(makeIntSlice(20000))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = c.ToPrettyJSON()
	}
}

// control benchmark – raw json
func BenchmarkRawJSON(b *testing.B) {
	data := makeIntSlice(50000)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(data)
	}
}

//
// ───────────────────────── CHAINED PIPELINE ─────────────────────────
//

func BenchmarkChain(b *testing.B) {
	c := New(makeIntSlice(200000))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		out := c.
			Filter(func(v int) bool { return v%2 == 0 }).
			Map(func(v int) int { return v * 3 }).
			Sort(func(a, b int) bool { return a < b })

		if out.IsEmpty() && len(out.items) != 0 {
			b.Fatalf("sanity check failed")
		}
	}
}

//
// ───────────────────────── CHUNK / MULTIPLY / REDUCE ─────────────────────────
//

func BenchmarkChunk(b *testing.B) {
	c := New(makeIntSlice(200000))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = c.Chunk(100)
	}
}

func BenchmarkMultiply(b *testing.B) {
	c := New(makeIntSlice(20000))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = c.Multiply(10)
	}
}

func BenchmarkReduce(b *testing.B) {
	c := New(makeIntSlice(200_000))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = c.Reduce(0, func(acc, v int) int {
			return acc + v
		})
	}
}

//
// ───────────────────────── PLUCK / MAPTO ─────────────────────────
//

func BenchmarkMapTo(b *testing.B) {
	c := New(makeIntSlice(200000))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = MapTo(c, func(v int) int { return v * 2 })
	}
}

func BenchmarkPluck(b *testing.B) {
	c := New(makeIntSlice(200000))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = Pluck(c, func(v int) int { return v })
	}
}

//
// ───────────────────────── TAKE / TAKEUNTIL ─────────────────────────
//

func BenchmarkTake(b *testing.B) {
	c := New(makeIntSlice(200000))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = c.Take(50000)
	}
}

func BenchmarkTakeUntilFn(b *testing.B) {
	c := New(makeIntSlice(200000))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = c.TakeUntilFn(func(v int) bool { return v > 100000 })
	}
}
