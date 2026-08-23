package main

import (
	"errors"
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
