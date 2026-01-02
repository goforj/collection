package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/goforj/collection"
	"github.com/samber/lo"
)

const (
	benchStart = "<!-- bench:embed:start -->"
	benchEnd   = "<!-- bench:embed:end -->"

	hotPathIters = 10_000

	benchInner = 8
)

type benchResult struct {
	name        string
	nsPerOp     float64
	bytesPerOp  int64
	allocsPerOp int64
	impl        string
}

func main() {
	onlyFlag := flag.String("only", "", "Run only benchmarks matching the name (comma-separated, case-insensitive)")
	flag.Parse()

	start := time.Now()
	only := parseOnly(*onlyFlag)
	borrowResults := runBenches(only, benchBorrow)
	condensed := renderCondensedTables(borrowResults)
	rawTable := renderTable(borrowResults)

	if err := updateReadme(condensed); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	if err := updateBenchmarksFile(rawTable); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Printf(
		"✔ Benchmarks updated in README.md (elapsed %s)\n",
		time.Since(start).Truncate(time.Millisecond),
	)
}

// ----------------------------------------------------------------------------
// Benchmark runner
// ----------------------------------------------------------------------------

type benchMode string

const (
	benchBorrow benchMode = "borrow"
)

var (
	ctorInts       func([]int) *collection.Collection[int]
	ctorNumericInt func([]int) *collection.NumericCollection[int]
	currentMode    benchMode
)

func setBenchMode(mode benchMode) {
	currentMode = mode
	switch mode {
	default:
		ctorInts = collection.New[int]
		ctorNumericInt = collection.NewNumeric[int]
	}
}

func runBenches(only map[string]struct{}, mode benchMode) []benchResult {
	setBenchMode(mode)
	cases := []struct {
		name string
		col  func(*testing.B)
		lo   func(*testing.B)
	}{
		{"Pipeline F→M→T→R", benchPipelineCollection, benchPipelineLo},
		{"All", benchAllCollection, benchAllLo},
		{"Any", benchAnyCollection, benchAnyLo},
		{"None", benchNoneCollection, benchNoneLo},
		{"First", benchFirstCollection, benchFirstLo},
		{"Last", benchLastCollection, benchLastLo},
		{"IndexWhere", benchIndexWhereCollection, benchIndexWhereLo},
		{"Each", benchEachCollection, benchEachLo},
		{"Map", benchMapCollection, benchMapLo},
		{"Reduce (sum)", benchReduceCollection, benchReduceLo},
		{"Filter", benchFilterCollection, benchFilterLo},
		{"Chunk", benchChunkCollection, benchChunkLo},
		{"Take", benchTakeCollection, benchTakeLo},
		{"Contains", benchContainsCollection, benchContainsLo},
		{"FirstWhere", benchFindCollection, benchFindLo},
		{"GroupBySlice", benchGroupByCollection, benchGroupByLo},
		{"CountBy", benchCountByCollection, benchCountByLo},
		{"CountByValue", benchCountByValueCollection, benchCountByValueLo},
		{"Skip", benchSkipCollection, benchSkipLo},
		{"SkipLast", benchSkipLastCollection, benchSkipLastLo},
		{"Reverse", benchReverseCollection, benchReverseLo},
		{"Shuffle", benchShuffleCollection, benchShuffleLo},
		{"Zip", benchZipCollection, benchZipLo},
		{"ZipWith", benchZipWithCollection, benchZipWithLo},
		{"Unique", benchUniqueCollection, benchUniqueLo},
		{"UniqueBy", benchUniqueByCollection, benchUniqueByLo},
		{"Union", benchUnionCollection, benchUnionLo},
		{"Intersect", benchIntersectCollection, benchIntersectLo},
		{"Difference", benchDifferenceCollection, benchDifferenceLo},
		{"ToMap", benchToMapCollection, benchToMapLo},
		{"Sum", benchSumCollection, benchSumLo},
		{"Min", benchMinCollection, benchMinLo},
		{"Max", benchMaxCollection, benchMaxLo},
	}

	var results []benchResult
	for _, c := range cases {
		if len(only) > 0 {
			if _, ok := only[strings.ToLower(c.name)]; !ok {
				continue
			}
		}

		results = append(
			results,
			measure(c.name, "collection", c.col),
			measure(c.name, "lo", c.lo),
		)
	}
	return results
}

