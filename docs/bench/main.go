package main

import (
	"bytes"
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
)

type benchResult struct {
	name        string
	nsPerOp     float64
	bytesPerOp  int64
	allocsPerOp int64
	impl        string
}

func main() {
	start := time.Now()
	results := runBenches()
	table := renderTable(results)

	if err := updateReadme(table); err != nil {
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

func runBenches() []benchResult {
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
		{"Find", benchFindCollection, benchFindLo},
		{"GroupBy", benchGroupByCollection, benchGroupByLo},
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
		results = append(
			results,
			measure(c.name, "collection", c.col),
			measure(c.name, "lo", c.lo),
		)
	}
	return results
}

func measure(name, impl string, fn func(*testing.B)) benchResult {
	res := testing.Benchmark(func(b *testing.B) {
		b.ReportAllocs()
		fn(b)
	})

	return benchResult{
		name:        name,
		impl:        impl,
		nsPerOp:     float64(res.NsPerOp()),
		bytesPerOp:  res.AllocedBytesPerOp(),
		allocsPerOp: res.AllocsPerOp(),
	}
}

// ----------------------------------------------------------------------------
// Bench cases
// ----------------------------------------------------------------------------

const (
	benchSize        = 200
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
		copy(workA, benchInts)

		_ = collection.New(workA).
			Filter(func(v int) bool { return v%2 == 0 }).
			Map(func(v int) int { return v * v }).
			Take(benchPipelineLen).
			Reduce(0, func(acc, v int) int { return acc + v })
	}
}

func benchPipelineLo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copy(workB, benchInts)

		out := lo.Filter(workB, func(v int, _ int) bool { return v%2 == 0 })
		out2 := lo.Map(out, func(v int, _ int) int { return v * v })
		out3 := lo.Subset(out2, 0, benchPipelineLen)
		_ = lo.Reduce(out3, func(acc int, v int, _ int) int { return acc + v }, 0)
	}
}

func benchAllCollection(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = collection.New(benchInts).All(func(v int) bool { return v < benchSize+1 })
	}
}

func benchAllLo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = lo.EveryBy(benchInts, func(v int) bool { return v < benchSize+1 })
	}
}

func benchAnyCollection(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = collection.New(benchInts).Any(func(v int) bool { return v == benchSize-1 })
	}
}

func benchAnyLo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = lo.SomeBy(benchInts, func(v int) bool { return v == benchSize-1 })
	}
}

func benchNoneCollection(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = collection.New(benchInts).None(func(v int) bool { return v < 0 })
	}
}

func benchNoneLo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = lo.NoneBy(benchInts, func(v int) bool { return v < 0 })
	}
}

func benchFirstCollection(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = collection.New(benchInts).First()
	}
}

func benchFirstLo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = lo.First(benchInts)
	}
}

func benchLastCollection(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = collection.New(benchInts).Last()
	}
}

func benchLastLo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = lo.Last(benchInts)
	}
}

func benchIndexWhereCollection(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = collection.New(benchInts).IndexWhere(func(v int) bool { return v == benchSize-1 })
	}
}

func benchIndexWhereLo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _ = lo.FindIndexOf(benchInts, func(v int) bool { return v == benchSize-1 })
	}
}

func benchEachCollection(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sum := 0
		collection.New(benchInts).Each(func(v int) { sum += v })
	}
}

func benchEachLo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sum := 0
		lo.ForEach(benchInts, func(v int, _ int) { sum += v })
	}
}

func benchMapCollection(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copy(workA, benchInts)

		_ = collection.New(workA).Map(func(v int) int { return v * 3 })
	}
}

func benchMapLo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copy(workB, benchInts)

		_ = lo.Map(workB, func(v int, _ int) int { return v * 3 })
	}
}

func benchReduceCollection(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = collection.New(benchInts).Reduce(0, func(acc, v int) int { return acc + v })
	}
}

func benchReduceLo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = lo.Reduce(benchInts, func(acc int, v int, _ int) int { return acc + v }, 0)
	}
}

func benchFilterCollection(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copy(workA, benchInts)

		_ = collection.New(workA).Filter(func(v int) bool { return v%3 == 0 })
	}
}

func benchFilterLo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copy(workB, benchInts)

		_ = lo.Filter(workB, func(v int, _ int) bool { return v%3 == 0 })
	}
}

func benchChunkCollection(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = collection.New(benchInts).Chunk(benchChunkSize)
	}
}

func benchChunkLo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = lo.Chunk(benchInts, benchChunkSize)
	}
}

func benchTakeCollection(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = collection.New(benchInts).Take(benchTakeN)
	}
}

func benchTakeLo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = lo.Subset(benchInts, 0, uint(benchTakeN))
	}
}

func benchContainsCollection(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = collection.New(benchInts).Contains(func(v int) bool { return v == benchSize-1 })
	}
}

func benchContainsLo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = lo.ContainsBy(benchInts, func(v int) bool { return v == benchSize-1 })
	}
}

func benchFindCollection(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = collection.New(benchInts).FirstWhere(func(v int) bool { return v == benchSize-1 })
	}
}

func benchFindLo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = lo.Find(benchInts, func(v int) bool { return v == benchSize-1 })
	}
}

func benchGroupByCollection(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = collection.GroupBy(collection.New(benchInts), func(v int) int { return v % benchGroupByMod })
	}
}

func benchGroupByLo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = lo.GroupBy(benchInts, func(v int) int { return v % benchGroupByMod })
	}
}

func benchCountByCollection(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = collection.CountBy(collection.New(benchIntsDup), func(v int) int { return v })
	}
}

func benchCountByLo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = lo.CountValuesBy(benchIntsDup, func(v int) int { return v })
	}
}

