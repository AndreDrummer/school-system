package controller

import (
	"school-system/cmd/models"
	"school-system/cmd/repository"
)

var classRoomInstance = &models.ClassRoom{
	Students:            make(map[int]models.Student),
	StudentsQty:         0,
	MinimumPassingGrade: 60,
}

func AllStudents() ([]models.Student, error) {
	return repository.GetAllStudents()
}

func AddStudent(student models.Student) (bool, error) {
	newStudent, err := repository.AddStudent(student)

	if err != nil {
		return false, err
	}

	return classRoomInstance.AddStudent(newStudent)
}

func AddGrade(studentID, grade int) (bool, error) {
	return classRoomInstance.AddGrade(studentID, grade)
}

func UpdateStudent(student models.Student) (bool, error) {
	return classRoomInstance.UpdateStudent(student)
}

func RemoveStudent(studentID int) (bool, error) {
	return classRoomInstance.RemoveStudent(studentID)
}

func GetStudentByID(studentID int) (models.Student, bool) {
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

func ClearAll() (bool, error) {
	return classRoomInstance.ClearAll()
}
