package main

import "fmt"

func main() {

	var age int
	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)
	if age < 18 {
		fmt.Println("You will not be geting any discount as you are a minor")
	} else if age >= 18 && age <= 60 {
		fmt.Println("You will be getting a discount of 30%")
	} else {
		fmt.Println("You'll be getting a discount of 50%")
	}
}
