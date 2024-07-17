package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

type TransformFunc func(*bytes.Buffer) *bytes.Buffer

func RenderBuffer(input *bytes.Buffer, transforms ...TransformFunc) *bytes.Buffer {
	for _, transform := range transforms {
		input = transform(input)
	}
	return input
}

func addLineNumbers(input *bytes.Buffer) *bytes.Buffer {
	output := &bytes.Buffer{}
	scanner := bufio.NewScanner(input)
	lineNumber := 1
	for scanner.Scan() {
		fmt.Fprintf(output, "%d. %s\n", lineNumber, scanner.Text())
		lineNumber++
	}
	return output
}

func main() {
	input := bytes.NewBufferString("my test string")
	result := RenderBuffer(input, addLineNumbers)
	_, err := io.Copy(os.Stdout, result)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing output: %v\n", err)
		os.Exit(1)
	}
}
