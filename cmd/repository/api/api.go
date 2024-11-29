package api

import (
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
