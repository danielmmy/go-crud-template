package tools

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/danielmmy/templates/gocrud/api"
)

type mockDb struct{}

var mockedUserdata = map[string]*api.Userdata{
	"user1": {Username: "user1", Name: "User1", Age: 10},
	"user2": {Username: "user2", Name: "User2", Age: 20},
}

// setup mocks connection to db.
// it returns a random error 1/20 times.
func (m *mockDb) setup() error {
	// add random error
	n := rand.Intn(20)
	if n == 1 {
		return errors.New("random error")
	}

	return nil
}

// ListUser mocks a multi search call to the database.
// this simple mocked implementation does not support filters or paging.
// a delay is added to simulate the db query.
// it returns a list of user information.
func (m *mockDb) ListUsers() []*api.Userdata {
	// add delay
	ticker := time.NewTicker(1 * time.Second)
	<-ticker.C

	// convert mocked data to slice
	users := make([]*api.Userdata, 0, len(mockedUserdata))
	for _, v := range mockedUserdata {
		users = append(users, v)
	}

	return users
}

// GetUser mocks a query for one element of the database by its key.
// it return the user info if found, nil otherwise.
func (m *mockDb) GetUser(username string) *api.Userdata {
	user, ok := mockedUserdata[username]
	if !ok {
		return nil
	}

	return user
}

// AddUser mocks a db insertion.
// it returns an http code and any possible error.
func (m *mockDb) AddUser(user *api.Userdata) (int, error) {
	if _, ok := mockedUserdata[user.Username]; ok {
		return http.StatusBadRequest, fmt.Errorf("user %s exists", user.Username)
	}

	mockedUserdata[user.Username] = user
	return http.StatusCreated, nil
}

// UpdateUser mocks a db update.
// it returns an http code and any possible error.
func (m *mockDb) UpdateUser(user *api.Userdata) (int, error) {
	localUser, ok := mockedUserdata[user.Username]
	if !ok {
		return http.StatusNotFound, fmt.Errorf("user %s not found", user.Username)
	}

	localUser.Name = user.Name
	localUser.Age = user.Age
	return http.StatusOK, nil
}

// DeleteUser mocks a db element exclusion by its key
func (m *mockDb) DeleteUser(username string) {
	delete(mockedUserdata, username)
}
