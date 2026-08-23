package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"slices"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/goforj/collection/v3"
	"github.com/samber/lo"
	"github.com/samber/lo/mutable"
)

const (
	benchStart = "<!-- bench:embed:start -->"
	benchEnd   = "<!-- bench:embed:end -->"

	benchSamples        = 7
	benchSampleDuration = 100 * time.Millisecond
)

type benchResult struct {
	name        string
	nsPerOp     float64
	bytesPerOp  int64
	allocsPerOp int64
	impl        string
	uncertain   bool
}

// main runs both ownership modes and updates the benchmark reports.
func main() {
	onlyFlag := flag.String("only", "", "Run only benchmarks matching the name (comma-separated, case-insensitive)")
	flag.Parse()

	start := time.Now()
	only := parseOnly(*onlyFlag)
	borrowResults := runBenches(only, benchBorrow)
	copyResults := runBenches(only, benchCopy)
	condensed := renderCondensedTables(borrowResults, benchBorrow)
	rawBorrow := renderTable(borrowResults, benchBorrow)

	if err := updateReadme(condensed); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	if err := updateBenchmarksFile(rawBorrow, renderTable(copyResults, benchCopy)); err != nil {
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
	benchCopy   benchMode = "copy"
)

var (
	currentMode benchMode
)

// setBenchMode selects whether collection construction borrows or clones input.
func setBenchMode(mode benchMode) {
	currentMode = mode
}

// collectionBenchmark selects a statically specialized body before timing starts.
func collectionBenchmark(borrow, copy func(*testing.B)) func(*testing.B) {
	if currentMode == benchCopy {
		return copy
	}
	return borrow
}

// runBenches executes the selected benchmark cases for one ownership mode.
func runBenches(only map[string]struct{}, mode benchMode) []benchResult {
	testing.Init()
	if err := flag.Set("test.benchtime", benchSampleDuration.String()); err != nil {
		panic(fmt.Sprintf("configure benchmark sample duration: %v", err))
	}
	setBenchMode(mode)
	cases := []struct {
		name string
		col  func(*testing.B)
		lo   func(*testing.B)
	}{
		{"Pipeline F→M→T→R", collectionBenchmark(benchPipelineCollectionBorrow, benchPipelineCollectionCopy), benchPipelineLo},
		{"All", collectionBenchmark(benchAllCollectionBorrow, benchAllCollectionCopy), benchAllLo},
		{"Any", collectionBenchmark(benchAnyCollectionBorrow, benchAnyCollectionCopy), benchAnyLo},
		{"None", collectionBenchmark(benchNoneCollectionBorrow, benchNoneCollectionCopy), benchNoneLo},
		{"First", collectionBenchmark(benchFirstCollectionBorrow, benchFirstCollectionCopy), benchFirstLo},
		{"Last", collectionBenchmark(benchLastCollectionBorrow, benchLastCollectionCopy), benchLastLo},
		{"IndexWhere", collectionBenchmark(benchIndexWhereCollectionBorrow, benchIndexWhereCollectionCopy), benchIndexWhereLo},
		{"Each", collectionBenchmark(benchEachCollectionBorrow, benchEachCollectionCopy), benchEachLo},
		{"Map", collectionBenchmark(benchMapCollectionBorrow, benchMapCollectionCopy), benchMapLo},
		{"Transform", collectionBenchmark(benchTransformCollectionBorrow, benchTransformCollectionCopy), benchTransformLo},
		{"Reduce (sum)", collectionBenchmark(benchReduceCollectionBorrow, benchReduceCollectionCopy), benchReduceLo},
		{"Filter", collectionBenchmark(benchFilterCollectionBorrow, benchFilterCollectionCopy), benchFilterLo},
		{"Retain", collectionBenchmark(benchRetainCollectionBorrow, benchRetainCollectionCopy), benchRetainLo},
		{"Chunk", collectionBenchmark(benchChunkCollectionBorrow, benchChunkCollectionCopy), benchChunkLo},
		{"Take", collectionBenchmark(benchTakeCollectionBorrow, benchTakeCollectionCopy), benchTakeLo},
		{"Contains", collectionBenchmark(benchContainsCollectionBorrow, benchContainsCollectionCopy), benchContainsLo},
		{"FirstWhere", collectionBenchmark(benchFindCollectionBorrow, benchFindCollectionCopy), benchFindLo},
		{"GroupBy", collectionBenchmark(benchGroupByCollectionBorrow, benchGroupByCollectionCopy), benchGroupByLo},
		{"CountBy", collectionBenchmark(benchCountByCollectionBorrow, benchCountByCollectionCopy), benchCountByLo},
		{"CountByValue", collectionBenchmark(benchCountByValueCollectionBorrow, benchCountByValueCollectionCopy), benchCountByValueLo},
		{"Skip", collectionBenchmark(benchSkipCollectionBorrow, benchSkipCollectionCopy), benchSkipLo},
		{"SkipLast", collectionBenchmark(benchSkipLastCollectionBorrow, benchSkipLastCollectionCopy), benchSkipLastLo},
		{"Reverse", collectionBenchmark(benchReverseCollectionBorrow, benchReverseCollectionCopy), benchReverseLo},
		{"Shuffle", collectionBenchmark(benchShuffleCollectionBorrow, benchShuffleCollectionCopy), benchShuffleLo},
		{"Zip", collectionBenchmark(benchZipCollectionBorrow, benchZipCollectionCopy), benchZipLo},
		{"ZipWith", collectionBenchmark(benchZipWithCollectionBorrow, benchZipWithCollectionCopy), benchZipWithLo},
		{"Unique", collectionBenchmark(benchUniqueCollectionBorrow, benchUniqueCollectionCopy), benchUniqueLo},
		{"UniqueBy", collectionBenchmark(benchUniqueByCollectionBorrow, benchUniqueByCollectionCopy), benchUniqueByLo},
		{"Union", collectionBenchmark(benchUnionCollectionBorrow, benchUnionCollectionCopy), benchUnionLo},
		{"Intersect", collectionBenchmark(benchIntersectCollectionBorrow, benchIntersectCollectionCopy), benchIntersectLo},
		{"Difference", collectionBenchmark(benchDifferenceCollectionBorrow, benchDifferenceCollectionCopy), benchDifferenceLo},
		{"ToMap", collectionBenchmark(benchToMapCollectionBorrow, benchToMapCollectionCopy), benchToMapLo},
		{"Sum", collectionBenchmark(benchSumCollectionBorrow, benchSumCollectionCopy), benchSumLo},
		{"Min", collectionBenchmark(benchMinCollectionBorrow, benchMinCollectionCopy), benchMinLo},
		{"Max", collectionBenchmark(benchMaxCollectionBorrow, benchMaxCollectionCopy), benchMaxLo},
	}

	var results []benchResult
	for caseIndex, c := range cases {
		if len(only) > 0 {
			if _, ok := only[strings.ToLower(c.name)]; !ok {
				continue
			}
		}

		collectionResult, loResult := measurePair(c.name, c.col, c.lo, caseIndex%2 == 0)
		results = append(results, collectionResult, loResult)
	}
	return results
}

// parseOnly normalizes the optional benchmark-name filter.
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

// measurePair alternates implementations and returns paired median measurements.
func measurePair(name string, collectionFn, loFn func(*testing.B), collectionFirst bool) (benchResult, benchResult) {
	collectionSamples := make([]benchmarkMeasurement, 0, benchSamples)
	loSamples := make([]benchmarkMeasurement, 0, benchSamples)
	ratios := make([]float64, 0, benchSamples)
	for sampleIndex := range benchSamples {
		var collectionSample, loSample benchmarkMeasurement
		if (sampleIndex%2 == 0) == collectionFirst {
			collectionSample = measureOnce(collectionFn)
			loSample = measureOnce(loFn)
		} else {
			loSample = measureOnce(loFn)
			collectionSample = measureOnce(collectionFn)
		}
		collectionSamples = append(collectionSamples, collectionSample)
		loSamples = append(loSamples, loSample)
		ratios = append(ratios, loSample.nsPerOp/collectionSample.nsPerOp)
	}

	uncertain := ratioIsUncertain(ratios)
	return summarizeMeasurements(name, "collection", collectionSamples, uncertain),
		summarizeMeasurements(name, "lo", loSamples, uncertain)
}

// benchmarkMeasurement holds one timing and allocation sample.
type benchmarkMeasurement struct {
	nsPerOp     float64
	bytesPerOp  int64
	allocsPerOp int64
}

// measureOnce collects one benchmark sample.
func measureOnce(fn func(*testing.B)) benchmarkMeasurement {
	result := testing.Benchmark(func(b *testing.B) {
		b.ReportAllocs()
		fn(b)
	})
	return benchmarkMeasurement{
		nsPerOp:     float64(result.T) / float64(result.N),
		bytesPerOp:  result.AllocedBytesPerOp(),
		allocsPerOp: result.AllocsPerOp(),
	}
}

// summarizeMeasurements returns median values for one implementation.
func summarizeMeasurements(name, impl string, samples []benchmarkMeasurement, uncertain bool) benchResult {
	nsPerOpSamples := make([]float64, 0, len(samples))
	bytesPerOpSamples := make([]int64, 0, len(samples))
	allocsPerOpSamples := make([]int64, 0, len(samples))
	for _, sample := range samples {
		nsPerOpSamples = append(nsPerOpSamples, sample.nsPerOp)
		bytesPerOpSamples = append(bytesPerOpSamples, sample.bytesPerOp)
		allocsPerOpSamples = append(allocsPerOpSamples, sample.allocsPerOp)
	}
	return benchResult{
		name:        name,
		impl:        impl,
		nsPerOp:     medianFloat64(nsPerOpSamples),
		bytesPerOp:  medianInt64(bytesPerOpSamples),
		allocsPerOp: medianInt64(allocsPerOpSamples),
		uncertain:   uncertain,
	}
}

// ratioIsUncertain reports whether paired samples cross the equivalence boundary or disagree on direction.
func ratioIsUncertain(ratios []float64) bool {
	direction := 0
	for _, ratio := range ratios {
		if ratio >= 1-equivalentEpsilon && ratio <= 1+equivalentEpsilon {
			return true
		}
		currentDirection := 1
		if ratio < 1-equivalentEpsilon {
			currentDirection = -1
		}
		if direction != 0 && currentDirection != direction {
			return true
		}
		direction = currentDirection
	}
	return false
}

// medianFloat64 returns the middle sample after sorting a copy of values.
func medianFloat64(values []float64) float64 {
	sorted := append([]float64(nil), values...)
	sort.Float64s(sorted)
	return sorted[len(sorted)/2]
}

// medianInt64 returns the middle sample after sorting a copy of values.
func medianInt64(values []int64) int64 {
	sorted := append([]int64(nil), values...)
	sort.Slice(sorted, func(i, j int) bool { return sorted[i] < sorted[j] })
	return sorted[len(sorted)/2]
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
	benchRetainInts []int
	unionLeft       []int
	unionRight      []int
	intersectLeft   []int
	intersectRight  []int
	differenceLeft  []int
	differenceRight []int
	workA           []int
	workB           []int
)

// init prepares shared inputs and scratch storage for the benchmark cases.
func init() {
	benchInts = make([]int, benchSize)
	for i := 0; i < benchSize; i++ {
		benchInts[i] = i
	}

	benchIntsDup = make([]int, benchSize)
	for i := 0; i < benchSize; i++ {
		benchIntsDup[i] = i % 128
	}

	benchRetainInts = make([]int, benchSize)
	for i := 0; i < benchSize; i++ {
		benchRetainInts[i] = i % 2
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

// benchIntBoolResult preserves value-and-presence results without type erasure.
type benchIntBoolResult struct {
	value int
	ok    bool
}

// benchIndexWhereResult preserves IndexWhere results without type erasure.
type benchIndexWhereResult struct {
	index int
	found bool
}

// benchPipelineCollectionBorrowHelper executes Pipeline with the CollectionBorrow implementation.
//
//go:noinline
func benchPipelineCollectionBorrowHelper(input []int) int {
	return collection.New(input).Filter(func(v int) bool { return v%2 == 0 }).Map(func(v int) int { return v * v }).Take(benchPipelineLen).Reduce(0, func(acc, v int) int { return acc + v })
}

// benchPipelineCollectionCopyHelper executes Pipeline with the CollectionCopy implementation.
//
//go:noinline
func benchPipelineCollectionCopyHelper(input []int) int {
	return collection.New(input).Clone().Filter(func(v int) bool { return v%2 == 0 }).Map(func(v int) int { return v * v }).Take(benchPipelineLen).Reduce(0, func(acc, v int) int { return acc + v })
}

// benchPipelineLoHelper executes Pipeline with the Lo implementation.
//
//go:noinline
func benchPipelineLoHelper(input []int) int {
	return lo.Reduce(lo.Subset(lo.Map(lo.Filter(input, func(v int, _ int) bool { return v%2 == 0 }), func(v int, _ int) int { return v * v }), 0, benchPipelineLen), func(acc int, v int, _ int) int { return acc + v }, 0)
}

// benchPipelineCollectionBorrow measures Pipeline with the CollectionBorrow implementation.
func benchPipelineCollectionBorrow(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchPipelineCollectionBorrowHelper(benchInts)
		_ = result
	}
}

// benchPipelineCollectionCopy measures Pipeline with the CollectionCopy implementation.
func benchPipelineCollectionCopy(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchPipelineCollectionCopyHelper(benchInts)
		_ = result
	}
}

// benchPipelineLo measures Pipeline with the Lo implementation.
func benchPipelineLo(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchPipelineLoHelper(benchInts)
		_ = result
	}
}

