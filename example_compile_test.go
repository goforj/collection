//go:build ignore
// +build ignore

package collection_test

import (
	"bytes"
	"fmt"
	"go/parser"
	"go/token"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

type ExampleCase struct {
	Name   string
	Code   string
	Output string // optional — for future matching
}

func TestDocExamplesCompile(t *testing.T) {
	fset := token.NewFileSet()

	// Parse THIS module's source directory, not the test directory.
	pkgs, err := parser.ParseDir(fset, ".", nil, parser.ParseComments)
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}

	var cases []ExampleCase

	for filename, file := range pkgs["collection"].Files {
		for _, cg := range file.Comments {
			extracted := extractFromComment(filename, cg.Text())
			cases = append(cases, extracted...)
		}
	}

	if len(cases) == 0 {
		t.Fatalf("no doc examples found to test")
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			ok, stderr := tryCompileExternal(tc.Code)
			if !ok {
				t.Fatalf("Compilation failed:\nExample Code:\n%s\n\nCompiler error:\n%s",
					indent(tc.Code), indent(stderr))
			}
		})
	}
}

func extractFromComment(filename, comment string) []ExampleCase {
	var out []ExampleCase

	lines := strings.Split(comment, "\n")
	var block []string
	in := false
	startLine := 0

	for i, raw := range lines {
		line := strings.TrimSpace(raw)

		if strings.HasPrefix(line, "Example:") {
			// close previous block if any
			if in && len(block) > 0 {
				out = append(out, makeCase(filename, startLine, block))
			}
			in = true
			block = nil
			startLine = i + 1
			continue
		}

		if !in {
			continue
		}

		// Strip `//`
		trim := strings.TrimSpace(strings.TrimPrefix(raw, "//"))

		if trim == "" || strings.HasPrefix(trim, "//") {
			if len(block) > 0 {
				out = append(out, makeCase(filename, startLine, block))
			}
			in = false
			block = nil
			continue
		}

		block = append(block, trim)
	}

	// trailing block
	if in && len(block) > 0 {
		out = append(out, makeCase(filename, startLine, block))
	}

	return out
}

func makeCase(filename string, line int, block []string) ExampleCase {
	return ExampleCase{
		Name: filepath.Base(filename) + ":" + strconv.Itoa(line),
		Code: strings.Join(block, "\n"),
	}
}

// Compile as an EXTERNAL user program.
func tryCompileExternal(src string) (bool, string) {
	tmp := fmt.Sprintf(`package main
import "github.com/goforj/collection"
func main() {
%s
}`, src)

	f, err := os.CreateTemp("", "example_*.go")
	if err != nil {
		return false, "cannot create temp file"
	}
	defer os.Remove(f.Name())
	f.WriteString(tmp)
	f.Close()

	// Use go build — safe, normal, resolves modules.
	cmd := exec.Command("go", "build", "-o", os.DevNull, f.Name())

	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err = cmd.Run()

	return err == nil, stderr.String()
}

func indent(s string) string {
	lines := strings.Split(s, "\n")
	for i := range lines {
		lines[i] = "    " + lines[i]
	}
	return strings.Join(lines, "\n")
}
