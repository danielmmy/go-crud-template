package handlers

import (
	"fmt"
	"net/http"

	"github.com/danielmmy/templates/gocrud/internal/tools"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

func (app *App) getUser(w http.ResponseWriter, r *http.Request) {
	// get username from url params
	username := chi.URLParam(r, "username")

	// connect to db
	db, err := tools.NewDatabaseConn()
	if err != nil {
		logrus.Error(err)
		app.writeResponse(w, http.StatusInternalServerError, errInternalServerError)
		return
	}

	// query for user
	user := (*db).GetUser(username)

	// if empty respond not found
	if user == nil {
		if err := app.writeResponse(w, http.StatusNotFound, fmt.Errorf("user %s not found", username)); err != nil {
			logrus.Error(err)
		}
		return
	}

	// respond with user
	if err := app.writeResponse(w, http.StatusOK, user); err != nil {
		logrus.Error(err)
	}
}
