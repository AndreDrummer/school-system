package repository

import (
	"log"
	"log/slog"
	"school-system/cmd/models"
	"school-system/cmd/repository/api"
	"school-system/cmd/repository/db"
	dbutils "school-system/cmd/repository/db/utils"
	"strconv"
	"strings"
)

var dbInstance = db.GetDB()

func GetAllStudents() ([]models.Student, error) {
	content, err := dbInstance.GetAll()

	if err != nil || len(content) == 0 {
		students, err := api.GetAll()

		if err != nil {
			return []models.Student{}, err
		}

		documents := make([]db.Document, len(students))
		for i, v := range students {
			documents[i] = db.Document(v)
		}

		err = dbInstance.InsertAll(documents)

		if err != nil {
			slog.Warn("Was not possible cache list of students.")
		}

		return students, nil
	} else {
		list := make([]models.Student, len(content))

		if len(content) > 0 {
			for i, v := range content {
				studentIDString := strings.Split(v, " ")[0]
				studentID, err := strconv.Atoi(studentIDString)
				if v == "" {
					continue
				}
				studentName, grades := dbutils.GetStudentNameAndGrades(v)

				if err != nil {
					log.Fatal(err)
				}

				newStudent := models.Student{
					ID:     studentID,
					Grades: dbutils.ConvertGradesToIntSlice(grades),
					Name:   studentName,
				}

				list[i] = newStudent
			}
		}

		return list, nil
	}
}

func AddStudent(student models.Student) (models.Student, error) {
	newAddedStudent, err := api.AddStudent(student)
	emptyStudent := models.Student{}

	if err != nil {
		return emptyStudent, err
	}

	err = dbInstance.Insert(newAddedStudent)

	if err != nil {
		return newAddedStudent, nil
	}

	return emptyStudent, err
}

func UpdateStudent(student models.Student) error {
	err := api.UpdateStudent(student)

	if err != nil {
		return err
	}

	return dbInstance.Update(student.ID, student)
}

func RemoveStudent(studentID int) error {
	err := api.RemoveStudent(studentID)

	if err != nil {
		return err
	}

	return dbInstance.Delete(studentID)
}
