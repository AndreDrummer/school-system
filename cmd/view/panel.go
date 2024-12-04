package view

import (
	"fmt"
	"log/slog"
	"os"
	"school-system/cmd/controller"
	"school-system/cmd/models"
	"school-system/cmd/utils"

	"strings"

	"time"
)

func Display() {
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

func AddStudent() {
	fmt.Print("\nEnter student name: ")

	studentName, _ := inputRead.ReadString('\n')
	studentName = strings.TrimSpace(studentName)
	nameIsEmpty := studentName == ""

	if nameIsEmpty {
		utils.ClearConsole()

		fmt.Println(" ** Invalid name **, please try again.")
		for nameIsEmpty {
			fmt.Print("\nEnter student name: ")
			studentName, _ = inputRead.ReadString('\n')
			studentName = strings.TrimSpace(studentName)
			nameIsEmpty = studentName == ""
		}
	}

	newStudent := &models.Student{
		ID:     0,
		Grades: make([]int, 0),
		Name:   studentName,
	}

	if err := controller.AddStudent(*newStudent); err == nil {
		utils.ClearConsole()
		utils.SetSuccessMsg(fmt.Sprintf("Student %v Added!", newStudent.Name))
	} else {
		slog.Error(err.Error())
	}
}

func AddGrade() {
	if ok, _, _ := areThereStudentsRegistered(); ok {
		fmt.Print("What student would you like to add a grade?\n\n")
		studentID := readStudentID()
		student := getStudentByID(studentID)

		if student != nil {
			grade := readGrade()
			if grade >= 0 {
				if err := controller.AddGrade(student, grade); err == nil {
					utils.SetSuccessMsg(fmt.Sprintf("Grade %v added to %v!", grade, student.Name))
				} else {
					slog.Error(err.Error())
				}
			} else {
				utils.ClearConsole()
			}
		}
	}
}

func RemoveStudent() {
	if ok, _, _ := areThereStudentsRegistered(); ok {
		studentID := readStudentID()
		student := getStudentByID(studentID)

		if err := controller.RemoveStudent(studentID); err == nil {
			utils.SetSuccessMsg(fmt.Sprintf("Student %v removed!", student.Name))
		} else {
			slog.Error(err.Error())
		}
	}
}

func CalculateAverage() {
	if ok, _, _ := areThereStudentsRegistered(); ok {
		studentID := readStudentID()
		student := controller.GetStudentByID(studentID)

		if student == nil {
			slog.Error("Student not found")
		} else {
			avg := student.GetAverage()
			utils.PressEnterToGoBack(fmt.Sprintf("\nThe average of %s is %v.\n", student.Name, avg))
		}
	}
}

func CheckPassOrFail() {
	if ok, _, _ := areThereStudentsRegistered(); ok {
		studentID := readStudentID()
		student := controller.GetStudentByID(studentID)

		if student == nil {
			slog.Error("Student not found")
		} else {
			approved := controller.CheckPassOrFail(*student)
			var resultMsg string

			if approved {
				resultMsg = "has been approved! :)"

			} else {
				resultMsg = "has failed :(!"
			}

			utils.PressEnterToGoBack(fmt.Sprintf("\n%s %v.\n", student.Name, resultMsg))
		}
	}
}

func DisplayAll(params *displayAllParams) error {
	var msg string

	if params == nil {
		params = &displayAllParams{}
	}

	if params.displayMsg == "" {
		msg = "\n\nPress Enter to go back. "
	} else {

		msg = params.displayMsg
	}

	if ok, students, err := areThereStudentsRegistered(); ok {
		tempSliceToSort := make([]string, 0)

		if err != nil {
			return err
		}

		for _, v := range students {
			var line string
			if len(v.Grades) == 0 {
				line = fmt.Sprintf("%d - %s -- No grades recorded.", v.ID, v.Name)
				tempSliceToSort = append(tempSliceToSort, line)
			} else {
				line = fmt.Sprintf("%d - %s --Grades-> %v", v.ID, v.Name, v.Grades)
				tempSliceToSort = append(tempSliceToSort, line)
			}

		}
		utils.SortSliceStringByID(tempSliceToSort, "-")

		for _, v := range tempSliceToSort {
			fmt.Println(v)
		}

		if params.readInput != nil {
			fmt.Print(msg)
			fmt.Scanln(params.readInput)
			utils.ClearConsole()
		} else {
			utils.PressEnterToGoBack("")
		}

	}

	return nil
}

func Clear() {
	answer := readYesOrNo("This will delete all data save in the database. Are you sure? ")
	if answer {
		if err := controller.ClearAll(); err == nil {
			utils.SetSuccessMsg("\n** Operação realizada com sucesso! **")
		} else {
			slog.Error(err.Error())
		}
	} else {
		slog.Error("** Operação não realizada! **")
	}
}
