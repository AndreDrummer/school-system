package routes

import (
	"fmt"
	"net/http"
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
	fmt.Println("Guenta que vai ta tudo ai..")

	httputils.SendResponse(
		w,
		httputils.Response{Data: "Guenta molekão.."},
		http.StatusOK,
	)
}
