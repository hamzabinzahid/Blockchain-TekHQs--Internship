package main

import (
	"CLICalculator/Calculator"
	"fmt"
)

func main() {
	var n1, n2 int
	var operation int
	operations := []string{"+", "-", "*", "/"}
	var result int
	var name string

	fmt.Println("WELCOME TO THE CALCULATOR")
	fmt.Print("WHAT SHOULD I CALL YOU : ")
	fmt.Scanln(&name)

	for {

		fmt.Printf("\n------ GREAT %s! LETS MOVE FORWARD ------\n", name)

		fmt.Print("Enter first number : ")
		fmt.Scanln(&n1)

		fmt.Print("Enter second number : ")
		fmt.Scanln(&n2)

		fmt.Println("\n------ GREAT! LETS MOVE ONTO THE OPERATION ------\n")

		fmt.Println("Press 1 to do ADD Operation")
		fmt.Println("Press 2 to do SUBTRACT Operation")
		fmt.Println("Press 3 to do MULTIPLY Operation")
		fmt.Println("Press 4 to do DIVIDE Operation")
		fmt.Print("Choose Your Desired Operation : ")
		fmt.Scanln(&operation)

		if operation != 1 && operation != 2 && operation != 3 && operation != 4 {
			fmt.Println("Invalid operation")
		} else {
			dataToSend := Calculator.Calculator{
				Num1:      n1,
				Num2:      n2,
				Operation: operations[operation-1],
			}

			switch operation {
			case 1:
				result = Calculator.Add(dataToSend)
				fmt.Printf("The result of %d %s %d = %d\n", n1, operations[operation-1], n2, result)
			case 2:
				result = Calculator.Subtract(dataToSend)
				fmt.Printf("The result of %d %s %d = %d\n", n1, operations[operation-1], n2, result)
			case 3:
				result = Calculator.Multiply(dataToSend)
				fmt.Printf("The result of %d %s %d = %d\n", n1, operations[operation-1], n2, result)
			case 4:
				result = Calculator.Divide(dataToSend)
				fmt.Printf("The result of %d %s %d = %d\n", n1, operations[operation-1], n2, result)
			default:
				fmt.Println("Invalid operation")
				// continue
			}
		}

		var choice string
		fmt.Printf("\nGreat %s! Would you like to make another calculation? (Press 1 for YES and any other key for NO) : ", name)
		fmt.Scanln(&choice)

		if choice != "1" {
			fmt.Printf("\n------ IT WAS NICE WORKING WITH YOU! GOOD LUCK👋 ------\n")
			break
		}
	}

}
