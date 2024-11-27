package schoolsystem

import (
	"log"
	"school-system/cmd/app/db"
	dbutils "school-system/cmd/app/db/utils"
	"school-system/cmd/app/domain"
	"strconv"
	"strings"
)

var Instance = &domain.ClassRoom{
	Students:            make(map[int]*domain.Student),
	StudentsQty:         0,
	MinimumPassingGrade: 60,
}

func GetAllStudents() []*domain.Student {
	content := db.GetAll()
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

	return list
}

func LoadStudentsFromDB() {
	students := GetAllStudents()

	for _, student := range students {
		Instance.StudentsQty++
		Instance.Students[student.ID] = student
	}
}
