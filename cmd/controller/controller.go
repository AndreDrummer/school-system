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

func GetStudentByID(studentID int) (*models.Student, bool) {
	student, ok := classRoomInstance.Students[studentID]

	return student, ok
}

func CalculateAverage(studentID int) (int, error) {
	avg, err := classRoomInstance.CalculateAverage(studentID)

	if err != nil {
		student, ok := GetStudentByID(studentID)
		if !ok {
			return 0, err
		}

		avg = student.GetAverage()
	}

	return avg, nil
}

func CheckPassOrFail(studentID int) bool {
	return classRoomInstance.CheckPassOrFail(studentID)
}

func ClearAll() error {
	return classRoomInstance.ClearAll()
}
