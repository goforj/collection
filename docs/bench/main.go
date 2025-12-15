package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"time"

	"github.com/goforj/collection"
	"github.com/samber/lo"
)

const (
	benchStart = "<!-- bench:embed:start -->"
	benchEnd   = "<!-- bench:embed:end -->"
)

type benchResult struct {
	name        string
	nsPerOp     float64
	bytesPerOp  int64
	allocsPerOp int64
	impl        string
}

func main() {
	results := runBenches()
	table := renderTable(results)

	if err := updateReadme(table); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println("✔ Benchmarks updated in README.md")
}

func runBenches() []benchResult {
	cases := []struct {
		name string
		col  func()
		lo   func()
	}{
		{
			name: "Pipeline Filter→Map→Take→Reduce",
			col:  benchPipelineCollection,
			lo:   benchPipelineLo,
		},
		{
			name: "Map",
			col:  benchMapCollection,
			lo:   benchMapLo,
		},
		{
			name: "Filter",
			col:  benchFilterCollection,
			lo:   benchFilterLo,
		},
		{
			name: "Unique",
			col:  benchUniqueCollection,
			lo:   benchUniqueLo,
		},
		{
			name: "Chunk",
			col:  benchChunkCollection,
			lo:   benchChunkLo,
		},
	}

	var results []benchResult

	for _, c := range cases {
		colRes := measure(c.name, "collection", c.col)
		loRes := measure(c.name, "lo", c.lo)
		results = append(results, colRes, loRes)
	}

	return results
}

const benchIterations = 200

func measure(name, impl string, fn func()) benchResult {
	runtime.GC()

	var start runtime.MemStats
	runtime.ReadMemStats(&start)

	startTime := time.Now()
	allocs := testingAllocsPerRun(fn, benchIterations)
	duration := time.Since(startTime)

	var end runtime.MemStats
	runtime.ReadMemStats(&end)

	bytesPerOp := int64(0)
	if end.TotalAlloc > start.TotalAlloc {
		bytesPerOp = int64((end.TotalAlloc - start.TotalAlloc) / uint64(benchIterations))
	}

	return benchResult{
		name:        name,
		impl:        impl,
		nsPerOp:     float64(duration.Nanoseconds()) / float64(benchIterations),
		bytesPerOp:  bytesPerOp,
		allocsPerOp: allocs,
	}
}

// testing.AllocsPerRun is unexported; reimplement minimal variant to avoid testing.B overhead.
func testingAllocsPerRun(fn func(), runs int) int64 {
	var totalAlloc uint64
	for i := 0; i < runs; i++ {
		runtime.GC()
		var m1, m2 runtime.MemStats
		runtime.ReadMemStats(&m1)
		fn()
		runtime.ReadMemStats(&m2)
		if m2.Mallocs > m1.Mallocs {
			totalAlloc += m2.Mallocs - m1.Mallocs
		}
	}
	return int64(totalAlloc / uint64(runs))
}

// ----------------------------------------------------------------------------
// Bench cases
// ----------------------------------------------------------------------------

const (
	benchSize        = 5_000
	benchPipelineLen = 250
	benchChunkSize   = 50
)

var (
	benchInts    []int
	benchIntsDup []int
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
}

func benchPipelineCollection() {
	_ = collection.New(benchInts).
		Filter(func(v int) bool { return v%2 == 0 }).
		Map(func(v int) int { return v * v }).
		Take(benchPipelineLen).
		Reduce(0, func(acc, v int) int { return acc + v })
}

func benchPipelineLo() {
	out := lo.Filter(benchInts, func(v int, _ int) bool { return v%2 == 0 })
	out2 := lo.Map(out, func(v int, _ int) int { return v * v })
	out3 := lo.Subset(out2, 0, benchPipelineLen)
	_ = lo.Reduce(out3, func(acc int, v int, _ int) int { return acc + v }, 0)
}

func benchMapCollection() {
	_ = collection.New(benchInts).Map(func(v int) int { return v * 3 })
}

func benchMapLo() {
	_ = lo.Map(benchInts, func(v int, _ int) int { return v * 3 })
}

func benchFilterCollection() {
	_ = collection.New(benchInts).Filter(func(v int) bool { return v%3 == 0 })
}

func benchFilterLo() {
	_ = lo.Filter(benchInts, func(v int, _ int) bool { return v%3 == 0 })
}

func benchUniqueCollection() {
	_ = collection.UniqueBy(collection.New(benchIntsDup), func(v int) int { return v })
}

func benchUniqueLo() {
	_ = lo.Uniq(benchIntsDup)
}

func benchChunkCollection() {
	_ = collection.New(benchInts).Chunk(benchChunkSize)
}

func benchChunkLo() {
	_ = lo.Chunk(benchInts, benchChunkSize)
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
	buf.WriteString("| Operation | collection ns/op | lo ns/op | speedup (×) | collection B/op | lo B/op | mem (×) | collection allocs/op | lo allocs/op | allocs (×) |\n")
	buf.WriteString("|-----------|-----------------:|---------:|------------:|-----------------:|--------:|--------:|---------------------:|--------------:|-----------:|\n")

	names := make([]string, 0, len(byName))
	for name := range byName {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		col, okCol := byName[name]["collection"]
		loRes, okLo := byName[name]["lo"]
		if !okCol || !okLo {
			continue
		}

		buf.WriteString(fmt.Sprintf("| %s | %s | %s | %s | %d | %d | %s | %d | %d | %s |\n",
			name,
			formatNs(col.nsPerOp),
			formatNs(loRes.nsPerOp),
			formatRatio(loRes.nsPerOp, col.nsPerOp),
			col.bytesPerOp,
			loRes.bytesPerOp,
			formatRatioInt(loRes.bytesPerOp, col.bytesPerOp),
			col.allocsPerOp,
			loRes.allocsPerOp,
			formatRatioInt(loRes.allocsPerOp, col.allocsPerOp),
		))
	}

	return strings.TrimSpace(buf.String())
}

func formatNs(ns float64) string {
	if ns >= 1e6 {
		return fmt.Sprintf("%.1fms", ns/1e6)
	}
	if ns >= 1e3 {
		return fmt.Sprintf("%.1fµs", ns/1e3)
	}
	return fmt.Sprintf("%.0fns", ns)
}

func formatRatio(lo, col float64) string {
	if col == 0 {
		if lo == 0 {
			return "1.0x"
		}
		return "∞"
	}
	return fmt.Sprintf("%.2fx", lo/col)
}

func formatRatioInt(lo, col int64) string {
	if col == 0 {
		if lo == 0 {
			return "1.0x"
		}
		return "∞"
	}
	return fmt.Sprintf("%.2fx", float64(lo)/float64(col))
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
			// hit non-empty content before module line
			return false
		}
	}

	return false
}
