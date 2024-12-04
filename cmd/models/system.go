package models

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

func (c *ClassRoom) CheckPassOrFail(student Student) bool {
	return student.GetAverage() > c.MinimumPassingGrade
}

func (c *ClassRoom) RemoveStudent(studentID int) {
	delete(c.Students, studentID)
}

func (c *ClassRoom) ClearAll() {
	for k := range c.Students {
		delete(c.Students, k)
	}
}
