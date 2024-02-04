package main

import (
	"fmt"
)

func inputNumber(NumberType string) int {
	var number int

	switch NumberType {
	case "student":
		for {
			var err error
			if _, err = fmt.Scanln(&number); err != nil {
				fmt.Println("Error:", err)
			} else {
				if number >= 60 || number <= 0 {
					fmt.Println("Invalid Number, minimum number is 1, Maximum number is 60")

				} else {
					break
				}
			}

		}

	case "grade":
		for {
			var err error
			if _, err = fmt.Scanln(&number); err != nil {
				fmt.Println("Error:", err)
			} else {
				break
			}
		}

	}
	return number

}

func inputGrades(sequence *[]int, student int) {
	// const maxNumbers = 10
	for index := 0; index < student; index++ {

		fmt.Print("Please enter a number for grades: ")
		*sequence = append(*sequence, inputNumber("grade"))
		// fmt.Println("Sequence after adding a number:", *sequence)
		if index == student-1 {
			fmt.Println("You've entered maximum (&student) number of integers.")
		}

	}
}
func gradingStudents(grade int, finalgrades *[]int) {
	// var s int = 5
	if grade < 38 {
		*finalgrades = append(*finalgrades, grade)
	} else {
		if grade%5 < 3 {
			*finalgrades = append(*finalgrades, grade)
		} else {
			*finalgrades = append(*finalgrades, grade+5-grade%5)
		}
	}

}

func main() {

	fmt.Print("Please enter a number of student: ")
	var student = inputNumber("student")
	fmt.Println(student)
	grades := make([]int, 0, student)
	finalgrades := make([]int, 0, student)
	inputGrades(&grades, student)
	fmt.Println(grades)
	// var wg sync.WaitGroup
	for _, grade := range grades {
		// wg.Add(1)
		gradingStudents(grade, &finalgrades)
		// fmt.Println("This is the data", data)
	}
	fmt.Println(finalgrades)

}
