// File: profile.go
// File Created: 07 Mar 2019 16:10
// By Maxence Moutoussamy <maxence.moutoussamy1@gmail.com>

package api

import (
	"encoding/json"
	"net/http"
	"server/internal"
	"server/internal/dao"
	"strconv"

	sessions "github.com/goincremental/negroni-sessions"
	"github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

type ProfileService struct {
	manager *dao.ProfileManager
}

type profile struct {
	Login     string `json:"login"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Birthday  string `json:"birthday"`
	City      string `json:"city"`
	Country   string `json:"country"`
}

func NewProfileService(manager *dao.ProfileManager) *ProfileService {
	return &ProfileService{manager}
}

func (s *ProfileService) GetByID(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "id must be a number", http.StatusBadRequest)
		return
	}

	if _, err := s.manager.GetProfileWithID(id); err != nil {
		logrus.WithError(err).Error("cannot get profile with id")
		http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rw.Write([]byte("Everything ok"))
}

func (s *ProfileService) GetCurrent(rw http.ResponseWriter, req *http.Request) {
	session := sessions.GetSession(req)

	login, ok := session.Get(internal.LoginKey).(string)
	if !ok {
		internalServerError(rw, "cannot get login from session (invalid type)", nil)
		return
	}

	daoProfile, err := s.manager.GetByLogin(login)
	if err != nil {
		if err == dao.ErrNotFound {
			http.NotFound(rw, req)
			return
		}
		internalServerError(rw, "cannot GetByLogin", err)
		return
	}

	profile := profile{
		Login:     daoProfile.Login,
		FirstName: daoProfile.Firstname,
		LastName:  daoProfile.Lastname,
		Birthday:  daoProfile.Birthday,
		City:      daoProfile.City,
		Country:   daoProfile.Country,
	}

	datas, err := json.Marshal(profile)
	if err != nil {
		internalServerError(rw, "cannot marshal profile type", err)
		return
	}

	rw.Header().Set("Content-type", "application/json")
	rw.Write(datas)
}
