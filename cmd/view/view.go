package view

import (
	"bufio"
	"fmt"
	"os"
	"school-system/cmd/controller"
	"school-system/cmd/models"
	"school-system/cmd/utils"

	"sort"
	"strings"
)

var inputRead *bufio.Reader = bufio.NewReader(os.Stdin)

type displayAllParams struct {
	displayMsg string
	readInput  interface{}
}

func areThereStudentsRegistered() (bool, error) {
	students, err := controller.AllStudents()

	if err != nil {
		return false, err
	}

	return len(students) > 0, nil
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

func getNextAvailableID() (int, error) {
	studentIDs := make([]int, 0)
	students, err := controller.AllStudents()

	if err != nil {
		return 0, err
	}

	for _, student := range students {
		studentID := student.ID
		studentIDs = append(studentIDs, studentID)
	}

	sort.Ints(studentIDs)
	startID := 1
	for _, ID := range studentIDs {
		if ID-startID == 0 {
			startID++
			continue
		} else {
			return startID, nil
		}
	}

	return len(studentIDs) + 1, nil
}