package controller

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	utils "school-system/cmd/app/Utils"
	"school-system/cmd/app/models"

	"sort"
	"strconv"
	"strings"
)

var inputRead *bufio.Reader = bufio.NewReader(os.Stdin)

var SystemInstance = &models.System{
	Students:            make(map[int]*models.Student),
	StudentsQty:         0,
	MinimumPassingGrade: 60,
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
		ID:     getNextAvailableID(),
		Grades: make([]int, 0),
		Name:   studentName,
	}

	if SystemInstance.AddStudent(newStudent) {
		utils.ClearConsole()
		utils.SetSuccessMsg(fmt.Sprintf("Student %v Added!", newStudent.Name))
	}

}

func AddGrade() {
	if areThereStudentsRegistered() {
		fmt.Print("What student would you like to add a grade?\n\n")
		studentID := readStudentID()
		student, studentExists := getStudentByID(studentID)

		if studentExists {
			grade := readGrade()

			if grade >= 0 {
				if SystemInstance.AddGrade(studentID, grade) {
					utils.SetSuccessMsg(fmt.Sprintf("Grade %v added to %v!", grade, student.Name))
				}
			} else {
				utils.ClearConsole()
			}
		}
	} else {
		utils.PressEnterToGoBack("\n** Empty! No student registered.")
	}
}

func RemoveStudent() {
	if areThereStudentsRegistered() {
		studentID := readStudentID()
		student, studentExists := getStudentByID(studentID)

		if studentExists && SystemInstance.RemoveStudent(studentID) {
			utils.SetSuccessMsg(fmt.Sprintf("Student %v removed!", student.Name))
		}
	} else {
		utils.PressEnterToGoBack("\n** Empty! No student registered.")
	}
}

func CalculateAverage() {
	if areThereStudentsRegistered() {
		studentID := readStudentID()
		avg, err := SystemInstance.CalculateAverage(studentID)

		if err != nil {
			slog.Error(err.Error())
		} else {
			student := SystemInstance.Students[studentID]
			utils.PressEnterToGoBack(fmt.Sprintf("\nThe average of %s is %v.\n", student.Name, avg))
		}
	} else {
		utils.PressEnterToGoBack("\n** Empty! No student registered.")
	}
}

func CheckPassOrFail() {
	if areThereStudentsRegistered() {
		studentID := readStudentID()

		approved := SystemInstance.CheckPassOrFail(studentID)
		var resultMsg string

		if approved {
			resultMsg = "has been approved! :)"

		} else {
			resultMsg = "has failed :(!"
		}

		student := SystemInstance.Students[studentID]
		utils.PressEnterToGoBack(fmt.Sprintf("\n%s %v.\n", student.Name, resultMsg))

	} else {
		utils.PressEnterToGoBack("\n** Empty! No student registered.")
	}
}

func DisplayAll(params *displayAllParams) {
	var msg string

	if params == nil {
		params = &displayAllParams{}
	}

	if params.displayMsg == "" {
		msg = "\n\nPress Enter to go back. "
	} else {

		msg = params.displayMsg
	}

	if areThereStudentsRegistered() {
		tempSliceToSort := make([]string, 0)
		students := SystemInstance.Students

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

	} else {
		utils.PressEnterToGoBack("\n** Empty! No student registered.")
	}
}

type displayAllParams struct {
	displayMsg string
	readInput  interface{}
}

func areThereStudentsRegistered() bool {
	return len(SystemInstance.Students) > 0
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

func GetStudentNameAndGrades(studentInfo string) (string, string) {

	parts := strings.Fields(studentInfo)
	var gradeStartIndex int

	for i := 0; i < len(parts); i++ {
		if _, err := strconv.Atoi(parts[i]); err == nil {
			gradeStartIndex = i
			break
		}
	}

	var studentName, grades string

	if gradeStartIndex > 0 {
		studentName = strings.Join(parts[1:gradeStartIndex], " ")
		grades = strings.Join(parts[gradeStartIndex:], " ")
	} else {
		studentName = strings.Join(parts[1:], " ")
		grades = ""
	}

	return studentName, grades
}

func getStudentByID(studentID int) (*models.Student, bool) {
	student, exists := SystemInstance.Students[studentID]
	return student, exists
}

func getNextAvailableID() int {
	studentIDs := make([]int, 0)
	students := SystemInstance.Students

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
			return startID
		}
	}

	return len(studentIDs) + 1
}

func Clear() {
	answer := readYesOrNo("This will delete all data save in the database. Are you sure? ")
	if answer && SystemInstance.ClearAll() {
		utils.SetSuccessMsg("\n** Operação realizada com sucesso! **")
	} else {
		slog.Error("** Operação não realizada! **")
	}
}
