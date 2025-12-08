package collection

import (
	"bytes"
	"os"
	"os/exec"
	"testing"
)

func TestDump_ReturnsCollection(t *testing.T) {
	c := New([]int{1, 2, 3})

	out := c

	if out.Items()[0] != 1 || out.Items()[2] != 3 {
		t.Fatalf("Dump() should return original collection; got %v", out.Items())
	}
}

func TestDd_TriggersExit(t *testing.T) {
	if os.Getenv("TEST_DD") == "1" {
		c := New([]int{10, 20})
		c.Dd() // should exit(1)
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestDd_TriggersExit")
	cmd.Env = append(os.Environ(), "TEST_DD=1")

	err := cmd.Run()

	if err == nil {
		t.Fatalf("Dd() should have exited the subprocess, but it did not")
	}

	if exit, ok := err.(*exec.ExitError); ok {
		if exit.ExitCode() != 1 {
			t.Fatalf("Dd() exit code = %d, want 1", exit.ExitCode())
		}
	} else {
		t.Fatalf("unexpected error type: %v", err)
	}
}

func TestDumpStr_ReturnsString(t *testing.T) {
	c := New([]int{5, 6, 7})

	out := c.DumpStr()

	if out == "" {
		t.Fatalf("DumpStr() returned empty string")
	}

	if !bytes.Contains([]byte(out), []byte("5")) {
		t.Fatalf("DumpStr() missing expected content: %s", out)
	}
}

func TestDdStr_ReturnsStringAndTriggersExit(t *testing.T) {
	if os.Getenv("TEST_DDSTR") == "1" {
		c := New([]int{9})
		_ = c.DdStr() // should exit(1)
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestDdStr_ReturnsStringAndTriggersExit")
	cmd.Env = append(os.Environ(), "TEST_DDSTR=1")

	err := cmd.Run()

	if err == nil {
		t.Fatalf("DdStr() should have exited but did not")
	}

	if exit, ok := err.(*exec.ExitError); ok {
		if exit.ExitCode() != 1 {
			t.Fatalf("DdStr() exit code = %d, want 1", exit.ExitCode())
		}
	} else {
		t.Fatalf("unexpected error type: %v", err)
	}
}

func TestDump_ReturnsSameCollection(t *testing.T) {
	c := New([]int{1, 2, 3})
	out := c.Dump()

	if out != c {
		t.Fatalf("Dump() should return the same collection")
	}
}

func TestDdStr_ReturnsOutputAndTriggersExit(t *testing.T) {
	called := false

	// Override exit behavior for test
	origExit := exitFunc
	exitFunc = func(v interface{}) { called = true }
	defer func() { exitFunc = origExit }()

	c := New([]int{5})
	out := c.DdStr()

	if out == "" {
		t.Fatalf("DdStr() should return non-empty dump output")
	}

	if !called {
		t.Fatalf("DdStr() did not trigger exitFunc")
	}
}
