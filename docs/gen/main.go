package main

import (
	"bytes"
	"fmt"
	"github.com/goforj/godump"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

func main() {
	if err := run(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println("✔ Examples generated in ./examples/")
}

func run() error {
	root, err := findRoot()
	if err != nil {
		return err
	}

	examplesDir := filepath.Join(root, "examples")
	_ = os.RemoveAll(examplesDir)
	if err := os.MkdirAll(examplesDir, 0o755); err != nil {
		return err
	}

	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, root, nil, parser.ParseComments)
	if err != nil {
		return err
	}

	pkg, ok := pkgs["collection"]
	if !ok {
		return fmt.Errorf(`package "collection" not found in %s`, root)
	}

	groups := map[string][]Example{}

	for filename, file := range pkg.Files {
		exs := extractExamples(fset, filename, file)
		for _, ex := range exs {
			groups[ex.FuncName] = append(groups[ex.FuncName], ex)
		}
	}

	for funcName, list := range groups {
		sort.Slice(list, func(i, j int) bool { return list[i].Line < list[j].Line })
		if err := writeMain(examplesDir, funcName, list); err != nil {
			return err
		}
	}

	return nil
}

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

func fileExists(p string) bool { _, err := os.Stat(p); return err == nil }

// ------------------------------------------------------------
// Example extraction
// ------------------------------------------------------------

type Example struct {
	FuncName string
	File     string
	Line     int
	Code     string
}

var exampleHeader = regexp.MustCompile(`(?i)^Example:`)

// docLine keeps a processed doc line plus its position in the source file.
type docLine struct {
	text string
	pos  token.Pos
}

func extractExamples(fset *token.FileSet, filename string, file *ast.File) []Example {
	var out []Example

	for _, decl := range file.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if !ok || fn.Doc == nil {
			continue
		}

		funcName := fn.Name.Name
		blocks := extractBlocks(fset, filename, funcName, fn)
		out = append(out, blocks...)
	}

	return out
}

// docLines returns the doc comment lines with the leading '//' removed,
// preserving indentation after that, and tracking their source positions.
func docLines(group *ast.CommentGroup) []docLine {
	if group == nil {
		return nil
	}

	var lines []docLine
	for _, c := range group.List {
		text := c.Text

		// Line comments: // ...
		if strings.HasPrefix(text, "//") {
			line := strings.TrimPrefix(text, "//")

			// Strip *one* leading space (common in prose).
			if strings.HasPrefix(line, " ") {
				line = line[1:]
			}

			// Strip a single leading tab used as a Go doc
			// "code block" marker. This avoids an extra
			// indent level when we embed under func main().
			if strings.HasPrefix(line, "\t") {
				line = line[1:]
			}

			lines = append(lines, docLine{
				text: line,
				pos:  c.Slash, // position of the leading '//'
			})
			continue
		}

		// Very simple handling for /* ... */ doc comments (if ever used).
		if strings.HasPrefix(text, "/*") {
			body := strings.TrimPrefix(text, "/*")
			body = strings.TrimSuffix(body, "*/")
			for _, l := range strings.Split(body, "\n") {
				lines = append(lines, docLine{
					text: l,
					pos:  c.Slash,
				})
			}
		}
	}

	return lines
}

func extractBlocks(fset *token.FileSet, filename, funcName string, fn *ast.FuncDecl) []Example {
	var out []Example

	lines := docLines(fn.Doc)
	if len(lines) == 0 {
		return out
	}

	var code []string
	var output []string

	inExample := false
	startLine := 0

	flush := func() {
		if len(code) == 0 {
			return
		}

		var combined []string
		combined = append(combined, code...)

		// IMPORTANT: do NOT inject an extra blank line between code and output.
		// We want:
		//   Dump(...)
		//   // expected output
		// not:
		//   Dump(...)
		//
		//   // expected output
		if len(output) > 0 {
			combined = append(combined, output...)
		}

		out = append(out, Example{
			FuncName: funcName,
			File:     filename,
			Line:     startLine,
			Code:     strings.Join(combined, "\n"),
		})

		code = nil
		output = nil
		inExample = false
	}

	for _, dl := range lines {
		rawLine := dl.text
		trimmed := strings.TrimSpace(rawLine)

		// Start of a new Example: block
		if exampleHeader.MatchString(trimmed) {
			flush()
			inExample = true
			// Use real file position for the "from X.go:line" annotation
			startLine = fset.Position(dl.pos).Line
			continue
		}

		if !inExample {
			continue
		}

		// Blank line: keep it, but don't switch sections.
		if trimmed == "" {
			if len(output) > 0 {
				output = append(output, "")
			} else {
				code = append(code, "")
			}
			continue
		}

		// OUTPUT lines: inside an Example, after we've seen some code,
		// and the line (ignoring leading spaces/tabs) starts with "//".
		// In the doc, these are written as:
		//   //  // 4.000000 #float64
		// After stripping the outer "//", rawLine begins with "// 4.000..."
		if len(code) > 0 && strings.HasPrefix(strings.TrimLeft(rawLine, " \t"), "//") {
			output = append(output, rawLine)
			continue
		}

		// Otherwise it's CODE; preserve indentation exactly.
		code = append(code, rawLine)
	}

	flush()
	return out
}

// ------------------------------------------------------------
// Write ./examples/<func>/main.go
// ------------------------------------------------------------

func writeMain(base, funcName string, list []Example) error {
	dir := filepath.Join(base, strings.ToLower(funcName))
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}

	var buf bytes.Buffer

	buf.WriteString("package main\n\n")
	buf.WriteString(`import "github.com/goforj/collection"` + "\n\n")
	buf.WriteString("func main() {\n")

	for _, ex := range list {

		// Header comment for each example
		//buf.WriteString(fmt.Sprintf("\t// Example %d (from %s:%d)\n",
		//	i+1, filepath.Base(ex.File), ex.Line))

		// Optional cleanup hook you had; keep if useful
		if strings.Contains(ex.File, "avg") {
			godump.Dump(ex.Code)
		}

		ex.Code = strings.ReplaceAll(ex.Code, "\n\n\n\t", "\n\n\t")

		for _, line := range strings.Split(ex.Code, "\n") {
			if strings.TrimSpace(line) == "" {
				// Preserve blank lines as blank
				buf.WriteString("\n")
				continue
			}

			// Indent every non-empty line (code or comment) once under func main().
			buf.WriteString("\t" + line + "\n")
		}
	}

	buf.WriteString("}\n")

	return os.WriteFile(filepath.Join(dir, "main.go"), buf.Bytes(), 0o644)
}
