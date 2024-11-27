package view

import (
	"fmt"
	"os"
	utils "school-system/cmd/app/Utils"

	"time"
)

func Run() {
	utils.ClearConsole()

	fmt.Print("Welcome to the Student Management System!\n")

	for {
		fmt.Print("\nChoose an option below: \n\n")
		fmt.Println("1 - Add a new Student")
		fmt.Println("2 - Add a grade to a Student")
		fmt.Println("3 - Remove a Student")
		fmt.Println("4 - Calculate average score of a Student")
		fmt.Println("5 - Check if a student passed or failed")
		fmt.Println("6 - Display all students and their grades")
		fmt.Println("7 - Apagar tudo")
		fmt.Println("0 - Exit")
		fmt.Print("\nEnter your choice: ")

		var choice int
		fmt.Scanln(&choice)
		handleChoice(choice)
	}
}

func handleChoice(choice int) {
	utils.ClearConsole()

	switch choice {
	case 0:
		fmt.Printf("\n\n ** Goodbye! **\n\n")
		time.Sleep(750 * time.Millisecond)
		os.Exit(0)
	case 1:
		AddStudent()
	case 2:
		AddGrade()
	case 3:
		RemoveStudent()
	case 4:
		CalculateAverage()
	case 5:
		CheckPassOrFail()
	case 6:
		DisplayAll(nil)
	case 7:
		Clear()
	default:
		fmt.Println("Invalid choice. Try again.")
	}
}