// benchAllCollectionBorrowHelper executes All with the CollectionBorrow implementation.
//
//go:noinline
func benchAllCollectionBorrowHelper(input []int) bool {
	return collection.New(input).All(func(v int) bool { return v < benchSize+1 })
}

// benchAllCollectionCopyHelper executes All with the CollectionCopy implementation.
//
//go:noinline
func benchAllCollectionCopyHelper(input []int) bool {
	return collection.New(input).Clone().All(func(v int) bool { return v < benchSize+1 })
}

// benchAllLoHelper executes All with the Lo implementation.
//
//go:noinline
func benchAllLoHelper(input []int) bool {
	return lo.EveryBy(input, func(v int) bool { return v < benchSize+1 })
}

// benchAllCollectionBorrow measures All with the CollectionBorrow implementation.
func benchAllCollectionBorrow(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchAllCollectionBorrowHelper(benchInts)
		_ = result
	}
}

// benchAllCollectionCopy measures All with the CollectionCopy implementation.
func benchAllCollectionCopy(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchAllCollectionCopyHelper(benchInts)
		_ = result
	}
}

// benchAllLo measures All with the Lo implementation.
func benchAllLo(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchAllLoHelper(benchInts)
		_ = result
	}
}

// benchAnyCollectionBorrowHelper executes Any with the CollectionBorrow implementation.
//
//go:noinline
func benchAnyCollectionBorrowHelper(input []int) bool {
	return collection.New(input).Any(func(v int) bool { return v == benchSize-1 })
}

// benchAnyCollectionCopyHelper executes Any with the CollectionCopy implementation.
//
//go:noinline
func benchAnyCollectionCopyHelper(input []int) bool {
	return collection.New(input).Clone().Any(func(v int) bool { return v == benchSize-1 })
}

// benchAnyLoHelper executes Any with the Lo implementation.
//
//go:noinline
func benchAnyLoHelper(input []int) bool {
	return lo.SomeBy(input, func(v int) bool { return v == benchSize-1 })
}

// benchAnyCollectionBorrow measures Any with the CollectionBorrow implementation.
func benchAnyCollectionBorrow(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchAnyCollectionBorrowHelper(benchInts)
		_ = result
	}
}

// benchAnyCollectionCopy measures Any with the CollectionCopy implementation.
func benchAnyCollectionCopy(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchAnyCollectionCopyHelper(benchInts)
		_ = result
	}
}

// benchAnyLo measures Any with the Lo implementation.
func benchAnyLo(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchAnyLoHelper(benchInts)
		_ = result
	}
}

// benchNoneCollectionBorrowHelper executes None with the CollectionBorrow implementation.
//
//go:noinline
func benchNoneCollectionBorrowHelper(input []int) bool {
	return collection.New(input).None(func(v int) bool { return v < 0 })
}

// benchNoneCollectionCopyHelper executes None with the CollectionCopy implementation.
//
//go:noinline
func benchNoneCollectionCopyHelper(input []int) bool {
	return collection.New(input).Clone().None(func(v int) bool { return v < 0 })
}

// benchNoneLoHelper executes None with the Lo implementation.
//
//go:noinline
func benchNoneLoHelper(input []int) bool {
	return lo.NoneBy(input, func(v int) bool { return v < 0 })
}

// benchNoneCollectionBorrow measures None with the CollectionBorrow implementation.
func benchNoneCollectionBorrow(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchNoneCollectionBorrowHelper(benchInts)
		_ = result
	}
}

// benchNoneCollectionCopy measures None with the CollectionCopy implementation.
func benchNoneCollectionCopy(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchNoneCollectionCopyHelper(benchInts)
		_ = result
	}
}

// benchNoneLo measures None with the Lo implementation.
func benchNoneLo(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchNoneLoHelper(benchInts)
		_ = result
	}
}

// benchFirstCollectionBorrowHelper executes First with the CollectionBorrow implementation.
//
//go:noinline
func benchFirstCollectionBorrowHelper(input []int) benchIntBoolResult {
	return func() benchIntBoolResult {
		value, ok := collection.New(input).First()
		return benchIntBoolResult{value, ok}
	}()
}

