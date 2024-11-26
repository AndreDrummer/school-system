package dbutils

import (
	"log"
	"strconv"
	"strings"
)

func ConvertGradesToInt(grades string) []int {
	gradeStringSlice := strings.Fields(grades)
	gradeIntSlice := make([]int, 0)

	for _, v := range gradeStringSlice {
		gradeInt, err := strconv.Atoi(v)

		if err != nil {
			log.Fatal(err)
		}

		gradeIntSlice = append(gradeIntSlice, gradeInt)
	}

	return gradeIntSlice
}

func GetStudentNameAndGrades(studentInfo string) (string, string) {

	parts := strings.Fields(studentInfo)
	var gradeStartIndex int

	for i := 1; i < len(parts); i++ {
		if _, err := strconv.Atoi(parts[i]); err == nil {
			gradeStartIndex = i
			break
		}
	}

	var studentName, grades string

	if gradeStartIndex > 0 {
		studentName = strings.Join(parts[1:gradeStartIndex], " ")
		grades = strings.Join(parts[gradeStartIndex:], " ")
	} else {
		studentName = strings.Join(parts[1:], " ")
		grades = ""
	}

	return studentName, grades
}
