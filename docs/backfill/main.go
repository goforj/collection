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
	fmt.Println("✔ @behavior, @chainable, and @terminal annotations backfilled")
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
			if !ok || fn.Doc == nil || !ast.IsExported(fn.Name.Name) {
				continue
			}

			// Chainable/terminal backfill (boolean) – always refresh to keep canonical.
			stripChainable(fn.Doc) // legacy cleanup
			stripFluent(fn.Doc)
			stripTerminal(fn.Doc)
			insertChainable(fn.Doc, inferChainable(fn))
			insertTerminal(fn.Doc, inferTerminal(fn))
			changed = true

			// Respect explicit annotation
			if hasBehavior(fn.Doc) {
				continue
			}

			behavior, ok := inferBehavior(fn)
			if !ok {
				continue // ambiguous → leave for human review
			}

			insertBehavior(fn.Doc, behavior)
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
// Behavior inference
// ------------------------------------------------------------
//

func inferBehavior(fn *ast.FuncDecl) (string, bool) {
	name := fn.Name.Name

	// 1. Known mutators (authoritative)
	if knownMutators[name] {
		return "mutable", true
	}

	// 2. Methods returning *Collection[T] default to immutable
	if returnsCollection(fn) {
		if mutatesReceiverItems(fn) {
			return "mutable", true
		}
		return "immutable", true
	}

	// 3. Non-collection return types → readonly
	if fn.Type.Results != nil {
		return "readonly", true
	}

	return "", false
}

func inferChainable(fn *ast.FuncDecl) string {
	if fn.Type.Results == nil || len(fn.Type.Results.List) != 1 {
		return "false"
	}

	if returnsCollectionLike(fn) {
		return "true"
	}

	return "false"
}

func inferTerminal(fn *ast.FuncDecl) string {
	if inferChainable(fn) == "true" {
		return "false"
	}
	return "true"
}

var knownMutators = map[string]bool{
	"Push":     true,
	"Pop":      true,
	"PopN":     true,
	"Append":   true,
	"Prepend":  true,
	"Merge":    true,
	"Multiply": true,
	"Reverse":  true,
	"Shuffle":  true,
	"Sort":     true,
	"Transform": true,
}

func returnsCollection(fn *ast.FuncDecl) bool {
	if fn.Type.Results == nil || len(fn.Type.Results.List) != 1 {
		return false
	}

	result := fn.Type.Results.List[0].Type
	star, ok := result.(*ast.StarExpr)
	if !ok {
		return false
	}

	sel, ok := star.X.(*ast.IndexExpr)
	if !ok {
		return false
	}

	ident, ok := sel.X.(*ast.Ident)
	return ok && ident.Name == "Collection"
}

func returnsCollectionLike(fn *ast.FuncDecl) bool {
	if fn.Type.Results == nil || len(fn.Type.Results.List) == 0 {
		return false
	}

	if len(fn.Type.Results.List) != 1 {
		return returnsCollectionResult(fn.Type.Results.List[len(fn.Type.Results.List)-1].Type)
	}

	return returnsCollectionResult(fn.Type.Results.List[0].Type)
}

func returnsCollectionResult(expr ast.Expr) bool {
	star, ok := expr.(*ast.StarExpr)
	if !ok {
		return false
	}

	sel, ok := star.X.(*ast.IndexExpr)
	if !ok {
		return false
	}

	ident, ok := sel.X.(*ast.Ident)
	if !ok {
		return false
	}

	return ident.Name == "Collection" || ident.Name == "NumericCollection"
}

//
// ------------------------------------------------------------
// AST mutation detection (veto)
// ------------------------------------------------------------
//

