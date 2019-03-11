// File: profile.go
// File Created: 07 Mar 2019 16:10
// By Maxence Moutoussamy <maxence.moutoussamy1@gmail.com>

package api

import (
	"net/http"
	"server/internal/dao"
	"strconv"

	"github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

type ProfileService struct {
	manager *dao.ProfileManager
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
