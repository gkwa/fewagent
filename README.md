# fewagent

FewAgent is a Go application that demonstrates a flexible buffer rendering system with customizable transformations.

## Purpose

The main purpose of this application is to showcase a function `RenderBuffer` that can apply a series of transformations to an input stream. It includes an example transformation `addLineNumbers` that adds line numbers to each line of the input.

## Example Usage

```bash
go run main.go
```

This will run the application with a hardcoded input string "my test string" and apply the `addLineNumbers` transformation.

Expected output:
```
1. my test string
```

## How to Run

1. Ensure you have Go installed on your system.
2. Clone this repository:
   ```bash
   git clone https://github.com/yourusername/fewagent.git
   cd fewagent
   ```
3. Run the application:
   ```bash
   go run main.go
   ```

## How to Test

To run the tests:

```bash
go test ./...
```

This will run two tests:
1. `TestRenderBufferWithAddLineNumbers`: Tests the `addLineNumbers` transformation.
2. `TestRenderBufferWithNoTransforms`: Tests `RenderBuffer` with no transformations applied.

## Install fewagent

on macOS/Linux:
```bash
brew install gkwa/homebrew-tools/fewagent
```

on Windows:
```powershell
TBD
```

## Extending the Application

You can extend this application by adding more transformation functions. To do so:

1. Define a new function that conforms to the `TransformFunc` type.
2. Add your new transformation to the `RenderBuffer` call in the `main` function.

Example:
```go
func upperCase(input io.Reader, output io.Writer) error {
    // Implementation here
}

// In main():
result, err := RenderBuffer(input, addLineNumbers, upperCase)
```

This flexibility allows for easy chaining of multiple transformations on the input stream.
```
