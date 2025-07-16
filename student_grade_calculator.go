package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Enter your name: ")
    name, _ := reader.ReadString('\n')
    name = strings.TrimSpace(name)

    var numSubjects int
    for {
        fmt.Print("Enter the number of subjects: ")
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)
        n, err := strconv.Atoi(input)
        if err == nil && n > 0 {
            numSubjects = n
            break
        }
        fmt.Println("Please enter a valid positive integer.")
    }

    subjects := make([]string, numSubjects)
    grades := make([]float64, numSubjects)

    for i := 0; i < numSubjects; i++ {
        fmt.Printf("Enter name of subject #%d: ", i+1)
        subject, _ := reader.ReadString('\n')
        subject = strings.TrimSpace(subject)
        subjects[i] = subject

        for {
            fmt.Printf("Enter grade for %s (0-100): ", subject)
            gradeInput, _ := reader.ReadString('\n')
            gradeInput = strings.TrimSpace(gradeInput)
            grade, err := strconv.ParseFloat(gradeInput, 64)
            if err == nil && grade >= 0 && grade <= 100 {
                grades[i] = grade
                break
            }
            fmt.Println("Invalid grade. Please enter a number between 0 and 100.")
        }
    }

    avg := calculateAverage(grades)

    fmt.Printf("\nStudent Name: %s\n", name)
    fmt.Println("Subject Grades:")
    for i := 0; i < numSubjects; i++ {
        fmt.Printf("  %s: %.2f\n", subjects[i], grades[i])
    }
    fmt.Printf("Average Grade: %.2f\n", avg)
}

func calculateAverage(grades []float64) float64 {
    sum := 0.0
    for _, grade := range grades {
        sum += grade
    }
    return sum / float64(len(grades))
}