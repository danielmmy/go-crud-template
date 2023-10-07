package handlers

import (
	"net/http"

	"github.com/danielmmy/templates/gocrud/api"
	"github.com/danielmmy/templates/gocrud/internal/tools"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

// updateUser handles user update requests
func (app *App) updateUser(w http.ResponseWriter, r *http.Request) {
	// read user data from body
	user := new(api.Userdata)
	if err := app.readJson(w, r, user); err != nil {
		logrus.Error(err)
		app.writeResponse(w, http.StatusBadRequest, err)
		return
	}

	// get username
	user.Username = chi.URLParam(r, "username")

	// connect to db
	db, err := tools.NewDatabaseConn()
	if err != nil {
		logrus.Error(err)
		app.writeResponse(w, http.StatusInternalServerError, errInternalServerError)
		return
	}

	// update user and respond
	code, err := (*db).UpdateUser(user)
	if err := app.writeResponse(w, code, err); err != nil {
		logrus.Error(err)
	}
}
