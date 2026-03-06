# Contributing to ProntoGUI Go Examples

Thank you for your interest in contributing to the ProntoGUI Go Examples project! This repository contains demo applications showcasing the [ProntoGUI](https://github.com/prontogui) library for Go.

## Getting Started

### Prerequisites

- Go 1.23 or later
- A working Go development environment

### Building and Running

Clone the repository and run any demo:

```bash
git clone https://github.com/prontogui/goexamples.git
cd goexamples
go run . -demo bingo
```

Available demos: `frame`, `list1`, `list2`, `table1`, `table2`, `table3`, `bingo`, `textfield`, `timer`, `misc`, `blank`, `snackbar`, `image`, `simple`, `file`.

## How to Contribute

### Reporting Issues

- Use the GitHub issue tracker to report bugs or suggest new demos.
- Include the Go version, OS, and steps to reproduce the issue.

### Submitting Changes

1. Fork the repository.
2. Create a feature branch from `main`:
   ```bash
   git checkout -b my-new-demo
   ```
3. Make your changes (see guidelines below).
4. Ensure the project builds cleanly:
   ```bash
   go build ./...
   go vet ./...
   ```
5. Commit your changes with a clear, descriptive message.
6. Push to your fork and open a pull request against `main`.

### Adding a New Demo

1. Create a new `.go` file in the project root (e.g., `mydemo.go`).
2. Implement a function that returns a `golib.ProntoGUI` instance configured with your demo.
3. Register the demo in `main.go` by adding it to the demo selection logic.
4. Include the standard license header at the top of your file (see existing files for the format).

## Code Guidelines

- Follow standard Go conventions and formatting (`gofmt`).
- Keep demos self-contained and focused on demonstrating specific ProntoGUI features.
- Use clear variable and function names that help readers understand the example.
- Include the BSD 3-Clause license header in all new Go source files.

## License

By contributing, you agree that your contributions will be licensed under the [BSD 3-Clause License](LICENSE) that covers this project.

## Questions?

If you have questions about contributing, feel free to open an issue for discussion.
