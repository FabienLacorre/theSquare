// File: connection.go
// File Created: 09 Mar 2019 09:06
// By Maxence Moutoussamy <maxence.moutoussamy1@gmail.com>

package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"server/internal"
	"server/internal/dao"
	"strings"

	"github.com/sirupsen/logrus"

	sessions "github.com/goincremental/negroni-sessions"
)

type ConnectionService struct {
	profileManager *dao.ProfileManager
}

type signInForm struct {
	Login     string `json:"login"`
	Password  string `json:"password"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	BirthDate string `json:"birthDate"`
	Country   string `json:"country"`
	City      string `json:"city"`
}

func NewConnectionService(profileManager *dao.ProfileManager) *ConnectionService {
	return &ConnectionService{profileManager}
}

func (s *ConnectionService) Login(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(http.StatusOK)
}

func (s *ConnectionService) Logout(rw http.ResponseWriter, req *http.Request) {
	session := sessions.GetSession(req)
	session.Set(internal.ConnectedKey, false)
	rw.WriteHeader(http.StatusOK)
}

func (s *ConnectionService) SignIn(rw http.ResponseWriter, req *http.Request) {
	var form signInForm

	defer req.Body.Close()

	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logrus.WithError(err).Error("cannot read body data")
		http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if err := json.Unmarshal(data, &form); err != nil {
		http.Error(rw, fmt.Sprintf("invalid json data: %v", err), http.StatusBadRequest)
		return
	}

	if err := form.verify(); err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	if isExist, err := s.profileManager.IsAccountExists(form.Login); err != nil {
		logrus.
			WithError(err).
			WithField("login", form.Login).
			Error("cannot verify if an account exists")
		http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	} else if isExist {
		rw.Write([]byte(fmt.Sprintf("login '%s' already used", form.Login)))
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	p := &dao.Profile{
		Login:     form.Login,
		Password:  form.Password,
		Firstname: form.Name,
		Lastname:  form.Surname,
		Birthday:  form.BirthDate,
		Country:   form.Country,
		City:      form.City,
	}

	if err := s.profileManager.Create(p); err != nil {
		logrus.
			WithError(err).
			Error("cannot create account")
		http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (f *signInForm) verify() error {
	errF := func(field string) error { return fmt.Errorf("the %s is empty", field) }

	f.BirthDate = strings.TrimSpace(f.BirthDate)
	if len(f.BirthDate) == 0 {
		return errF("birth date")
	}

	f.Login = strings.TrimSpace(f.Login)
	if len(f.Login) == 0 {
		return errF("login")
	}

	f.Password = strings.TrimSpace(f.Password)
	if len(f.Password) == 0 {
		return errF("password")
	}

	f.Name = strings.TrimSpace(f.Name)
	if len(f.Name) == 0 {
		return errF("name")
	}

	f.Surname = strings.TrimSpace(f.Surname)
	if len(f.Surname) == 0 {
		return errF("surname")
	}

	f.Country = strings.TrimSpace(f.Country)
	if len(f.Country) == 0 {
		return errF("country")
	}

	f.City = strings.TrimSpace(f.City)
	if len(f.City) == 0 {
		return errF("city")
	}

	return nil
}
