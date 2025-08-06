# Go CLI Calculator

A lightweight, cross-platform command-line calculator written in Go, designed to perform basic arithmetic operations with robust error handling.

## Features

- **Basic Arithmetic Operations**:
  - Addition (`+`)
  - Subtraction (`-`)
  - Multiplication (`*`)
  - Division (`/`)
- Input validation for numbers
- Protection against division by zero
- Clean, formatted output
- Cross-platform compatibility (Windows, macOS, Linux)
- Unit tests for reliability

## Installation

To get started, ensure you have the following prerequisites:

- [Go](https://go.dev/doc/install) (version 1.16 or higher recommended)

Follow these steps to set up the project:

1. Clone the repository:

   ```bash
   git clone https://github.com/shabbirtoha/cli_calc.git

   ```

2. Navigate to the project directory:

   ```bash
   cd cli_calc

   ```

## Usage

Run the calculator using the following command syntax:

```bash
go run calc.go [number1] [operator] [number2]

```

### Example Commands

- **Addition**:

  ```bash
  go run calc.go 5 + 3

  ```

  Output: `5 + 3 = 8`

- **Subtraction**:

  ```bash
  go run calc.go 10 - 4

  ```

  Output: `10 - 4 = 6`

- **Multiplication**:

  ```bash
  go run calc.go 7 * 3

  ```

  Output: `7 * 3 = 21`

- **Division**:

  ```bash
  go run calc.go 15 / 5

  ```

  Output: `15 / 5 = 3`

## Building an Executable

To create a standalone binary for easier use:

1. Build the executable:

   ```bash
   go build -o calculator calc.go

   ```

2. Run the binary directly:

   ```bash
   ./calculator 12 + 5

   ```

   Output: `12 + 5 = 17`

On Windows, use `calculator.exe` instead of `./calculator`.

## Error Handling

The calculator handles the following errors with clear messages:

- Invalid number inputs (e.g., non-numeric values)
- Unsupported operators (only `+`, `-`, `*`, `/` are supported)
- Division by zero
- Incorrect number of arguments

### Error Examples

- Invalid operator:

  ```bash
  go run calc.go 5 x 3

  ```

  Output: `Error: Invalid operator. Supported operators: +, -, *, /`

- Division by zero:

  ```bash
  go run calc.go 10 / 0

  ```

  Output: `Error: Cannot divide by zero`

- Missing arguments:

  ```bash
  go run calc.go 5 +

  ```

  Output: `Error: Missing arguments. Usage: [number1] [operator] [number2]`

- Invalid number input:

  ```bash
  go run calc.go a + 3

  ```

  Output: `Error: Invalid number input for first argument`

## Development

### Prerequisites

- Go 1.16 or higher
- Basic familiarity with command-line tools

### Running Tests

The project includes unit tests to ensure reliability. Run them with:

```bash
go test

```

This executes the test cases defined in `calc_test.go`, covering valid operations and error scenarios.

## Project Structure

```
go-cli-calculator/
├── calc.go         # Main calculator implementation
├── calc_test.go    # Unit tests for the calculator
└── README.md       # Project documentation

```

## Roadmap

Future enhancements planned for the project:

- Support for floating-point numbers with configurable precision
- Modulo operation (`%`) support
- Interactive mode for continuous calculations
- Calculation history feature
