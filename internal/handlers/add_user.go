package handlers

import (
	"net/http"

	"github.com/danielmmy/templates/gocrud/api"
	"github.com/danielmmy/templates/gocrud/internal/tools"
	"github.com/sirupsen/logrus"
)

// addUser handles user insertion requests
func (app *App) addUser(w http.ResponseWriter, r *http.Request) {
	// read user data from body
	user := new(api.Userdata)
	if err := app.readJson(w, r, user); err != nil {
		logrus.Error(err)
		app.writeResponse(w, http.StatusBadRequest, err)
		return
	}

	// connect to db
	db, err := tools.NewDatabaseConn()
	if err != nil {
		logrus.Error(err)
		app.writeResponse(w, http.StatusInternalServerError, errInternalServerError)
		return
	}

	// insert user and respond
	code, err := (*db).AddUser(user)
	if err := app.writeResponse(w, code, err); err != nil {
		logrus.Error(err)
	}
}
