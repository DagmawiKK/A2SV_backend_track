package main

import (
	"fmt"
)

func calculateAverage(grades []float64) float64 {
	var sum float64
	for _, grade := range grades {
		sum += grade
	}
	return sum / float64(len(grades))
}

func main() {
	var name string
	var numSubjects int

	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)

	for {
		fmt.Print("Enter the number of subjects: ")
		_, err := fmt.Scanln(&numSubjects)
		if err == nil && numSubjects > 0 {
			break
		}
		fmt.Println("Please enter a valid number of subjects.")
	}

	subjects := make([]string, numSubjects)
	grades := make([]float64, numSubjects)

	for i := 0; i < numSubjects; i++ {
		fmt.Printf("\nEnter the name of subject %d: ", i+1)
		fmt.Scanln(&subjects[i])

		var grade float64
		for {
			fmt.Printf("Enter the grade for %s (0 to 100): ", subjects[i])
			_, err := fmt.Scanln(&grade)
			if err == nil && grade >= 0 && grade <= 100 {
				grades[i] = grade
				break
			}
			fmt.Println("Please enter a valid grade between 0 and 100.")
		}
	}

	average := calculateAverage(grades)

	fmt.Printf("\nStudent: %s\n", name)
	fmt.Println("Subject Grades:")

	for i := 0; i < numSubjects; i++ {
		fmt.Printf("- %s: %.2f\n", subjects[i], grades[i])
	}

	fmt.Printf("\nAverage Grade: %.2f\n", average)

	var gradeLetter string
	switch {
	case average >= 90:
		gradeLetter = "A"
	case average >= 80:
		gradeLetter = "B"
	case average >= 70:
		gradeLetter = "C"
	case average >= 60:
		gradeLetter = "D"
	default:
		gradeLetter = "F"
	}

	fmt.Printf("Your overall grade is: %s\n", gradeLetter)
}
