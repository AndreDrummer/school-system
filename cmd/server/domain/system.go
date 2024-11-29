package domain

import (
	"errors"
	"school-system/cmd/server/db"
)

type ClassRoom struct {
	Students            map[int]*Student
	StudentsQty         int
	MinimumPassingGrade int
}

func (c *ClassRoom) AddStudent(newStudent *Student) (bool, error) {
	c.Students[newStudent.ID] = newStudent
	c.StudentsQty++

	return db.Insert(*newStudent)
}

func (c *ClassRoom) AddGrade(studentID, grade int) (bool, error) {
	student := c.Students[studentID]
	student.AddGrade(grade)

	return db.Update(studentID, *student)
}

func (c *ClassRoom) UpdateStudent(student *Student) (bool, error) {
	return db.Update(student.ID, *student)
}

func (c *ClassRoom) RemoveStudent(studentID int) (bool, error) {
	delete(c.Students, studentID)

	return db.Delete(studentID)
}

func (c *ClassRoom) CalculateAverage(studentID int) (int, error) {
	student, ok := c.Students[studentID]

	if ok {
		return student.GetAverage(), nil
	}

	return 0, errors.New("not found... searching on DB")
}

func (c *ClassRoom) CheckPassOrFail(studentID int) bool {
	student, ok := c.Students[studentID]

	if ok {
		return student.GetAverage() > c.MinimumPassingGrade
	}

	return false
}

func (c *ClassRoom) ClearAll() (bool, error) {
	return db.Clear()
}
