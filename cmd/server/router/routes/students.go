package routes

import (
	"log/slog"
	"net/http"
	schoolsystem "school-system/cmd/app/controller"
	httputils "school-system/cmd/server/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func HandleRequests(w http.ResponseWriter, r *http.Request) {

}

func Students(r chi.Router) {
	r.Get("/students", listAll)
	r.Get("/students/{id:[0-9]+}", getByID)
}

var listAll = func(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	students, err := schoolsystem.AllStudents()

	if err != nil {
		slog.Error(err.Error())
		httputils.SendResponse(
			w,
			httputils.Response{Error: "Could not get the students list."},
			http.StatusInternalServerError,
		)
		return
	}

	httputils.SendResponse(
		w,
		httputils.Response{Data: students},
		http.StatusOK,
	)
}

var getByID = func(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	studentID := chi.URLParam(r, "id")
	studentIDInt, err := strconv.Atoi(studentID)

	if err != nil {
		slog.Error(err.Error())
		httputils.SendResponse(
			w,
			httputils.Response{Error: "A problem occured tryna get user"},
			http.StatusBadRequest,
		)
		return
	}

	student, ok := schoolsystem.GetStudentByID(studentIDInt)

	if !ok {
		httputils.SendResponse(
			w,
			httputils.Response{Data: "Student not found"},
			http.StatusNotFound,
		)
		return
	}

	httputils.SendResponse(
		w,
		httputils.Response{Data: student},
		http.StatusOK,
	)
}