// benchFirstCollectionCopyHelper executes First with the CollectionCopy implementation.
//
//go:noinline
func benchFirstCollectionCopyHelper(input []int) benchIntBoolResult {
	return func() benchIntBoolResult {
		value, ok := collection.New(input).Clone().First()
		return benchIntBoolResult{value, ok}
	}()
}

// benchFirstLoHelper executes First with the Lo implementation.
//
//go:noinline
func benchFirstLoHelper(input []int) benchIntBoolResult {
	return func() benchIntBoolResult { value, ok := lo.First(input); return benchIntBoolResult{value, ok} }()
}

// benchFirstCollectionBorrow measures First with the CollectionBorrow implementation.
func benchFirstCollectionBorrow(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchFirstCollectionBorrowHelper(benchInts)
		_ = result
	}
}

// benchFirstCollectionCopy measures First with the CollectionCopy implementation.
func benchFirstCollectionCopy(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchFirstCollectionCopyHelper(benchInts)
		_ = result
	}
}

// benchFirstLo measures First with the Lo implementation.
func benchFirstLo(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchFirstLoHelper(benchInts)
		_ = result
	}
}

// benchLastCollectionBorrowHelper executes Last with the CollectionBorrow implementation.
//
//go:noinline
func benchLastCollectionBorrowHelper(input []int) benchIntBoolResult {
	return func() benchIntBoolResult {
		value, ok := collection.New(input).Last()
		return benchIntBoolResult{value, ok}
	}()
}

// benchLastCollectionCopyHelper executes Last with the CollectionCopy implementation.
//
//go:noinline
func benchLastCollectionCopyHelper(input []int) benchIntBoolResult {
	return func() benchIntBoolResult {
		value, ok := collection.New(input).Clone().Last()
		return benchIntBoolResult{value, ok}
	}()
}

// benchLastLoHelper executes Last with the Lo implementation.
//
//go:noinline
func benchLastLoHelper(input []int) benchIntBoolResult {
	return func() benchIntBoolResult { value, ok := lo.Last(input); return benchIntBoolResult{value, ok} }()
}

// benchLastCollectionBorrow measures Last with the CollectionBorrow implementation.
func benchLastCollectionBorrow(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchLastCollectionBorrowHelper(benchInts)
		_ = result
	}
}

// benchLastCollectionCopy measures Last with the CollectionCopy implementation.
func benchLastCollectionCopy(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchLastCollectionCopyHelper(benchInts)
		_ = result
	}
}

// benchLastLo measures Last with the Lo implementation.
func benchLastLo(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchLastLoHelper(benchInts)
		_ = result
	}
}

// benchIndexWhereCollectionBorrowHelper executes IndexWhere with the CollectionBorrow implementation.
//
//go:noinline
func benchIndexWhereCollectionBorrowHelper(input []int) benchIndexWhereResult {
	return func() benchIndexWhereResult {
		index, found := collection.New(input).IndexWhere(func(v int) bool { return v == benchSize-1 })
		return benchIndexWhereResult{index, found}
	}()
}

// benchIndexWhereCollectionCopyHelper executes IndexWhere with the CollectionCopy implementation.
//
//go:noinline
func benchIndexWhereCollectionCopyHelper(input []int) benchIndexWhereResult {
	return func() benchIndexWhereResult {
		index, found := collection.New(input).Clone().IndexWhere(func(v int) bool { return v == benchSize-1 })
		return benchIndexWhereResult{index, found}
	}()
}

// benchIndexWhereLoHelper executes IndexWhere with the Lo implementation.
//
//go:noinline
func benchIndexWhereLoHelper(input []int) benchIndexWhereResult {
	return func() benchIndexWhereResult {
		index, _, found := lo.FindIndexOf(input, func(v int) bool { return v == benchSize-1 })
		return benchIndexWhereResult{index, found}
	}()
}

// benchIndexWhereCollectionBorrow measures IndexWhere with the CollectionBorrow implementation.
func benchIndexWhereCollectionBorrow(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchIndexWhereCollectionBorrowHelper(benchInts)
		_ = result
	}
}

// benchIndexWhereCollectionCopy measures IndexWhere with the CollectionCopy implementation.
func benchIndexWhereCollectionCopy(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchIndexWhereCollectionCopyHelper(benchInts)
		_ = result
	}
}

// benchIndexWhereLo measures IndexWhere with the Lo implementation.
func benchIndexWhereLo(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchIndexWhereLoHelper(benchInts)
		_ = result
	}
}

// benchEachCollectionBorrowHelper executes Each with the CollectionBorrow implementation.
//
//go:noinline
func benchEachCollectionBorrowHelper(input []int) int {
	return func() int { sum := 0; collection.New(input).Each(func(v int) { sum += v }); return sum }()
}

// benchEachCollectionCopyHelper executes Each with the CollectionCopy implementation.
//
//go:noinline
func benchEachCollectionCopyHelper(input []int) int {
	return func() int { sum := 0; collection.New(input).Clone().Each(func(v int) { sum += v }); return sum }()
}

// benchEachLoHelper executes Each with the Lo implementation.
//
//go:noinline
func benchEachLoHelper(input []int) int {
	return func() int { sum := 0; lo.ForEach(input, func(v int, _ int) { sum += v }); return sum }()
}

// benchEachCollectionBorrow measures Each with the CollectionBorrow implementation.
func benchEachCollectionBorrow(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchEachCollectionBorrowHelper(benchInts)
		_ = result
	}
}

// benchEachCollectionCopy measures Each with the CollectionCopy implementation.
func benchEachCollectionCopy(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchEachCollectionCopyHelper(benchInts)
		_ = result
	}
}

// benchEachLo measures Each with the Lo implementation.
func benchEachLo(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchEachLoHelper(benchInts)
		_ = result
	}
}

// benchMapCollectionBorrowHelper executes Map with the CollectionBorrow implementation.
//
//go:noinline
func benchMapCollectionBorrowHelper(input []int) collection.Slice[int] {
	return collection.New(input).Map(func(v int) int { return v * 3 })
}

// benchMapCollectionCopyHelper executes Map with the CollectionCopy implementation.
//
//go:noinline
func benchMapCollectionCopyHelper(input []int) collection.Slice[int] {
	return collection.New(input).Clone().Map(func(v int) int { return v * 3 })
}

// benchMapLoHelper executes Map with the Lo implementation.
//
//go:noinline
func benchMapLoHelper(input []int) []int {
	return lo.Map(input, func(v int, _ int) int { return v * 3 })
}

// benchMapCollectionBorrow measures Map with the CollectionBorrow implementation.
func benchMapCollectionBorrow(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchMapCollectionBorrowHelper(benchInts)
		_ = result
	}
}

// benchMapCollectionCopy measures Map with the CollectionCopy implementation.
func benchMapCollectionCopy(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchMapCollectionCopyHelper(benchInts)
		_ = result
	}
}

// benchMapLo measures Map with the Lo implementation.
func benchMapLo(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchMapLoHelper(benchInts)
		_ = result
	}
}

// benchTransformCollectionBorrowHelper executes Transform with the CollectionBorrow implementation.
//
//go:noinline
func benchTransformCollectionBorrowHelper(input []int) collection.Slice[int] {
	return collection.New(input).Transform(func(v int) int { return v * 3 })
}

// benchTransformCollectionCopyHelper executes Transform with the CollectionCopy implementation.
//
//go:noinline
func benchTransformCollectionCopyHelper(input []int) collection.Slice[int] {
	return collection.New(input).Clone().Transform(func(v int) int { return v * 3 })
}

// benchTransformLoHelper executes Transform with lo's mutable implementation.
//
//go:noinline
func benchTransformLoHelper(input []int) []int {
	mutable.Map(input, func(v int) int { return v * 3 })
	return input
}

// benchTransformCollectionBorrow measures Transform with the CollectionBorrow implementation.
func benchTransformCollectionBorrow(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		copy(workA, benchInts)
		result := benchTransformCollectionBorrowHelper(workA)
		_ = result
	}
}

// benchTransformCollectionCopy measures Transform with the CollectionCopy implementation.
func benchTransformCollectionCopy(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchTransformCollectionCopyHelper(benchInts)
		_ = result
	}
}

// benchTransformLo measures Transform with lo's mutable implementation.
func benchTransformLo(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		copy(workB, benchInts)
		result := benchTransformLoHelper(workB)
		_ = result
	}
}

// benchReduceCollectionBorrowHelper executes Reduce with the CollectionBorrow implementation.
//
//go:noinline
func benchReduceCollectionBorrowHelper(input []int) int {
	return collection.New(input).Reduce(0, func(acc, v int) int { return acc + v })
}

// benchReduceCollectionCopyHelper executes Reduce with the CollectionCopy implementation.
//
//go:noinline
func benchReduceCollectionCopyHelper(input []int) int {
	return collection.New(input).Clone().Reduce(0, func(acc, v int) int { return acc + v })
}

