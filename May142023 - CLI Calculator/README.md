# Go CLI Calculator

This project is a simple **Command Line Calculator** built using **Golang** to practice core concepts such as:

- Structs
- Functions
- Packages and Imports
- Creating and using custom modules
- User input handling and validation
- Looping and control flow with conditions and switches

---

## Features
- Take two numbers as input from the user
- Allow the user to choose an operation:
  - Addition
  - Subtraction
  - Multiplication
  - Division
- Use a **struct** to hold operation-related data
- Perform the calculation using dedicated **functions** in a separate module (`Calculator`)
- Display the result of the operation
- Ask the user if they want to perform another transaction
- Handle invalid operations or incorrect inputs gracefully

---

## Code Structure

```plaintext
CLICalculator/
│
├── Calculator/             # Custom package for arithmetic operations
│   └── operations.go       # Contains Calculator struct and operation functions
│
├── go.mod                  # Go module file
└── main.go                 # Main CLI logic and user interaction
