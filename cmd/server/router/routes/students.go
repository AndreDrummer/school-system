package routes

import (
	"log/slog"
	"net/http"
	schoolsystem "school-system/cmd/app/controller"
	httputils "school-system/cmd/server/http"

	"github.com/go-chi/chi/v5"
)

func HandleRequests(w http.ResponseWriter, r *http.Request) {

}

func Students(r chi.Router) {
	r.Get("/students", listAll)
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