func parseOnly(raw string) map[string]struct{} {
	only := make(map[string]struct{})
	for _, part := range strings.Split(raw, ",") {
		name := strings.ToLower(strings.TrimSpace(part))
		if name == "" {
			continue
		}
		only[name] = struct{}{}
	}
	return only
}

func measure(name, impl string, fn func(*testing.B)) benchResult {
	res := testing.Benchmark(func(b *testing.B) {
		b.ReportAllocs()
		fn(b)
	})

	nsPerOp := float64(res.NsPerOp())
	bytesPerOp := res.AllocedBytesPerOp()
	allocsPerOp := res.AllocsPerOp()
	if benchInner > 1 {
		nsPerOp = nsPerOp / float64(benchInner)
		bytesPerOp = bytesPerOp / int64(benchInner)
		allocsPerOp = allocsPerOp / int64(benchInner)
	}

	return benchResult{
		name:        name,
		impl:        impl,
		nsPerOp:     nsPerOp,
		bytesPerOp:  bytesPerOp,
		allocsPerOp: allocsPerOp,
	}
}

// ----------------------------------------------------------------------------
// Bench cases
// ----------------------------------------------------------------------------

const (
	benchSize        = 1000
	benchPipelineLen = 40
	benchChunkSize   = 20
	benchSkipN       = 40
	benchGroupByMod  = 10
	benchTakeN       = 40
)

var (
	benchInts       []int
	benchIntsDup    []int
	unionLeft       []int
	unionRight      []int
	intersectLeft   []int
	intersectRight  []int
	differenceLeft  []int
	differenceRight []int
	workA           []int
	workB           []int
)

func init() {
	benchInts = make([]int, benchSize)
	for i := 0; i < benchSize; i++ {
		benchInts[i] = i
	}

	benchIntsDup = make([]int, benchSize)
	for i := 0; i < benchSize; i++ {
		benchIntsDup[i] = i % 128
	}

	// overlapping ranges to exercise set ops
	unionLeft = benchIntsDup
	unionRight = benchInts
	intersectLeft = benchIntsDup
	intersectRight = benchInts
	differenceLeft = benchInts
	differenceRight = benchIntsDup

	workA = make([]int, benchSize)
	workB = make([]int, benchSize)
}

func benchPipelineCollection(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchInner; j++ {
			input := collectionInputForMutating(benchInts)
			_ = ctorInts(input).
				Filter(func(v int) bool { return v%2 == 0 }).
				Map(func(v int) int { return v * v }).
				Take(benchPipelineLen).
				Reduce(0, func(acc, v int) int { return acc + v })

		}
	}
}

func benchPipelineLo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchInner; j++ {
			input := benchInts

			out := lo.Filter(input, func(v int, _ int) bool { return v%2 == 0 })
			out2 := lo.Map(out, func(v int, _ int) int { return v * v })
			out3 := lo.Subset(out2, 0, benchPipelineLen)
			_ = lo.Reduce(out3, func(acc int, v int, _ int) int { return acc + v }, 0)

		}
	}
}

func benchAllCollection(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchInner; j++ {
			_ = ctorInts(benchInts).All(func(v int) bool { return v < benchSize+1 })

		}
	}
}

func benchAllLo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchInner; j++ {
			_ = lo.EveryBy(benchInts, func(v int) bool { return v < benchSize+1 })

		}
	}
}

func benchAnyCollection(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchInner; j++ {
			_ = ctorInts(benchInts).Any(func(v int) bool { return v == benchSize-1 })

		}
	}
}

func benchAnyLo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchInner; j++ {
			_ = lo.SomeBy(benchInts, func(v int) bool { return v == benchSize-1 })

		}
	}
}

func benchNoneCollection(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchInner; j++ {
			_ = ctorInts(benchInts).None(func(v int) bool { return v < 0 })

		}
	}
}

