package collection

import (
	"reflect"
	"testing"
)

func TestChunk_EvenChunks(t *testing.T) {
	c := New([]int{1, 2, 3, 4})

	chunks := c.Chunk(2)

	expected := [][]int{
		{1, 2},
		{3, 4},
	}

	if !reflect.DeepEqual(chunks, expected) {
		t.Fatalf("expected %v, got %v", expected, chunks)
	}
}

func TestChunk_UnevenChunks(t *testing.T) {
	c := New([]int{1, 2, 3, 4, 5})

	chunks := c.Chunk(2)

	expected := [][]int{
		{1, 2},
		{3, 4},
		{5},
	}

	if !reflect.DeepEqual(chunks, expected) {
		t.Fatalf("expected %v, got %v", expected, chunks)
	}
}

func TestChunk_SizeLargerThanCollection(t *testing.T) {
	c := New([]int{1, 2, 3})

	chunks := c.Chunk(10)

	expected := [][]int{
		{1, 2, 3},
	}

	if !reflect.DeepEqual(chunks, expected) {
		t.Fatalf("expected %v, got %v", expected, chunks)
	}
}

func TestChunk_SizeOne(t *testing.T) {
	c := New([]int{1, 2, 3})

	chunks := c.Chunk(1)

	expected := [][]int{
		{1},
		{2},
		{3},
	}

	if !reflect.DeepEqual(chunks, expected) {
		t.Fatalf("expected %v, got %v", expected, chunks)
	}
}

func TestChunk_EmptyCollection(t *testing.T) {
	c := New([]int{})

	chunks := c.Chunk(3)

	// Should return an empty slice, not nil.
	if chunks == nil {
		t.Fatalf("expected empty slice, got nil")
	}

	if len(chunks) != 0 {
		t.Fatalf("expected 0 chunks, got %d", len(chunks))
	}
}

func TestChunk_InvalidSize(t *testing.T) {
	c := New([]int{1, 2, 3})

	chunks := c.Chunk(0)

	if chunks != nil {
		t.Fatalf("expected nil for size <= 0, got %v", chunks)
	}

	chunks = c.Chunk(-5)
	if chunks != nil {
		t.Fatalf("expected nil for size <= 0, got %v", chunks)
	}
}
