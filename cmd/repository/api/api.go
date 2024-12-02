package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	apperrors "school-system/cmd/errors"
	"school-system/cmd/models"
)

const (
	baseURL = "http://localhost:8080/api/v1"
)

type StudentListResponse struct {
	Data []models.Student `json:"data"`
}

type StudentResponse struct {
	Data models.Student `json:"data"`
}

func GetAll() ([]models.Student, error) {
	url := baseURL + "/students"
	emptyStudentList := []models.Student{}

	response, err := http.Get(url)

	if err != nil {
		slog.Error(fmt.Sprintf("making request: %s", err.Error()))
		return emptyStudentList, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		slog.Error(fmt.Sprintf("reading response: %s", err.Error()))
		return emptyStudentList, err
	}

	var students StudentListResponse
	err = json.Unmarshal(body, &students)

	if err != nil {
		slog.Error(fmt.Sprintf("Error reading response: %s", err.Error()))
		return emptyStudentList, &apperrors.JsonDecodingError{
			Type: fmt.Sprintf("%T", students),
			Err:  err,
		}
	}

	return students.Data, nil
}
func AddStudent(student models.Student) (models.Student, error) {
	emptyStudent := models.Student{}

	studentJson, err := json.Marshal(student)

	if err != nil {
		jsonEncondingError := &apperrors.JsonEncodingError{
			Type: fmt.Sprintf("%T", student),
			Err:  err,
		}

		slog.Error(jsonEncondingError.Error())
		return emptyStudent, jsonEncondingError
	}

	url := baseURL + "/students"
	response, err := http.Post(
		url,
		"application/json",
		bytes.NewBuffer([]byte(studentJson)),
	)

	if err != nil {
		slog.Error("adding new student")
		return emptyStudent, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		slog.Error("reading post response body")
		return emptyStudent, err
	}

	var justCreatedStudent models.Student
	err = json.Unmarshal(body, &justCreatedStudent)

	if err != nil {
		jsonDecodingError := &apperrors.JsonDecodingError{
			Type: fmt.Sprintf("%T", justCreatedStudent),
			Err:  err,
		}
		return emptyStudent, jsonDecodingError
	}

	return justCreatedStudent, nil
}
