package main

import (
	"net/http"
	"os"

	"github.com/danielmmy/templates/gocrud/internal/handlers"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetReportCaller(true)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app := handlers.NewApp(handlers.WithPort(port))
	srv := http.Server{
		Addr:    app.GetAddr(),
		Handler: app.Handle(),
	}

	logrus.Printf("stating server on %s...", app.GetAddr())
	logrus.Error(srv.ListenAndServe())
}
