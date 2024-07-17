package main

import (
	"bufio"
	"fmt"
	"strings"
)

type TransformFunc func(string) string

func RenderBuffer(input string, transforms ...TransformFunc) string {
	for _, transform := range transforms {
		input = transform(input)
	}
	return input
}

func addLineNumbers(input string) string {
	var result strings.Builder
	scanner := bufio.NewScanner(strings.NewReader(input))
	lineNumber := 1
	for scanner.Scan() {
		result.WriteString(fmt.Sprintf("%d. %s\n", lineNumber, scanner.Text()))
		lineNumber++
	}
	return strings.TrimSuffix(result.String(), "\n")
}

func main() {
	input := "my test string"
	result := RenderBuffer(input, addLineNumbers)
	fmt.Println(result)
}
