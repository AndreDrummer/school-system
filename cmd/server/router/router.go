package router

import (
	"net/http"
	"school-system/cmd/server/router/routes"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Handler() http.Handler {
	handler := chi.NewMux()

	handler.Use(middleware.Recoverer)
	handler.Use(middleware.RequestID)
	handler.Use(middleware.Logger)

	handler.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			routes.Students(r)
		})
	})

	return handler
}
