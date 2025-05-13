package main

import (
	"fmt"
	"strings"
)

func countWords(sentence string) map[string]int {
	//Conver the sentence to lower case
	sentence = strings.ToLower(sentence)

	//Split the sentence into words by space
	words := strings.Fields(sentence)

	//Making an empty map to store the words record
	wordsRecord := make(map[string]int)

	for _, word := range words {
		wordsRecord[word]++
	}

	return wordsRecord
}

func filterEvenNumbers(numbers []int) []int {
	newSlice := []int{}

	for _, num := range numbers {
		if num%2 == 0 {
			newSlice = append(newSlice, num)
		}
	}

	return newSlice
}

// Student Grade Calculation
type Student struct {
	Name  string
	Marks []float64
}

func averageGrade(student Student) float64 {
	var sum float64 = 0

	for _, mark := range student.Marks {
		sum += float64(mark)
	}

	average := sum / float64(len(student.Marks))
	return float64(average)
}

func main() {

	// Counting the words in a sentence
	// sentence := "The quick brown fox jumps over the lazy dog and the quick blue hare"
	// modifiedData := countWords(sentence)

	// for word, count := range modifiedData {
	// 	fmt.Printf("%s appears %d times\n", word, count)
	// }

	// // Filtering even numbers from a slice
	// numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// evenNumbers := filterEvenNumbers(numbers)
	// for index, num := range evenNumbers {
	// 	fmt.Printf("Even number at position %d is %d\n", index+1, num)
	// }

	// Student Grade Calculation
	var name string
	var totalSubjects int
	var marks []float64

	fmt.Print("Please enter your name : ")
	fmt.Scanln(&name)

	fmt.Print("Please enter the number of subjects you want to calculate the average of : ")
	fmt.Scanln(&totalSubjects)

	for i := 0; i < totalSubjects; i++ {
		var mark float64
		fmt.Printf("Enter the marks of subject %d : ", i+1)
		fmt.Scanln(&mark)
		marks = append(marks, mark)
	}

	avg := averageGrade(Student{
		Name:  name,
		Marks: marks,
	})

	fmt.Printf("The average grade of student %s is %.2f\n", name, avg)

}
