package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

type TransformFunc func(io.Reader, io.Writer) error

func RenderBuffer(input io.Reader, transforms ...TransformFunc) (io.Reader, error) {
	var err error
	for _, transform := range transforms {
		var output bytes.Buffer
		err = transform(input, &output)
		if err != nil {
			return nil, err
		}
		input = &output
	}
	return input, nil
}

func addLineNumbers(input io.Reader, output io.Writer) error {
	scanner := bufio.NewScanner(input)
	lineNumber := 1
	for scanner.Scan() {
		_, err := fmt.Fprintf(output, "%d. %s\n", lineNumber, scanner.Text())
		if err != nil {
			return err
		}
		lineNumber++
	}
	return scanner.Err()
}

func main() {
	input := strings.NewReader("my test string")
	result, err := RenderBuffer(input, addLineNumbers)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	_, err = io.Copy(os.Stdout, result)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing output: %v\n", err)
		os.Exit(1)
	}
}