func benchNoneLo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchInner; j++ {
			_ = lo.NoneBy(benchInts, func(v int) bool { return v < 0 })

		}
	}
}

func benchFirstCollection(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchInner; j++ {
			_, _ = ctorInts(benchInts).First()

		}
	}
}

func benchFirstLo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchInner; j++ {
			_, _ = lo.First(benchInts)

		}
	}
}

func benchLastCollection(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchInner; j++ {
			_, _ = ctorInts(benchInts).Last()

		}
	}
}

func benchLastLo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchInner; j++ {
			_, _ = lo.Last(benchInts)

		}
	}
}

func benchIndexWhereCollection(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchInner; j++ {
			_, _ = ctorInts(benchInts).IndexWhere(func(v int) bool { return v == benchSize-1 })

		}
	}
}

func benchIndexWhereLo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchInner; j++ {
			_, _, _ = lo.FindIndexOf(benchInts, func(v int) bool { return v == benchSize-1 })

		}
	}
}

func benchEachCollection(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchInner; j++ {
			sum := 0
			ctorInts(benchInts).Each(func(v int) { sum += v })

		}
	}
}

func benchEachLo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchInner; j++ {
			sum := 0
			lo.ForEach(benchInts, func(v int, _ int) { sum += v })

		}
	}
}

func benchMapCollection(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchInner; j++ {
			input := benchInts
			_ = ctorInts(input).Map(func(v int) int { return v * 3 })

		}
	}
}

func benchMapLo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchInner; j++ {
			input := benchInts
			_ = lo.Map(input, func(v int, _ int) int { return v * 3 })

		}
	}
}

func benchReduceCollection(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchInner; j++ {
			_ = ctorInts(benchInts).Reduce(0, func(acc, v int) int { return acc + v })

		}
	}
}

func benchReduceLo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchInner; j++ {
			_ = lo.Reduce(benchInts, func(acc int, v int, _ int) int { return acc + v }, 0)

		}
	}
}

func benchFilterCollection(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchInner; j++ {
			input := collectionInputForMutating(benchInts)
			_ = ctorInts(input).Filter(func(v int) bool { return v%3 == 0 })

		}
	}
}

func benchFilterLo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchInner; j++ {
			input := benchInts
			_ = lo.Filter(input, func(v int, _ int) bool { return v%3 == 0 })

		}
	}
}

func benchChunkCollection(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchInner; j++ {
			_ = ctorInts(benchInts).Chunk(benchChunkSize)

		}
	}
}

func benchChunkLo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchInner; j++ {
			_ = lo.Chunk(benchInts, benchChunkSize)

		}
	}
}

func benchTakeCollection(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchInner; j++ {
			_ = ctorInts(benchInts).Take(benchTakeN)

		}
	}
}

func benchTakeLo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchInner; j++ {
			_ = lo.Subset(benchInts, 0, uint(benchTakeN))

		}
	}
}

func benchContainsCollection(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchInner; j++ {
			_ = collection.Contains(ctorInts(benchInts), benchSize-1)

		}
	}
}

func benchContainsLo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchInner; j++ {
			_ = lo.ContainsBy(benchInts, func(v int) bool { return v == benchSize-1 })

		}
	}
}

func benchFindCollection(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchInner; j++ {
			_, _ = ctorInts(benchInts).FirstWhere(func(v int) bool { return v == benchSize-1 })

		}
	}
}

func benchFindLo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchInner; j++ {
			_, _ = lo.Find(benchInts, func(v int) bool { return v == benchSize-1 })

		}
	}
}

func benchGroupByCollection(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchInner; j++ {
			_ = collection.GroupBySlice(ctorInts(benchInts), func(v int) int { return v % benchGroupByMod })

		}
	}
}

func benchGroupByLo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchInner; j++ {
			_ = lo.GroupBy(benchInts, func(v int) int { return v % benchGroupByMod })

		}
	}
}

func benchCountByCollection(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchInner; j++ {
			_ = collection.CountBy(ctorInts(benchIntsDup), func(v int) int { return v })

		}
	}
}

