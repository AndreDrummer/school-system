package view

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"school-system/cmd/controller"
	"school-system/cmd/models"
	"school-system/cmd/utils"

	"strings"
)

var inputRead *bufio.Reader = bufio.NewReader(os.Stdin)

type displayAllParams struct {
	displayMsg string
	readInput  interface{}
}

func areThereStudentsRegistered() (bool, []models.Student, error) {
	students, err := controller.AllStudents()
	if err != nil {
		slog.Error(err.Error())
		return false, students, err
	}

	if len(students) == 0 {
		utils.PressEnterToGoBack("\n** Empty! No student registered.")
		return false, students, nil
	}

	return len(students) > 0, students, nil
}

func readStudentID() int {
	var studentID int

	for {
		DisplayAll(&displayAllParams{
			displayMsg: "\nEnter the student ID: ",
			readInput:  &studentID,
		})
		_, exists := getStudentByID(studentID)

		if exists {
			break
		}

		utils.ClearConsole()
		fmt.Print("\nPlease enter a valid student ID!\n\n")
	}

	return studentID
}

func readGrade() int {
	var grade int

	fmt.Print("Enter grade (0-100): ")
	_, err := fmt.Scanf("%d", &grade)

	for err != nil || grade < 0 || grade > 100 {
		fmt.Print("Invalid grade. Enter a number between 0 and 100: ")
		_, err = fmt.Scanf("%d", &grade)
	}

	return grade
}

func readYesOrNo(msg string) bool {
	fmt.Print(msg)
	asnwer, _ := inputRead.ReadString('\n')
	asnwer = strings.TrimSpace(asnwer)
	asnwerIsEmpty := asnwer == ""
	acceptableYesAnswers := []string{"YES", "Y", "yes", "y"}
	acceptableNOAnswers := []string{"NO", "N", "no", "n"}

	acceptableAnswers := []string{}
	acceptableAnswers = append(acceptableAnswers, acceptableNOAnswers...)
	acceptableAnswers = append(acceptableAnswers, acceptableYesAnswers...)

	if asnwerIsEmpty || !utils.Contains(acceptableAnswers, asnwer) {
		utils.ClearConsole()

		fmt.Println(" ** Invalid entry **, please try again.")
		for asnwerIsEmpty {
			fmt.Print("\nEnter student name: ")
			asnwer, _ = inputRead.ReadString('\n')
			asnwer = strings.TrimSpace(asnwer)
			asnwerIsEmpty = asnwer == ""
		}
	}

	if asnwer == "YES" || asnwer == "Y" || asnwer == "yes" || asnwer == "y" {
		return true
	} else if asnwer == "NO" || asnwer == "N" || asnwer == "no" || asnwer == "n" {
		return false
	}

	return false
}

func getStudentByID(studentID int) (models.Student, bool) {
	student, exists := controller.GetStudentByID(studentID)
	return student, exists
}
