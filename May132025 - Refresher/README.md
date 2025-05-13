# Golang Practice – Daily Learning Summary(May 13 2025)
This project documents my learning and practice of the Go (Golang) programming language. The focus was on understanding core concepts from beginner to intermediate level through hands-on examples and small coding exercises.

---

## Topics Covered

### 1. User Input using fmt.Scanln
- Learned how to take input from the user such as name, number of subjects, and their marks.
```go
var name string
fmt.Scanln(&name)
```

---

### 2. Working with Slices
- Used slices to store dynamic lists of data such as multiple marks.
```go
var marks []float64
marks = append(marks, 89.5)
```

---

### 3. Looping in Go
- Practiced different types of loops like for loop, for-range loop, and while-style loop.
```go
for i := 0; i < totalSubjects; i++ {
    fmt.Scanln(&mark)
    marks = append(marks, mark)
}

for _, value := range marks {
    fmt.Println(value)
}
```

---

### 4. If/Else Conditions
- Implemented logic using conditional statements to calculate grades based on average marks.
```go
if avg >= 90 {
    fmt.Println("Grade: A")
} else if avg >= 75 {
    fmt.Println("Grade: B")
} else {
    fmt.Println("Grade: C")
}
```

---

### 5. Structs
- Defined custom data types using structs, for example a Student struct to store name and marks.
```go
type Student struct {
    Name  string
    Marks []float64
}
```

---

### 6. Functions
- Created reusable functions to perform operations like calculating the average of student marks.
```go
func averageGrade(student Student) float64 {
    var sum float64
    for _, mark := range student.Marks {
        sum += mark
    }
    return sum / float64(len(student.Marks))
}
```

---

### 7. Maps in Go
- Used maps to store and count the frequency of words from a sentence.
```go
words := strings.Fields(sentence)
wordCount := make(map[string]int)

for _, word := range words {
    wordCount[word]++
}
```

---


## Key Takeaways

- Structs in Go are used in place of classes to define structured data.
- Pointers are required with Scanln and Decoder to store data in variables.
- Maps automatically initialize missing keys with a zero value.
- Looping and functions helped organize logic in a clean and readable way.

---

This summary reflects everything I practiced and understood during today's session.