// benchReduceLoHelper executes Reduce with the Lo implementation.
//
//go:noinline
func benchReduceLoHelper(input []int) int {
	return lo.Reduce(input, func(acc int, v int, _ int) int { return acc + v }, 0)
}

// benchReduceCollectionBorrow measures Reduce with the CollectionBorrow implementation.
func benchReduceCollectionBorrow(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchReduceCollectionBorrowHelper(benchInts)
		_ = result
	}
}

// benchReduceCollectionCopy measures Reduce with the CollectionCopy implementation.
func benchReduceCollectionCopy(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchReduceCollectionCopyHelper(benchInts)
		_ = result
	}
}

// benchReduceLo measures Reduce with the Lo implementation.
func benchReduceLo(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchReduceLoHelper(benchInts)
		_ = result
	}
}

// benchFilterCollectionBorrowHelper executes Filter with the CollectionBorrow implementation.
//
//go:noinline
func benchFilterCollectionBorrowHelper(input []int) collection.Slice[int] {
	return collection.New(input).Filter(func(v int) bool { return v%3 == 0 })
}

// benchFilterCollectionCopyHelper executes Filter with the CollectionCopy implementation.
//
//go:noinline
func benchFilterCollectionCopyHelper(input []int) collection.Slice[int] {
	return collection.New(input).Clone().Filter(func(v int) bool { return v%3 == 0 })
}

// benchFilterLoHelper executes Filter with the Lo implementation.
//
//go:noinline
func benchFilterLoHelper(input []int) []int {
	return lo.Filter(input, func(v int, _ int) bool { return v%3 == 0 })
}

// benchFilterCollectionBorrow measures Filter with the CollectionBorrow implementation.
func benchFilterCollectionBorrow(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchFilterCollectionBorrowHelper(benchInts)
		_ = result
	}
}

// benchFilterCollectionCopy measures Filter with the CollectionCopy implementation.
func benchFilterCollectionCopy(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchFilterCollectionCopyHelper(benchInts)
		_ = result
	}
}

// benchFilterLo measures Filter with the Lo implementation.
func benchFilterLo(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchFilterLoHelper(benchInts)
		_ = result
	}
}

// benchRetainCollectionBorrowHelper executes Retain with the CollectionBorrow implementation.
//
//go:noinline
func benchRetainCollectionBorrowHelper(input []int) collection.Slice[int] {
	return collection.New(input).Retain(func(v int) bool { return v != 0 })
}

// benchRetainCollectionCopyHelper executes Retain with the CollectionCopy implementation.
//
//go:noinline
func benchRetainCollectionCopyHelper(input []int) collection.Slice[int] {
	return collection.New(input).Clone().Retain(func(v int) bool { return v != 0 })
}

// benchRetainLoHelper executes Retain with lo's mutable implementation.
//
//go:noinline
func benchRetainLoHelper(input []int) []int {
	result := mutable.Filter(input, func(v int) bool { return v != 0 })
	clear(input[len(result):])
	return result
}

// benchRetainCollectionBorrow measures Retain with the CollectionBorrow implementation.
func benchRetainCollectionBorrow(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		copy(workA, benchRetainInts)
		result := benchRetainCollectionBorrowHelper(workA)
		_ = result
	}
}

// benchRetainCollectionCopy measures Retain with the CollectionCopy implementation.
func benchRetainCollectionCopy(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchRetainCollectionCopyHelper(benchRetainInts)
		_ = result
	}
}

// benchRetainLo measures Retain with lo's mutable implementation.
func benchRetainLo(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		copy(workB, benchRetainInts)
		result := benchRetainLoHelper(workB)
		_ = result
	}
}

// benchChunkCollectionBorrowHelper executes Chunk with the CollectionBorrow implementation.
//
//go:noinline
func benchChunkCollectionBorrowHelper(input []int) [][]int {
	return collection.New(input).Chunk(benchChunkSize)
}

// benchChunkCollectionCopyHelper executes Chunk with the CollectionCopy implementation.
//
//go:noinline
func benchChunkCollectionCopyHelper(input []int) [][]int {
	return collection.New(input).Clone().Chunk(benchChunkSize)
}

// benchChunkLoHelper executes Chunk with the Lo implementation.
//
//go:noinline
func benchChunkLoHelper(input []int) [][]int {
	return lo.Chunk(input, benchChunkSize)
}

// benchChunkCollectionBorrow measures Chunk with the CollectionBorrow implementation.
func benchChunkCollectionBorrow(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchChunkCollectionBorrowHelper(benchInts)
		_ = result
	}
}

// benchChunkCollectionCopy measures Chunk with the CollectionCopy implementation.
func benchChunkCollectionCopy(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchChunkCollectionCopyHelper(benchInts)
		_ = result
	}
}

// benchChunkLo measures Chunk with the Lo implementation.
func benchChunkLo(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchChunkLoHelper(benchInts)
		_ = result
	}
}

// benchTakeCollectionBorrowHelper executes Take with the CollectionBorrow implementation.
//
//go:noinline
func benchTakeCollectionBorrowHelper(input []int) collection.Slice[int] {
	return collection.New(input).Take(benchTakeN)
}

// benchTakeCollectionCopyHelper executes Take with the CollectionCopy implementation.
//
//go:noinline
func benchTakeCollectionCopyHelper(input []int) collection.Slice[int] {
	return collection.New(input).Clone().Take(benchTakeN)
}

// benchTakeLoHelper executes Take with the Lo implementation.
//
//go:noinline
func benchTakeLoHelper(input []int) []int {
	return lo.Subset(input, 0, uint(benchTakeN))
}

// benchTakeCollectionBorrow measures Take with the CollectionBorrow implementation.
func benchTakeCollectionBorrow(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchTakeCollectionBorrowHelper(benchInts)
		_ = result
	}
}

// benchTakeCollectionCopy measures Take with the CollectionCopy implementation.
func benchTakeCollectionCopy(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchTakeCollectionCopyHelper(benchInts)
		_ = result
	}
}

// benchTakeLo measures Take with the Lo implementation.
func benchTakeLo(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchTakeLoHelper(benchInts)
		_ = result
	}
}

// benchContainsCollectionBorrowHelper executes Contains with the CollectionBorrow implementation.
//
//go:noinline
func benchContainsCollectionBorrowHelper(input []int) bool {
	return slices.Contains(collection.New(input), benchSize-1)
}

// benchContainsCollectionCopyHelper executes Contains with the CollectionCopy implementation.
//
//go:noinline
func benchContainsCollectionCopyHelper(input []int) bool {
	return slices.Contains(collection.New(input).Clone(), benchSize-1)
}

// benchContainsLoHelper executes Contains with the Lo implementation.
//
//go:noinline
func benchContainsLoHelper(input []int) bool {
	return lo.Contains(input, benchSize-1)
}

// benchContainsCollectionBorrow measures Contains with the CollectionBorrow implementation.
func benchContainsCollectionBorrow(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchContainsCollectionBorrowHelper(benchInts)
		_ = result
	}
}

// benchContainsCollectionCopy measures Contains with the CollectionCopy implementation.
func benchContainsCollectionCopy(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchContainsCollectionCopyHelper(benchInts)
		_ = result
	}
}

// benchContainsLo measures Contains with the Lo implementation.
func benchContainsLo(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchContainsLoHelper(benchInts)
		_ = result
	}
}

// benchFindCollectionBorrowHelper executes Find with the CollectionBorrow implementation.
//
//go:noinline
func benchFindCollectionBorrowHelper(input []int) benchIntBoolResult {
	return func() benchIntBoolResult {
		value, ok := collection.New(input).FirstWhere(func(v int) bool { return v == benchSize-1 })
		return benchIntBoolResult{value, ok}
	}()
}

// benchFindCollectionCopyHelper executes Find with the CollectionCopy implementation.
//
//go:noinline
func benchFindCollectionCopyHelper(input []int) benchIntBoolResult {
	return func() benchIntBoolResult {
		value, ok := collection.New(input).Clone().FirstWhere(func(v int) bool { return v == benchSize-1 })
		return benchIntBoolResult{value, ok}
	}()
}

// benchFindLoHelper executes Find with the Lo implementation.
//
//go:noinline
func benchFindLoHelper(input []int) benchIntBoolResult {
	return func() benchIntBoolResult {
		value, ok := lo.Find(input, func(v int) bool { return v == benchSize-1 })
		return benchIntBoolResult{value, ok}
	}()
}

// benchFindCollectionBorrow measures Find with the CollectionBorrow implementation.
func benchFindCollectionBorrow(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchFindCollectionBorrowHelper(benchInts)
		_ = result
	}
}

// benchFindCollectionCopy measures Find with the CollectionCopy implementation.
func benchFindCollectionCopy(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchFindCollectionCopyHelper(benchInts)
		_ = result
	}
}

// benchFindLo measures Find with the Lo implementation.
func benchFindLo(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchFindLoHelper(benchInts)
		_ = result
	}
}

