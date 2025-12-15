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
		{"Map", benchMapCollection, benchMapLo},
		{"Filter", benchFilterCollection, benchFilterLo},
		{"Chunk", benchChunkCollection, benchChunkLo},
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
	benchSize        = 50_000
	benchPipelineLen = 5_000
	benchChunkSize   = 100
)

var (
	benchInts []int
	workA     []int
	workB     []int
)

func init() {
	benchInts = make([]int, benchSize)
	for i := 0; i < benchSize; i++ {
		benchInts[i] = i
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
	buf.WriteString("| Op | ns/op (col/lo ×) | allocs/op (col/lo) | 10k iters Δ (time / allocs) |\n")
	buf.WriteString("|----|------------------|--------------------|-----------------------------|\n")

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

		timeDelta := (loRes.nsPerOp - col.nsPerOp) * hotPathIters
		allocDelta := (loRes.allocsPerOp - col.allocsPerOp) * hotPathIters

		deltaCell := fmt.Sprintf(
			"%s / %s",
			formatDurationNs(timeDelta),
			formatInt(allocDelta),
		)

		buf.WriteString(fmt.Sprintf(
			"| %s | %s | %s | %s |\n",
			name,
			nsCell,
			allocCell,
			deltaCell,
		))
	}

	buf.WriteString("\n> **Hot-path context**  \n")
	buf.WriteString("> `10k iters Δ` is a derived estimate showing total time and allocation savings over sustained workloads (e.g. worker pools).")

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
