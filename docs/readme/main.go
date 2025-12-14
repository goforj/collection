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
// Parsing
// ------------------------------------------------------------
//

type FuncDoc struct {
	Name  string
	Group string
}

var groupHeader = regexp.MustCompile(`(?i)^\s*@group\s+(.+)$`)

func parseFuncs(root string) ([]FuncDoc, error) {
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

	seen := map[string]FuncDoc{}

	for _, file := range pkg.Files {
		for _, decl := range file.Decls {
			fn, ok := decl.(*ast.FuncDecl)
			if !ok || fn.Doc == nil {
				continue
			}

			name := fn.Name.Name
			if !ast.IsExported(name) {
				continue
			}

			if _, exists := seen[name]; exists {
				continue
			}

			seen[name] = FuncDoc{
				Name:  name,
				Group: extractGroup(fn.Doc),
			}
		}
	}

	out := make([]FuncDoc, 0, len(seen))
	for _, fd := range seen {
		out = append(out, fd)
	}

	return out, nil
}

func extractGroup(group *ast.CommentGroup) string {
	for _, c := range group.List {
		text := strings.TrimSpace(strings.TrimPrefix(c.Text, "//"))
		if m := groupHeader.FindStringSubmatch(text); m != nil {
			return strings.TrimSpace(m[1])
		}
	}
	return "Other"
}

//
// ------------------------------------------------------------
// Rendering
// ------------------------------------------------------------
//

func renderAPI(funcs []FuncDoc) string {
	groups := map[string][]string{}

	for _, fn := range funcs {
		groups[fn.Group] = append(groups[fn.Group], fn.Name)
	}

	groupNames := make([]string, 0, len(groups))
	for g := range groups {
		groupNames = append(groupNames, g)
	}
	sort.Strings(groupNames)

	var buf bytes.Buffer

	for _, group := range groupNames {
		names := groups[group]
		sort.Strings(names)

		buf.WriteString("### " + group + "\n")
		for _, name := range names {
			buf.WriteString("- `" + name + "`\n")
		}
		buf.WriteString("\n")
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
