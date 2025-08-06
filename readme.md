Go CLI Calculator
A lightweight, cross-platform command-line calculator written in Go, designed to perform basic arithmetic operations with robust error handling.
Features

Basic Arithmetic Operations:
Addition (+)
Subtraction (-)
Multiplication (\*)
Division (/)

Input validation for numbers
Protection against division by zero
Clean, formatted output
Cross-platform compatibility (Windows, macOS, Linux)
Unit tests for reliability

Installation
To get started, ensure you have the following prerequisites:

Go (version 1.16 or higher recommended)

Follow these steps to set up the project:

Clone the repository:git clone https://github.com/yourusername/go-cli-calculator.git

Navigate to the project directory:cd go-cli-calculator

Usage
Run the calculator using the following command syntax:
go run calc.go [number1] [operator] [number2]

Example Commands

Addition:
go run calc.go 5 + 3

Output: 5 + 3 = 8

Subtraction:
go run calc.go 10 - 4

Output: 10 - 4 = 6

Multiplication:
go run calc.go 7 \* 3

Output: 7 \* 3 = 21

Division:
go run calc.go 15 / 5

Output: 15 / 5 = 3

Building an Executable
To create a standalone binary for easier use:

Build the executable:go build -o calculator calc.go

Run the binary directly:./calculator 12 + 5

Output: 12 + 5 = 17

On Windows, use calculator.exe instead of ./calculator.
Error Handling
The calculator handles the following errors with clear messages:

Invalid number inputs (e.g., non-numeric values)
Unsupported operators (only +, -, \*, / are supported)
Division by zero
Incorrect number of arguments

Error Examples

Invalid operator:
go run calc.go 5 x 3

Output: Error: Invalid operator. Supported operators: +, -, \*, /

Division by zero:
go run calc.go 10 / 0

Output: Error: Cannot divide by zero

Missing arguments:
go run calc.go 5 +

Output: Error: Missing arguments. Usage: [number1] [operator] [number2]

Invalid number input:
go run calc.go a + 3

Output: Error: Invalid number input for first argument

Development
Prerequisites

Go 1.16 or higher
Basic familiarity with command-line tools

Running Tests
The project includes unit tests to ensure reliability. Run them with:
go test

This executes the test cases defined in calc_test.go, covering valid operations and error scenarios.
Project Structure
go-cli-calculator/
├── calc.go # Main calculator implementation
├── calc_test.go # Unit tests for the calculator
└── README.md # Project documentation

Contributing
Contributions are welcome! To contribute:

Report Bugs: Open an issue on the GitHub Issues page.
Suggest Features: Start a discussion in the GitHub Discussions.
Submit Pull Requests: Fork the repository, make your changes, and submit a pull request.

Please ensure your contributions:

Adhere to Go coding standards
Include relevant tests in calc_test.go
Update this README.md if necessary

Roadmap
Future enhancements planned for the project:

Support for floating-point numbers with configurable precision
Modulo operation (%) support
Interactive mode for continuous calculations
Calculation history feature

License
This project is licensed under the MIT License. See the LICENSE file for details.