func benchCountByLo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchInner; j++ {
			_ = lo.CountValuesBy(benchIntsDup, func(v int) int { return v })

		}
	}
}

func benchCountByValueCollection(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchInner; j++ {
			_ = collection.CountByValue(ctorInts(benchIntsDup))

		}
	}
}

func benchCountByValueLo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchInner; j++ {
			_ = lo.CountValues(benchIntsDup)

		}
	}
}

func benchSkipCollection(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchInner; j++ {
			_ = ctorInts(benchInts).Skip(benchSkipN)

		}
	}
}

func benchSkipLo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchInner; j++ {
			_ = lo.Drop(benchInts, benchSkipN)

		}
	}
}

func benchSkipLastCollection(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchInner; j++ {
			_ = ctorInts(benchInts).SkipLast(benchSkipN)

		}
	}
}

func benchSkipLastLo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchInner; j++ {
			_ = lo.DropRight(benchInts, benchSkipN)

		}
	}
}

func benchReverseCollection(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchInner; j++ {
			input := collectionInputForMutating(benchInts)
			_ = ctorInts(input).Reverse()

		}
	}
}

func benchReverseLo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchInner; j++ {
			copy(workB, benchInts)
			_ = lo.Reverse(workB)

		}
	}
}

func benchShuffleCollection(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchInner; j++ {
			input := collectionInputForMutating(benchInts)
			_ = ctorInts(input).Shuffle()

		}
	}
}

func benchShuffleLo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchInner; j++ {
			copy(workB, benchInts)
			_ = lo.Shuffle(workB)

		}
	}
}

func benchZipCollection(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchInner; j++ {
			_ = collection.Zip(ctorInts(benchInts), ctorInts(benchIntsDup))

		}
	}
}

func benchZipLo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchInner; j++ {
			_ = lo.Zip2(benchInts, benchIntsDup)

		}
	}
}

func benchZipWithCollection(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchInner; j++ {
			_ = collection.ZipWith(ctorInts(benchInts), ctorInts(benchIntsDup), func(a, b int) int {
				return a + b
			})

		}
	}
}

func benchZipWithLo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchInner; j++ {
			_ = lo.ZipBy2(benchInts, benchIntsDup, func(a, b int) int {
				return a + b
			})

		}
	}
}

func benchUniqueCollection(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchInner; j++ {
			_ = collection.UniqueComparable(ctorInts(benchIntsDup))

		}
	}
}

func benchUniqueLo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchInner; j++ {
			_ = lo.Uniq(benchIntsDup)

		}
	}
}

func benchUniqueByCollection(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchInner; j++ {
			_ = collection.UniqueBy(ctorInts(benchIntsDup), func(v int) int { return v })

		}
	}
}

func benchUniqueByLo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchInner; j++ {
			_ = lo.UniqBy(benchIntsDup, func(v int) int { return v })

		}
	}
}

func benchUnionCollection(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchInner; j++ {
			_ = collection.Union(ctorInts(unionLeft), ctorInts(unionRight))

		}
	}
}

func benchUnionLo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchInner; j++ {
			_ = lo.Union(unionLeft, unionRight)

		}
	}
}

func benchIntersectCollection(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchInner; j++ {
			_ = collection.Intersect(ctorInts(intersectLeft), ctorInts(intersectRight))

		}
	}
}

func benchIntersectLo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchInner; j++ {
			_ = lo.Intersect(intersectLeft, intersectRight)

		}
	}
}

func benchDifferenceCollection(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchInner; j++ {
			_ = collection.Difference(ctorInts(differenceLeft), ctorInts(differenceRight))

		}
	}
}

func benchDifferenceLo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchInner; j++ {
			_, _ = lo.Difference(differenceLeft, differenceRight)

		}
	}
}

func benchToMapCollection(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchInner; j++ {
			_ = collection.ToMap(ctorInts(benchInts), func(v int) int { return v }, func(v int) int { return v })

		}
	}
}

