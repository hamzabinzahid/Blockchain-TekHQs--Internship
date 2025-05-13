package main

import "fmt"

func main() {
	// For Loop
	num := 7
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", num, i, num*i)
	}

	// While Loop Using For Keyword
	var number int = 1
	for number <= 100 {
		number = number * 2
	}
	fmt.Printf("The number at th end is %d\n", number)

	// Using For Range Loop
	fruits := []string{"Apple", "Banana", "Pineapple", "Melon"}
	for index, fruit := range fruits {
		fmt.Printf("Fruit at index %d is %s\n", index, fruit)
	}

}
