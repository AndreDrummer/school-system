package view

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	schoolsystem "school-system/cmd/app/controller"
	"school-system/cmd/app/domain"
	utils "school-system/cmd/app/utils"
	"sort"
	"strings"
)

var inputRead *bufio.Reader = bufio.NewReader(os.Stdin)

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

	studentID, err := getNextAvailableID()

	if err != nil {
		slog.Error(err.Error())
		return
	}

	newStudent := &domain.Student{
		ID:     studentID,
		Grades: make([]int, 0),
		Name:   studentName,
	}

	if ok, err := schoolsystem.AddStudent(newStudent); ok {
		utils.ClearConsole()
		utils.SetSuccessMsg(fmt.Sprintf("Student %v Added!", newStudent.Name))
	} else {
		slog.Error(err.Error())
	}

}

func AddGrade() {
	if ok, err := areThereStudentsRegistered(); ok {
		fmt.Print("What student would you like to add a grade?\n\n")
		studentID := readStudentID()
		student, studentExists := getStudentByID(studentID)

		if studentExists {
			grade := readGrade()
			if grade >= 0 {
				if ok, err := schoolsystem.AddGrade(studentID, grade); ok {
					utils.SetSuccessMsg(fmt.Sprintf("Grade %v added to %v!", grade, student.Name))
				} else {
					slog.Error(err.Error())
				}
			} else {
				utils.ClearConsole()
			}
		}
	} else {
		if err != nil {
			slog.Error(err.Error())
		} else {
			utils.PressEnterToGoBack("\n** Empty! No student registered.")
		}
	}
}

func RemoveStudent() {
	if ok, err := areThereStudentsRegistered(); ok {
		studentID := readStudentID()
		student, _ := getStudentByID(studentID)

		if ok, err := schoolsystem.RemoveStudent(studentID); ok {
			utils.SetSuccessMsg(fmt.Sprintf("Student %v removed!", student.Name))
		} else {
			slog.Error(err.Error())
		}
	} else {
		if err != nil {
			slog.Error(err.Error())
		} else {
			utils.PressEnterToGoBack("\n** Empty! No student registered.")
		}
	}
}

func CalculateAverage() {
	if ok, err := areThereStudentsRegistered(); ok {
		studentID := readStudentID()
		avg, err := schoolsystem.CalculateAverage(studentID)

		if err != nil {
			slog.Error(err.Error())
		} else {
			student, _ := schoolsystem.GetStudentByID(studentID)
			utils.PressEnterToGoBack(fmt.Sprintf("\nThe average of %s is %v.\n", student.Name, avg))
		}
	} else {
		if err != nil {
			slog.Error(err.Error())
		} else {
			utils.PressEnterToGoBack("\n** Empty! No student registered.")
		}
	}
}

func CheckPassOrFail() {
	if ok, err := areThereStudentsRegistered(); ok {
		studentID := readStudentID()

		approved := schoolsystem.CheckPassOrFail(studentID)
		var resultMsg string

		if approved {
			resultMsg = "has been approved! :)"

		} else {
			resultMsg = "has failed :(!"
		}

		student, _ := schoolsystem.GetStudentByID(studentID)
		utils.PressEnterToGoBack(fmt.Sprintf("\n%s %v.\n", student.Name, resultMsg))

	} else {
		if err != nil {
			slog.Error(err.Error())
		} else {
			utils.PressEnterToGoBack("\n** Empty! No student registered.")
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

	if ok, err := areThereStudentsRegistered(); ok {
		tempSliceToSort := make([]string, 0)
		students, err := schoolsystem.AllStudents()

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

	} else {
		if err != nil {
			slog.Error(err.Error())
		} else {
			utils.PressEnterToGoBack("\n** Empty! No student registered.")
		}
	}

	return nil
}

func Clear() {
	answer := readYesOrNo("This will delete all data save in the database. Are you sure? ")
	if answer {
		if ok, err := schoolsystem.ClearAll(); ok {
			utils.SetSuccessMsg("\n** Operação realizada com sucesso! **")
		} else {
			slog.Error(err.Error())
		}
	} else {
		slog.Error("** Operação não realizada! **")
	}
}

type displayAllParams struct {
	displayMsg string
	readInput  interface{}
}

func areThereStudentsRegistered() (bool, error) {
	students, err := schoolsystem.AllStudents()

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

func getStudentByID(studentID int) (*domain.Student, bool) {
	student, exists := schoolsystem.GetStudentByID(studentID)
	return student, exists
}

func getNextAvailableID() (int, error) {
	studentIDs := make([]int, 0)
	students, err := schoolsystem.AllStudents()

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
