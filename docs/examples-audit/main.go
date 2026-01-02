//go:build ignore
// +build ignore

package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"unicode"
)

var (
	ansiRegexp        = regexp.MustCompile(`\x1b\\[[0-9;]*m`)
	errExampleRunFail = errors.New("example run failed")
)

type result struct {
	name       string
	expected   []string
	actual     []string
	runErr     error
	hasOutput  bool
	exitWanted bool
	nondet     bool
}

func main() {
	examplesDir := "examples"
	entries, err := os.ReadDir(examplesDir)
	if err != nil {
		fatalf("cannot read examples directory: %v", err)
	}

	var results []result
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}

		name := e.Name()
		path := filepath.Join(examplesDir, name, "main.go")
		src, err := os.ReadFile(path)
		if err != nil {
			fatalf("read %s: %v", path, err)
		}

		expected := extractExpectedOutput(src)
		exitWanted := expectsExit(src)
		nondet := isNonDeterministicExample(src)
		actual, runErr := runExample(name, src)
		hasOutput := len(expected) > 0

		results = append(results, result{
			name:       name,
			expected:   expected,
			actual:     actual,
			runErr:     runErr,
			hasOutput:  hasOutput,
			exitWanted: exitWanted,
			nondet:     nondet,
		})
	}

	var failed []result
	var missing []result

	for _, res := range results {
		if res.nondet {
			continue
		}
		if res.runErr != nil {
			failed = append(failed, res)
			continue
		}

		if !res.hasOutput {
			if len(res.actual) > 0 {
				missing = append(missing, res)
			}
			continue
		}

		if !equalLines(res.expected, res.actual) {
			failed = append(failed, res)
		}
	}

	if len(failed) == 0 && len(missing) == 0 {
		fmt.Println("✔ Examples output matches comments (normalized).")
		return
	}

	if len(failed) > 0 {
		fmt.Println("✖ Output mismatches:")
		for _, res := range failed {
			fmt.Printf("\n[%s]\n", res.name)
			if res.runErr != nil {
				fmt.Printf("run error: %v\n", res.runErr)
				continue
			}
			fmt.Println("expected:")
			fmt.Println(joinLines(res.expected))
			fmt.Println("actual:")
			fmt.Println(joinLines(res.actual))
		}
	}

	if len(missing) > 0 {
		fmt.Println("\n✖ Missing expected output comments (output was produced):")
		for _, res := range missing {
			fmt.Printf("\n[%s]\n", res.name)
			fmt.Println("actual:")
			fmt.Println(joinLines(res.actual))
		}
	}

	os.Exit(1)
}

func fatalf(format string, args ...any) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(1)
}

func runExample(name string, src []byte) ([]string, error) {
	clean := stripBuildTags(src)

	tmpDir, err := os.MkdirTemp("", "example-audit-*")
	if err != nil {
		return nil, fmt.Errorf("mkdir temp dir: %w", err)
	}
	defer os.RemoveAll(tmpDir)

	orig := filepath.Join("examples", name, "main.go")
	tmpFile := filepath.Join(tmpDir, "main.go")
	if err := os.WriteFile(tmpFile, clean, 0o600); err != nil {
		return nil, fmt.Errorf("write temp main.go: %w", err)
	}

	overlay := map[string]any{
		"Replace": map[string]string{
			abs(orig): abs(tmpFile),
		},
	}
	overlayJSON, err := json.Marshal(overlay)
	if err != nil {
		return nil, fmt.Errorf("marshal overlay: %w", err)
	}

	overlayPath := filepath.Join(tmpDir, "overlay.json")
	if err := os.WriteFile(overlayPath, overlayJSON, 0o600); err != nil {
		return nil, fmt.Errorf("write overlay: %w", err)
	}

	cmd := exec.Command("go", "run", "-overlay", overlayPath, "./examples/"+name)
	cmd.Env = append(os.Environ(), "NO_COLOR=1")

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()
	if err != nil {
		if expectsExit(src) && exitCodeIs(err, 1) {
			return normalizeOutput(stdout.String()), nil
		}
		if stderr.Len() > 0 {
			return normalizeOutput(stdout.String()), fmt.Errorf("%w: %s", errExampleRunFail, strings.TrimSpace(stderr.String()))
		}
		return normalizeOutput(stdout.String()), fmt.Errorf("%w: %v", errExampleRunFail, err)
	}

	return normalizeOutput(stdout.String()), nil
}

func exitCodeIs(err error, code int) bool {
	var exitErr *exec.ExitError
	if !errors.As(err, &exitErr) {
		return false
	}
	return exitErr.ExitCode() == code
}

func abs(p string) string {
	a, err := filepath.Abs(p)
	if err != nil {
		panic(err)
	}
	return a
}

func stripBuildTags(src []byte) []byte {
	lines := strings.Split(string(src), "\n")

	i := 0
	for i < len(lines) {
		line := strings.TrimSpace(lines[i])

		if strings.HasPrefix(line, "//go:build") ||
			strings.HasPrefix(line, "// +build") ||
			line == "" {
			i++
			continue
		}

		break
	}

	return []byte(strings.Join(lines[i:], "\n"))
}

func extractExpectedOutput(src []byte) []string {
	lines := strings.Split(string(src), "\n")
	var out []string
	for _, line := range lines {
		payload, ok := commentPayload(line)
		if !ok {
			continue
		}
		if isOutputLine(payload) {
			out = append(out, normalizeLine(payload))
		}
	}
	return normalizeMapBlocks(out)
}

