package handlers

import (
	"net/http"

	"github.com/danielmmy/templates/gocrud/internal/tools"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

// deleteUser handles user deletion requests
func (app *App) deleteUser(w http.ResponseWriter, r *http.Request) {
	// get username
	username := chi.URLParam(r, "username")

	// connect to db
	db, err := tools.NewDatabaseConn()
	if err != nil {
		logrus.Error(err)
		app.writeResponse(w, http.StatusInternalServerError, errInternalServerError)
		return
	}

	// delete user and respond
	(*db).DeleteUser(username)
	if err := app.writeResponse(w, http.StatusOK, nil); err != nil {
		logrus.Error(err)
	}
}
