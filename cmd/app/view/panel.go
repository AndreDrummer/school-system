package view

import (
	"fmt"
	"os"
	utils "school-system/cmd/app/Utils"
	"school-system/cmd/app/controller"

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
		controller.AddStudent()
	case 2:
		controller.AddGrade()
	case 3:
		controller.RemoveStudent()
	case 4:
		controller.CalculateAverage()
	case 5:
		controller.CheckPassOrFail()
	case 6:
		controller.DisplayAll(nil)
	case 7:
		controller.Clear()
	default:
		fmt.Println("Invalid choice. Try again.")
	}
}
