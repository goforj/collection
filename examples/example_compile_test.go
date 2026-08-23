package examples

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

var errGoBuildFailed = errors.New("go build failed")

// TestExamplesBuild verifies that every generated example compiles without its ignore build tag.
func TestExamplesBuild(t *testing.T) {
	t.Parallel()

	examplesDir := "."

	entries, err := os.ReadDir(examplesDir)
	if err != nil {
		t.Fatalf("cannot read examples directory: %v", err)
	}

	for _, e := range entries {
		if !e.IsDir() {
			continue
		}

		name := e.Name()
		path := filepath.Join(examplesDir, name)
		mainFile := filepath.Join(path, "main.go")
		if _, err := os.Stat(mainFile); errors.Is(err, os.ErrNotExist) {
			continue
		}
		src, err := os.ReadFile(mainFile)
		if err != nil {
			t.Fatalf("read example %q: %v", name, err)
		}
		if !bytes.Contains(src, []byte("github.com/goforj/collection/v4")) {
			t.Fatalf("example %q does not import the v4 module path", name)
		}

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			if err := buildExampleWithoutTags(path); err != nil {
				t.Fatalf("example %q failed to build:\n%s", name, err)
			}
		})
	}
}

// abs resolves paths because Go overlay replacements require absolute filenames.
func abs(p string) string {
	a, err := filepath.Abs(p)
	if err != nil {
		panic(err)
	}
	return a
}

// buildExampleWithoutTags compiles an ignored generated example through a temporary overlay.
func buildExampleWithoutTags(exampleDir string) error {
	orig := filepath.Join(exampleDir, "main.go")

	src, err := os.ReadFile(orig)
	if err != nil {
		return fmt.Errorf("read main.go: %w", err)
	}

	clean := stripBuildTags(src)

	tmpDir, err := os.MkdirTemp("", "example-overlay-*")
	if err != nil {
		return fmt.Errorf("mkdir temp dir: %w", err)
	}
	defer os.RemoveAll(tmpDir)

	tmpFile := filepath.Join(tmpDir, "main.go")
	err = os.WriteFile(tmpFile, clean, 0o600)
	if err != nil {
		return fmt.Errorf("write temp main.go: %w", err)
	}

	overlay := map[string]any{
		"Replace": map[string]string{
			abs(orig): abs(tmpFile),
		},
	}

	overlayJSON, err := json.Marshal(overlay)
	if err != nil {
		return fmt.Errorf("marshal overlay: %w", err)
	}

	overlayPath := filepath.Join(tmpDir, "overlay.json")
	err = os.WriteFile(overlayPath, overlayJSON, 0o600)
	if err != nil {
		return fmt.Errorf("write overlay: %w", err)
	}

	cmd := exec.Command(
		"go", "build",
		"-overlay", overlayPath,
		"-o", os.DevNull,
		"./"+filepath.Base(exampleDir),
	)
	cmd.Dir = filepath.Dir(exampleDir)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("%w: %s", errGoBuildFailed, stderr.String())
	}

	return nil
}

// stripBuildTags removes the leading constraints that keep generated commands out of normal builds.
func stripBuildTags(src []byte) []byte {
	lines := strings.Split(string(src), "\n")

	i := 0
	for i < len(lines) {
		line := strings.TrimSpace(lines[i])

		if strings.HasPrefix(line, "//go:build") ||
			strings.HasPrefix(line, "// +build") ||
			line == "" {
			i++
			continue
		}

		break
	}

	return []byte(strings.Join(lines[i:], "\n"))
}
