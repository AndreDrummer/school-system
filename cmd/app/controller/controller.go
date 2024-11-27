package schoolsystem

import (
	"fmt"
	"log"
	"log/slog"
	"school-system/cmd/app/db"
	dbutils "school-system/cmd/app/db/utils"
	"school-system/cmd/app/domain"
	"strconv"
	"strings"
)

var instance = &domain.ClassRoom{
	Students:            make(map[int]*domain.Student),
	StudentsQty:         0,
	MinimumPassingGrade: 60,
}

func Init() {
	students, err := AllStudents()

	if err != nil {
		slog.Error(fmt.Sprintf("error %v initializing system", err.Error()))
		return
	}

	for _, student := range students {
		instance.StudentsQty++
		instance.Students[student.ID] = student
	}
}

func AllStudents() ([]*domain.Student, error) {
	content, err := db.GetAll()

	if err != nil {
		return []*domain.Student{}, err
	}

	list := make([]*domain.Student, len(content))

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

			newStudent := &domain.Student{
				ID:     studentID,
				Grades: dbutils.ConvertGradesToIntSlice(grades),
				Name:   studentName,
			}

			list[i] = newStudent
		}
	}

	return list, nil
}

func AddStudent(student *domain.Student) (bool, error) {
	return instance.AddStudent(student)
}

func AddGrade(studentID, grade int) (bool, error) {
	return instance.AddGrade(studentID, grade)
}

func RemoveStudent(studentID int) (bool, error) {
	return instance.RemoveStudent(studentID)
}

func CalculateAverage(studentID int) (int, error) {
	return instance.CalculateAverage(studentID)
}

func GetStudentByID(studentID int) (*domain.Student, bool) {
	student, ok := instance.Students[studentID]

	return student, ok
}

func CheckPassOrFail(studentID int) bool {
	return instance.CheckPassOrFail(studentID)
}

func ClearAll() (bool, error) {
	return instance.ClearAll()
}
