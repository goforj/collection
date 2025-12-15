package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"

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
		col  func(*testing.B)
		lo   func(*testing.B)
	}{
		{
			name: "Pipeline F→M→T→R",
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
	benchSize        = 50_000
	benchPipelineLen = 5_000
	benchChunkSize   = 100
)

var (
	benchInts    []int
	benchIntsDup []int
	workA        []int
	workB        []int
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

	workA = make([]int, benchSize)
	workB = make([]int, benchSize)
}

func benchPipelineCollection(b *testing.B) {
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
	for i := 0; i < b.N; i++ {
		copy(workB, benchInts)
		out := lo.Filter(workB, func(v int, _ int) bool { return v%2 == 0 })
		out2 := lo.Map(out, func(v int, _ int) int { return v * v })
		out3 := lo.Subset(out2, 0, benchPipelineLen)
		_ = lo.Reduce(out3, func(acc int, v int, _ int) int { return acc + v }, 0)
	}
}

func benchMapCollection(b *testing.B) {
	for i := 0; i < b.N; i++ {
		copy(workA, benchInts)
		_ = collection.New(workA).Map(func(v int) int { return v * 3 })
	}
}

func benchMapLo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		copy(workB, benchInts)
		_ = lo.Map(workB, func(v int, _ int) int { return v * 3 })
	}
}

func benchFilterCollection(b *testing.B) {
	for i := 0; i < b.N; i++ {
		copy(workA, benchInts)
		_ = collection.New(workA).Filter(func(v int) bool { return v%3 == 0 })
	}
}

func benchFilterLo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		copy(workB, benchInts)
		_ = lo.Filter(workB, func(v int, _ int) bool { return v%3 == 0 })
	}
}

func benchUniqueCollection(b *testing.B) {
	// not used
}

func benchUniqueLo(b *testing.B) {
	// not used
}

func benchChunkCollection(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = collection.New(benchInts).Chunk(benchChunkSize)
	}
}

func benchChunkLo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = lo.Chunk(benchInts, benchChunkSize)
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
	buf.WriteString("| Op | ns/op (col/lo, ×) | B/op (col/lo, ×) | allocs/op (col/lo, ×) |\n")
	buf.WriteString("|---|-------------------|------------------|-----------------------|\n")

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

		nsCell := fmt.Sprintf("%s / %s (%s)", formatNs(col.nsPerOp), formatNs(loRes.nsPerOp), formatRatio(loRes.nsPerOp, col.nsPerOp))
		bCell := fmt.Sprintf("%d / %d (%s)", col.bytesPerOp, loRes.bytesPerOp, formatRatioInt(loRes.bytesPerOp, col.bytesPerOp))
		allocCell := fmt.Sprintf("%d / %d (%s)", col.allocsPerOp, loRes.allocsPerOp, formatRatioInt(loRes.allocsPerOp, col.allocsPerOp))

		buf.WriteString(fmt.Sprintf("| %s | %s | %s | %s |\n", name, nsCell, bCell, allocCell))
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
