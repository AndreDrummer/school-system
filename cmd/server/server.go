package server

import (
	"log/slog"
	"net/http"
	"school-system/cmd/server/router"
	"time"
)

func Run() {
	serverHanlder := router.Handler()
	server := &http.Server{
		Addr:         ":8080",
		Handler:      serverHanlder,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  1 * time.Minute,
	}

	err := server.ListenAndServe()

	if err != nil {
		slog.Error("Could not initialize server..")
	}
}
