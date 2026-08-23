package main

import (
	"errors"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
	"testing"
)

// TestCountTestOutputIgnoresStderr verifies fresh-cache diagnostics cannot
// corrupt the structured test event stream.
func TestCountTestOutputIgnoresStderr(t *testing.T) {
	stdout := []byte("{\"Action\":\"run\",\"Test\":\"TestOne\"}\n{\"Action\":\"pass\",\"Test\":\"TestOne\"}\n{\"Action\":\"run\",\"Test\":\"TestTwo\"}\n")
	stderr := []byte("go: downloading example.com/dependency v1.0.0\n")

	got, err := countTestOutput(stdout, stderr, nil)
	if err != nil {
		t.Fatalf("count test output: %v", err)
	}
	if got != 2 {
		t.Fatalf("count test output = %d, want 2", got)
	}
}

// TestCountTestOutputReportsCommandFailure verifies command diagnostics remain
// available when the test process fails.
func TestCountTestOutputReportsCommandFailure(t *testing.T) {
	_, err := countTestOutput(nil, []byte("download failed"), errors.New("exit status 1"))
	if err == nil {
		t.Fatal("count test output succeeded, want error")
	}
	if !strings.Contains(err.Error(), "download failed") {
		t.Fatalf("count test output error = %q, want stderr diagnostics", err)
	}
}

// TestCountTestOutputReportsMalformedStream verifies invalid trailing output is
// rejected with enough context to diagnose the generator failure.
func TestCountTestOutputReportsMalformedStream(t *testing.T) {
	stdout := []byte("{\"Action\":\"run\",\"Test\":\"TestOne\"}\nnot-json\n")
	stderr := []byte("go: diagnostic context\n")

	_, err := countTestOutput(stdout, stderr, nil)
	if err == nil {
		t.Fatal("count test output succeeded, want malformed stream error")
	}
	for _, want := range []string{"not-json", "go: diagnostic context"} {
		if !strings.Contains(err.Error(), want) {
			t.Fatalf("count test output error = %q, want %q", err, want)
		}
	}
}

// TestExtractDescriptionKeepsPostMetadataProse verifies that metadata can
// appear before a function's ownership contract without hiding that contract.
func TestExtractDescriptionKeepsPostMetadataProse(t *testing.T) {
	group := commentGroup(t, `// Clone returns an independent copy of the collection.
// @group Construction
// @behavior immutable
// @chainable true
// @terminal false
//
// The returned Slice owns a new backing array, while its elements remain shallow copies.
//
// Use Clone when later mutations must not be shared with the original collection.
//
// Example: clone integers
//
// 	cloned := values.Clone()`)

	const want = "Clone returns an independent copy of the collection.\n\nThe returned Slice owns a new backing array, while its elements remain shallow copies.\n\nUse Clone when later mutations must not be shared with the original collection."
	if got := extractDescription(group); got != want {
		t.Fatalf("extractDescription() = %q, want %q", got, want)
	}
}

// TestRenderAPIIndexLabelsMethods verifies that the index describes both
// package functions and methods.
func TestRenderAPIIndexLabelsMethods(t *testing.T) {
	got := renderAPI([]*FuncDoc{
		{Name: "New", Group: "Construction"},
		{Name: "Clone", Group: "Construction"},
	})
	if !strings.Contains(got, "| Group | Functions and methods |") {
		t.Fatalf("renderAPI() index heading missing functions and methods: %q", got)
	}
	if !strings.Contains(got, "[Clone](#clone) · [New](#new)") {
		t.Fatalf("renderAPI() index does not separate sibling links with a middle dot: %q", got)
	}
}

// commentGroup parses source so tests exercise the same AST representation as
// the README generator.
func commentGroup(t *testing.T, source string) *ast.CommentGroup {
	t.Helper()

	file, err := parser.ParseFile(token.NewFileSet(), "doc.go", source+"\npackage test", parser.ParseComments)
	if err != nil {
		t.Fatalf("parse comment group: %v", err)
	}
	if len(file.Comments) != 1 {
		t.Fatalf("comment groups = %d, want 1", len(file.Comments))
	}
	return file.Comments[0]
}
