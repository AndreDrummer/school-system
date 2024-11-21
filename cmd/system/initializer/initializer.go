package student_system_initializer

import (
	"fmt"
	"log"
	"os"
	"school-system/cmd/Utils/file_handler"
	"school-system/cmd/structs"
	panel "school-system/cmd/system"
	"school-system/cmd/system/controller"
	"strconv"
	"strings"
)

var (
	systemInstance *controller.System
)

func initSystem() {
	systemInstance = controller.NewSystem()
}

func convertGradesToInt(grades string) []int {
	gradeStringSlice := strings.Fields(grades)
	gradeIntSlice := make([]int, 0)

	for _, v := range gradeStringSlice {
		gradeInt, err := strconv.Atoi(v)

		if err != nil {
			log.Fatal(err)
		}

		gradeIntSlice = append(gradeIntSlice, gradeInt)
	}

	return gradeIntSlice
}

func loadStudentsFromDB() {
	dbFile, err := os.OpenFile(controller.DBFilename, os.O_RDWR, 0644)

	if err != nil {
		log.Fatal(err)
	}

	dbFileContent := file_handler.GetFileContent(dbFile)

	// Remove any empty line that may exists.
	file_handler.OverrideFileContent(dbFile, dbFileContent)

	if len(dbFileContent) > 0 {
		for _, v := range dbFileContent {
			studentIDString := strings.Split(v, ".")[0]
			studentID, err := strconv.Atoi(studentIDString)
			if v == "" {
				continue
			}
			studentName, grades := controller.GetStudentNameAndGrades(v)

			if err != nil {
				log.Fatal(err)
			}

			newStudent := &structs.Student{
				ID:     studentID,
				Grades: convertGradesToInt(grades),
				Name:   studentName,
			}

			systemInstance.StudentsQty++
			systemInstance.Students[studentID] = newStudent
		}
	}
}

func createDBFile(filename string) error {
	_, err := os.OpenFile(filename, os.O_CREATE, 0644)

	if err != nil {
		fmt.Printf("ERROR creating file %v\n", filename)
		return err
	}

	return nil
}

// Fake DB: All is based on files
func initDB() {

	_, errorReadingFile := os.ReadFile(controller.DBFilename)

	if errorReadingFile != nil {
		errorCreatingFile := createDBFile(controller.DBFilename)

		if errorCreatingFile != nil {
			log.Fatal(errorCreatingFile)
		}

		errorReadingFile = nil
	}
}

func Initialize() {
	initSystem()

	initDB()
	loadStudentsFromDB()

	panel.Start(systemInstance)
}