// benchGroupByCollectionBorrowHelper executes GroupBy with the CollectionBorrow implementation.
//
//go:noinline
func benchGroupByCollectionBorrowHelper(input []int) map[int][]int {
	return collection.New(input).GroupBy(func(v int) int { return v % benchGroupByMod })
}

// benchGroupByCollectionCopyHelper executes GroupBy with the CollectionCopy implementation.
//
//go:noinline
func benchGroupByCollectionCopyHelper(input []int) map[int][]int {
	return collection.New(input).Clone().GroupBy(func(v int) int { return v % benchGroupByMod })
}

// benchGroupByLoHelper executes GroupBy with the Lo implementation.
//
//go:noinline
func benchGroupByLoHelper(input []int) map[int][]int {
	return lo.GroupBy(input, func(v int) int { return v % benchGroupByMod })
}

// benchGroupByCollectionBorrow measures GroupBy with the CollectionBorrow implementation.
func benchGroupByCollectionBorrow(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchGroupByCollectionBorrowHelper(benchInts)
		_ = result
	}
}

// benchGroupByCollectionCopy measures GroupBy with the CollectionCopy implementation.
func benchGroupByCollectionCopy(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchGroupByCollectionCopyHelper(benchInts)
		_ = result
	}
}

// benchGroupByLo measures GroupBy with the Lo implementation.
func benchGroupByLo(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchGroupByLoHelper(benchInts)
		_ = result
	}
}

// benchCountByCollectionBorrowHelper executes CountBy with the CollectionBorrow implementation.
//
//go:noinline
func benchCountByCollectionBorrowHelper(input []int) map[int]int {
	return collection.New(input).CountBy(func(v int) int { return v })
}

// benchCountByCollectionCopyHelper executes CountBy with the CollectionCopy implementation.
//
//go:noinline
func benchCountByCollectionCopyHelper(input []int) map[int]int {
	return collection.New(input).Clone().CountBy(func(v int) int { return v })
}

// benchCountByLoHelper executes CountBy with the Lo implementation.
//
//go:noinline
func benchCountByLoHelper(input []int) map[int]int {
	return lo.CountValuesBy(input, func(v int) int { return v })
}

// benchCountByCollectionBorrow measures CountBy with the CollectionBorrow implementation.
func benchCountByCollectionBorrow(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchCountByCollectionBorrowHelper(benchIntsDup)
		_ = result
	}
}

// benchCountByCollectionCopy measures CountBy with the CollectionCopy implementation.
func benchCountByCollectionCopy(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchCountByCollectionCopyHelper(benchIntsDup)
		_ = result
	}
}

// benchCountByLo measures CountBy with the Lo implementation.
func benchCountByLo(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchCountByLoHelper(benchIntsDup)
		_ = result
	}
}

// benchCountByValueCollectionBorrowHelper executes CountByValue with the CollectionBorrow implementation.
//
//go:noinline
func benchCountByValueCollectionBorrowHelper(input []int) map[int]int {
	return collection.CountByValue(collection.New(input))
}

// benchCountByValueCollectionCopyHelper executes CountByValue with the CollectionCopy implementation.
//
//go:noinline
func benchCountByValueCollectionCopyHelper(input []int) map[int]int {
	return collection.CountByValue(collection.New(input).Clone())
}

// benchCountByValueLoHelper executes CountByValue with the Lo implementation.
//
//go:noinline
func benchCountByValueLoHelper(input []int) map[int]int {
	return lo.CountValues(input)
}

// benchCountByValueCollectionBorrow measures CountByValue with the CollectionBorrow implementation.
func benchCountByValueCollectionBorrow(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchCountByValueCollectionBorrowHelper(benchIntsDup)
		_ = result
	}
}

// benchCountByValueCollectionCopy measures CountByValue with the CollectionCopy implementation.
func benchCountByValueCollectionCopy(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchCountByValueCollectionCopyHelper(benchIntsDup)
		_ = result
	}
}

// benchCountByValueLo measures CountByValue with the Lo implementation.
func benchCountByValueLo(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchCountByValueLoHelper(benchIntsDup)
		_ = result
	}
}

// benchSkipCollectionBorrowHelper executes Skip with the CollectionBorrow implementation.
//
//go:noinline
func benchSkipCollectionBorrowHelper(input []int) collection.Slice[int] {
	return collection.New(input).Skip(benchSkipN)
}

// benchSkipCollectionCopyHelper executes Skip with the CollectionCopy implementation.
//
//go:noinline
func benchSkipCollectionCopyHelper(input []int) collection.Slice[int] {
	return collection.New(input).Clone().Skip(benchSkipN)
}

// benchSkipLoHelper executes Skip with the Lo implementation.
//
//go:noinline
func benchSkipLoHelper(input []int) []int {
	return lo.Drop(input, benchSkipN)
}

// benchSkipCollectionBorrow measures Skip with the CollectionBorrow implementation.
func benchSkipCollectionBorrow(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchSkipCollectionBorrowHelper(benchInts)
		_ = result
	}
}

// benchSkipCollectionCopy measures Skip with the CollectionCopy implementation.
func benchSkipCollectionCopy(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchSkipCollectionCopyHelper(benchInts)
		_ = result
	}
}

// benchSkipLo measures Skip with the Lo implementation.
func benchSkipLo(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchSkipLoHelper(benchInts)
		_ = result
	}
}

// benchSkipLastCollectionBorrowHelper executes SkipLast with the CollectionBorrow implementation.
//
//go:noinline
func benchSkipLastCollectionBorrowHelper(input []int) collection.Slice[int] {
	return collection.New(input).SkipLast(benchSkipN)
}

// benchSkipLastCollectionCopyHelper executes SkipLast with the CollectionCopy implementation.
//
//go:noinline
func benchSkipLastCollectionCopyHelper(input []int) collection.Slice[int] {
	return collection.New(input).Clone().SkipLast(benchSkipN)
}

// benchSkipLastLoHelper executes SkipLast with the Lo implementation.
//
//go:noinline
func benchSkipLastLoHelper(input []int) []int {
	return lo.DropRight(input, benchSkipN)
}

// benchSkipLastCollectionBorrow measures SkipLast with the CollectionBorrow implementation.
func benchSkipLastCollectionBorrow(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchSkipLastCollectionBorrowHelper(benchInts)
		_ = result
	}
}

// benchSkipLastCollectionCopy measures SkipLast with the CollectionCopy implementation.
func benchSkipLastCollectionCopy(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchSkipLastCollectionCopyHelper(benchInts)
		_ = result
	}
}

// benchSkipLastLo measures SkipLast with the Lo implementation.
func benchSkipLastLo(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchSkipLastLoHelper(benchInts)
		_ = result
	}
}

// benchReverseCollectionBorrowHelper executes Reverse with the CollectionBorrow implementation.
//
//go:noinline
func benchReverseCollectionBorrowHelper(input []int) collection.Slice[int] {
	return collection.New(input).Reverse()
}

// benchReverseCollectionCopyHelper executes Reverse with the CollectionCopy implementation.
//
//go:noinline
func benchReverseCollectionCopyHelper(input []int) collection.Slice[int] {
	return collection.New(input).Clone().Reverse()
}

// benchReverseLoHelper executes Reverse with the Lo implementation.
//
//go:noinline
func benchReverseLoHelper(input []int) []int {
	return lo.Reverse(input)
}

// benchReverseCollectionBorrow measures Reverse with the CollectionBorrow implementation.
func benchReverseCollectionBorrow(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		copy(workA, benchInts)
		result := benchReverseCollectionBorrowHelper(workA)
		_ = result
	}
}

// benchReverseCollectionCopy measures Reverse with the CollectionCopy implementation.
func benchReverseCollectionCopy(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchReverseCollectionCopyHelper(benchInts)
		_ = result
	}
}

// benchReverseLo measures Reverse with the Lo implementation.
func benchReverseLo(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		copy(workB, benchInts)
		result := benchReverseLoHelper(workB)
		_ = result
	}
}

// benchShuffleCollectionBorrowHelper executes Shuffle with the CollectionBorrow implementation.
//
//go:noinline
func benchShuffleCollectionBorrowHelper(input []int) collection.Slice[int] {
	return collection.New(input).Shuffle()
}

// benchShuffleCollectionCopyHelper executes Shuffle with the CollectionCopy implementation.
//
//go:noinline
func benchShuffleCollectionCopyHelper(input []int) collection.Slice[int] {
	return collection.New(input).Clone().Shuffle()
}

// benchShuffleLoHelper executes Shuffle with the Lo implementation.
//
//go:noinline
func benchShuffleLoHelper(input []int) []int {
	return lo.Shuffle(input)
}

// benchShuffleCollectionBorrow measures Shuffle with the CollectionBorrow implementation.
func benchShuffleCollectionBorrow(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		copy(workA, benchInts)
		result := benchShuffleCollectionBorrowHelper(workA)
		_ = result
	}
}

