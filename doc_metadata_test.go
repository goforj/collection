package collection

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"
	"testing"
)

func TestExportedAPIsHaveDocMetadata(t *testing.T) {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, ".", func(info os.FileInfo) bool {
		return !strings.HasSuffix(info.Name(), "_test.go")
	}, parser.ParseComments)
	if err != nil {
		t.Fatalf("parse package: %v", err)
	}

	pkg := pkgs["collection"]
	if pkg == nil {
		t.Fatalf("collection package not found")
	}

	required := []string{"@group", "@behavior", "@chainable", "@terminal"}

	for _, file := range pkg.Files {
		for _, decl := range file.Decls {
			fn, ok := decl.(*ast.FuncDecl)
			if !ok || !ast.IsExported(fn.Name.Name) {
				continue
			}

			if fn.Doc == nil {
				t.Fatalf("%s is missing doc comment", fn.Name.Name)
			}

			doc := fn.Doc.Text()
			for _, tag := range required {
				if !strings.Contains(doc, tag) {
					t.Fatalf("%s is missing %s metadata", fn.Name.Name, tag)
				}
			}
		}
	}
}