func benchToMapLo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchInner; j++ {
			_ = lo.SliceToMap(benchInts, func(v int) (int, int) { return v, v })

		}
	}
}

func benchSumCollection(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchInner; j++ {
			_ = ctorNumericInt(benchInts).Sum()

		}
	}
}

func benchSumLo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchInner; j++ {
			_ = lo.Sum(benchInts)

		}
	}
}

func benchMinCollection(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchInner; j++ {
			_, _ = ctorNumericInt(benchInts).Min()

		}
	}
}

func benchMinLo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchInner; j++ {
			_ = lo.Min(benchInts)

		}
	}
}

func benchMaxCollection(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchInner; j++ {
			_, _ = ctorNumericInt(benchInts).Max()

		}
	}
}

func benchMaxLo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchInner; j++ {
			_ = lo.Max(benchInts)

		}
	}
}

// ----------------------------------------------------------------------------
// Rendering
// ----------------------------------------------------------------------------

func renderTable(results []benchResult) string {
	byName := map[string]map[string]benchResult{}
	for _, r := range results {
		if _, ok := byName[r.name]; !ok {
			byName[r.name] = map[string]benchResult{}
		}
		byName[r.name][r.impl] = r
	}

	var buf bytes.Buffer
	buf.WriteString("| Op | ns/op (vs lo) | × (faster) | bytes/op (vs lo) | × (less memory) | allocs/op (vs lo) |\n")
	buf.WriteString("|---:|----------------|:--:|------------------|:--:|--------------------|\n")

	names := make([]string, 0, len(byName))
	for name := range byName {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		col := byName[name]["collection"]
		loRes := byName[name]["lo"]

		nsCell := fmt.Sprintf(
			"%s / %s",
			formatNs(col.nsPerOp),
			formatNs(loRes.nsPerOp),
		)
		ratioCell := formatRatio(loRes.nsPerOp, col.nsPerOp)

		bytesCell := fmt.Sprintf(
			"%s / %s",
			formatBytes(col.bytesPerOp),
			formatBytes(loRes.bytesPerOp),
		)
		bytesRatioCell := formatRatioBytes(loRes.bytesPerOp, col.bytesPerOp)

		allocCell := fmt.Sprintf("%d / %d", col.allocsPerOp, loRes.allocsPerOp)

		buf.WriteString(fmt.Sprintf(
			"| **%s** | %s | %s | %s | %s | %s |\n",
			name,
			nsCell,
			ratioCell,
			bytesCell,
			bytesRatioCell,
			allocCell,
		))
	}

	return strings.TrimSpace(buf.String())
}

type benchGroup struct {
	name string
	ops  []string
}

func renderCondensedTables(results []benchResult) string {
	byName := map[string]map[string]benchResult{}
	for _, r := range results {
		if _, ok := byName[r.name]; !ok {
			byName[r.name] = map[string]benchResult{}
		}
		byName[r.name][r.impl] = r
	}

	groups := []benchGroup{
		{
			name: "Read-only scalar ops (wrapper overhead only)",
			ops: []string{
				"All",
				"Any",
				"None",
				"First",
				"Last",
				"FirstWhere",
				"IndexWhere",
				"Contains",
				"Reduce (sum)",
				"Sum",
				"Min",
				"Max",
				"Each",
			},
		},
		{
			name: "Transforming ops",
			ops: []string{
				"Map",
				"Chunk",
				"Take",
				"Skip",
				"SkipLast",
				"Zip",
				"ZipWith",
				"Unique",
				"UniqueBy",
				"Union",
				"Intersect",
				"Difference",
				"GroupBySlice",
				"CountBy",
				"CountByValue",
				"ToMap",
			},
		},
		{
			name: "Pipelines",
			ops: []string{
				"Pipeline F→M→T→R",
			},
		},
		{
			name: "Mutating ops",
			ops: []string{
				"Filter",
				"Reverse",
				"Shuffle",
			},
		},
	}

	var buf bytes.Buffer
	buf.WriteString("Full raw tables: see `BENCHMARKS.md`.\n\n")

	for _, group := range groups {
		rows := make([]string, 0, len(group.ops))
		for _, name := range group.ops {
			entry, ok := byName[name]
			if !ok {
				continue
			}
			col := entry["collection"]
			loRes := entry["lo"]

			wrapperOnly := group.name == "Read-only scalar ops (wrapper overhead only)"
			allowBold := group.name == "Pipelines" || group.name == "Transforming ops" || group.name == "Mutating ops"
			speed := formatSpeed(loRes.nsPerOp, col.nsPerOp, allowBold, wrapperOnly)
			mem := formatDeltaBytes(col.bytesPerOp, loRes.bytesPerOp)
			allocs := formatDeltaAllocs(col.allocsPerOp, loRes.allocsPerOp)

			rows = append(rows, fmt.Sprintf("| **%s** | %s | %s | %s |", name, speed, mem, allocs))
		}

		if len(rows) == 0 {
			continue
		}

		buf.WriteString(fmt.Sprintf("#### %s\n\n", group.name))
		buf.WriteString("| Op | Speed vs lo | Memory | Allocs |\n")
		buf.WriteString("|---:|:-----------:|:------:|:------:|\n")
		buf.WriteString(strings.Join(rows, "\n"))
		buf.WriteString("\n\n")
	}

	return strings.TrimSpace(buf.String())
}

