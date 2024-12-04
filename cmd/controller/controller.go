package controller

import (
	"log/slog"
	"school-system/cmd/models"
	"school-system/cmd/repository"
)

var classRoomInstance = &models.ClassRoom{
	Students:            make(map[int]*models.Student),
	StudentsQty:         0,
	MinimumPassingGrade: 60,
}

func AllStudents() ([]models.Student, error) {
	students := classRoomInstance.AllStudents()

	if len(students) == 0 {
		slog.Info("Looking into the database for data...\n\n")
		apiStudentsList, err := repository.GetAllStudents()

		students = apiStudentsList

		classRoomInstance.AddAllStudents(students)

		if err != nil {
			return students, err
		}
	}

	return students, nil
}

func AddStudent(student models.Student) error {
	newStudent, err := repository.AddStudent(student)

	if err != nil {
		return err
	}

	classRoomInstance.AddStudent(newStudent)
	return nil
}

func AddGrade(student *models.Student, grade int) error {
	student.AddGrade(grade)
	return repository.UpdateStudent(*student)
}

func RemoveStudent(studentID int) error {
	classRoomInstance.RemoveStudent(studentID)
	return repository.RemoveStudent(studentID)
}

func GetStudentByID(studentID int) *models.Student {
	student := classRoomInstance.Students[studentID]

	return student
}

func CheckPassOrFail(student models.Student) bool {
	return classRoomInstance.CheckPassOrFail(student)
}

func ClearAll() error {
	classRoomInstance.ClearAll()

	return repository.ClearAll()
}