func commentPayload(line string) (string, bool) {
	idx := strings.Index(line, "//")
	if idx == -1 {
		return "", false
	}
	if strings.TrimSpace(line[:idx]) != "" {
		return "", false
	}
	payload := line[idx+2:]
	if strings.HasPrefix(payload, " ") && (len(payload) == 1 || payload[1] != ' ') {
		payload = payload[1:]
	}
	return payload, true
}

func isOutputLine(line string) bool {
	trimmed := strings.TrimLeftFunc(line, unicode.IsSpace)
	if trimmed == "" {
		return false
	}

	lower := strings.ToLower(trimmed)
	if strings.HasPrefix(lower, "example:") ||
		strings.HasPrefix(lower, "output:") ||
		strings.HasPrefix(lower, "+build") ||
		strings.HasPrefix(lower, "go:build") ||
		strings.HasPrefix(lower, "process finished") {
		return false
	}

	if strings.HasPrefix(trimmed, "#") ||
		strings.HasPrefix(trimmed, "<#") ||
		strings.HasPrefix(trimmed, "+") ||
		strings.HasPrefix(trimmed, "-") ||
		strings.HasPrefix(trimmed, "{") ||
		strings.HasPrefix(trimmed, "}") ||
		strings.HasPrefix(trimmed, "[") ||
		strings.HasPrefix(trimmed, "]") ||
		strings.HasPrefix(trimmed, "\"") ||
		strings.Contains(trimmed, "=>") {
		return true
	}

	first := []rune(trimmed)[0]
	if unicode.IsDigit(first) {
		return true
	}

	return strings.HasPrefix(trimmed, "true") ||
		strings.HasPrefix(trimmed, "false") ||
		strings.HasPrefix(trimmed, "nil")
}

func expectsExit(src []byte) bool {
	return bytes.Contains(src, []byte(".Dd(")) || bytes.Contains(src, []byte("collection.Dd("))
}

func isNonDeterministicExample(src []byte) bool {
	return bytes.Contains(src, []byte("Shuffle("))
}

func normalizeOutput(raw string) []string {
	raw = strings.ReplaceAll(raw, "\r\n", "\n")
	raw = ansiRegexp.ReplaceAllString(raw, "")

	lines := strings.Split(raw, "\n")
	var out []string
	for _, line := range lines {
		if strings.HasPrefix(line, "<#dump ") || strings.HasPrefix(line, "<#diff ") {
			continue
		}
		out = append(out, normalizeLine(line))
	}

	for len(out) > 0 && out[0] == "" {
		out = out[1:]
	}
	for len(out) > 0 && out[len(out)-1] == "" {
		out = out[:len(out)-1]
	}
	return normalizeMapBlocks(out)
}

func normalizeLine(line string) string {
	line = strings.TrimRight(line, " \t")
	if line == "" {
		return line
	}

	i := 0
	spaces := 0
	for i < len(line) {
		if line[i] == ' ' {
			spaces++
			i++
			continue
		}
		if line[i] == '\t' {
			spaces += 2
			i++
			continue
		}
		break
	}

	if spaces == 0 {
		return line
	}

	normalized := (spaces / 2) * 2
	return strings.Repeat(" ", normalized) + strings.TrimLeft(line[i:], " \t")
}

type mapBlock struct {
	key   string
	lines []string
}

func normalizeMapBlocks(lines []string) []string {
	var out []string
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		trimmed := strings.TrimLeft(line, " ")
		if strings.HasPrefix(trimmed, "#map") && strings.HasSuffix(strings.TrimSpace(line), "{") {
			mapIndent := leadingSpaces(line)
			out = append(out, line)
			var blocks []mapBlock

			j := i + 1
			for j < len(lines) {
				if strings.TrimSpace(lines[j]) == "}" && leadingSpaces(lines[j]) == mapIndent {
					break
				}

				if isMapEntryLine(lines[j]) {
					indent := leadingSpaces(lines[j])
					start := j
					j++
					for j < len(lines) {
						if strings.TrimSpace(lines[j]) == "}" && leadingSpaces(lines[j]) == mapIndent {
							break
						}
						if isMapEntryLine(lines[j]) && leadingSpaces(lines[j]) == indent {
							break
						}
						j++
					}

					blockLines := append([]string(nil), lines[start:j]...)
					key := entryKey(blockLines[0])
					blocks = append(blocks, mapBlock{key: key, lines: blockLines})
					continue
				}

				out = append(out, lines[j])
				j++
			}

			sort.Slice(blocks, func(a, b int) bool {
				return blocks[a].key < blocks[b].key
			})
			for _, block := range blocks {
				out = append(out, block.lines...)
			}

			if j < len(lines) && strings.TrimSpace(lines[j]) == "}" {
				out = append(out, lines[j])
				i = j
				continue
			}
		}

		out = append(out, line)
	}
	return out
}

func isMapEntryLine(line string) bool {
	return strings.Contains(line, "=>")
}

func entryKey(line string) string {
	parts := strings.SplitN(strings.TrimSpace(line), "=>", 2)
	return strings.TrimSpace(parts[0])
}

func leadingSpaces(line string) int {
	count := 0
	for i := 0; i < len(line); i++ {
		if line[i] != ' ' {
			break
		}
		count++
	}
	return count
}

func equalLines(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func joinLines(lines []string) string {
	if len(lines) == 0 {
		return "(no output)"
	}
	return strings.Join(lines, "\n")
}
