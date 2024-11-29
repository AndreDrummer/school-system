package controller

import "school-system/cmd/models"

var instance = &models.ClassRoom{
	Students:            make(map[int]models.Student),
	StudentsQty:         0,
	MinimumPassingGrade: 60,
}

func AllStudents() ([]models.Student, error) {
	list := make([]models.Student, 0)

	return list, nil
}

func AddStudent(student models.Student) (bool, error) {
	return false, nil
}

func AddGrade(studentID, grade int) (bool, error) {
	return instance.AddGrade(studentID, grade)
}

func UpdateStudent(student models.Student) (bool, error) {
	return instance.UpdateStudent(student)
}

func RemoveStudent(studentID int) (bool, error) {
	return instance.RemoveStudent(studentID)
}

func GetStudentByID(studentID int) (models.Student, bool) {
	student, ok := instance.Students[studentID]

	return student, ok
}

func CalculateAverage(studentID int) (int, error) {
	avg, err := instance.CalculateAverage(studentID)

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
	return instance.CheckPassOrFail(studentID)
}

func ClearAll() (bool, error) {
	return instance.ClearAll()
}
