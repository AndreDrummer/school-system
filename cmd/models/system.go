package models

import (
	"errors"
)

type ClassRoom struct {
	Students            map[int]*Student
	StudentsQty         int
	MinimumPassingGrade int
}

func (c *ClassRoom) AllStudents() []Student {
	studentsMap := c.Students
	studentsList := make([]Student, 0)

	for _, v := range studentsMap {
		studentsList = append(studentsList, *v)
	}

	return studentsList
}

func (c *ClassRoom) AddAllStudents(students []Student) {
	for _, student := range students {
		c.AddStudent(student)
	}
}

func (c *ClassRoom) AddStudent(newStudent Student) {
	c.StudentsQty++
	c.Students[newStudent.ID] = &newStudent
}

func (c *ClassRoom) RemoveStudent(studentID int) {
	delete(c.Students, studentID)
}

func (c *ClassRoom) CalculateAverage(studentID int) (int, error) {
	student, ok := c.Students[studentID]

	if ok {
		return student.GetAverage(), nil
	}

	return 0, errors.New("not found")
}

func (c *ClassRoom) CheckPassOrFail(studentID int) bool {
	student, ok := c.Students[studentID]

	if ok {
		return student.GetAverage() > c.MinimumPassingGrade
	}

	return false
}

func (c *ClassRoom) ClearAll() error {
	return nil
}