// benchShuffleCollectionCopy measures Shuffle with the CollectionCopy implementation.
func benchShuffleCollectionCopy(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchShuffleCollectionCopyHelper(benchInts)
		_ = result
	}
}

// benchShuffleLo measures Shuffle with the Lo implementation.
func benchShuffleLo(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		copy(workB, benchInts)
		result := benchShuffleLoHelper(workB)
		_ = result
	}
}

// benchZipCollectionBorrowHelper executes Zip with the CollectionBorrow implementation.
//
//go:noinline
func benchZipCollectionBorrowHelper(left, right []int) []collection.Pair[int, int] {
	return collection.New(left).Zip(right)
}

// benchZipCollectionCopyHelper executes Zip with the CollectionCopy implementation.
//
//go:noinline
func benchZipCollectionCopyHelper(left, right []int) []collection.Pair[int, int] {
	return collection.New(left).Clone().Zip(collection.New(right).Clone())
}

// benchZipLoHelper executes Zip with the Lo implementation.
//
//go:noinline
func benchZipLoHelper(left, right []int) []lo.Tuple2[int, int] {
	return lo.Zip2(left, right)
}

// benchZipCollectionBorrow measures Zip with the CollectionBorrow implementation.
func benchZipCollectionBorrow(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchZipCollectionBorrowHelper(benchInts, benchIntsDup)
		_ = result
	}
}

// benchZipCollectionCopy measures Zip with the CollectionCopy implementation.
func benchZipCollectionCopy(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchZipCollectionCopyHelper(benchInts, benchIntsDup)
		_ = result
	}
}

// benchZipLo measures Zip with the Lo implementation.
func benchZipLo(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchZipLoHelper(benchInts, benchIntsDup)
		_ = result
	}
}

// benchZipWithCollectionBorrowHelper executes ZipWith with the CollectionBorrow implementation.
//
//go:noinline
func benchZipWithCollectionBorrowHelper(left, right []int) collection.Slice[int] {
	return collection.New(left).ZipWith(right, func(a, b int) int { return a + b })
}

// benchZipWithCollectionCopyHelper executes ZipWith with the CollectionCopy implementation.
//
//go:noinline
func benchZipWithCollectionCopyHelper(left, right []int) collection.Slice[int] {
	return collection.New(left).Clone().ZipWith(collection.New(right).Clone(), func(a, b int) int { return a + b })
}

// benchZipWithLoHelper executes ZipWith with the Lo implementation.
//
//go:noinline
func benchZipWithLoHelper(left, right []int) []int {
	return lo.ZipBy2(left, right, func(a, b int) int { return a + b })
}

// benchZipWithCollectionBorrow measures ZipWith with the CollectionBorrow implementation.
func benchZipWithCollectionBorrow(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchZipWithCollectionBorrowHelper(benchInts, benchIntsDup)
		_ = result
	}
}

// benchZipWithCollectionCopy measures ZipWith with the CollectionCopy implementation.
func benchZipWithCollectionCopy(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchZipWithCollectionCopyHelper(benchInts, benchIntsDup)
		_ = result
	}
}

// benchZipWithLo measures ZipWith with the Lo implementation.
func benchZipWithLo(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchZipWithLoHelper(benchInts, benchIntsDup)
		_ = result
	}
}

// benchUniqueCollectionBorrowHelper executes Unique with the CollectionBorrow implementation.
//
//go:noinline
func benchUniqueCollectionBorrowHelper(input []int) collection.Slice[int] {
	return collection.UniqueComparable(collection.New(input))
}

// benchUniqueCollectionCopyHelper executes Unique with the CollectionCopy implementation.
//
//go:noinline
func benchUniqueCollectionCopyHelper(input []int) collection.Slice[int] {
	return collection.UniqueComparable(collection.New(input).Clone())
}

// benchUniqueLoHelper executes Unique with the Lo implementation.
//
//go:noinline
func benchUniqueLoHelper(input []int) []int {
	return lo.Uniq(input)
}

// benchUniqueCollectionBorrow measures Unique with the CollectionBorrow implementation.
func benchUniqueCollectionBorrow(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchUniqueCollectionBorrowHelper(benchIntsDup)
		_ = result
	}
}

// benchUniqueCollectionCopy measures Unique with the CollectionCopy implementation.
func benchUniqueCollectionCopy(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchUniqueCollectionCopyHelper(benchIntsDup)
		_ = result
	}
}

// benchUniqueLo measures Unique with the Lo implementation.
func benchUniqueLo(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchUniqueLoHelper(benchIntsDup)
		_ = result
	}
}

// benchUniqueByCollectionBorrowHelper executes UniqueBy with the CollectionBorrow implementation.
//
//go:noinline
func benchUniqueByCollectionBorrowHelper(input []int) collection.Slice[int] {
	return collection.New(input).UniqueBy(func(v int) int { return v })
}

// benchUniqueByCollectionCopyHelper executes UniqueBy with the CollectionCopy implementation.
//
//go:noinline
func benchUniqueByCollectionCopyHelper(input []int) collection.Slice[int] {
	return collection.New(input).Clone().UniqueBy(func(v int) int { return v })
}

// benchUniqueByLoHelper executes UniqueBy with the Lo implementation.
//
//go:noinline
func benchUniqueByLoHelper(input []int) []int {
	return lo.UniqBy(input, func(v int) int { return v })
}

// benchUniqueByCollectionBorrow measures UniqueBy with the CollectionBorrow implementation.
func benchUniqueByCollectionBorrow(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchUniqueByCollectionBorrowHelper(benchIntsDup)
		_ = result
	}
}

// benchUniqueByCollectionCopy measures UniqueBy with the CollectionCopy implementation.
func benchUniqueByCollectionCopy(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchUniqueByCollectionCopyHelper(benchIntsDup)
		_ = result
	}
}

// benchUniqueByLo measures UniqueBy with the Lo implementation.
func benchUniqueByLo(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchUniqueByLoHelper(benchIntsDup)
		_ = result
	}
}

// benchUnionCollectionBorrowHelper executes Union with the CollectionBorrow implementation.
//
//go:noinline
func benchUnionCollectionBorrowHelper(left, right []int) collection.Slice[int] {
	return collection.Union(collection.New(left), collection.New(right))
}

// benchUnionCollectionCopyHelper executes Union with the CollectionCopy implementation.
//
//go:noinline
func benchUnionCollectionCopyHelper(left, right []int) collection.Slice[int] {
	return collection.Union(collection.New(left).Clone(), collection.New(right).Clone())
}

// benchUnionLoHelper executes Union with the Lo implementation.
//
//go:noinline
func benchUnionLoHelper(left, right []int) []int {
	return lo.Union(left, right)
}

// benchUnionCollectionBorrow measures Union with the CollectionBorrow implementation.
func benchUnionCollectionBorrow(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchUnionCollectionBorrowHelper(unionLeft, unionRight)
		_ = result
	}
}

// benchUnionCollectionCopy measures Union with the CollectionCopy implementation.
func benchUnionCollectionCopy(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchUnionCollectionCopyHelper(unionLeft, unionRight)
		_ = result
	}
}

// benchUnionLo measures Union with the Lo implementation.
func benchUnionLo(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchUnionLoHelper(unionLeft, unionRight)
		_ = result
	}
}

// benchIntersectCollectionBorrowHelper executes Intersect with the CollectionBorrow implementation.
//
//go:noinline
func benchIntersectCollectionBorrowHelper(left, right []int) collection.Slice[int] {
	return collection.Intersect(collection.New(left), collection.New(right))
}

// benchIntersectCollectionCopyHelper executes Intersect with the CollectionCopy implementation.
//
//go:noinline
func benchIntersectCollectionCopyHelper(left, right []int) collection.Slice[int] {
	return collection.Intersect(collection.New(left).Clone(), collection.New(right).Clone())
}

// benchIntersectLoHelper executes Intersect with the Lo implementation.
//
//go:noinline
func benchIntersectLoHelper(left, right []int) []int {
	return lo.Intersect(left, right)
}

// benchIntersectCollectionBorrow measures Intersect with the CollectionBorrow implementation.
func benchIntersectCollectionBorrow(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchIntersectCollectionBorrowHelper(intersectLeft, intersectRight)
		_ = result
	}
}

// benchIntersectCollectionCopy measures Intersect with the CollectionCopy implementation.
func benchIntersectCollectionCopy(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchIntersectCollectionCopyHelper(intersectLeft, intersectRight)
		_ = result
	}
}

// benchIntersectLo measures Intersect with the Lo implementation.
func benchIntersectLo(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchIntersectLoHelper(intersectLeft, intersectRight)
		_ = result
	}
}

// benchDifferenceCollectionBorrowHelper executes Difference with the CollectionBorrow implementation.
//
//go:noinline
func benchDifferenceCollectionBorrowHelper(left, right []int) collection.Slice[int] {
	return collection.Difference(collection.New(left), collection.New(right))
}

