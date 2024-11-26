package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Students(r chi.Router) {

	r.Get("/students", listAll)
}

var listAll = func(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
}