func benchCountByValueCollection(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = collection.CountByValue(collection.New(benchIntsDup))
	}
}

func benchCountByValueLo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = lo.CountValues(benchIntsDup)
	}
}

func benchSkipCollection(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = collection.New(benchInts).Skip(benchSkipN)
	}
}

func benchSkipLo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = lo.Drop(benchInts, benchSkipN)
	}
}

func benchSkipLastCollection(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = collection.New(benchInts).SkipLast(benchSkipN)
	}
}

func benchSkipLastLo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = lo.DropRight(benchInts, benchSkipN)
	}
}

func benchReverseCollection(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copy(workA, benchInts)

		_ = collection.New(workA).Reverse()
	}
}

func benchReverseLo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copy(workB, benchInts)

		_ = lo.Reverse(workB)
	}
}

func benchShuffleCollection(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copy(workA, benchInts)

		_ = collection.New(workA).Shuffle()
	}
}

func benchShuffleLo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copy(workB, benchInts)

		_ = lo.Shuffle(workB)
	}
}

func benchZipCollection(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = collection.Zip(collection.New(benchInts), collection.New(benchIntsDup))
	}
}

func benchZipLo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = lo.Zip2(benchInts, benchIntsDup)
	}
}

func benchZipWithCollection(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = collection.ZipWith(collection.New(benchInts), collection.New(benchIntsDup), func(a, b int) int {
			return a + b
		})
	}
}

func benchZipWithLo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = lo.ZipBy2(benchInts, benchIntsDup, func(a, b int) int {
			return a + b
		})
	}
}

func benchUniqueCollection(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = collection.UniqueComparable(collection.New(benchIntsDup))
	}
}

func benchUniqueLo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = lo.Uniq(benchIntsDup)
	}
}

func benchUniqueByCollection(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = collection.UniqueBy(collection.New(benchIntsDup), func(v int) int { return v })
	}
}

func benchUniqueByLo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = lo.UniqBy(benchIntsDup, func(v int) int { return v })
	}
}

func benchUnionCollection(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = collection.Union(collection.New(unionLeft), collection.New(unionRight))
	}
}

func benchUnionLo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = lo.Union(unionLeft, unionRight)
	}
}

func benchIntersectCollection(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = collection.Intersect(collection.New(intersectLeft), collection.New(intersectRight))
	}
}

func benchIntersectLo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = lo.Intersect(intersectLeft, intersectRight)
	}
}

func benchDifferenceCollection(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = collection.Difference(collection.New(differenceLeft), collection.New(differenceRight))
	}
}

func benchDifferenceLo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = lo.Difference(differenceLeft, differenceRight)
	}
}

func benchToMapCollection(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = collection.ToMap(collection.New(benchInts), func(v int) int { return v }, func(v int) int { return v })
	}
}

func benchToMapLo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = lo.SliceToMap(benchInts, func(v int) (int, int) { return v, v })
	}
}

func benchSumCollection(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = collection.NewNumeric(benchInts).Sum()
	}
}

func benchSumLo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = lo.Sum(benchInts)
	}
}

func benchMinCollection(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = collection.NewNumeric(benchInts).Min()
	}
}

func benchMinLo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = lo.Min(benchInts)
	}
}

func benchMaxCollection(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = collection.NewNumeric(benchInts).Max()
	}
}

func benchMaxLo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = lo.Max(benchInts)
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
	buf.WriteString("### Performance Benchmarks\n\n")
	buf.WriteString("| Op | ns/op (col/lo ×) | allocs/op (col/lo) |\n")
	buf.WriteString("|----|------------------|--------------------|\n")

	names := make([]string, 0, len(byName))
	for name := range byName {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		col := byName[name]["collection"]
		loRes := byName[name]["lo"]

		nsCell := fmt.Sprintf(
			"%s / %s (%s)",
			formatNs(col.nsPerOp),
			formatNs(loRes.nsPerOp),
			formatRatio(loRes.nsPerOp, col.nsPerOp),
		)

		allocCell := fmt.Sprintf("%d / %d", col.allocsPerOp, loRes.allocsPerOp)

		buf.WriteString(fmt.Sprintf(
			"| %s | %s | %s |\n",
			name,
			nsCell,
			allocCell,
		))
	}

	return strings.TrimSpace(buf.String())
}

func formatNs(ns float64) string {
	switch {
	case ns >= 1e6:
		return fmt.Sprintf("%.1fms", ns/1e6)
	case ns >= 1e3:
		return fmt.Sprintf("%.1fµs", ns/1e3)
	default:
		return fmt.Sprintf("%.0fns", ns)
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

func formatRatio(lo, col float64) string {
	if col == 0 {
		return "∞"
	}
	return fmt.Sprintf("%.2fx", lo/col)
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

// ----------------------------------------------------------------------------
// README injection
// ----------------------------------------------------------------------------

func updateReadme(table string) error {
	root, err := findRoot()
	if err != nil {
		return err
	}

	readmePath := filepath.Join(root, "README.md")
	data, err := os.ReadFile(readmePath)
	if err != nil {
		return err
	}

	out, err := replaceSection(string(data), table)
	if err != nil {
		return err
	}

	return os.WriteFile(readmePath, []byte(out), 0o644)
}

func replaceSection(readme, content string) (string, error) {
	start := strings.Index(readme, benchStart)
	end := strings.Index(readme, benchEnd)
	if start == -1 || end == -1 || end < start {
		return "", fmt.Errorf("benchmark anchors not found or malformed")
	}

	var buf bytes.Buffer
	buf.WriteString(readme[:start+len(benchStart)])
	buf.WriteString("\n\n")
	buf.WriteString(content)
	buf.WriteString("\n")
	buf.WriteString(readme[end:])
	return buf.String(), nil
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