// benchDifferenceCollectionCopyHelper executes Difference with the CollectionCopy implementation.
//
//go:noinline
func benchDifferenceCollectionCopyHelper(left, right []int) collection.Slice[int] {
	return collection.Difference(collection.New(left).Clone(), collection.New(right).Clone())
}

// benchDifferenceLoHelper executes Difference with the Lo implementation.
//
//go:noinline
func benchDifferenceLoHelper(left, right []int) []int {
	return func() []int { result, _ := lo.Difference(left, right); return result }()
}

// benchDifferenceCollectionBorrow measures Difference with the CollectionBorrow implementation.
func benchDifferenceCollectionBorrow(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchDifferenceCollectionBorrowHelper(differenceLeft, differenceRight)
		_ = result
	}
}

// benchDifferenceCollectionCopy measures Difference with the CollectionCopy implementation.
func benchDifferenceCollectionCopy(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchDifferenceCollectionCopyHelper(differenceLeft, differenceRight)
		_ = result
	}
}

// benchDifferenceLo measures Difference with the Lo implementation.
func benchDifferenceLo(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchDifferenceLoHelper(differenceLeft, differenceRight)
		_ = result
	}
}

// benchToMapCollectionBorrowHelper executes ToMap with the CollectionBorrow implementation.
//
//go:noinline
func benchToMapCollectionBorrowHelper(input []int) map[int]int {
	return collection.New(input).ToMap(func(v int) int { return v }, func(v int) int { return v })
}

// benchToMapCollectionCopyHelper executes ToMap with the CollectionCopy implementation.
//
//go:noinline
func benchToMapCollectionCopyHelper(input []int) map[int]int {
	return collection.New(input).Clone().ToMap(func(v int) int { return v }, func(v int) int { return v })
}

// benchToMapLoHelper executes ToMap with the Lo implementation.
//
//go:noinline
func benchToMapLoHelper(input []int) map[int]int {
	return lo.SliceToMap(input, func(v int) (int, int) { return v, v })
}

// benchToMapCollectionBorrow measures ToMap with the CollectionBorrow implementation.
func benchToMapCollectionBorrow(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchToMapCollectionBorrowHelper(benchInts)
		_ = result
	}
}

// benchToMapCollectionCopy measures ToMap with the CollectionCopy implementation.
func benchToMapCollectionCopy(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchToMapCollectionCopyHelper(benchInts)
		_ = result
	}
}

// benchToMapLo measures ToMap with the Lo implementation.
func benchToMapLo(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchToMapLoHelper(benchInts)
		_ = result
	}
}

// benchSumCollectionBorrowHelper executes Sum with the CollectionBorrow implementation.
//
//go:noinline
func benchSumCollectionBorrowHelper(input []int) int {
	return collection.Sum(input)
}

// benchSumCollectionCopyHelper executes Sum with the CollectionCopy implementation.
//
//go:noinline
func benchSumCollectionCopyHelper(input []int) int {
	return collection.Sum(append([]int(nil), input...))
}

// benchSumLoHelper executes Sum with the Lo implementation.
//
//go:noinline
func benchSumLoHelper(input []int) int {
	return lo.Sum(input)
}

// benchSumCollectionBorrow measures Sum with the CollectionBorrow implementation.
func benchSumCollectionBorrow(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchSumCollectionBorrowHelper(benchInts)
		_ = result
	}
}

// benchSumCollectionCopy measures Sum with the CollectionCopy implementation.
func benchSumCollectionCopy(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchSumCollectionCopyHelper(benchInts)
		_ = result
	}
}

// benchSumLo measures Sum with the Lo implementation.
func benchSumLo(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchSumLoHelper(benchInts)
		_ = result
	}
}

// benchMinCollectionBorrowHelper executes Min with the CollectionBorrow implementation.
//
//go:noinline
func benchMinCollectionBorrowHelper(input []int) benchIntBoolResult {
	return func() benchIntBoolResult { value, ok := collection.Min(input); return benchIntBoolResult{value, ok} }()
}

// benchMinCollectionCopyHelper executes Min with the CollectionCopy implementation.
//
//go:noinline
func benchMinCollectionCopyHelper(input []int) benchIntBoolResult {
	return func() benchIntBoolResult {
		value, ok := collection.Min(append([]int(nil), input...))
		return benchIntBoolResult{value, ok}
	}()
}

// benchMinLoHelper executes Min with the Lo implementation.
//
//go:noinline
func benchMinLoHelper(input []int) benchIntBoolResult {
	return benchIntBoolResult{value: lo.Min(input), ok: true}
}

// benchMinCollectionBorrow measures Min with the CollectionBorrow implementation.
func benchMinCollectionBorrow(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchMinCollectionBorrowHelper(benchInts)
		_ = result
	}
}

// benchMinCollectionCopy measures Min with the CollectionCopy implementation.
func benchMinCollectionCopy(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchMinCollectionCopyHelper(benchInts)
		_ = result
	}
}

// benchMinLo measures Min with the Lo implementation.
func benchMinLo(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchMinLoHelper(benchInts)
		_ = result
	}
}

// benchMaxCollectionBorrowHelper executes Max with the CollectionBorrow implementation.
//
//go:noinline
func benchMaxCollectionBorrowHelper(input []int) benchIntBoolResult {
	return func() benchIntBoolResult { value, ok := collection.Max(input); return benchIntBoolResult{value, ok} }()
}

// benchMaxCollectionCopyHelper executes Max with the CollectionCopy implementation.
//
//go:noinline
func benchMaxCollectionCopyHelper(input []int) benchIntBoolResult {
	return func() benchIntBoolResult {
		value, ok := collection.Max(append([]int(nil), input...))
		return benchIntBoolResult{value, ok}
	}()
}

// benchMaxLoHelper executes Max with the Lo implementation.
//
//go:noinline
func benchMaxLoHelper(input []int) benchIntBoolResult {
	return benchIntBoolResult{value: lo.Max(input), ok: true}
}

// benchMaxCollectionBorrow measures Max with the CollectionBorrow implementation.
func benchMaxCollectionBorrow(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchMaxCollectionBorrowHelper(benchInts)
		_ = result
	}
}

// benchMaxCollectionCopy measures Max with the CollectionCopy implementation.
func benchMaxCollectionCopy(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchMaxCollectionCopyHelper(benchInts)
		_ = result
	}
}

// benchMaxLo measures Max with the Lo implementation.
func benchMaxLo(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		result := benchMaxLoHelper(benchInts)
		_ = result
	}
}

// ----------------------------------------------------------------------------
// Rendering
// ----------------------------------------------------------------------------

// renderTable formats detailed results for one ownership mode.
func renderTable(results []benchResult, mode benchMode) string {
	byName := map[string]map[string]benchResult{}
	for _, r := range results {
		if _, ok := byName[r.name]; !ok {
			byName[r.name] = map[string]benchResult{}
		}
		byName[r.name][r.impl] = r
	}

	var buf bytes.Buffer
	buf.WriteString("| Op | ns/op (vs lo) | Timing | bytes/op (vs lo) | × (less memory) | allocs/op (vs lo) |\n")
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
		ratioCell := formatBenchmarkRatio(name, mode, loRes.nsPerOp, col.nsPerOp, col.uncertain || loRes.uncertain)

		bytesCell := fmt.Sprintf(
			"%s / %s",
			formatBytes(col.bytesPerOp),
			formatBytes(loRes.bytesPerOp),
		)
		bytesRatioCell := formatRatioBytes(loRes.bytesPerOp, col.bytesPerOp)
		allocCell := fmt.Sprintf("%d / %d", col.allocsPerOp, loRes.allocsPerOp)
		if isDifferentWorkBenchmark(name) {
			nsCell, ratioCell, bytesCell, bytesRatioCell, allocCell = "different work", "API trade-off", "different work", "API trade-off", "API trade-off"
		} else if mode == benchBorrow && isViewBenchmark(name) {
			bytesRatioCell, allocCell = "ownership trade-off", "ownership trade-off"
		}

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
	if mode == benchBorrow {
		buf.WriteString("\nChunk, Skip, and SkipLast return collection views while lo returns copied slices. Their rows describe ownership and allocation trade-offs, not equal-work speed superiority. Difference returns one-sided output while lo returns both sides, so its rows are an API trade-off.\n")
	}

	return strings.TrimSpace(buf.String())
}

type benchGroup struct {
	name string
	ops  []string
}