func formatNs(ns float64) string {
	switch {
	case ns < 1:
		return "<1ns"
	case ns >= 1e6:
		return fmt.Sprintf("%.1fms", ns/1e6)
	case ns >= 1e3:
		return fmt.Sprintf("%.1fµs", ns/1e3)
	default:
		return fmt.Sprintf("%.0fns", ns)
	}
}

func formatBytes(bytes int64) string {
	switch {
	case bytes >= 1_000_000:
		return fmt.Sprintf("%.1fMB", float64(bytes)/1_000_000)
	case bytes >= 1_000:
		return fmt.Sprintf("%.1fKB", float64(bytes)/1_000)
	default:
		return fmt.Sprintf("%dB", bytes)
	}
}

func formatDurationNs(ns float64) string {
	switch {
	case ns >= 1e9:
		return fmt.Sprintf("%.2fs", ns/1e9)
	case ns >= 1e6:
		return fmt.Sprintf("%.0fms", ns/1e6)
	case ns >= 1e3:
		return fmt.Sprintf("%.0fµs", ns/1e3)
	default:
		return fmt.Sprintf("%.0fns", ns)
	}
}

const (
	wrapperEpsilon     = 0.10 // ±10% wrapper overhead tolerance
	benchRatioNoiseNs  = 50.0
	wrapperOnlyEpsilon = 0.15
)

func formatRatio(lo, col float64) string {
	if lo < benchRatioNoiseNs && col < benchRatioNoiseNs {
		return "≈"
	}
	if col == 0 {
		return "∞"
	}

	ratio := lo / col

	// Treat small deltas as equivalent (wrapper overhead, measurement noise)
	if ratio >= 1-wrapperEpsilon && ratio <= 1+wrapperEpsilon {
		return "≈"
	}

	out := fmt.Sprintf("%.2fx", ratio)
	if ratio > 1 {
		return fmt.Sprintf("**%s**", out)
	}
	return out
}

func formatSpeed(lo, col float64, allowBold bool, wrapperOnly bool) string {
	if lo < benchRatioNoiseNs && col < benchRatioNoiseNs {
		return "≈"
	}
	if col == 0 {
		return "∞"
	}

	ratio := lo / col
	if wrapperOnly && ratio >= 1-wrapperOnlyEpsilon && ratio <= 1+wrapperOnlyEpsilon {
		return "≈"
	}
	if ratio >= 1-wrapperEpsilon && ratio <= 1+wrapperEpsilon {
		return "≈"
	}

	out := fmt.Sprintf("%.2fx", ratio)
	if ratio > 1 && allowBold {
		return fmt.Sprintf("**%s**", out)
	}
	return out
}