func mutatesReceiverItems(fn *ast.FuncDecl) bool {
	if fn.Recv == nil || len(fn.Recv.List) == 0 {
		return false
	}

	if len(fn.Recv.List[0].Names) == 0 {
		return false
	}

	receiver := fn.Recv.List[0].Names[0].Name

	mutates := false

	ast.Inspect(fn.Body, func(n ast.Node) bool {
		switch x := n.(type) {

		// c.items = ...
		case *ast.AssignStmt:
			for _, lhs := range x.Lhs {
				if isReceiverItems(lhs, receiver) {
					mutates = true
					return false
				}
			}

		// append(c.items, ...)
		case *ast.CallExpr:
			if isMutatingCall(x, receiver) {
				mutates = true
				return false
			}
		}
		return true
	})

	return mutates
}

func isReceiverItems(expr ast.Expr, receiver string) bool {
	sel, ok := expr.(*ast.SelectorExpr)
	if !ok {
		return false
	}

	x, ok := sel.X.(*ast.Ident)
	return ok && x.Name == receiver && sel.Sel.Name == "items"
}

func isMutatingCall(call *ast.CallExpr, receiver string) bool {
	ident, ok := call.Fun.(*ast.Ident)
	if !ok {
		return false
	}

	if ident.Name != "append" && ident.Name != "copy" {
		return false
	}

	if len(call.Args) == 0 {
		return false
	}

	return isReceiverItems(call.Args[0], receiver)
}

//
// ------------------------------------------------------------
// Doc helpers
// ------------------------------------------------------------
//

func hasBehavior(doc *ast.CommentGroup) bool {
	for _, c := range doc.List {
		if strings.Contains(c.Text, "@behavior") {
			return true
		}
	}
	return false
}

func insertBehavior(doc *ast.CommentGroup, behavior string) {
	for i, c := range doc.List {
		text := strings.TrimSpace(strings.TrimPrefix(c.Text, "//"))
		if strings.HasPrefix(text, "Example:") || text == "" {
			doc.List = append(
				doc.List[:i],
				append([]*ast.Comment{
					{Text: fmt.Sprintf("// @behavior %s", behavior)},
				}, doc.List[i:]...)...,
			)
			return
		}
	}

	doc.List = append(doc.List, &ast.Comment{
		Text: fmt.Sprintf("// @behavior %s", behavior),
	})
}

func insertChainable(doc *ast.CommentGroup, val string) {
	for i, c := range doc.List {
		text := strings.TrimSpace(strings.TrimPrefix(c.Text, "//"))
		if strings.HasPrefix(text, "Example:") || text == "" {
			doc.List = append(
				doc.List[:i],
				append([]*ast.Comment{
					{Text: fmt.Sprintf("// @chainable %s", val)},
				}, doc.List[i:]...)...,
			)
			return
		}
	}

	doc.List = append(doc.List, &ast.Comment{
		Text: fmt.Sprintf("// @chainable %s", val),
	})
}

func insertTerminal(doc *ast.CommentGroup, val string) {
	for i, c := range doc.List {
		text := strings.TrimSpace(strings.TrimPrefix(c.Text, "//"))
		if strings.HasPrefix(text, "Example:") || text == "" {
			doc.List = append(
				doc.List[:i],
				append([]*ast.Comment{
					{Text: fmt.Sprintf("// @terminal %s", val)},
				}, doc.List[i:]...)...,
			)
			return
		}
	}

	doc.List = append(doc.List, &ast.Comment{
		Text: fmt.Sprintf("// @terminal %s", val),
	})
}

func stripChainable(doc *ast.CommentGroup) {
	out := doc.List[:0]
	for _, c := range doc.List {
		if strings.Contains(c.Text, "@chainable") {
			continue
		}
		out = append(out, c)
	}
	doc.List = out
}

func stripFluent(doc *ast.CommentGroup) {
	out := doc.List[:0]
	for _, c := range doc.List {
		if strings.Contains(c.Text, "@fluent") {
			continue
		}
		out = append(out, c)
	}
	doc.List = out
}

func stripTerminal(doc *ast.CommentGroup) {
	out := doc.List[:0]
	for _, c := range doc.List {
		if strings.Contains(c.Text, "@terminal") {
			continue
		}
		out = append(out, c)
	}
	doc.List = out
}

//
// ------------------------------------------------------------
// File IO
// ------------------------------------------------------------
//

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
