package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

const (
	apiStart       = "<!-- api:embed:start -->"
	apiEnd         = "<!-- api:embed:end -->"
	testCountStart = "<!-- test-count:embed:start -->"
	testCountEnd   = "<!-- test-count:embed:end -->"
)

func main() {
	if err := run(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println("✔ API section updated in README.md")
}

func run() error {
	root, err := findRoot()
	if err != nil {
		return err
	}

	testsCount, err := countTests(root)
	if err != nil {
		return fmt.Errorf("count tests: %w", err)
	}

	funcs, err := parseFuncs(root)
	if err != nil {
		return err
	}

	api := renderAPI(funcs)

	readmePath := filepath.Join(root, "README.md")
	data, err := os.ReadFile(readmePath)
	if err != nil {
		return err
	}

	out, err := replaceAPISection(string(data), api)
	if err != nil {
		return err
	}

	out, err = updateTestsSection(out, testsCount)
	if err != nil {
		return err
	}

	return os.WriteFile(readmePath, []byte(out), 0o644)
}

//
// ------------------------------------------------------------
// Data model
// ------------------------------------------------------------
//

type FuncDoc struct {
	Name        string
	Group       string
	Behavior    string
	Chainable   string
	Terminal    string
	Description string
	Examples    []Example
}

type Example struct {
	Label string
	Code  string
	Line  int
}

//
// ------------------------------------------------------------
// Parsing
// ------------------------------------------------------------
//

var (
	groupHeader     = regexp.MustCompile(`(?i)^\s*@group\s+(.+)$`)
	behaviorHeader  = regexp.MustCompile(`(?i)^\s*@behavior\s+(.+)$`)
	chainableHeader = regexp.MustCompile(`(?i)^\s*@chainable\s+(.+)$`)
	terminalHeader  = regexp.MustCompile(`(?i)^\s*@terminal\s+(.+)$`)
	exampleHeader   = regexp.MustCompile(`(?i)^\s*Example:\s*(.*)$`)
)

func parseFuncs(root string) ([]*FuncDoc, error) {
	fset := token.NewFileSet()

	pkgs, err := parser.ParseDir(
		fset,
		root,
		func(info os.FileInfo) bool {
			return !strings.HasSuffix(info.Name(), "_test.go")
		},
		parser.ParseComments,
	)
	if err != nil {
		return nil, err
	}

	pkgName, err := selectPackage(pkgs)
	if err != nil {
		return nil, err
	}

	pkg, ok := pkgs[pkgName]
	if !ok {
		return nil, fmt.Errorf(`package %q not found`, pkgName)
	}

	funcs := map[string]*FuncDoc{}
	packageFuncs := packageFunctionNames(pkg)

	for _, file := range pkg.Files {
		for _, decl := range file.Decls {
			fn, ok := decl.(*ast.FuncDecl)
			if !ok || fn.Doc == nil {
				continue
			}

			if !ast.IsExported(fn.Name.Name) {
				continue
			}

			fd := &FuncDoc{
				Name:        funcDocName(fn, packageFuncs),
				Group:       extractGroup(fn.Doc),
				Behavior:    extractBehavior(fn.Doc),
				Chainable:   extractChainable(fn.Doc),
				Terminal:    extractTerminal(fn.Doc),
				Description: extractDescription(fn.Doc),
				Examples:    extractExamples(fset, fn),
			}

			if existing, ok := funcs[fd.Name]; ok {
				existing.Examples = append(existing.Examples, fd.Examples...)
			} else {
				funcs[fd.Name] = fd
			}
		}
	}

	out := make([]*FuncDoc, 0, len(funcs))
	for _, fd := range funcs {
		sort.Slice(fd.Examples, func(i, j int) bool {
			return fd.Examples[i].Line < fd.Examples[j].Line
		})
		out = append(out, fd)
	}

	return out, nil
}

// packageFunctionNames returns package-level function names that may collide with methods.
func packageFunctionNames(pkg *ast.Package) map[string]bool {
	names := make(map[string]bool)
	for _, file := range pkg.Files {
		for _, decl := range file.Decls {
			fn, ok := decl.(*ast.FuncDecl)
			if ok && fn.Recv == nil {
				names[fn.Name.Name] = true
			}
		}
	}
	return names
}

// funcDocName qualifies a method only when a package function has the same name.
func funcDocName(fn *ast.FuncDecl, packageFuncs map[string]bool) string {
	if fn.Recv == nil || len(fn.Recv.List) == 0 || !packageFuncs[fn.Name.Name] {
		return fn.Name.Name
	}
	receiver := fn.Recv.List[0].Type
	for {
		switch value := receiver.(type) {
		case *ast.ParenExpr:
			receiver = value.X
		case *ast.StarExpr:
			receiver = value.X
		case *ast.IndexExpr:
			receiver = value.X
		case *ast.IndexListExpr:
			receiver = value.X
		case *ast.Ident:
			return value.Name + "." + fn.Name.Name
		default:
			return fn.Name.Name
		}
	}
}

func extractGroup(group *ast.CommentGroup) string {
	for _, c := range group.List {
		line := strings.TrimSpace(strings.TrimPrefix(c.Text, "//"))
		if m := groupHeader.FindStringSubmatch(line); m != nil {
			return strings.TrimSpace(m[1])
		}
	}
	return "Other"
}

func extractBehavior(group *ast.CommentGroup) string {
	for _, c := range group.List {
		line := strings.TrimSpace(strings.TrimPrefix(c.Text, "//"))
		if m := behaviorHeader.FindStringSubmatch(line); m != nil {
			return strings.ToLower(strings.TrimSpace(m[1]))
		}
	}
	return ""
}

func extractChainable(group *ast.CommentGroup) string {
	for _, c := range group.List {
		line := strings.TrimSpace(strings.TrimPrefix(c.Text, "//"))
		if m := chainableHeader.FindStringSubmatch(line); m != nil {
			return strings.ToLower(strings.TrimSpace(m[1]))
		}
	}
	return ""
}

func extractTerminal(group *ast.CommentGroup) string {
	for _, c := range group.List {
		line := strings.TrimSpace(strings.TrimPrefix(c.Text, "//"))
		if m := terminalHeader.FindStringSubmatch(line); m != nil {
			return strings.ToLower(strings.TrimSpace(m[1]))
		}
	}
	return ""
}

func extractDescription(group *ast.CommentGroup) string {
	var lines []string

	for _, c := range group.List {
		line := strings.TrimSpace(strings.TrimPrefix(c.Text, "//"))

		if exampleHeader.MatchString(line) ||
			groupHeader.MatchString(line) ||
			behaviorHeader.MatchString(line) ||
			chainableHeader.MatchString(line) ||
			terminalHeader.MatchString(line) {
			break
		}

		if len(lines) == 0 && line == "" {
			continue
		}

		lines = append(lines, line)
	}

	return strings.TrimSpace(strings.Join(lines, "\n"))
}

func extractExamples(fset *token.FileSet, fn *ast.FuncDecl) []Example {
	var out []Example
	var current []string
	var label string
	var start int
	inExample := false

	flush := func() {
		if len(current) == 0 {
			return
		}

		out = append(out, Example{
			Label: label,
			Code:  strings.Join(normalizeIndent(current), "\n"),
			Line:  start,
		})

		current = nil
		label = ""
		inExample = false
	}

	for _, c := range fn.Doc.List {
		raw := strings.TrimPrefix(c.Text, "//")
		line := strings.TrimSpace(raw)

		if m := exampleHeader.FindStringSubmatch(line); m != nil {
			flush()
			inExample = true
			label = strings.TrimSpace(m[1])
			start = fset.Position(c.Slash).Line
			continue
		}

		if !inExample {
			continue
		}

		current = append(current, raw)
	}

	flush()
	return out
}

// selectPackage picks the primary package to document.
// Strategy:
//  1. If only one package exists, use it.
//  2. Prefer the non-"main" package with the most files.
//  3. Fall back to the first package alphabetically.
func selectPackage(pkgs map[string]*ast.Package) (string, error) {
	if len(pkgs) == 0 {
		return "", fmt.Errorf("no packages found")
	}

	if len(pkgs) == 1 {
		for name := range pkgs {
			return name, nil
		}
	}

	type candidate struct {
		name  string
		count int
	}

	candidates := make([]candidate, 0, len(pkgs))
	for name, pkg := range pkgs {
		candidates = append(candidates, candidate{
			name:  name,
			count: len(pkg.Files),
		})
	}

	sort.Slice(candidates, func(i, j int) bool {
		if candidates[i].count == candidates[j].count {
			return candidates[i].name < candidates[j].name
		}
		return candidates[i].count > candidates[j].count
	})

	for _, cand := range candidates {
		if cand.name != "main" {
			return cand.name, nil
		}
	}

	return candidates[0].name, nil
}

//
// ------------------------------------------------------------
// Rendering
// ------------------------------------------------------------
//

func renderAPI(funcs []*FuncDoc) string {
	byGroup := map[string][]*FuncDoc{}

	for _, fd := range funcs {
		byGroup[fd.Group] = append(byGroup[fd.Group], fd)
	}

	groupNames := make([]string, 0, len(byGroup))
	for g := range byGroup {
		groupNames = append(groupNames, g)
	}
	sort.Strings(groupNames)

	var buf bytes.Buffer

	// ---------------- Index ----------------
	buf.WriteString("# API Index\n\n")
	buf.WriteString("| Group | Functions |\n")
	buf.WriteString("|------:|-----------|\n")

	for _, group := range groupNames {
		sort.Slice(byGroup[group], func(i, j int) bool {
			return byGroup[group][i].Name < byGroup[group][j].Name
		})

		var links []string
		for _, fn := range byGroup[group] {
			links = append(links, fmt.Sprintf("[%s](#%s)", fn.Name, strings.ToLower(fn.Name)))
		}

		buf.WriteString(fmt.Sprintf("| **%s** | %s |\n",
			group,
			strings.Join(links, " - "),
		))
	}

	buf.WriteString("\n\n")

	// ---------------- Details ----------------
	for _, group := range groupNames {
		buf.WriteString("## " + group + "\n\n")

		for _, fn := range byGroup[group] {
			anchor := strings.ToLower(fn.Name)

			header := fn.Name
			if fn.Behavior != "" {
				header += " - " + fn.Behavior
			}
			if fn.Chainable == "true" {
				header += " - chainable"
			}
			if fn.Terminal == "true" {
				header += " - terminal"
			}

			buf.WriteString(fmt.Sprintf("### <a id=\"%s\"></a>%s\n\n", anchor, header))

			if fn.Description != "" {
				buf.WriteString(fn.Description + "\n\n")
			}

			for _, ex := range fn.Examples {
				if ex.Label != "" && len(fn.Examples) > 1 {
					buf.WriteString(fmt.Sprintf("_Example: %s_\n\n", ex.Label))
				}

				buf.WriteString("```go\n")
				buf.WriteString(strings.TrimSpace(ex.Code))
				buf.WriteString("\n```\n\n")
			}
		}
	}

	return strings.TrimRight(buf.String(), "\n")
}

//
// ------------------------------------------------------------
// README replacement
// ------------------------------------------------------------
//

func replaceAPISection(readme, api string) (string, error) {
	start := strings.Index(readme, apiStart)
	end := strings.Index(readme, apiEnd)

	if start == -1 || end == -1 || end < start {
		return "", fmt.Errorf("API anchors not found or malformed")
	}

	var out bytes.Buffer
	out.WriteString(readme[:start+len(apiStart)])
	out.WriteString("\n\n")
	out.WriteString(api)
	out.WriteString("\n")
	out.WriteString(readme[end:])

	return out.String(), nil
}

// countTests runs the test suite and counts its structured test events.
func countTests(root string) (int, error) {
	cmd := exec.Command("go", "test", "./...", "-run", "Test", "-count=1", "-json")
	cmd.Dir = root

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	runErr := cmd.Run()
	return countTestOutput(stdout.Bytes(), stderr.Bytes(), runErr)
}

// countTestOutput parses only the JSON stream written to stdout because the Go
// command may emit dependency download diagnostics to stderr on a fresh cache.
func countTestOutput(stdout, stderr []byte, runErr error) (int, error) {
	if runErr != nil {
		return 0, fmt.Errorf("go test -json: %w\nstdout:\n%s\nstderr:\n%s", runErr, stdout, stderr)
	}

	var total int
	dec := json.NewDecoder(bytes.NewReader(stdout))

	for {
		var event struct {
			Action string `json:"Action"`
			Test   string `json:"Test"`
		}
		err := dec.Decode(&event)
		if err == io.EOF {
			break
		}
		if err != nil {
			return 0, fmt.Errorf("decode go test JSON: %w\nstdout:\n%s\nstderr:\n%s", err, stdout, stderr)
		}
		if event.Action == "run" && event.Test != "" {
			total++
		}
	}

	return total, nil
}

var testsBadgePattern = regexp.MustCompile(`tests-\d+-brightgreen`)

func updateTestsSection(readme string, tests int) (string, error) {
	start := strings.Index(readme, testCountStart)
	end := strings.Index(readme, testCountEnd)

	if start == -1 || end == -1 || end < start {
		return "", fmt.Errorf("test count anchors not found or malformed")
	}

	before := readme[:start+len(testCountStart)]
	body := readme[start+len(testCountStart) : end]
	after := readme[end:]

	leading := ""
	if strings.HasPrefix(body, "\n") {
		leading = "\n"
	}

	badge := fmt.Sprintf("%s    <img src=\"https://img.shields.io/badge/tests-%d-brightgreen\" alt=\"Tests\">\n", leading, tests)

	return before + badge + after, nil
}

//
// ------------------------------------------------------------
// Helpers
// ------------------------------------------------------------
//

func findRoot() (string, error) {
	wd, _ := os.Getwd()
	if fileExists(filepath.Join(wd, "go.mod")) {
		return wd, nil
	}
	parent := filepath.Join(wd, "..")
	if fileExists(filepath.Join(parent, "go.mod")) {
		return filepath.Clean(parent), nil
	}
	return "", fmt.Errorf("could not find project root")
}

func fileExists(p string) bool {
	_, err := os.Stat(p)
	return err == nil
}

func normalizeIndent(lines []string) []string {
	min := -1

	for _, l := range lines {
		if strings.TrimSpace(l) == "" {
			continue
		}
		n := len(l) - len(strings.TrimLeft(l, " \t"))
		if min == -1 || n < min {
			min = n
		}
	}

	if min <= 0 {
		return lines
	}

	out := make([]string, len(lines))
	for i, l := range lines {
		if len(l) >= min {
			out[i] = l[min:]
		} else {
			out[i] = strings.TrimLeft(l, " \t")
		}
	}

	return out
}
