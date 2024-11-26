package routes

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Students(r chi.Router) {
	r.Get("/students", listAll)
}

var listAll = func(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	fmt.Println("Guenta que vai ta tudo ai..")
	w.WriteHeader(200)
	w.Write([]byte("Guenta molekão.."))
}