func formatRatioBytes(lo, col int64) string {
	switch {
	case lo == 0 && col == 0:
		return "≈"
	case col == 0:
		return "**∞x less**"
	case lo == 0:
		return "∞x more"
	}

	ratio := float64(lo) / float64(col)
	if ratio >= 0.90 && ratio <= 1.10 {
		return "≈"
	}

	out := fmt.Sprintf("%.2fx", ratio)
	if ratio > 1 {
		return fmt.Sprintf("**%s less**", out)
	}
	return fmt.Sprintf("%s more", out)
}

func formatInt(v int64) string {
	if v < 0 {
		v = -v
	}
	switch {
	case v >= 1_000_000:
		return fmt.Sprintf("%.1fM", float64(v)/1_000_000)
	case v >= 1_000:
		return fmt.Sprintf("%dk", v/1_000)
	default:
		return fmt.Sprintf("%d", v)
	}
}

func formatDeltaBytes(col, lo int64) string {
	if col == lo {
		return "≈"
	}
	diff := col - lo
	if diff > 0 {
		return fmt.Sprintf("+%s", formatBytes(diff))
	}
	return fmt.Sprintf("-%s", formatBytes(-diff))
}

func formatDeltaAllocs(col, lo int64) string {
	if col == lo {
		return "≈"
	}
	diff := col - lo
	if diff > 0 {
		return fmt.Sprintf("+%d", diff)
	}
	return fmt.Sprintf("%d", diff)
}

func collectionInputForMutating(src []int) []int {
	if currentMode == benchBorrow {
		copy(workA, src)
		return workA
	}
	return src
}

// ----------------------------------------------------------------------------
// README injection
// ----------------------------------------------------------------------------

func updateReadme(condensed string) error {
	root, err := findRoot()
	if err != nil {
		return err
	}

	readmePath := filepath.Join(root, "README.md")
	data, err := os.ReadFile(readmePath)
	if err != nil {
		return err
	}

	out, err := replaceSection(string(data), condensed)
	if err != nil {
		return err
	}

	return os.WriteFile(readmePath, []byte(out), 0o644)
}

func replaceSection(readme, condensed string) (string, error) {
	start := strings.Index(readme, benchStart)
	end := strings.Index(readme, benchEnd)
	if start == -1 || end == -1 || end < start {
		return "", fmt.Errorf("benchmark anchors not found or malformed")
	}

	section := readme[start+len(benchStart) : end]
	updated, err := replaceBenchTable(section, condensed)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	buf.WriteString(readme[:start+len(benchStart)])
	buf.WriteString(updated)
	buf.WriteString(readme[end:])
	return buf.String(), nil
}

func replaceBenchTable(section, condensed string) (string, error) {
	trimmed := strings.TrimSpace(condensed)
	if trimmed == "" {
		return "", fmt.Errorf("condensed benchmark content is empty")
	}
	return "\n\n" + trimmed + "\n", nil
}

func updateBenchmarksFile(rawTable string) error {
	root, err := findRoot()
	if err != nil {
		return err
	}

	path := filepath.Join(root, "BENCHMARKS.md")
	var buf bytes.Buffer
	buf.WriteString("# Benchmarks\n\n")
	buf.WriteString("Raw results for `collection.New` (borrowed) vs `lo`.\n\n")
	buf.WriteString(rawTable)
	buf.WriteString("\n")
	return os.WriteFile(path, buf.Bytes(), 0o644)
}

// ----------------------------------------------------------------------------
// Helpers
// ----------------------------------------------------------------------------

const projectModule = "module github.com/goforj/collection"

func findRoot() (string, error) {
	dir, _ := os.Getwd()
	for {
		gm := filepath.Join(dir, "go.mod")
		if fileExists(gm) && isProjectModule(gm) {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	return "", fmt.Errorf("could not find project root")
}

func fileExists(p string) bool {
	_, err := os.Stat(p)
	return err == nil
}

func isProjectModule(path string) bool {
	data, err := os.ReadFile(path)
	if err != nil {
		return false
	}
	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "module ") {
			return line == projectModule
		}
		if line != "" && !strings.HasPrefix(line, "//") {
			return false
		}
	}
	return false
}
