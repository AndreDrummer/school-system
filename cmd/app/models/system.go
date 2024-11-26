package models

import (
	"fmt"
	"school-system/cmd/app/db"
)

type System struct {
	Students            map[int]*Student
	StudentsQty         int
	MinimumPassingGrade int
}

func (system *System) AddStudent(newStudent *Student) bool {
	system.Students[newStudent.ID] = newStudent
	system.StudentsQty++

	return db.Insert(*newStudent)
}

func (system *System) AddGrade(studentID, grade int) bool {
	student := system.Students[studentID]
	student.AddGrade(grade)

	return db.Update(studentID, *student)
}

func (system *System) RemoveStudent(studentID int) bool {
	delete(system.Students, studentID)

	return db.Delete(studentID)
}

func (system *System) CalculateAverage(studentID int) (int, error) {
	student, ok := system.Students[studentID]

	if ok {
		return student.GetAverage(), nil
	}

	return 0, fmt.Errorf("Student of ID %d does not exists", studentID)
}

func (system *System) CheckPassOrFail(studentID int) bool {
	student, ok := system.Students[studentID]

	if ok {
		return student.GetAverage() > system.MinimumPassingGrade
	}

	return false
}

func (system *System) ClearAll() bool {
	return db.Clear()
}
