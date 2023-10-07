package tools

import (
	"github.com/danielmmy/templates/gocrud/api"
	"github.com/sirupsen/logrus"
)

// DatabaseInterface is the interface for user db operations.
type DatabaseInterface interface {
	setup() error
	ListUsers() []*api.Userdata
	GetUser(string) *api.Userdata
	AddUser(*api.Userdata) (int, error)
	UpdateUser(*api.Userdata) (int, error)
	DeleteUser(string)
}

// NewDatabaseConn mocks a new db connection.
// it returns a db interface and any possible errors.
func NewDatabaseConn() (*DatabaseInterface, error) {
	var db DatabaseInterface = new(mockDb)
	if err := db.setup(); err != nil {
		logrus.Error(err)
		return nil, err
	}

	return &db, nil
}
