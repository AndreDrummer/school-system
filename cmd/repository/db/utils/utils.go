package dbutils

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
)

func ConvertGradesToIntSlice(grades string) []int {
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

func ConvertStructToString(s interface{}) string {
	structValue := reflect.ValueOf(s)
	structType := reflect.TypeOf(s)
	var builder strings.Builder
	for i := 0; i < structType.NumField(); i++ {
		value := structValue.Field(i)
		if value.CanInt() || value.CanConvert(reflect.TypeOf(string(""))) {
			builder.WriteString(fmt.Sprintf("%v ", value))
		}
		if value.CanConvert(reflect.TypeOf([]int{})) {
			convertedValue := value.Convert(reflect.TypeOf([]int{}))
			result, ok := convertedValue.Interface().([]int)
			if ok {
				for _, v := range result {
					builder.WriteString(fmt.Sprintf("%v ", strconv.Itoa(v)))
				}
			}
		}
	}
	return builder.String()
}
