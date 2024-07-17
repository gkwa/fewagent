package main

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestRenderBufferWithAddLineNumbers(t *testing.T) {
	input := strings.NewReader("line1\nline2\nline3")
	expected := "1. line1\n2. line2\n3. line3\n"

	result, err := RenderBuffer(input, addLineNumbers)
	if err != nil {
		t.Fatalf("RenderBuffer returned an error: %v", err)
	}

	var output bytes.Buffer
	_, err = io.Copy(&output, result)
	if err != nil {
		t.Fatalf("Error copying result: %v", err)
	}

	if output.String() != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, output.String())
	}
}

func TestRenderBufferWithNoTransforms(t *testing.T) {
	input := strings.NewReader("test content")
	expected := "test content"

	result, err := RenderBuffer(input)
	if err != nil {
		t.Fatalf("RenderBuffer returned an error: %v", err)
	}

	var output bytes.Buffer
	_, err = io.Copy(&output, result)
	if err != nil {
		t.Fatalf("Error copying result: %v", err)
	}

	if output.String() != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, output.String())
	}
}
