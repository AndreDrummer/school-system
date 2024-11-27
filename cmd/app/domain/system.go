package domain

import (
	"fmt"
	"school-system/cmd/app/db"
)

type ClassRoom struct {
	Students            map[int]*Student
	StudentsQty         int
	MinimumPassingGrade int
}

func (system *ClassRoom) AddStudent(newStudent *Student) (bool, error) {
	system.Students[newStudent.ID] = newStudent
	system.StudentsQty++

	return db.Insert(*newStudent)
}

func (system *ClassRoom) AddGrade(studentID, grade int) (bool, error) {
	student := system.Students[studentID]
	student.AddGrade(grade)

	return db.Update(studentID, *student)
}

func (system *ClassRoom) RemoveStudent(studentID int) (bool, error) {
	delete(system.Students, studentID)

	return db.Delete(studentID)
}

func (system *ClassRoom) CalculateAverage(studentID int) (int, error) {
	student, ok := system.Students[studentID]

	if ok {
		return student.GetAverage(), nil
	}

	return 0, fmt.Errorf("Student of ID %d does not exists", studentID)
}

func (system *ClassRoom) CheckPassOrFail(studentID int) bool {
	student, ok := system.Students[studentID]

	if ok {
		return student.GetAverage() > system.MinimumPassingGrade
	}

	return false
}

func (system *ClassRoom) ClearAll() (bool, error) {
	return db.Clear()
}
