//go:build ignore
// +build ignore

package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

const (
	apiStart = "<!-- api:embed:start -->"
	apiEnd   = "<!-- api:embed:end -->"
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
	groupHeader   = regexp.MustCompile(`(?i)^\s*@group\s+(.+)$`)
	exampleHeader = regexp.MustCompile(`(?i)^\s*Example:\s*(.*)$`)
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

	pkg, ok := pkgs["collection"]
	if !ok {
		return nil, fmt.Errorf(`package "collection" not found`)
	}

	funcs := map[string]*FuncDoc{}

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
				Name:        fn.Name.Name,
				Group:       extractGroup(fn.Doc),
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

func extractGroup(group *ast.CommentGroup) string {
	for _, c := range group.List {
		line := strings.TrimSpace(strings.TrimPrefix(c.Text, "//"))
		if m := groupHeader.FindStringSubmatch(line); m != nil {
			return strings.TrimSpace(m[1])
		}
	}
	return "Other"
}

func extractDescription(group *ast.CommentGroup) string {
	var lines []string

	for _, c := range group.List {
		line := strings.TrimSpace(strings.TrimPrefix(c.Text, "//"))

		if exampleHeader.MatchString(line) || groupHeader.MatchString(line) {
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

		normalized := normalizeIndent(current)

		out = append(out, Example{
			Label: label,
			Code:  strings.Join(normalized, "\n"),
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
	buf.WriteString("### Index\n\n")
	buf.WriteString("| Group | Functions |\n")
	buf.WriteString("|------:|-----------|\n")

	for _, group := range groupNames {
		sort.Slice(byGroup[group], func(i, j int) bool {
			return byGroup[group][i].Name < byGroup[group][j].Name
		})

		var links []string
		for _, fn := range byGroup[group] {
			anchor := strings.ToLower(fn.Name)
			links = append(links, fmt.Sprintf("[%s](#%s)", fn.Name, anchor))
		}

		buf.WriteString(fmt.Sprintf("| **%s** | %s |\n",
			group,
			strings.Join(links, " "),
		))
	}

	buf.WriteString("\n\n")

	// ---------------- Details ----------------
	for _, group := range groupNames {
		buf.WriteString("## " + group + "\n\n")

		for _, fn := range byGroup[group] {
			buf.WriteString(fmt.Sprintf("### %s\n", fn.Name))

			if fn.Description != "" {
				buf.WriteString(fn.Description + "\n\n")
			}

			for _, ex := range fn.Examples {
				if ex.Label != "" {
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
	minIndent := -1

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}

		indent := len(line) - len(strings.TrimLeft(line, " \t"))
		if minIndent == -1 || indent < minIndent {
			minIndent = indent
		}
	}

	if minIndent <= 0 {
		return lines
	}

	out := make([]string, len(lines))
	for i, line := range lines {
		if len(line) >= minIndent {
			out[i] = line[minIndent:]
		} else {
			out[i] = strings.TrimLeft(line, " \t")
		}
	}

	return out
}
