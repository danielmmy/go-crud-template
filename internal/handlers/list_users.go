package handlers

import (
	"net/http"

	"github.com/danielmmy/templates/gocrud/internal/tools"
	"github.com/sirupsen/logrus"
)

// listUsers handles user list requests.
func (app *App) listUsers(w http.ResponseWriter, r *http.Request) {
	// connect to db
	db, err := tools.NewDatabaseConn()
	if err != nil {
		logrus.Error(err)
		app.writeResponse(w, http.StatusInternalServerError, errInternalServerError)
		return
	}

	// recover user list
	users := (*db).ListUsers()

	// respond
	if err := app.writeResponse(w, http.StatusOK, users); err != nil {
		logrus.Error(err)
	}
}
