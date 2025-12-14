//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if err := run(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println("✔ @group annotations added where possible")
}

func run() error {
	root, err := findRoot()
	if err != nil {
		return err
	}

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
		return err
	}

	pkg, ok := pkgs["collection"]
	if !ok {
		return fmt.Errorf(`package "collection" not found`)
	}

	for filename, file := range pkg.Files {
		changed := false

		for _, decl := range file.Decls {
			fn, ok := decl.(*ast.FuncDecl)
			if !ok || fn.Doc == nil {
				continue
			}

			if !ast.IsExported(fn.Name.Name) {
				continue
			}

			if hasGroup(fn.Doc) {
				continue
			}

			group, ok := inferredGroups[fn.Name.Name]
			if !ok {
				continue // ambiguous, leave for human review
			}

			insertGroup(fn.Doc, group)
			changed = true
		}

		if changed {
			if err := writeFile(fset, filename, file); err != nil {
				return err
			}
		}
	}

	return nil
}

//
// ------------------------------------------------------------
// Group inference
// ------------------------------------------------------------
//

var inferredGroups = map[string]string{
	// Querying
	"First":      "Querying",
	"Last":       "Querying",
	"At":         "Querying",
	"IndexWhere": "Querying",
	"FindWhere":  "Querying",
	"FirstWhere": "Querying",
	"LastWhere":  "Querying",
	"Contains":   "Querying",
	"Any":        "Querying",
	"All":        "Querying",
	"None":       "Querying",
	"IsEmpty":    "Querying",

	// Slicing
	"Take":     "Slicing",
	"TakeLast": "Slicing",
	"Skip":     "Slicing",
	"SkipLast": "Slicing",
	"Chunk":    "Slicing",
	"Pop":      "Slicing",
	"PopN":     "Slicing",

	// Ordering
	"Sort":    "Ordering",
	"Reverse": "Ordering",
	"Shuffle": "Ordering",
	"Before":  "Ordering",
	"After":   "Ordering",

	// Transformation
	"Map":       "Transformation",
	"MapTo":     "Transformation",
	"Each":      "Transformation",
	"Transform": "Transformation",
	"Pluck":     "Transformation",
	"Pipe":      "Transformation",
	"Tap":       "Transformation",
	"Times":     "Transformation",
	"Append":    "Transformation",
	"Prepend":   "Transformation",
	"Push":      "Transformation",
	"Concat":    "Transformation",
	"Merge":     "Transformation",
	"Multiply":  "Transformation",

	// Aggregation
	"Reduce":       "Aggregation",
	"Count":        "Aggregation",
	"CountBy":      "Aggregation",
	"CountByValue": "Aggregation",
	"Sum":          "Aggregation",
	"Avg":          "Aggregation",
	"Min":          "Aggregation",
	"Max":          "Aggregation",
	"Median":       "Aggregation",
	"Mode":         "Aggregation",

	// Maps
	"FromMap": "Maps",
	"ToMap":   "Maps",
	"ToMapKV": "Maps",

	// Construction / Core
	"New":        "Construction",
	"NewNumeric": "Construction",
	"Items":      "Construction",

	// Debugging / Output
	"Dump":         "Debugging",
	"DumpStr":      "Debugging",
	"Dd":           "Debugging",
	"ToJSON":       "Debugging",
	"ToPrettyJSON": "Debugging",

	// Grouping / Sets
	"GroupBy":  "Grouping",
	"Unique":   "Set Operations",
	"UniqueBy": "Set Operations",
}

//
// ------------------------------------------------------------
// AST helpers
// ------------------------------------------------------------
//

func hasGroup(doc *ast.CommentGroup) bool {
	for _, c := range doc.List {
		if strings.Contains(c.Text, "@group") {
			return true
		}
	}
	return false
}

func insertGroup(doc *ast.CommentGroup, group string) {
	for i, c := range doc.List {
		text := strings.TrimSpace(strings.TrimPrefix(c.Text, "//"))

		// Insert before Example: or blank separator
		if strings.HasPrefix(text, "Example:") || text == "" {
			doc.List = append(
				doc.List[:i],
				append(
					[]*ast.Comment{
						{Text: fmt.Sprintf("// @group %s", group)},
					},
					doc.List[i:]...,
				)...,
			)
			return
		}
	}

	// Fallback: append to end of doc block
	doc.List = append(doc.List, &ast.Comment{
		Text: fmt.Sprintf("// @group %s", group),
	})
}

func writeFile(fset *token.FileSet, filename string, file *ast.File) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	cfg := &printer.Config{
		Mode:     printer.TabIndent | printer.UseSpaces,
		Tabwidth: 8,
	}

	return cfg.Fprint(f, fset, file)
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
