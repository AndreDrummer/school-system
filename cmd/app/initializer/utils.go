package initializer

import (
	"log"
	"school-system/cmd/app/controller"
	"school-system/cmd/app/db"
	"school-system/cmd/app/models"
	"strconv"
	"strings"
)

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
	content := db.GetAll()

	if len(content) > 0 {
		for _, v := range content {
			studentIDString := strings.Split(v, " ")[0]
			studentID, err := strconv.Atoi(studentIDString)
			if v == "" {
				continue
			}
			studentName, grades := controller.GetStudentNameAndGrades(v)

			if err != nil {
				log.Fatal(err)
			}

			newStudent := &models.Student{
				ID:     studentID,
				Grades: convertGradesToInt(grades),
				Name:   studentName,
			}

			controller.SystemInstance.StudentsQty++
			controller.SystemInstance.Students[studentID] = newStudent
		}
	}
}