// renderCondensedTables formats the summary tables embedded in the README.
func renderCondensedTables(results []benchResult, mode benchMode) string {
	groups := []benchGroup{
		{
			name: "Read-only scalar ops",
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
				"Chunk",
				"Filter",
				"Map",
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
				"GroupBy",
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
				"Retain",
				"Reverse",
				"Shuffle",
				"Transform",
			},
		},
	}

	var buf bytes.Buffer
	buf.WriteString("Full raw tables: see `BENCHMARKS.md`.\n\n")
	byName := map[string]map[string]benchResult{}
	for _, r := range results {
		if _, ok := byName[r.name]; !ok {
			byName[r.name] = map[string]benchResult{}
		}
		byName[r.name][r.impl] = r
	}

	for _, group := range groups {
		rows := make([]string, 0, len(group.ops))
		for _, name := range group.ops {
			entry, ok := byName[name]
			if !ok {
				continue
			}
			col := entry["collection"]
			loRes := entry["lo"]

			scalarOnly := group.name == "Read-only scalar ops"
			allowBold := group.name == "Pipelines" || group.name == "Transforming ops" || group.name == "Mutating ops"
			speed := formatBenchmarkSpeed(name, mode, loRes.nsPerOp, col.nsPerOp, col.uncertain || loRes.uncertain, allowBold, scalarOnly)
			mem := formatDeltaBytes(col.bytesPerOp, loRes.bytesPerOp)
			allocs := formatDeltaAllocs(col.allocsPerOp, loRes.allocsPerOp)
			if isDifferentWorkBenchmark(name) {
				speed, mem, allocs = "different work", "API trade-off", "API trade-off"
			} else if mode == benchBorrow && isViewBenchmark(name) {
				mem, allocs = "ownership trade-off", "ownership trade-off"
			}

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

// formatNs renders nanoseconds per operation at a readable precision.
func formatNs(ns float64) string {
	switch {
	case ns < 1:
		return "<1ns"
	case ns >= 1e6:
		return fmt.Sprintf("%.1fms", ns/1e6)
	case ns >= 1e3:
		return fmt.Sprintf("%.1fµs", ns/1e3)
	default:
		return fmt.Sprintf("%.1fns", ns)
	}
}

// formatBytes renders a byte count with binary units.
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

// formatDurationNs renders a nanosecond duration with an appropriate unit.
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
	equivalentEpsilon = 0.10 // ±10% equivalence tolerance
	benchRatioNoiseNs = 50.0
	scalarOnlyEpsilon = 0.15
)

// formatRatio compares lo timing with collection timing.
func formatRatio(lo, col float64) string {
	if lo < benchRatioNoiseNs && col < benchRatioNoiseNs {
		return "≈"
	}
	if col == 0 {
		return "∞"
	}

	ratio := lo / col

	// Treat small deltas as equivalent measurement noise.
	if ratio >= 1-equivalentEpsilon && ratio <= 1+equivalentEpsilon {
		return "≈"
	}

	if ratio > 1 {
		return "**faster**"
	}
	return "slower"
}

// formatBenchmarkRatio avoids presenting view creation as a faster equivalent copy.
// formatBenchmarkRatio applies benchmark-specific timing comparison rules.
func formatBenchmarkRatio(name string, mode benchMode, lo, col float64, uncertain bool) string {
	if mode == benchBorrow && isViewBenchmark(name) {
		return "view trade-off"
	}
	if mode == benchBorrow && isCodeEquivalentBenchmark(name) {
		return "same loop"
	}
	if uncertain {
		return formatUncertainTiming(lo, col, equivalentEpsilon)
	}
	return formatRatio(lo, col)
}

// formatSpeed renders the relative speed between lo and collection.
func formatSpeed(lo, col float64, allowBold bool, scalarOnly bool) string {
	if lo < benchRatioNoiseNs && col < benchRatioNoiseNs {
		return "≈"
	}
	if col == 0 {
		return "∞"
	}

	ratio := lo / col
	if scalarOnly && ratio >= 1-scalarOnlyEpsilon && ratio <= 1+scalarOnlyEpsilon {
		return "≈"
	}
	if ratio >= 1-equivalentEpsilon && ratio <= 1+equivalentEpsilon {
		return "≈"
	}

	if ratio > 1 && allowBold {
		return "**faster**"
	}
	if ratio > 1 {
		return "faster"
	}
	return "slower"
}

// formatBenchmarkSpeed marks borrowed views as a semantic trade-off.
// formatBenchmarkSpeed applies benchmark-specific speed presentation rules.
func formatBenchmarkSpeed(name string, mode benchMode, lo, col float64, uncertain, allowBold, scalarOnly bool) string {
	if mode == benchBorrow && isViewBenchmark(name) {
		return "view trade-off"
	}
	if mode == benchBorrow && isCodeEquivalentBenchmark(name) {
		return "same loop"
	}
	if uncertain {
		epsilon := equivalentEpsilon
		if scalarOnly {
			epsilon = scalarOnlyEpsilon
		}
		return formatUncertainTiming(lo, col, epsilon)
	}
	return formatSpeed(lo, col, allowBold, scalarOnly)
}

// formatUncertainTiming distinguishes equivalent medians from inconsistent samples.
func formatUncertainTiming(lo, col, epsilon float64) string {
	if lo < benchRatioNoiseNs && col < benchRatioNoiseNs {
		return "≈"
	}
	if col == 0 {
		return "inconclusive"
	}
	ratio := lo / col
	if ratio >= 1-epsilon && ratio <= 1+epsilon {
		return "≈"
	}
	return "inconclusive"
}

// isCodeEquivalentBenchmark reports operations whose compared implementations compile to the same loop.
func isCodeEquivalentBenchmark(name string) bool {
	return name == "FirstWhere"
}

// isViewBenchmark reports operations whose collection result intentionally shares backing storage.
// isViewBenchmark reports whether a case returns a borrowed collection view.
func isViewBenchmark(name string) bool {
	switch name {
	case "Chunk", "Skip", "SkipLast":
		return true
	default:
		return false
	}
}

// isDifferentWorkBenchmark reports APIs that produce materially different outputs.
// isDifferentWorkBenchmark reports whether implementations perform intentionally different work.
func isDifferentWorkBenchmark(name string) bool {
	return name == "Difference"
}

// formatRatioBytes compares lo allocation size with collection allocation size.
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

// formatInt renders an integer with thousands separators.
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

// formatDeltaBytes renders the collection allocation-size difference from lo.
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

// formatDeltaAllocs renders the collection allocation-count difference from lo.
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

// ----------------------------------------------------------------------------
// README injection
// ----------------------------------------------------------------------------

// updateReadme replaces the generated benchmark summary in the project README.
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

// replaceSection replaces the benchmark embed region in README content.
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

// replaceBenchTable replaces the generated table within the benchmark section.
func replaceBenchTable(section, condensed string) (string, error) {
	trimmed := strings.TrimSpace(condensed)
	if trimmed == "" {
		return "", fmt.Errorf("condensed benchmark content is empty")
	}
	return "\n\n" + trimmed + "\n", nil
}

// updateBenchmarksFile writes detailed borrow and copy results to the benchmark report.
func updateBenchmarksFile(rawBorrowTable, rawCopyTable string) error {
	root, err := findRoot()
	if err != nil {
		return err
	}

	path := filepath.Join(root, "BENCHMARKS.md")
	var buf bytes.Buffer
	buf.WriteString("# Benchmarks\n\n")
	fmt.Fprintf(&buf, "Methodology: %s on %s/%s, GOMAXPROCS=%d; median of %d paired samples at %s each, alternating implementation order. Timing differences are shown only when every pair falls outside the ±%.0f%% equivalence band in the same direction. Medians inside the band are labeled `≈`; medians outside it without consistent paired evidence are labeled `inconclusive`. Mutable borrowed inputs are restored inside every timed iteration for both implementations.\n\n", runtime.Version(), runtime.GOOS, runtime.GOARCH, runtime.GOMAXPROCS(0), benchSamples, benchSampleDuration, equivalentEpsilon*100)
	buf.WriteString("Raw results for `collection.New` (borrowed) vs `lo`. For Chunk, Skip, and SkipLast, collection returns a view while lo returns a copy; those rows describe an ownership and allocation trade-off, not equal-work speed superiority. Difference returns one-sided output while lo returns both sides, so its rows are an API trade-off.\n\n")
	buf.WriteString("FirstWhere compiles to the same scan loop in both implementations. Its ratio is labeled `same loop` because binary placement can dominate the timing of such a small function in this in-process harness.\n\n")
	buf.WriteString(rawBorrowTable)
	buf.WriteString("\n\n")
	buf.WriteString("Raw results for `collection.New().Clone()` (explicit copy) vs `lo`. This section includes collection's explicit input-copy cost.\n\n")
	buf.WriteString(rawCopyTable)
	buf.WriteString("\n")
	return os.WriteFile(path, buf.Bytes(), 0o644)
}

// ----------------------------------------------------------------------------
// Helpers
// ----------------------------------------------------------------------------

const projectModule = "module github.com/goforj/collection/v3"

// findRoot locates the project module containing the benchmark reports.
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

// fileExists reports whether a path names an existing file.
func fileExists(p string) bool {
	_, err := os.Stat(p)
	return err == nil
}

// isProjectModule reports whether a module file belongs to this project.
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
